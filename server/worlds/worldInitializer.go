package worlds

import (
	"sync"

	pl "github.com/kainn9/grpc_game/server/player"
	r "github.com/kainn9/grpc_game/server/roles"
)

// WorldsConfigStruct holds server configuration and state information.
type WorldsConfigStruct struct {
	Mutex         sync.RWMutex
	worldsMap     map[int]*World
	ActivePlayers map[string]*pl.Player
	Roles         map[*r.Role]int32
	startingWorld *World
	randomSpawn   bool
}

// worldsStruct holds world objects.
type worldsStruct struct {
	worldOne             *World
	worldTwo             *World
	worldThree           *World
	landOfYohoPassageOne *World
	landOfYohoPassageTwo *World
	landOfYohoVillage    *World
}

var worlds = worldsStruct{
	worldOne:             NewWorld(848, 1600, IntroWorldBuilder, "introZone", 564, 418),
	worldTwo:             NewWorld(848, 3200, MainWorldBuilder, "desert passage", 1658, 500),
	worldThree:           NewWorld(4000, 6000, AltWorldBuilder, "land of poisoned water", 1298, 672),
	landOfYohoPassageOne: NewWorld(480, 960, LandOfYohoPassageOneBuilder, "land of yoho passage one", 430, 131),
	landOfYohoPassageTwo: NewWorld(756, 1100, LandOfYohoPassageTwoBuilder, "land of yoho passage two", 426, 488),
	landOfYohoVillage:    NewWorld(6000, 3278, LandOfYohoVillageBuilder, "land of yoho village", 1819, 153),
}

var WorldsConfig = WorldsConfigStruct{
	worldsMap:     make(map[int]*World),
	ActivePlayers: make(map[string]*pl.Player),
	Mutex:         sync.RWMutex{},
	Roles:         make(map[*r.Role]int32),
	randomSpawn:   false,
}

// initializer sets up initial configuration for the game.
func Initializer() {
	// Add worlds to worlds map
	// NOTE: YO BE CAREFUL ABOUT NOT SETTING A DUPE INDEX
	// IT WILL CREATE MUTEX DEADLOCK ON WORLD SWAPS!!!
	WorldsConfig.worldsMap[0] = worlds.worldOne
	worlds.worldOne.Index = 0

	WorldsConfig.worldsMap[1] = worlds.worldTwo
	worlds.worldTwo.Index = 1

	WorldsConfig.worldsMap[2] = worlds.worldThree
	worlds.worldThree.Index = 2

	WorldsConfig.worldsMap[3] = worlds.landOfYohoPassageOne
	worlds.landOfYohoPassageOne.Index = 3

	WorldsConfig.worldsMap[4] = worlds.landOfYohoPassageTwo
	worlds.landOfYohoPassageTwo.Index = 4

	WorldsConfig.worldsMap[5] = worlds.landOfYohoVillage
	worlds.landOfYohoVillage.Index = 5

	// intro world
	WorldsConfig.startingWorld = worlds.worldOne

	// set up roles
	WorldsConfig.Roles[r.Knight] = 0
	WorldsConfig.Roles[r.Monk] = 1
	WorldsConfig.Roles[r.Demon] = 2
	WorldsConfig.Roles[r.Werewolf] = 3
	WorldsConfig.Roles[r.Mage] = 4
	WorldsConfig.Roles[r.HeavyKnight] = 5

	// Start tick loops for each world.
	for _, w := range WorldsConfig.worldsMap {
		newTickLoop(w)
	}
}
