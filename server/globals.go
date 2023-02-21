package main

// TODO move these out of Global Scope and into "config" structs
import (
	"sync"

	"github.com/solarlune/resolv"
)

/*
	Putting global scope server vars and consts here to avoid confusion
*/

const (
	defaultFriction = 0.5
	defaultAccel    = 0.5 + defaultFriction
	defaultMaxSpeed = 4.0
	defaultJumpSpd  = 10.0
	defaultGravity  = 0.75
)

var (
	mutex         sync.RWMutex
	addr          string = ":50051"
	mainW                = NewWorld(848, 3200, MainWorldBuilder, "main")
	altW                 = NewWorld(4000, 6000, AltWorldBuilder, "alt")
	worldsMap            = make(map[string]*World)
	activePlayers        = make(map[string]*Player)
	AOTP                 = make(map[*resolv.Object]*Player) // map of Attack resolv objects to Player struct
)
