package main

import (
	"sync"

	r "github.com/kainn9/grpc_game/server/roles"
	"github.com/solarlune/resolv"
)

// gamePhysStruct holds default physics values for the game.
type gamePhysStruct struct {
	defaultFriction float64
	defaultAccel    float64
	defaultMaxSpeed float64
	defaultJumpSpd  float64
	defaultGravity  float64
}

// serverConfigStruct holds server configuration and state information.
type serverConfigStruct struct {
	mutex         sync.RWMutex        
	addr          string               
	worldsMap     map[string]*world
	activePlayers map[string]*player
	AOTP          map[*resolv.Object]*player // Map of Attack resolv objects to Player struct, eventually should be world scoped.
	OTA map[*resolv.Object]*r.Attack
	HTAP map[string]bool
}

// worldsStruct holds world objects.
type worldsStruct struct {
	main *world // Main world object.
	alt  *world // Alternative world object.
}


var worlds = worldsStruct{
	main: newWorld(848, 3200, mainWorldBuilder, "main"),
	alt:  newWorld(4000, 6000, altWorldBuilder, "alt"),
}

var serverConfig = serverConfigStruct{
	addr:          ":50051",
	worldsMap:     make(map[string]*world),
	activePlayers: make(map[string]*player),
	AOTP:          make(map[*resolv.Object]*player), // TODO: This should be world scoped.
	OTA:           make(map[*resolv.Object]*r.Attack), // TODO: This should be world scoped.
	HTAP:          make(map[string]bool), // TODO: This should be world scoped.
	mutex:         sync.RWMutex{},
}

var gamePhys = gamePhysStruct{
	defaultFriction: 0.5, // Default friction value.
	defaultMaxSpeed: 4.0, // Default max speed value.
	defaultJumpSpd:  12.0, // Default jump speed value.
	defaultGravity:  0.75, // Default gravity value.
}

// initializer sets up initial configuration for the game.
func initializer() {
	// Add main and alt worlds to serverConfig.worldsMap.
	serverConfig.worldsMap["main"] = worlds.main
	serverConfig.worldsMap["alt"] = worlds.alt

	// Compute default acceleration value based on friction value.
	gamePhys.defaultAccel = 0.5 + gamePhys.defaultFriction

	// Start tick loops for each world.
	for _, w := range serverConfig.worldsMap {
		newTickLoop(w)
	}
}



// TODO: Struct up this guy...
var (
	cellX = 16
	cellY = 8
)