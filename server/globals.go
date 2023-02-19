package main

import (
	"sync"

	"github.com/solarlune/resolv"
)


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

