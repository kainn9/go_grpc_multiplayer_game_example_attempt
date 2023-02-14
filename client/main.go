package main

import (
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	ut "github.com/kainn9/grpc_game/client/util"
	"golang.org/x/image/font"
)

const (
	ScreenWidth  = 880
	ScreenHeight = 480
)

var (
	ticks     int
	worldsMap = make(map[string]WorldData)
	game      = NewGame()
	addr      = "localhost:50051"
)

// dev mode stuff
var (
	rulerW         = ut.LoadImg("./sprites/rulers/wRuler.png")
	rulerH         = ut.LoadImg("./sprites/rulers/hRuler.png")
	devPreview     = false
	useHeightRuler = false
	devCamSpeed    = float64(2)
)

type Game struct {
	Debug        bool
	ShowHelpText bool
	Screen       *ebiten.Image
	FontFace     font.Face
	World        *World
	CurrentWorld string
}

func (g *Game) Layout(outsideScreenWidth, outsideScreenHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) IncrementTicks() {
	if ticks > 60 {
		ticks = 1
	} else {
		ticks++
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
	pc.Cam.Surface.Clear()

	g.World.Draw(screen)

	if freePlay {
		opts := &ebiten.DrawImageOptions{}
		opts = pc.Cam.GetTranslation(opts, pc.Cam.X, pc.Cam.Y)

		if useHeightRuler {
			pc.Cam.Surface.DrawImage(rulerH, opts)
		} else {
			pc.Cam.Surface.DrawImage(rulerW, opts)
		}

	}

	// Blit!
	pc.Cam.Blit(screen)

	// render debug stats after blit
	// to be on highest z-index
	msg := fmt.Sprintf("Arrow Keys to move, space to jump\nPress 1 to toggle freeplay/devMode\nPress 3 to turn on dev preview\nPress 4 to swap worlds\nTPS: %0.2f\n", ebiten.ActualTPS())

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
		msg += fmt.Sprintf("X:%v\nY:%v\n", (pc.Cam.X + (ScreenWidth / 2) + 185), pc.Cam.Y+(ScreenHeight/2)+1898-172)

	} else {
		msg += fmt.Sprintf("X:%v\nY:%v\n", pc.X, pc.Y)
	}
	ebitenutil.DebugPrint(screen, msg)
}

/*
Creates new game.
*/
func NewGame() *Game {


	worldsMap["main"] = *NewWorldData(4000, 6000, mainWorldBg)
	worldsMap["alt"] = *NewWorldData(848, 3200, altWorldBg)

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

	// TODO: MAKE THIS TOGGLE
	ebiten.SetFullscreen(false)
	/*
		RunGame starts the main loop and runs the game. game's
		Update function is called every tick to update the game logic.
		game's Draw function is called every frame to draw the screen.
		game's Layout function is called when necessary, and you can specify
		the logical screen size by the function.
		game's functions are called on the same goroutine.
	*/
	ebiten.RunGame(game)
}
