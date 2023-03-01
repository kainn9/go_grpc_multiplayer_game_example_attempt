package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	pb "github.com/kainn9/grpc_game/proto"
	ut "github.com/kainn9/grpc_game/util"
	camera "github.com/melonfunction/ebiten-camera"
	"github.com/pborman/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type PlayerController struct {
	stream pb.PlayersService_PlayerLocationClient
	world  *World
	pid    string
	x      float64
	y      float64
	playerCam
}

type playerCam struct {
	*camera.Camera
	playerCamData
}

type playerCamData struct {
	playerCXpos float64
	playerCYpos float64
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

	cam := playerCam{
		Camera: camera.NewCamera(clientConfig.screenWidth, clientConfig.screenHeight, 0, 0, 0, 1),
		playerCamData: playerCamData{
			playerCXpos: 0,
			playerCYpos: 0,
		},
	}

	p := &PlayerController{
		pid:       pid,
		playerCam: cam,
	}

	p.stream = p.NewStream()

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

	conn, err := grpc.Dial(clientConfig.addr, opts...)

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	clientConfig.connRef = conn

	c := pb.NewPlayersServiceClient(conn)

	// TODO: Delete OR Keep?
	maxSizeOption := grpc.MaxCallRecvMsgSize(32 * 10e9)

	md := metadata.Pairs("pid", p.pid)
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
		clientConfig.volume128--
	}
	if ebiten.IsKeyPressed(ebiten.KeyX) {
		clientConfig.volume128++
	}
	if clientConfig.volume128 < 0 {
		clientConfig.volume128 = 0
	}
	if 128 < clientConfig.volume128 {
		clientConfig.volume128 = 128
	}

	if clientConfig.audPlayer != nil {
		clientConfig.audPlayer.SetVolume(float64(clientConfig.volume128) / 128)
	}

}

/*
Listens for Player inputs during game update phase
*/
func (pc *PlayerController) InputListener() {


	// attack Hbox Tester
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if hitBoxTest.on {
			return
		}
		hitBoxSim(pc.world.bg, pc)
	}
	

	updateVolumeIfNeeded()

	if inpututil.IsKeyJustPressed(ebiten.Key0) {
		clientConfig.fullScreen = !clientConfig.fullScreen
		ebiten.SetFullscreen(clientConfig.fullScreen)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key4) {
		pc.inputHandler("swap")
	}

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		devConfig.freePlay = !devConfig.freePlay
	}


	// Free Play Cam
	// Also an example of a "Cam Hack"
	if devConfig.freePlay {
		cam := pc.playerCam

		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			moveX := float64(cam.X + devConfig.devCamSpeed)
			moveY := float64(cam.Y)
			cam.SetPosition(moveX, moveY)
		}

		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			moveX := float64(cam.X - devConfig.devCamSpeed)
			moveY := float64(cam.Y)
			cam.SetPosition(moveX, moveY)
		}

		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			moveX := float64(cam.X)
			moveY := float64(cam.Y - devConfig.devCamSpeed)
			cam.SetPosition(moveX, moveY)
		}

		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			moveX := float64(cam.X)
			moveY := float64(cam.Y + devConfig.devCamSpeed)
			cam.SetPosition(moveX, moveY)
		}

		if inpututil.IsKeyJustPressed(ebiten.Key3) {
			pc.world.Init(DevWorldBuilder)
			devConfig.devPreview = !devConfig.devPreview

			if !devConfig.devPreview {
				pc.world.space.Remove(pc.world.space.Objects()...)
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.Key2) {
			devConfig.useHeightRuler = !devConfig.useHeightRuler
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyW) {
			devConfig.devCamSpeed += 1
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyS) {
			devConfig.devCamSpeed -= 1
		}

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
		pc.inputHandler("primaryAtk")
		isPressing = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		pc.inputHandler("secondaryAtk")
		isPressing = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyC) {
		pc.inputHandler("tertAtk")
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
		req := pb.PlayerReq{Id: p.pid, Input: input}
		devConfig.reqT = time.Now()
		p.stream.Send(&req)
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
	gw := pc.world.Width
	gh := pc.world.Height

	ScreenWidth := float64(clientConfig.screenWidth)
	ScreenHeight := float64(clientConfig.screenHeight)

	if devConfig.freePlay {
		return
	}

	// used to in player#draw
	// to avoid player jitters in millisecond/micro-pixel
	// diff in player pos vs where they are being rendered on cam
	pc.playerCXpos = pc.x
	pc.playerCYpos= pc.y

	x := (pc.x / 2)
	y := (pc.y / 2)

	// edges of level where we want to 
	// stop centering the player in the cam
	// to avoid showing empty space
	xBoundLeft := (pc.x - ScreenWidth/2) < 0
	xBoundBottom := (pc.y + (ScreenHeight / 2)) > gh
	xBoundRight := (pc.x + ScreenWidth/2) > gw
	xBoundTop := (pc.y - ScreenHeight/2) < 0


	if xBoundLeft && xBoundBottom {

		yOff := (ScreenHeight / 2) - (gh - pc.y)

		ny := y - yOff
		nx := (ScreenWidth / 2) - x

		nx = ut.CamLerp(pc.playerCam.X, nx)
		ny = ut.CamLerp(pc.playerCam.Y, ny)

		pc.playerCam.SetPosition(nx, ny)

	} else if xBoundLeft && xBoundTop {

		nx := (ScreenWidth / 2) - x
		ny := (ScreenHeight / 2) - y

		nx = ut.CamLerp(pc.playerCam.X, nx)
		ny = ut.CamLerp(pc.playerCam.Y, ny)

		pc.playerCam.SetPosition(nx, ny)

	} else if xBoundRight && xBoundBottom {

		yOff := (ScreenHeight / 2) - (gh - pc.y)

		nx := x - ((ScreenWidth / 2) - (gw - pc.x))
		ny := y - yOff

		nx = ut.CamLerp(pc.playerCam.X, nx)
		ny = ut.CamLerp(pc.playerCam.Y, ny)

		pc.playerCam.SetPosition(nx, ny)

	} else if xBoundRight && xBoundTop {

		nx := x - ((ScreenWidth / 2) - (gw - pc.x))
		ny := (ScreenHeight / 2) - y

		nx = ut.CamLerp(pc.playerCam.X, nx)
		ny = ut.CamLerp(pc.playerCam.Y, ny)

		pc.playerCam.SetPosition(nx, ny)

	} else if xBoundLeft {

		nx := (ScreenWidth / 2) - x

		nx = ut.CamLerp(pc.playerCam.X, nx)
		ny := ut.CamLerp(pc.playerCam.Y, y)

		pc.playerCam.SetPosition(nx, ny)

	} else if xBoundRight {

		nx := x - ((ScreenWidth / 2) - (gw - pc.x))

		nx = ut.CamLerp(pc.playerCam.X, nx)
		ny := ut.CamLerp(pc.playerCam.Y, y)

		pc.playerCam.SetPosition(nx, ny)

	} else if xBoundBottom {

		yOff := (ScreenHeight / 2) - (gh - pc.y)

		ny := y - yOff

		nx := ut.CamLerp(pc.playerCam.X, x)
		ny = ut.CamLerp(pc.playerCam.Y, ny)

		pc.playerCam.SetPosition(nx, ny)

	} else if xBoundTop {

		ny := (ScreenHeight / 2) - y

		nx := ut.CamLerp(pc.playerCam.X, x)
		ny = ut.CamLerp(pc.playerCam.Y, ny)

		pc.playerCam.SetPosition(nx, ny)

	} else {
		nx := ut.CamLerp(pc.playerCam.X, x)
		ny := ut.CamLerp(pc.playerCam.Y, y)

		pc.playerCam.SetPosition(nx, ny)

	}
}

/*
Helper function for handling current player
state from server stream
*/
func CurrentPlayerHandler(pc *PlayerController, ps *pb.Player, p *Player) {
	cw := pc.world

	cw.playerController.x = ps.Lx
	cw.playerController.y = ps.Ly


	if clientConfig.game.CurrentWorld != ps.World {

		newData := clientConfig.worldsMap[ps.World]

		UpdateWorldData(cw, &newData, ps.World)
	}

	DrawPlayer(cw, p, true)
}

func (pc *PlayerController) SubscribeToState() {
	world := pc.world
	wTex := &world.worldTex
	if clientConfig.streamInit {
		return
	}

	go func() {
		clientConfig.streamInit = true

		for {
			res, err := pc.stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Setting World State Error %v\n", err)
				break
			}
			devConfig.ping = float64(time.Since(devConfig.reqT))
			
			// reg lock on insertion?
			wTex.Lock()
			world.state = res.Players
			wTex.Unlock()

		}
	}()
}

func (pc *PlayerController) health() int {
	p := pc.world.playerMap[pc.pid]
	if p != nil {
		return p.health
	}
	
	// TODO: idk what to do here yet
	return 0
}


