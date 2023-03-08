package main

import (
	"bytes"
	"fmt"
	_ "image/png"
	"io"
	"log"

	"net/http"
	_ "net/http/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	ut "github.com/kainn9/grpc_game/util"
	"golang.org/x/image/font"
)

type Game struct {
	Debug        bool
	ShowHelpText bool
	Screen       *ebiten.Image
	FontFace     font.Face
	World        *World
	CurrentWorld string
}

func (g *Game) InitMusic() {
	// TODO: Create Audio System

	go func() {
		clientConfig.volume128 = 128
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
		"PING(Calc is a bit busted rn): %v\nArrow Keys to move, space to jump(you can wall jump too)\nPress F to poke\nPress G for Dash Attack(WIP, takes one sec to fire and I have no windup anim yet/Attack part)\nPress T for 20 sec grav boost(2 min CD)\nPress 0 to toggle full-screen\nPress Z/X to controll volume\nCurr volume: %v\nPress 1 to toggle freeplay/devMode\nPress 3 to turn on dev preview\nPress 4 to swap worlds\nTPS: %0.2f\nhealth: %v\n",
		devConfig.ping,
		clientConfig.volume128,
		ebiten.ActualTPS(),
		pc.health(),
	)

	if devConfig.freePlay {

		msg += fmt.Sprintln("FREE PLAY ON!!!\nPress 2 to toggle rulers\nUse w/s to decrease/increase cam speed")
		msg += fmt.Sprintf("Cam Speed:%v\n", devConfig.devCamSpeed)

		/*
			This calc is scuffed.
			It only works if you dont move and
			needs 2 be redone if the spawn cords change...
			need to come up with the right maffz
			to make it consistent
		*/
		msg += fmt.Sprintf("X:%v\nY:%v\n", (pc.playerCam.X + (float64(clientConfig.screenWidth) / 2) + 185), pc.playerCam.Y+(float64(clientConfig.screenHeight)/2)+1898-172)

	} else {
		msg += fmt.Sprintf("X:%v\nY:%v\n", pc.x, pc.y)
	}
	ebitenutil.DebugPrint(screen, msg)
}

/*
Creates new game.
*/
func NewGame() *Game {

	clientConfig.worldsMap["main"] = *NewWorldData(848, 3200, assetsHelper.mainWorldBg)
	clientConfig.worldsMap["alt"] = *NewWorldData(4000, 6000, assetsHelper.altWorldBg)

	// Set window things.
	ebiten.SetWindowTitle("MultiPlayer Platformer!")
	ebiten.SetWindowSize(clientConfig.screenWidth, clientConfig.screenHeight)

	// set/init world
	w := NewWorld("main")

	// attach player controller to world
	w.playerController = NewPlayerController()

	// attach playerController to world
	w.playerController.world = w

	return &Game{
		ShowHelpText: true,
		World:        w,
		CurrentWorld: "main",
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

	ebiten.SetFullscreen(clientConfig.fullScreen)
	/*
		RunGame starts the main loop and runs the game. game's
		Update function is called every tick to update the game logic.
		game's Draw function is called every frame to draw the screen.
		game's Layout function is called when necessary, and you can specify
		the logical screen size by the function.
		game's functions are called on the same goroutine.
	*/

	// Need to make Async, as attributes to 90% of startup time rn....

	// so disabling music for now LOL
	clientConfig.game.InitMusic()
	ebiten.RunGame(clientConfig.game)

	// TODO:
	// does this work?
	// low key was never closing client connection on close...
	defer clientConfig.connRef.Close()
}
