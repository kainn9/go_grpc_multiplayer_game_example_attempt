package main

import (
	"github.com/solarlune/resolv"
)
type Player struct {
	Object         *resolv.Object
	SpeedX         float64
	SpeedY         float64
	OnGround       *resolv.Object
	OnPlayer 			*resolv.Object
	WallSliding    *resolv.Object
	FacingRight    bool
	IgnorePlatform *resolv.Object
	Pid string
	WorldKey string
}

/*
	Creates a resolve object from a player
	and attaches it to a resolv space(usually one per world)
*/
func AddPlayerToSpace(space *resolv.Space, p *Player) *Player {
	p.Object = resolv.NewObject(612, 500, 35, 48)
	p.Object.SetShape(resolv.NewRectangle(0, 0, p.Object.W, p.Object.H))
	
	space.Add(p.Object)
	return p
}

/* 
	Helper to disconnect player
*/
func DisconnectPlayer (pid string, w *World) {
	w.Space.Remove(w.Players[pid].Object)
	delete(w.Players, pid)
	delete(activePlayers, pid)
}

/* 
	Helper to change a players current world
*/
func ChangePlayersWorld(oldWorld *World, newWorld *World, cp *Player) {
	delete(oldWorld.Players, cp.Pid)
	oldWorld.Space.Remove(cp.Object)
	newWorld.Players[cp.Pid] = cp
	AddPlayerToSpace(newWorld.Space, cp)
	cp.WorldKey = newWorld.Name
}