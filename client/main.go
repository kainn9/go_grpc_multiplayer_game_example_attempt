package main

import (
	"bytes"
	"fmt"
	_ "image/png"
	"io/ioutil"
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
		volume128 = 128
		sampleRate := 32000
		songBytes, err := ut.LoadMusic("./audio/base.mp3")
		if err != nil {
			log.Fatalf("Error Loading Song: %v\n", err)
		}

		s, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(songBytes))
		if err != nil {
			log.Fatalf("Error decoding Song Bytes: %v\n", err)
		}
		b, _ := ioutil.ReadAll(s)
		audCtx := audio.NewContext(sampleRate)
		audPlayer = audCtx.NewPlayerFromBytes(b)
		audPlayer.Play()
	}()

}

func (g *Game) Layout(outsideScreenWidth, outsideScreenHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

/*
Game counter, capping at 60
to match ebiten TPS
*/
func (g *Game) IncrementTicks() {
	
	ticks++
	
	if ticks > 60 {
		ticks = 0
	}


	for k, a := range fixedAnims {
		
		p := game.World.playerMap[a.pid]
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
	pc := g.World.PlayerController

	// Clear the camera before drawing
	pc.PlayerCam.Surface.Clear()

	g.World.Draw(screen)

	/*
		TODO: clean this up/make a seperate dev client
		Dev-Tool/FreePlay stuff:
	*/
	if freePlay {
		opts := &ebiten.DrawImageOptions{}
		opts = pc.PlayerCam.GetTranslation(opts, pc.PlayerCam.X, pc.PlayerCam.Y)

		if useHeightRuler {
			pc.PlayerCam.Surface.DrawImage(rulerH, opts)
		} else {
			pc.PlayerCam.Surface.DrawImage(rulerW, opts)
		}

	}

	// Blit!
	pc.PlayerCam.Blit(screen)

	/*
		TODO: clean this up/make a seperate dev client
		render debug stats after blit
		to be on highest z-index
	*/
	msg := fmt.Sprintf(
		"PING(Calc is a bit busted rn): %v\nArrow Keys to move, space to jump(you can wall jump too)\nPress F to poke\nPress G for Dash Attack(WIP, takes one sec to fire and I have no windup anim yet/Attack part)\nPress T for 20 sec grav boost(2 min CD)\nPress 0 to toggle full-screen\nPress Z/X to controll volume\nCurr volume: %v\nPress 1 to toggle freeplay/devMode\nPress 3 to turn on dev preview\nPress 4 to swap worlds\nTPS: %0.2f\nhealth: %v\n",
		ping,
		volume128,
		ebiten.ActualTPS(),
		pc.health(),
	)

	if freePlay {

		msg += fmt.Sprintln("FREE PLAY ON!!!\nPress 2 to toggle rulers\nUse w/s to decrease/increase cam speed")
		msg += fmt.Sprintf("Cam Speed:%v\n", devCamSpeed)

		/*
			This calc is scuffed.
			It only works if you dont move and
			needs 2 be redone if the spawn cords change...
			need to come up with the right maffz
			to make it consistent
		*/
		msg += fmt.Sprintf("X:%v\nY:%v\n", (pc.PlayerCam.X + (ScreenWidth / 2) + 185), pc.PlayerCam.Y+(ScreenHeight/2)+1898-172)

	} else {
		msg += fmt.Sprintf("X:%v\nY:%v\n", pc.X, pc.Y)
	}
	ebitenutil.DebugPrint(screen, msg)
}

/*
Creates new game.
*/
func NewGame() *Game {

	worldsMap["main"] = *NewWorldData(848, 3200, mainWorldBg)
	worldsMap["alt"] = *NewWorldData(4000, 6000, altWorldBg)

	// Set window things.
	ebiten.SetWindowTitle("MultiPlayer Platformer!")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	// set/init world
	w := NewWorld("main")

	// attach player controller to world
	w.PlayerController = NewPlayerController()

	// attach playerController to world
	w.PlayerController.World = w

	return &Game{
		ShowHelpText: true,
		World:        w,
		CurrentWorld: "main",
	}
}

func main() {
	// PPROF HANDLER
	// Add a handler for the pprof endpoint at
	// http: //localhost:6060/debug/pprof/
	// enablePPROF toggled in global
	if enablePPROF {
		go func() {
			http.ListenAndServe("localhost:6060", nil)
		}()
	}

	ebiten.SetFullscreen(fullScreen)
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
	game.InitMusic()
	ebiten.RunGame(game)

	// TODO:
	// does this work?
	// low key was never closing client connection on close...
	defer connRef.Close()
}
