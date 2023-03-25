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
	space            *resolv.Space
	playerController *PlayerController
	state            []*pb.Player
	playerMap        map[string]*Player
	worldTex         sync.RWMutex
	worldData
}

type worldData struct {
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
		worldTex: sync.RWMutex{},
	}
	w.worldData = GetWorldData(key)

	return w
}

/*
Returns world data using worldsMap + world key
*/
func GetWorldData(worldKey string) worldData {
	return clientConfig.worldsMap[worldKey]
}

/*
Used to create "world data" that can be embedded in world struct
*/
func NewWorldData(height float64, width float64, bg *ebiten.Image) *worldData {
	wd := &worldData{
		Height: height,
		Width:  width,
		bg:     bg,
	}
	return wd
}

/*
TODO: clean this up/make a seperate dev client

Really only used for dev testing...
uses the builder function to place
resolv geometry that will get picked up
by world.draw() for previewing where to place
the resolv objects serverside
*/
func (world *World) Init(worldBuilder BuilderFunc) {

	if world.space != nil {
		log.Println("World Already Init'd...Skipping")
	}

	log.Println("World Init")
	gw := float64(world.Width)
	gh := float64(world.Height)

	// use this + freePlay to help build maps
	// with resolv until a real system is in place
	world.space = resolv.NewSpace(int(gw), int(gh), 8, 8)
	worldBuilder(world, gw, gh)
}

/*
Invokes world's update based receiver functions
*/
func Update(world *World) {
	cp := world.playerController

	cp.SubscribeToState()
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
	pc := w.playerController

	x := float64(pc.x)
	y := float64(pc.y)

	bgOpts := &ebiten.DrawImageOptions{}
	bgOpts = pc.playerCam.GetTranslation(bgOpts, -x/2, -y/2)
	pc.playerCam.Surface.DrawImage(w.bg, bgOpts)

	if devConfig.devPreview {

		for _, o := range w.space.Objects() {
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
Renders players from server response
*/
func (world *World) DrawPlayers() {

	newPlayerMap := make(map[string]*Player)

	wTex := &world.worldTex
	/*
		write only lock when rendering state
		from map, as it mutated in the go routine
		above and not thread safe otherwise
	*/
	wTex.RLock()
	for k := range world.state {

		ps := world.state[k]

		p := NewPlayer()

		p.speedX = ps.SpeedX
		p.speedY = ps.SpeedY
		p.x = ps.Lx
		p.y = ps.Ly
		p.facingRight = ps.FacingRight
		p.jumping = ps.Jumping
		p.currAttack = ps.CurrAttack
		p.cc = ps.CC
		p.windup = ps.Windup
		p.attackMovement = ps.AttackMovement
		p.id = ps.Id
		p.health = int(ps.Health)
		p.defending = ps.Defending
		p.Role = *clientConfig.roles[ps.Role]
		p.dead = ps.Dead

		newPlayerMap[ps.Id] = p

		if ps.Id == world.playerController.pid {
			CurrentPlayerHandler(world.playerController, ps, p)
		} else {
			DrawPlayer(world, p, false)
		}

	}
	wTex.RUnlock()

	world.playerMap = newPlayerMap
}

/*
Helper to change world data when
client receives the information that
the currentPlayer is a different world/level
then the current Game.CurrentWorld(string/key)
*/
func UpdateWorldData(w *World, new *worldData, key string) {
	w.worldData = clientConfig.worldsMap[key]
	clientConfig.game.CurrentWorld = key
}
