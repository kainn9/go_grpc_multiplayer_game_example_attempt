package main

import (
	"sync"

	r "github.com/kainn9/grpc_game/server/roles"
)

// serverConfigStruct holds server configuration and state information.
type serverConfigStruct struct {
	mutex         sync.RWMutex
	addr          string
	worldsMap     map[int]*world
	activePlayers map[string]*player
	roles         map[r.PlayerType]int32
	startingWorld *world
	randomSpawn   bool
}

// worldsStruct holds world objects.
type worldsStruct struct {
	worldOne             *world
	worldTwo             *world
	worldThree           *world
	landOfYohoPassageOne *world
	landOfYohoPassageTwo *world
	landOfYohoVillage    *world
}

var worlds = worldsStruct{
	worldOne:             newWorld(848, 1600, introWorldBuilder, "introZone", 22, 600),
	worldTwo:             newWorld(848, 3200, mainWorldBuilder, "desert passage", 612, 500),
	worldThree:           newWorld(4000, 6000, altWorldBuilder, "land of poisoned water", 1250, 3700),
	landOfYohoPassageOne: newWorld(480, 960, landOfYohoPassageOneBuilder, "land of yoho passage one", 100, 100),
	landOfYohoPassageTwo: newWorld(756, 1100, landOfYohoPassageTwoBuilder, "land of yoho passage two", 100, 100),
	landOfYohoVillage:    newWorld(6000, 3278, landOfYohoVillageBuilder, "land of yoho village", 100, 100),
}

var serverConfig = serverConfigStruct{
	addr:          ":50051",
	worldsMap:     make(map[int]*world),
	activePlayers: make(map[string]*player),
	mutex:         sync.RWMutex{},
	roles:         make(map[r.PlayerType]int32),
	randomSpawn:   false,
}

// initializer sets up initial configuration for the game.
func initializer() {
	// Add worlds to worlds map
	// NOTE: YO BE CAREFUL ABOUT NOT SETTING A DUPE INDEX
	// IT WILL CREATE MUTEX DEADLOCK ON WORLD SWAPS!!!
	serverConfig.worldsMap[0] = worlds.worldOne
	worlds.worldOne.index = 0

	serverConfig.worldsMap[1] = worlds.worldTwo
	worlds.worldTwo.index = 1

	serverConfig.worldsMap[2] = worlds.worldThree
	worlds.worldThree.index = 2

	serverConfig.worldsMap[3] = worlds.landOfYohoPassageOne
	worlds.landOfYohoPassageOne.index = 3

	serverConfig.worldsMap[4] = worlds.landOfYohoPassageTwo
	worlds.landOfYohoPassageTwo.index = 4

	serverConfig.worldsMap[5] = worlds.landOfYohoVillage
	worlds.landOfYohoVillage.index = 5

	// intro world
	serverConfig.startingWorld = worlds.worldOne

	// set up roles
	serverConfig.roles[r.Knight.RoleType] = 0
	serverConfig.roles[r.Monk.RoleType] = 1
	serverConfig.roles[r.Demon.RoleType] = 2
	serverConfig.roles[r.Werewolf.RoleType] = 3
	serverConfig.roles[r.Mage.RoleType] = 4

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
