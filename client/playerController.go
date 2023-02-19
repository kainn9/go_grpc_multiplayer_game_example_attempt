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
	"github.com/solarlune/resolv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type PlayerController struct {
	Stream    pb.PlayersService_PlayerLocationClient
	PlayerTex sync.RWMutex
	World     *World
	Pid       string
	X         float64
	Y         float64
	PlayerCam
}

type PlayerCam struct {
	*camera.Camera
	PlayerCamData
}

type PlayerCamData struct {
	yOff float64
	xOff float64
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

	pid := uuid.New()

	cam := PlayerCam{
		Camera: camera.NewCamera(ScreenWidth, ScreenHeight, 0, 0, 0, 1),
		PlayerCamData: PlayerCamData{
			yOff: 0,
			xOff: 0,
		},
	}

	p := &PlayerController{
		Pid:       pid,
		PlayerTex: sync.RWMutex{},
		PlayerCam: cam,
	}

	p.Stream = p.NewStream()

	return p
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

	connRef = conn

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
handles volume for now
*/
func updateVolumeIfNeeded() {
	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		volume128--
	}
	if ebiten.IsKeyPressed(ebiten.KeyX) {
		volume128++
	}
	if volume128 < 0 {
		volume128 = 0
	}
	if 128 < volume128 {
		volume128 = 128
	}

	if audPlayer != nil {
		audPlayer.SetVolume(float64(volume128) / 128)
	}

}

/*
Listens for Player inputs during game update phase
*/
func (pc *PlayerController) InputListener() {

	updateVolumeIfNeeded()

	if inpututil.IsKeyJustPressed(ebiten.Key0) {
		fullScreen = !fullScreen
		ebiten.SetFullscreen(fullScreen)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key4) {
		pc.inputHandler("swap")
	}

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		freePlay = !freePlay
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {

		// Notice the difference between left/Right ATK
		// X values is the width/2 of the hitbox itself
		// 43 - 38 = 5
		atkRHbox := resolv.NewObject(pc.X+43, pc.Y+15, 10, 5, "hitBox")
		atkLHbox := resolv.NewObject(pc.X-38, pc.Y+15, 10, 5, "hitBox")

		playerHitBox := resolv.NewObject(pc.X, pc.Y, 18, 44, "hitBox")

		hitBoxes := make([]*resolv.Object, 0)

		hitBoxes = append(hitBoxes, playerHitBox)
		hitBoxes = append(hitBoxes, atkRHbox)
		hitBoxes = append(hitBoxes, atkLHbox)

		pc.World.Init(DevWorldBuilder)

		for _, b := range hitBoxes {
			pc.World.Space.Add(b)
		}

	}

	// Free Play Cam
	// Also an example of a "Cam Hack"
	if freePlay {
		cam := pc.PlayerCam

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

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		pc.inputHandler("primaryAttack")
		isPressing = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		pc.inputHandler("gravBoost")
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
	TODO: SetCameraPosition() needs to be cleaned up...

	Logic to keep PlayerController camera
	following player w/o exposing level boundaries
	offsets are also used by DrawPlayer()
	to render non PC players w/o player
	jitters(TODO: it's a long story prob worth documenting).

	Note: there is still cam jitters issue...its just that
	the players used to also jitter

	// Perhaps some predictive camera/client side camera movement should be done
*/

func (pc *PlayerController) SetCameraPosition() {
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

		pc.yOff = (ScreenHeight / 2) - (gh - pc.Y)
		pc.xOff = (ScreenWidth / 2) - pc.X

		ny := y - pc.yOff
		nx := (ScreenWidth / 2) - x

		pc.PlayerCam.SetPosition(nx, ny)

	} else if xBoundLeft && xBoundTop {

		pc.yOff = pc.Y - (ScreenHeight / 2)
		pc.xOff = (ScreenWidth / 2) - pc.X

		nx := (ScreenWidth / 2) - x
		ny := (ScreenHeight / 2) - y

		pc.PlayerCam.SetPosition(nx, ny)

	} else if xBoundRight && xBoundBottom {

		pc.xOff = (gw - pc.X) - (ScreenWidth / 2)
		pc.yOff = (ScreenHeight / 2) - (gh - pc.Y)

		nx := x - ((ScreenWidth / 2) - (gw - pc.X))
		ny := y - pc.yOff

		pc.PlayerCam.SetPosition(nx, ny)

	} else if xBoundRight && xBoundTop {

		pc.yOff = pc.Y - (ScreenHeight / 2)
		pc.xOff = (gw - pc.X) - (ScreenWidth / 2)

		nx := x - ((ScreenWidth / 2) - (gw - pc.X))
		ny := (ScreenHeight / 2) - y

		pc.PlayerCam.SetPosition(nx, ny)

	} else if xBoundLeft {
		pc.yOff = 0
		pc.xOff = (ScreenWidth / 2) - pc.X

		nx := (ScreenWidth / 2) - x

		pc.PlayerCam.SetPosition(nx, y)

	} else if xBoundRight {
		pc.yOff = 0
		pc.xOff = (gw - pc.X) - (ScreenWidth / 2)

		nx := x - ((ScreenWidth / 2) - (gw - pc.X))

		pc.PlayerCam.SetPosition(nx, y)

	} else if xBoundBottom {
		pc.yOff = (ScreenHeight / 2) - (gh - pc.Y)
		pc.xOff = 0

		ny := y - pc.yOff

		pc.PlayerCam.SetPosition(x, ny)

	} else if xBoundTop {
		pc.yOff = pc.Y - (ScreenHeight / 2)
		pc.xOff = 0

		ny := (ScreenHeight / 2) - y

		pc.PlayerCam.SetPosition(x, ny)

	} else {
		pc.yOff = 0
		pc.xOff = 0

		pc.PlayerCam.SetPosition(x, y)

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

func (pc *PlayerController) SubscribeToState() {
	world := pc.World
	wTex := &world.WorldTex
	ptex := &pc.PlayerTex

	if streamInit {
		return
	}

	go func() {
		streamInit = true
		for {
			ptex.Lock()
			res, err := pc.Stream.Recv()
			ptex.Unlock()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Setting World State Error %v\n", err)
				break
			}

			// reg lock on insertion
			wTex.Lock()
			world.State = res.Players
			wTex.Unlock()
		}
	}()
}
