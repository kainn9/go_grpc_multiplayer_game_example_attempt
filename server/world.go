package main

import (
	"sync"
	"time"

	"github.com/solarlune/resolv"
)

type world struct {
	space           *resolv.Space      // A Resolv space for collision detection
	height          float64            // The height of the game's world
	width           float64            // The width of the game's world
	players         map[string]*player // A map of players currently in the world
	name            string             // The name of the world
	index           int                // index in worlds map
	events          []*event           // An queue of events to be processed(world scoped)
	eventsMutex     sync.RWMutex       // A mutex to lock down resources when necessary(world scoped)
	hitboxMutex     sync.RWMutex       // A mutex to lock down resources when necessary(world scoped)
	worldSpawnCords *worldSpawnCords
}

type worldSpawnCords struct {
	x int
	y int
}

// creates a new game world.
func newWorld(height float64, width float64, worldBuilder builderFunc, name string, spawnX int, spawnY int) *world {
	w := &world{
		name:        name,
		width:       width,
		height:      height,
		players:     make(map[string]*player),
		eventsMutex: sync.RWMutex{},
		hitboxMutex: sync.RWMutex{},
		events:      make([]*event, 0),
		worldSpawnCords: &worldSpawnCords{
			x: spawnX,
			y: spawnY,
		},
	}

	// Initialize the world with the specified builder function.
	w.Init(worldBuilder)
	return w
}

// initializes the game world.
func (world *world) Init(worldBuilder builderFunc) {
	gw := world.width
	gh := world.height

	// Define the world's Resolv Space. A Space is essentially a grid made up of 16x16 cells.
	// Each cell can have 0 or more Objects within it, and collisions can be found by checking the Space to see if the Cells at specific positions contain (or would contain) Objects.
	// This is a broad, simplified approach to collision detection.

	// Generally, you want cells to be the size of the smallest collide-able objects in your game,
	// and you want to move Objects at a maximum speed of one cell size per collision check to avoid
	// missing any possible collisions.

	world.space = resolv.NewSpace(int(gw), int(gh), cellX, cellY)

	// Construct the solid level geometry. Note that the simple approach of checking cells in a Space for collision works simply when the geometry is aligned with the cells.
	worldBuilder(world, gw, gh)
}

// A function to update the game world, where all game logic happens.
// The physics are basically a rip of the Resolv example: https://github.com/SolarLune/resolv/blob/master/examples/worldPlatformer.go.
func (world *world) Update(cp *player, input string) {

	if cp.dead {
		return
	}
	// Add the "player" tag to the player object if it doesn't already have it.
	if !cp.object.HasTags("player") {
		cp.object.AddTags("player")
	}

	// TODO: Move make an actual buff system/tracker
	// Implement a gravity boost/buff if the "gravBoost" input is received and the player doesn't have the buff already.
	if input == "gravBoost" && !cp.gravBoost {
		cp.jumpSpd = 15
		cp.gravBoost = true
		time.AfterFunc(20*time.Second, func() { cp.jumpSpd = cp.Role.Phys.DefaultJumpSpd })
		time.AfterFunc(120*time.Second, func() { cp.gravBoost = false })
	}

	// Handle player world transfers.
	// ATM, this is spammable w/e for dev purposes
	cp.worldTransferHandler(input)

	// Can't do reg movement when attacking
	if !cp.canAcceptInputs() {
		cp.speedX = 0
	} else {
		// Handle player inputs
		cp.horizontalMovementListener(input)
		cp.jumpHandler(input)
		cp.attackHandler(input, world)
		cp.defenseHandler(input)
	}

	if cp.attackMovementActive() {
		cp.movementPhase(cp.currAttack)
	}

	if cp.defending {
		cp.handleDefenseMovement()
	}

	// Handle player getting attacked.
	cp.attackedHandler()

	// Handle player Phys and collisions.

	cp.gravityHandler()
	cp.horizontalMovementHandler(input, world.width)
	cp.verticalMovmentHandler(input, world)

	// IDK where to put this yet...
	// its wall slide stuff
	// is it vertical is it horizontal?
	// I think its technically vertical lol
	wallNext := 1.0
	if !cp.facingRight {
		wallNext = -1
	}
	// If the wall next to the Player runs out, stop wall sliding
	if c := cp.object.Check(wallNext, 0, "solid"); cp.wallSliding != nil && c == nil {
		cp.wallSliding = nil
	}
	world.hitboxMutex.Lock() // TODO: maybe??
	cp.object.Update()       // Update the player's position in the space.
	world.hitboxMutex.Unlock()

}
