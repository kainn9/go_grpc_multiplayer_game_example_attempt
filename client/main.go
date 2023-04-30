package main

import (
	"bytes"
	"fmt"
	_ "image/png"
	"io"
	"log"
	"math"

	"net/http"
	_ "net/http/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	ut "github.com/kainn9/grpc_game/client_util"
	"golang.org/x/image/font"
)

type Game struct {
	Debug        bool
	ShowHelpText bool
	Screen       *ebiten.Image
	FontFace     font.Face
	World        *World
	CurrentWorld int
}

func (g *Game) InitMusic() {
	// TODO: Create Audio System

	go func() {
		clientConfig.volume128 = 32
		sampleRate := 32000
		songBytes, err := ut.LoadMusic("./audio/base.mp3")
		if err != nil {
			log.Fatalf("Error Loading Song: %v\n", err)
		}

		s, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(songBytes))
		if err != nil {
			log.Fatalf("Error decoding Song Bytes: %v\n", err)
		}
		b, _ := io.ReadAll(s)
		audCtx := audio.NewContext(sampleRate)
		clientConfig.audPlayer = audCtx.NewPlayerFromBytes(b)
		clientConfig.audPlayer.Play()
	}()

}

func (g *Game) Layout(outsideScreenWidth, outsideScreenHeight int) (int, int) {
	return clientConfig.screenWidth, clientConfig.screenHeight
}

/*
Game counter, capping at 60
to match ebiten TPS
*/
func (g *Game) IncrementTicks() {

	clientConfig.ticks++

	if clientConfig.ticks > 60 {
		clientConfig.ticks = 0
	}

	for k, a := range fixedAnims {

		p := clientConfig.game.World.playerMap[a.pid]
		if p == nil {

			continue
		}

		ca := p.currentAnimation
		if ca == nil {
			continue
		}

		if p.currentAnimation.Name != a.animName {

			delete(fixedAnims, k)
		} else {
			a.ticks++
		}
	}
}

func (g *Game) Update() error {
	g.IncrementTicks()

	Update(g.World)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	pc := g.World.playerController

	// Clear the camera before drawing
	pc.playerCam.Surface.Clear()

	g.World.Draw(screen)

	/*
		TODO: clean this up/make a seperate dev client
		Dev-Tool/FreePlay stuff:
	*/
	if devConfig.freePlay {
		opts := &ebiten.DrawImageOptions{}
		opts = pc.playerCam.GetTranslation(opts, pc.playerCam.X, pc.playerCam.Y)

		if devConfig.useHeightRuler {
			pc.playerCam.Surface.DrawImage(devConfig.rulerH, opts)
		} else {
			pc.playerCam.Surface.DrawImage(devConfig.rulerW, opts)
		}

	}

	// Blit!
	pc.playerCam.Blit(screen)

	/*
		TODO: clean this up/make a seperate dev client
		render debug stats after blit
		to be on highest z-index
	*/
	msg := fmt.Sprintf(
		"Press M to hide this menu!!!!\n\nPING(Calc is a bit busted rn): %v\nCurr volume: %v\nTPS: %0.2f\nhealth: %v\n\nBasic Player Controls: Arrow Keys for Movment(you can use down arrow to drop from platform)\nSpace to jump(you can wall jump too)\n\nClient Controls:\n0-key to toggle full-screen\nZ/X-keys to control volume\n\nAdmin/Dev Controls:\n1-key to toggle dev-camera\n3-key to turn on dev world builder preview\nL-key for hitbox mode(Note: this will clear background until client restart)\n4-key to swap worlds\n",
		devConfig.ping,
		clientConfig.volume128,
		ebiten.ActualTPS(),
		pc.health(),
	)

	if devConfig.freePlay {

		msg += fmt.Sprintln("FREE PLAY ON!!!\nPress 2 to toggle rulers\nUse w/s to decrease/increase cam speed")
		msg += fmt.Sprintf("Cam Speed:%v\n", devConfig.devCamSpeed)

		msg += fmt.Sprintf("X:%v\nY:%v\n", math.Round(pc.playerCam.X+(pc.x/2)), math.Round(pc.playerCam.Y+(pc.y/2)))

	} else {
		msg += fmt.Sprintf("X:%v\nY:%v\n", pc.x, pc.y)
	}
	if clientConfig.showHelp {
		ebitenutil.DebugPrint(screen, msg)
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Health: %v\n", pc.health()))
	}

}

/*
Creates new game.
*/
func NewGame() *Game {

	// Set window things.
	ebiten.SetWindowTitle("MultiPlayer Platformer!")
	ebiten.SetWindowSize(clientConfig.screenWidth, clientConfig.screenHeight)

	// set/init world
	w := NewWorld(0)

	// attach player controller to world
	w.playerController = NewPlayerController()

	// attach playerController to world
	w.playerController.world = w

	return &Game{
		ShowHelpText: true,
		World:        w,
		CurrentWorld: 0,
	}
}

func main() {
	initClient()

	// PPROF HANDLER
	// Add a handler for the pprof endpoint at
	// http: //localhost:6060/debug/pprof/
	// enablePPROF toggled in global
	if clientConfig.enablePPROF {
		go func() {
			http.ListenAndServe("localhost:6060", nil)
		}()
	}

	ebiten.SetWindowResizable(true)

	/*
		RunGame starts the main loop and runs the game. game's
		Update function is called every tick to update the game logic.
		game's Draw function is called every frame to draw the screen.
		game's Layout function is called when necessary, and you can specify
		the logical screen size by the function.
		game's functions are called on the same goroutine.
	*/

	clientConfig.game.InitMusic()

	ebiten.RunGame(clientConfig.game)

	defer clientConfig.connRef.Close()
}

func toggleFS() {
	if clientConfig.fullScreen {
		if clientConfig.defaultWindowPosX == 0 || clientConfig.defaultWindowPosY == 0 {
			clientConfig.setWindowDefaults()
		}

		screenWidth, screenHeight := ebiten.ScreenSizeInFullscreen()
		ebiten.SetWindowSize(screenWidth, screenHeight)
		ebiten.SetWindowPosition(0, 0)
	} else {
		ebiten.SetWindowSize(clientConfig.screenWidth, clientConfig.screenHeight)
		ebiten.SetWindowPosition(clientConfig.defaultWindowPosX, clientConfig.defaultWindowPosY)
	}
}
