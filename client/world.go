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

	if world.Space != nil {
		log.Println("World Already Init'd...Skipping")
	}

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
	cp := world.PlayerController
	cp.GetState()
	cp.InputListener()
	cp.SetCameraPosition()

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
	bgOpts = pc.PlayerCam.GetTranslation(bgOpts, -x/2, -y/2)
	pc.PlayerCam.Surface.DrawImage(w.bg, bgOpts)

	if devPreview {

		for _, o := range w.Space.Objects() {
			if o.HasTags("platform") {
				drawColor := color.RGBA{180, 100, 0, 255}
				ebitenutil.DrawRect(w.bg, o.X, o.Y, o.W, o.H, drawColor)
			} else if o.HasTags("hitBox") {
				drawColor := color.RGBA{180, 34, 50, 120}
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
		p.CurrAttack = ps.CurrAttack
		p.CC = ps.CC

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

		// bottom bounds
		resolv.NewObject(0, gh - 16, gw, 16, "solid"),
		
		// Village Plat
		resolv.NewObject(1166, 3912, 6000, 10, "platform"),

		// Y-ZONE Plats
		resolv.NewObject(1008, 3931, 110, 10, "platform"),
		resolv.NewObject(909, 3973, 110, 10, "platform"),
		resolv.NewObject(902, 3818, 120, 10, "platform"),
		resolv.NewObject(996, 3687, 160, 10, "platform"),
		resolv.NewObject(673, 3717, 320, 10, "platform"),
		resolv.NewObject(512, 3844, 300, 10, "platform"),
		resolv.NewObject(676, 3632, 130, 10, "platform"),
		resolv.NewObject(906, 3588, 130, 10, "platform"),
		resolv.NewObject(1006, 3544, 130, 10, "platform"),
		resolv.NewObject(896, 3438, 130, 10, "platform"),
		resolv.NewObject(994, 3304, 160, 10, "platform"),
		resolv.NewObject(906, 3212, 130, 10, "platform"),
		resolv.NewObject(1002, 3170, 130, 10, "platform"),
		resolv.NewObject(473, 3812, 135, 10, "platform"),
		resolv.NewObject(477, 3577, 120, 10, "platform"),
		resolv.NewObject(477, 3577, 120, 10, "platform"),
		resolv.NewObject(513, 3465, 320, 10, "platform"),
		resolv.NewObject(477, 3432, 130, 10, "platform"),
		resolv.NewObject(672, 3332, 320, 10, "platform"),
		resolv.NewObject(674, 3252, 130, 10, "platform"),
		resolv.NewObject(900, 3060, 130, 10, "platform"),
		resolv.NewObject(994, 2928, 160, 10, "platform"),
		resolv.NewObject(669, 2955, 320, 10, "platform"),
		resolv.NewObject(673, 2874, 130, 10, "platform"),
		resolv.NewObject(509, 3086, 310, 10, "platform"),
		resolv.NewObject(509, 3086, 310, 10, "platform"),
		resolv.NewObject(476, 3195, 110, 10, "platform"),
		resolv.NewObject(477, 3958, 120, 10, "platform"),
		resolv.NewObject(267, 3973, 120, 10, "platform"),
		resolv.NewObject(267, 3973, 120, 10, "platform"),
		resolv.NewObject(366, 3928, 120, 10, "platform"),
		resolv.NewObject(366, 3928, 120, 10, "platform"),
		resolv.NewObject(257, 3824, 160, 10, "platform"),
		resolv.NewObject(355, 3687, 160, 10, "platform"),
		resolv.NewObject(267, 3590, 120, 10, "platform"),
		resolv.NewObject(364, 3546, 120, 10, "platform"),
		resolv.NewObject(259, 3439, 130, 10, "platform"),
		resolv.NewObject(34, 3718, 320, 10, "platform"),
		resolv.NewObject(0, 3848, 170, 10, "platform"),
		resolv.NewObject(34, 3634, 120, 10, "platform"),
		resolv.NewObject(0, 3464, 170, 10, "platform"),
		resolv.NewObject(31, 3335, 320, 10, "platform"),
		resolv.NewObject(355, 3305, 160, 10, "platform"),
		resolv.NewObject(38, 3258, 120, 10, "platform"),
		resolv.NewObject(266, 3204, 120, 10, "platform"),
		resolv.NewObject(368, 3164, 110, 10, "platform"),
		resolv.NewObject(255, 3050, 160, 10, "platform"),
		resolv.NewObject(474, 3048, 140, 10, "platform"),
		resolv.NewObject(355, 2922, 160, 10, "platform"),
		resolv.NewObject(0, 3078, 160, 10, "platform"),
		resolv.NewObject(30, 2949, 320, 10, "platform"),

		// mid section blocker left
		resolv.NewObject(206, 2584, 2030, 150, "solid"),
		resolv.NewObject(206, 2574, 2030, 10, "platform"),


		// left blocker left
		resolv.NewObject(0, 2108, 60, 540, "solid"),
		resolv.NewObject(0, 2098, 60, 10, "platform"),


		// forrest floating plats
		resolv.NewObject(64, 2639, 60, 10, "platform"),
		resolv.NewObject(128, 2549, 150, 10, "platform"),

		resolv.NewObject(305, 2500, 125, 10, "platform"),
		resolv.NewObject(452, 2450, 125, 10, "platform"),
		resolv.NewObject(615, 2392, 125, 10, "platform"),
		resolv.NewObject(797, 2359, 130, 10, "platform"),
		resolv.NewObject(797, 2359, 130, 10, "platform"),
		resolv.NewObject(956, 2316, 130, 10, "platform"),
		resolv.NewObject(1127, 2265, 130, 10, "platform"),
		resolv.NewObject(1308, 2241, 85, 10, "platform"),

		// wood forrest plat left
		resolv.NewObject(694, 2529, 1370, 10, "platform"),

		// castle floating plats
		resolv.NewObject(2093, 2484, 70, 10, "platform"),
		resolv.NewObject(2196, 2466, 30, 10, "platform"),
		resolv.NewObject(2400, 2450, 63, 10, "platform"),
		resolv.NewObject(2516, 2448, 63, 10, "platform"),
		resolv.NewObject(2611, 2428, 63, 10, "platform"),
		resolv.NewObject(2293, 2453, 63, 10, "platform"),

		// sky-town wallStalk and floaters
		resolv.NewObject(1278, 811, 54, 10, "platform"),
		resolv.NewObject(1278, 821, 54, 1275, "solid"),

		resolv.NewObject(1428, 2184, 54, 45, "solid"),
		resolv.NewObject(1428, 2174, 54, 10, "platform"),

		resolv.NewObject(1346, 2110, 54, 45, "solid"),
		resolv.NewObject(1346, 2100, 54, 10, "platform"),


		// sky-town floor left
		resolv.NewObject(0, 872, 1192, 10, "platform"),
		
		// sky-town floor right
		resolv.NewObject(1371, 837, 650, 10, "platform"),


		// dungeon town wall right divider
		resolv.NewObject(1970, 0, 55, 1826, "solid"),

		// rock plats
		resolv.NewObject(0, 2848, 80, 10, "platform"),
		resolv.NewObject(143, 2760, 160, 10, "platform"),
		resolv.NewObject(152, 2668, 50, 10, "platform"),
	)

}
