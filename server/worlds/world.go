package worlds

import (
	"errors"
	"sync"
	"time"

	evt "github.com/kainn9/grpc_game/server/event"
	gc "github.com/kainn9/grpc_game/server/globalConstants"
	pl "github.com/kainn9/grpc_game/server/player"

	"github.com/kainn9/grpc_game/server/roles"
	ut "github.com/kainn9/grpc_game/util"
	"github.com/solarlune/resolv"
)

type World struct {
	space           *resolv.Space         // A Resolv space for collision detection
	height          float64               // The height of the game's world
	width           float64               // The width of the game's world
	Players         map[string]*pl.Player // A map of players currently in the world
	name            string                // The name of the world
	Index           int                   // index in worlds map
	events          []*evt.Event          // An queue of events to be processed(world scoped)
	eventsMutex     sync.RWMutex
	hitboxMutex     sync.RWMutex
	WPlayersMutex   sync.RWMutex
	WorldSpawnCords *worldSpawnCords
}

type worldSpawnCords struct {
	X int
	Y int
}

// creates a new game world.
func NewWorld(height float64, width float64, worldBuilder builderFunc, name string, spawnX int, spawnY int) *World {
	w := &World{
		name:          name,
		width:         width,
		height:        height,
		Players:       make(map[string]*pl.Player),
		eventsMutex:   sync.RWMutex{},
		hitboxMutex:   sync.RWMutex{},
		WPlayersMutex: sync.RWMutex{},
		events:        make([]*evt.Event, 0),
		WorldSpawnCords: &worldSpawnCords{
			Y: spawnY,
			X: spawnX,
		},
	}

	// Initialize the world with the specified builder function.
	w.Init(worldBuilder)
	return w
}

// initializes the game world.
func (world *World) Init(worldBuilder builderFunc) {
	gw := world.width
	gh := world.height

	// Define the world's Resolv Space. A Space is essentially a grid made up of 16x16 cells.
	// Each cell can have 0 or more Objects within it, and collisions can be found by checking the Space to see if the Cells at specific positions contain (or would contain) Objects.
	// This is a broad, simplified approach to collision detection.

	// Generally, you want cells to be the size of the smallest collide-able objects in your game,
	// and you want to move Objects at a maximum speed of one cell size per collision check to avoid
	// missing any possible collisions.

	world.space = resolv.NewSpace(int(gw), int(gh), gc.CELL_X, gc.CELL_Y)

	// Construct the solid level geometry. Note that the simple approach of checking cells in a Space for collision works simply when the geometry is aligned with the cells.
	worldBuilder(world, gw, gh)
}

// A function to update the game world, where all game logic happens.
// The physics are basically a rip of the Resolv example: https://github.com/SolarLune/resolv/blob/master/examples/worldPlatformer.go.
func (world *World) Update(cp *pl.Player, input string) {

	if !cp.DeathCallBackPending && cp.Dying {
		cp.DeathCallBackPending = true

		time.AfterFunc((time.Duration(1650))*time.Millisecond, func() {
			RemovePlayerFromWorld(cp.Pid, cp.CurrentWorld.(*World))
		})
		return
	}

	// need two guard clauses:
	// 1 for setting up pending death
	// 2 for disabling interaction as the player is dying
	if cp.Dying {
		return
	}

	// Add the "player" tag to the player object if it doesn't already have it.
	if !cp.Object.HasTags("player") {
		cp.Object.AddTags("player")
	}

	worldTransferHandler(cp, input)
	rotateRole(cp, input)

	// Can't do reg movement when attacking
	if !cp.CanAcceptInputs() {
		cp.SpeedX = 0
	} else {
		// Handle player inputs
		cp.HorizontalMovementListener(input)
		cp.JumpHandler(input)

		cp.AttackHandler(input, world)
		cp.DefenseHandler(input)
	}

	if cp.AttackMovementActive() {
		cp.MovementPhase(cp.CurrAttack)
	}

	if cp.Defending {
		cp.HandleDefenseMovement()
	}

	// Handle player getting attacked.
	cp.AttackedHandler()

	// Handle player Phys and collisions.

	cp.GravityHandler()
	cp.HorizontalMovementHandler(input)
	cp.VerticalMovmentHandler(input, world)

	// portal check
	if check := cp.Object.Check(cp.SpeedX, 0, "portal"); check != nil {

		portal := check.Objects[0]
		portalData := portalData(portal)

		oldWorld := cp.CurrentWorld
		newWorld := WorldsConfig.worldsMap[portalData.worldKey]

		changePlayersWorld(oldWorld.(*World), newWorld, cp, portalData.x, portalData.y)

	}

	// IDK where to put this yet...
	// its wall slide stuff
	// is it vertical is it horizontal?
	// I think its technically vertical lol
	wallNext := 1.0
	if !cp.FacingRight {
		wallNext = -1
	}
	// If the wall next to the Player runs out, stop wall sliding
	if c := cp.Object.Check(wallNext, 0, "solid"); cp.WallSliding != nil && c == nil {
		cp.WallSliding = nil
	}

	world.hitboxMutex.Lock()
	cp.Object.Update()
	world.hitboxMutex.Unlock()

}

// addPlayerToSpace adds a player to a Resolv space with the given coordinates
func AddPlayerToSpace(w *World, p *pl.Player, x float64, y float64) *pl.Player {

	p.Object = resolv.NewObject(x, y, p.HitBoxW, p.HitBoxH)

	p.Object.SetShape(resolv.NewRectangle(0, 0, p.Object.W, p.Object.H))

	pl.InitHitboxData(p.Object, p, nil)
	w.hitboxMutex.Lock()
	w.space.Add(p.Object)
	w.hitboxMutex.Unlock()
	return p
}

// removes a player from a world with the given unique identifier and world pointer
// player will respawn in default world if gRPC conn is still live
func RemovePlayerFromWorld(pid string, w *World) {

	w.WPlayersMutex.RLock()
	p := w.Players[pid]
	w.WPlayersMutex.RUnlock()

	// Stop server crash if client disconnects before fully loading/creating a player
	if p == nil {
		return
	}

	w.hitboxMutex.Lock()
	obj := w.Players[pid].Object
	w.space.Remove(obj)
	w.hitboxMutex.Unlock()

	w.WPlayersMutex.Lock()
	delete(w.Players, pid)
	w.WPlayersMutex.Unlock()

	WorldsConfig.Mutex.Lock()
	delete(WorldsConfig.ActivePlayers, pid)
	WorldsConfig.Mutex.Unlock()
}

// changePlayersWorld swaps a player from their old world to a new world,
// updating their position and worldKey in the process.
// pass 0 for optX/Y to use world default spawn cords
func changePlayersWorld(oldWorld *World, newWorld *World, cp *pl.Player, optX int, optY int) {

	var x int
	var y int

	if optX == 0 {
		x = newWorld.WorldSpawnCords.X
	} else {
		x = optX
	}

	if optY == 0 {
		y = newWorld.WorldSpawnCords.Y
	} else {
		y = optY
	}

	oldWorld.WPlayersMutex.Lock()
	delete(oldWorld.Players, cp.Pid)
	oldWorld.WPlayersMutex.Unlock()

	oldWorld.hitboxMutex.Lock()
	oldWorld.space.Remove(cp.Object)
	oldWorld.hitboxMutex.Unlock()

	newWorld.WPlayersMutex.Lock()
	newWorld.Players[cp.Pid] = cp
	newWorld.WPlayersMutex.Unlock()

	AddPlayerToSpace(newWorld, cp, float64(x), float64(y-20))
	cp.CurrentWorld = newWorld
}

func worldTransferHandler(cp *pl.Player, input string) {
	if input == "swap" {
		w, k := CurrentPlayerWorldOrDefault(cp.Pid)

		newWorldKey := k + 1

		if newWorldKey >= len(WorldsConfig.worldsMap) {
			newWorldKey = 0
		}

		changePlayersWorld(w, WorldsConfig.worldsMap[newWorldKey], cp, 0, 0)
		return
	}
}

func rotateRole(cp *pl.Player, input string) {
	if input == "roleSwap" {
		getKey := func(m map[*roles.Role]int32, target int32) *roles.Role {
			for key, value := range m {
				if value == target {
					return key
				}
			}
			return nil
		}

		currentRoleKey := WorldsConfig.Roles[cp.Role] + 1

		oldRoleHeight := cp.Role.HitBoxH

		maxKey := len(WorldsConfig.Roles)

		if currentRoleKey > int32(maxKey)-1 {
			currentRoleKey = 0
		}

		newRole := getKey(WorldsConfig.Roles, currentRoleKey)

		newRoleHeight := newRole.HitBoxH

		roleHeightDiff := oldRoleHeight - newRoleHeight
		if roleHeightDiff > 0 {
			roleHeightDiff = 0
		}

		cp.RotateRoleData(newRole)

		w := cp.CurrentWorld // Note/TODO: may not be concurrent safe

		w.GetHitboxMutex().Lock()
		w.GetSpace().Remove(cp.Object)
		w.GetHitboxMutex().Unlock()

		AddPlayerToSpace(w.(*World), cp, cp.Object.X, cp.Object.Y+(roleHeightDiff))
	}
}

/*
	There is no authN or persistance currently.
	Identity boils down to a randomly generated
	"Player ID"(pid) that is created when the
	client first connects to the server and the
	bidirectional stream is setup

	Anwyay, we use the following two functions
	to track which world the player is using these PIDs
*/

func LocateFromPID(pid string) (world *World, worldKey int, err error) {
	WorldsConfig.Mutex.RLock()
	player := WorldsConfig.ActivePlayers[pid]
	WorldsConfig.Mutex.RUnlock()

	if player != nil {
		w := player.CurrentWorld

		if w != nil {
			return w.(*World), player.CurrentWorld.GetIndex(), nil
		}

	}

	return nil, 0, errors.New("no match found")
}

/*
Helper that uses LocateFromPID
to return a world that a player
is currently attached to. Unlike
LocateFromPID this defaults to returning
the main/starting world
*/
func CurrentPlayerWorldOrDefault(pid string) (world *World, worldKey int) {

	// adding new player to default world
	// or setting current world to where
	// the active player is located
	w, k, err := LocateFromPID(pid)

	if err != nil {
		if WorldsConfig.randomSpawn {
			wKey := int(ut.RandomInt(int64(len(WorldsConfig.worldsMap))))
			return WorldsConfig.worldsMap[wKey], wKey
		}

		return WorldsConfig.startingWorld, WorldsConfig.startingWorld.Index
	}
	return w, k
}
