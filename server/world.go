package main

import (
	"time"

	pb "github.com/kainn9/grpc_game/proto"
	"github.com/solarlune/resolv"
	"github.com/tanema/gween"
)

type World struct {
	Space                 *resolv.Space
	FloatingPlatform      *resolv.Object
	FloatingPlatformTween *gween.Sequence
	State                 *pb.PlayerResp
	Height                float64
	Width                 float64
	Stream                pb.PlayersService_PlayerLocationServer
	Players               map[string]*Player
	Name                  string
}

/*
Creates a New World and Initializes it
*/
func NewWorld(height float64, width float64, worldBuilder BuilderFunc, name string) *World {
	w := &World{
		Width:   width,
		Height:  height,
		Players: make(map[string]*Player),
		Name:    name,
	}

	w.Init(worldBuilder)
	return w
}

/*
Initializes world physics using Resolv Lib,
*World's attrs, and a builder func
*/
func (world *World) Init(worldBuilder BuilderFunc) {
	gw := world.Width
	gh := world.Height

	// Define the world's Space. Here, a Space is essentially a grid (the game's width and height, or 640x360), made up of 16x16 cells. Each cell can have 0 or more Objects within it,
	// and collisions can be found by checking the Space to see if the Cells at specific positions contain (or would contain) Objects. This is a broad, simplified approach to collision
	// detection.
	world.Space = resolv.NewSpace(int(gw), int(gh), 8, 8)

	// Construct the solid level geometry. Note that the simple approach of checking cells in a Space for collision works simply when the geometry is aligned with the cells
	worldBuilder(world, gw, gh)
}

/*
Currently where all game logic happens
physics is basically a rip of the resolv example

https://github.com/SolarLune/resolv/blob/master/examples/worldPlatformer.go
*/
func (world *World) Update(cp *Player, input string) {

	if !cp.Object.HasTags("player") {
		cp.Object.AddTags("player")
	}

	// grav boost/buff
	// TODO: Move make an actual buff system/tracker
	if input == "gravBoost" && !cp.GravBoost {
		cp.jumpSpd = 15
		cp.GravBoost = true
		time.AfterFunc(20*time.Second, func() { cp.jumpSpd = defaultJumpSpd })
		time.AfterFunc(120*time.Second, func() { cp.GravBoost = false })
	}

	cp.WorldTransferHandler(input)

	cp.Gravity()
	cp.JumpHandler(input)
	cp.AttackHandler(input, world)
	cp.AttackedHandler()

	cp.HorizontalMovementHandler(input, world.Width)
	cp.VerticalMovmentHandler(input, world)

	// IDK where to put this yet...
	wallNext := 1.0
	if !cp.FacingRight {
		wallNext = -1
	}

	// If the wall next to the Player runs out, stop wall sliding.
	if c := cp.Object.Check(wallNext, 0, "solid"); cp.WallSliding != nil && c == nil {
		cp.WallSliding = nil
	}

	cp.Object.Update() // Update the player's position in the space.

}

func (w *World) removeAtk(a *resolv.Object) {
	w.Space.Remove(a)
	delete(AOTP, a)
}
