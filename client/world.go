package main

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	pb "github.com/kainn9/grpc_game/proto"
	"github.com/solarlune/resolv"
)

var (
	mainWorldBg = LoadImg("./backgrounds/mapMain.png")
	altWorldBg  = LoadImg("./backgrounds/mapAlt.png")
)

type World struct {
	Game             *Game
	Space            *resolv.Space
	PlayerController *PlayerController
	State            []*pb.Player
	WorldTex         sync.RWMutex
	WorldData
}

type WorldData struct {
	Height float64
	Width  float64
	bg     *ebiten.Image
}

/*
Builder func should call world.Space.Add()
to create geometry. When done clientSide,
its really just to preview where to place
the resolv objects /w freeplay and it won't
trigger collisions or anything but you can
than copy paste the builder function to
the server and setup a world to use it for
physics
*/
type BuilderFunc func(*World, float64, float64)

/*
Creates a New World
*/
func NewWorld(key string) *World {
	w := &World{
		WorldTex: sync.RWMutex{},
	}
	w.WorldData = GetWorldData(key)
	return w
}

func GetWorldData(worldKey string) WorldData {
	return worldsMap[worldKey]
}

func NewWorldData(height float64, width float64, bg *ebiten.Image) *WorldData {
	wd := &WorldData{
		Height: height,
		Width:  width,
		bg:     bg,
	}
	return wd
}

/*
Really only used for dev testing...
uses the builder function to place
resolv geometry that will get picked up
by world.draw() for previewing where to place
the resolv objects serverside
*/
func (world *World) Init(worldBuilder BuilderFunc) {
	log.Println("World Init")

	gw := float64(world.Width)
	gh := float64(world.Height)

	// use this + freePlay to help build maps
	// with resolv until a real system is in place
	world.Space = resolv.NewSpace(int(gw), int(gh), 8, 8)
	worldBuilder(world, gw, gh)
}

/*
Invokes world's update based receiver functions
*/
func Update(world *World) {
	world.PlayerController.inputListener()
	world.PlayerController.setCameraPosition()
	world.PlayerController.GetState()
}

/*
Invokes world's draw based receiver functions
*/
func (w *World) Draw(screen *ebiten.Image) {
	w.DrawBg()
	w.DrawPlayers()
}

/*
Renders the current world BG(based on worldData struct)
*/
func (w *World) DrawBg() {
	pc := w.PlayerController

	x := float64(pc.X)
	y := float64(pc.Y)

	bgOpts := &ebiten.DrawImageOptions{}
	bgOpts = pc.Cam.GetTranslation(bgOpts, -x/2, -y/2)
	pc.Cam.Surface.DrawImage(w.bg, bgOpts)

	if devPreview {

		for _, o := range w.Space.Objects() {
			if o.HasTags("platform") {
				drawColor := color.RGBA{180, 100, 0, 255}
				ebitenutil.DrawRect(w.bg, o.X, o.Y, o.W, o.H, drawColor)
			} else {
				drawColor := color.RGBA{60, 60, 60, 255}
				ebitenutil.DrawRect(w.bg, o.X, o.Y, o.W, o.H, drawColor)
			}
		}

	}
}

/*
Renders the players from server response
*/
func (world *World) DrawPlayers() {

	wTex := &world.WorldTex
	/*
		write only lock when rendering state
		from map, as it mutated in the go routine
		above and not thread safe otherwise
	*/
	wTex.RLock()
	for k := range world.State {
		ps := world.State[k]
		p := NewPlayer()
		p.SpeedX = ps.SpeedX
		p.SpeedY = ps.SpeedY
		p.X = ps.Lx
		p.Y = ps.Ly
		p.FacingRight = ps.FacingRight
		p.Jumping = ps.Jumping

		if ps.Id == world.PlayerController.Pid {
			CurrentPlayerHandler(world.PlayerController, ps, p)
		} else {
			DrawPlayer(world, p, false)
		}

	}
	wTex.RUnlock()
}

/*
Helper to change world data when
client receives the information that
the currentPlayer is a different world/level
then the current Game.CurrentWorld(string/key)
*/
func UpdateWorldData(w *World, new *WorldData, key string) {
	w.WorldData = worldsMap[key]
	game.CurrentWorld = key
}

/*
Fill in your own geometry here, toggle dev
mode to true, and use free play to figure out
where to place resolv objects on the serverside
(At least until I actually learn a system for level design haha)
*/
func DevWorldBuilder(world *World, gw float64, gh float64) {
	world.Space.Add(
		resolv.NewObject(0, 660, 800, 10, "solid"),
	)
}
