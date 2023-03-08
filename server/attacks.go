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

	if input == string(r.PrimaryAttackKey) {
		cp.attack(world, r.PrimaryAttackKey)
	}

	// NEW WAY(WIP)
	if input == string(r.SecondaryAttackKey) {
		cp.attack(world, r.SecondaryAttackKey)
	}

	if input == string(r.TertAttackKey) {
		cp.attack(world, r.TertAttackKey)
	}

	if input == string(r.QuaternaryAttackKey) {
		cp.attack(world, r.QuaternaryAttackKey)
	}
}

func (cp *player) attack(world *world, atKey r.AtKey) {

	atk := cp.Attacks[atKey]

	cp.windupPhase(atk, atKey)

}
