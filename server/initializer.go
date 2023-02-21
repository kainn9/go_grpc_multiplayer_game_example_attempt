package main

import (
	"log"
	"sync"
	"time"

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
	AOTP:          make(map[*resolv.Object]*player),
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


// Starts a new ticker loop that calls processEventsPerTick with the given world
func newTickLoop(w *world) {
	go func() {
		ticker := time.NewTicker(time.Second / 60)
		defer ticker.Stop()

		for range ticker.C {
			processEventsPerTick(w)
		}
	}()
}

// Process events in the given world, removing each event as it is processed
func processEventsPerTick(w *world) {
	// Log information about the number of events in the world, if it exceeds certain thresholds
	if len(w.events) > 25 {
		log.Printf("WORLD: %v\n", w.name)
		log.Printf("LEN! 25 %v\n", len(w.events) > 25)
		log.Printf("LEN! 50 %v\n", len(w.events) > 50)
		log.Printf("LEN! 100 %v\n", len(w.events) > 100)
	}

	// Iterate over the worlds event queue
	// process up to 100 events per tick
	for i := 0; i < 100; i++ {

		// If there are no events left or the current index is out of range, exit the loop
		if len(w.events) == 0 || i > len(w.events)-1 {
			break
		}

		// Get the current event
		ev := w.events[i]

		// If there is a player associated with the event, handle the event with the player and world
		if w.players[ev.Id] != nil {
			cp := w.players[ev.Id]
			requestHandler(cp, w, ev)
		}

		// Remove the event from the world's events queue
		w.mutex.RLock()
		w.events = append(w.events[:i], w.events[i+1:]...)
		defer w.mutex.RUnlock()
		i--
	}
}

// TODO: Struct up this guy...
var (
	cell = 8
)