package main

import (
	r "github.com/kainn9/grpc_game/server/roles"
)

// attackHandler handles player's attack inputs
func (cp *player) attackHandler(input string, world *world) {


	// can't attack while attacking yo
	if cp.currAttack != nil {
		return
	}

	if input == "primaryAtk" {
		cp.attack(world, r.PrimaryAttackKey)
	}

	// NEW WAY(WIP)
	if input == string("secondaryAtk") {
		cp.attack(world, r.SecondaryAttackKey)
	}

	if input == string("tertAtk") {
		cp.attack(world, r.TertAttackKey)
	}
}


func (cp *player) attack(world *world, atKey r.AtKey) {

	atk := cp.Attacks[atKey]

	cp.windupPhase(atk, atKey)

}

