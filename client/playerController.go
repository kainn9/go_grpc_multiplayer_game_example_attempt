package main

import (
	"context"
	"io"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	pb "github.com/kainn9/grpc_game/proto"
	camera "github.com/melonfunction/ebiten-camera"
	"github.com/pborman/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	freePlay = false
)

type PlayerController struct {
	Stream    pb.PlayersService_PlayerLocationClient
	PlayerTex sync.RWMutex
	World     *World
	Cam       *camera.Camera
	Pid       string
	X         float64
	Y         float64
}

/*
Establishes stream w/ server and
creates PlayerController with/by

  - loading player sprites
  - generating pid
  - inits camera
  - has a mutex
  - inits stream/connection with pid
*/
func NewPlayerController() *PlayerController {

	LoadPlayerControllerSprites()

	pid := uuid.New()

	p := &PlayerController{
		Pid:       pid,
		PlayerTex: sync.RWMutex{},
		Cam:       camera.NewCamera(ScreenWidth, ScreenHeight, 0, 0, 0, 1),
	}

	p.Stream = p.NewStream()

	return p
}

/*
Loads the default player sprites
*/
func LoadPlayerControllerSprites() {
	playerSpriteIdleLeft = LoadImg("./sprites/knight/knightIdleLeft.png")
	playerSpriteIdleRight = LoadImg("./sprites/knight/knightIdleRight.png")

	playerSpriteWalkingRight = LoadImg("./sprites/knight/knightRunningRight.png")
	playerSpriteWalkingLeft = LoadImg("./sprites/knight/knightRunningLeft.png")
	playerSpriteJumpLeft = LoadImg("./sprites/knight/knightJumpLeft.png")
	playerSpriteJumpRight = LoadImg("./sprites/knight/knightJumpRight.png")
}

/*
Initializes stream with PID
*/
func (p *PlayerController) NewStream() pb.PlayersService_PlayerLocationClient {
		/*
		disable to skip ssl
		or just run make genSSL
		change value in ssl/ssl.sh
	*/
	tls := false 
	opts := []grpc.DialOption{}

	if tls {
		certFile := "../ssl/ca.crt"

		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Errr getting client cert %v\n", err)
		}
		
		opts = append(opts, grpc.WithTransportCredentials(creds))

	} else {

		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}


	
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	c := pb.NewPlayersServiceClient(conn)

	// TODO: Delete OR Keep?
	maxSizeOption := grpc.MaxCallRecvMsgSize(32 * 10e9)

	md := metadata.Pairs("pid", p.Pid)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := c.PlayerLocation(ctx, maxSizeOption)

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	return stream
}

/*
Listens for Player inputs during game update phase
*/
func (pc *PlayerController) inputListener() {

	if inpututil.IsKeyJustPressed(ebiten.Key4) {
		pc.inputHandler("swap")
	}

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		freePlay = !freePlay
	}

	// Free Play Cam
	// Also an example of a "Cam Hack"
	if freePlay {
		cam := pc.Cam

		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			moveX := float64(cam.X + devCamSpeed)
			moveY := float64(cam.Y)
			cam.SetPosition(moveX, moveY)
		}

		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			moveX := float64(cam.X - devCamSpeed)
			moveY := float64(cam.Y)
			cam.SetPosition(moveX, moveY)
		}

		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			moveX := float64(cam.X)
			moveY := float64(cam.Y - devCamSpeed)
			cam.SetPosition(moveX, moveY)
		}

		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			moveX := float64(cam.X)
			moveY := float64(cam.Y + devCamSpeed)
			cam.SetPosition(moveX, moveY)
		}

		if inpututil.IsKeyJustPressed(ebiten.Key3) {
			pc.World.Init(DevWorldBuilder)
			devPreview = !devPreview

			if !devPreview {
				pc.World.Space.Remove(pc.World.Space.Objects()...)
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.Key2) {
			useHeightRuler = !useHeightRuler
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyW) {
			devCamSpeed += 1
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyS) {
			devCamSpeed -= 1
		}

		pc.inputHandler("freePlay")
		return
	}

	// Non Free Play Listner stuff:
	isPressing := false

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		pc.inputHandler("keyRight")
		isPressing = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		pc.inputHandler("keyLeft")
		isPressing = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		pc.inputHandler("keySpace")
		isPressing = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		pc.inputHandler("keyDown")
		isPressing = true
	}
	if !isPressing {
		pc.inputHandler("nada")
	}
}

/*
Used by Input listener/playerController
to stream the inputs to server
*/
func (p *PlayerController) inputHandler(input string) {
	go func() {
		req := pb.PlayerReq{Id: p.Pid, Input: input}
		p.Stream.Send(&req)
	}()
}

/*
Logic to keep PlayerController camera
following player w/o exposing level boundaries
(needs to be cleaned up)
*/
func (pc *PlayerController) setCameraPosition() {
	gw := pc.World.Width
	gh := pc.World.Height

	if freePlay {
		return
	}

	x := (pc.X / 2)
	y := (pc.Y / 2)

	xBoundLeft := (pc.X - ScreenWidth/2) < 0
	xBoundBottom := (pc.Y + (ScreenHeight / 2)) > gh
	xBoundRight := (pc.X + ScreenWidth/2) > gw
	xBoundTop := (pc.Y - ScreenHeight/2) < 0

	if xBoundLeft && xBoundBottom {
		pc.Cam.SetPosition((ScreenWidth/2)-x, (ScreenHeight/2)+(gh-pc.Y))

	} else if xBoundLeft && xBoundTop {
		pc.Cam.SetPosition((ScreenWidth/2)-x, (ScreenHeight/2)-y)

	} else if xBoundRight && xBoundBottom {
		pc.Cam.SetPosition(x, (ScreenHeight/2)+(gh-pc.Y))

	} else if xBoundRight && xBoundTop {
		xOff := float64(gw - pc.X)
		pc.Cam.SetPosition(x-((ScreenWidth/2)-xOff), (ScreenHeight/2)-y)

	} else if xBoundLeft {
		pc.Cam.SetPosition((ScreenWidth/2)-x, y)

	} else if xBoundRight {
		xOff := float64(gw - pc.X)
		pc.Cam.SetPosition(x-((ScreenWidth/2)-xOff), y)

	} else if xBoundBottom {
		yOff := float64(gh) - float64(pc.Y)
		pc.Cam.SetPosition(x, y-((ScreenHeight/2)-yOff))

	} else if xBoundTop {
		pc.Cam.SetPosition(x, (ScreenHeight/2)-y)
	} else {
		pc.Cam.SetPosition(x, y)

	}
}

/*
Helper function for handling current player
state from server stream
*/
func CurrentPlayerHandler(pc *PlayerController, ps *pb.Player, p *Player) {
	cw := pc.World
	cw.PlayerController.X = ps.Lx
	cw.PlayerController.Y = ps.Ly

	if game.CurrentWorld != ps.World {

		newData := worldsMap[ps.World]

		UpdateWorldData(cw, &newData, ps.World)
	}

	DrawPlayer(cw, p, true)
}


func (pc *PlayerController) GetState() {
	world := pc.World
	wTex := &world.WorldTex
	ptex := &pc.PlayerTex

	go func() {
		for {
			ptex.Lock()
			res, err := pc.Stream.Recv()
			ptex.Unlock()

			if err == io.EOF {
				break
			} 

			if err != nil {
				log.Fatalf("Setting World State Error %v\n", err);
				break
			}
			
			// reg lock on insertion
			wTex.Lock()
			world.State = res.Players
			wTex.Unlock()
		}
	}()
}