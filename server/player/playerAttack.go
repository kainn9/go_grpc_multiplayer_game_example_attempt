package player

import (
	r "github.com/kainn9/grpc_game/server/roles"
)

// attackHandler handles player's attack inputs
func (cp *Player) AttackHandler(input string, world World) {

	// can't attack while attacking yo
	if cp.CurrAttack != nil {
		return
	}

	if input == string(r.PrimaryAttackKey) {
		cp.attack(world, r.PrimaryAttackKey)
	}

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

func (cp *Player) attack(world World, atKey r.AtKey) {

	atk := cp.Attacks[atKey]

	cp.windupPhase(atk, atKey)

}
