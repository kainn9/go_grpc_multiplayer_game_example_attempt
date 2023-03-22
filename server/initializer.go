package main

import (
	"sync"

	r "github.com/kainn9/grpc_game/server/roles"
)

// serverConfigStruct holds server configuration and state information.
type serverConfigStruct struct {
	mutex         sync.RWMutex
	addr          string
	worldsMap     map[string]*world
	activePlayers map[string]*player
	roles         map[r.PlayerType]int32
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

// BIG TODO:
// all maps should be world scoped
// or find a way to remove maps in favor of
// references...for example storing what attack a hitbox
// belongs to on the hitbox itself
// or storing the player that owns the hitbox on the hitbox itself
var serverConfig = serverConfigStruct{
	addr:          ":50051",
	worldsMap:     make(map[string]*world),
	activePlayers: make(map[string]*player),
	mutex:         sync.RWMutex{},
	roles:         make(map[r.PlayerType]int32),
}

// initializer sets up initial configuration for the game.
func initializer() {
	// Add main and alt worlds to serverConfig.worldsMap.
	serverConfig.worldsMap["main"] = worlds.main
	serverConfig.worldsMap["alt"] = worlds.alt

	serverConfig.roles[r.Knight.RoleType] = 0
	serverConfig.roles[r.Monk.RoleType] = 1

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
