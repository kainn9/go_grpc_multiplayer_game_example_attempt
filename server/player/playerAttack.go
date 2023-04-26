package player

import (
	"time"

	r "github.com/kainn9/grpc_game/server/roles"
	ut "github.com/kainn9/grpc_game/util"
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

	if atk == nil {
		return
	}

	cp.CdStringMutex.RLock()
	onCd := string(cp.CdString[r.AtkOrderMap[atKey]]) == "1"
	cp.CdStringMutex.RUnlock()

	if onCd {
		return
	}

	if atk.Cooldown != 0 {

		cp.CdStringMutex.Lock()
		newCdString := ut.SetNthCharTo1(cp.CdString, r.AtkOrderMap[atKey])
		cp.CdString = newCdString
		cp.CdStringMutex.Unlock()

		time.AfterFunc((time.Duration(atk.Cooldown))*time.Millisecond, func() {
			cp.CdStringMutex.Lock()
			defer cp.CdStringMutex.Unlock()

			newCdString := ut.SetNthCharTo0(cp.CdString, r.AtkOrderMap[atKey])
			cp.CdString = newCdString
		})
	}

	if atk.InvincibleNoBox {
		cp.InvincibleNoBox = true
	}

	cp.windupPhase(atk, atKey)

}
