package player

import (
	"math"
	"time"

	ut "github.com/kainn9/grpc_game/util"
)

func (cp *Player) DefenseHandler(input string) {

	if cp.Role.Defense == nil {
		return
	}

	if input == "defense" && !cp.defenseOnCD() {

		delay := cp.Defense.Delay
		time.AfterFunc(time.Duration(delay)*time.Millisecond, func() {

			cp.Defending = true
			cp.SetDefCd()
			cp.handleDefenseCoolDown()

			if cp.Defense.DefenseDuration != 0 {
				cp.handleDefenseDuration()
			}

		})
	}

}

func (cp *Player) HandleDefenseMovement() {

	if cp.Defense.DefenseMovement == nil {
		return
	}

	movementSpeed := cp.Defense.Speed

	if cp.MovmentStartX == noMovmentStartSet {
		cp.MovmentStartX = int(cp.Object.X)
		cp.MaxSpeed = cp.Defense.Speed
	} else {

		if cp.FacingRight {
			cp.SpeedX = float64(movementSpeed)
		} else {
			cp.SpeedX = float64(-movementSpeed)
		}

		distTraveled := math.Abs(float64(cp.MovmentStartX) - cp.Object.X)

		maxDist := cp.Defense.Displacment

		if distTraveled > maxDist {
			cp.Defending = false
			cp.MovmentStartX = noMovmentStartSet

		}
	}
}

func (cp *Player) endDefenseMovement() {
	cp.Defending = false
	cp.MovmentStartX = noMovmentStartSet
}

func (cp *Player) handleDefenseCoolDown() {
	time.AfterFunc(time.Duration(cp.Defense.Cooldown)*time.Millisecond, func() {
		cp.CdStringMutex.Lock()
		defer cp.CdStringMutex.Unlock()

		newCdString := ut.SetNthCharTo0(cp.CdString, 4)
		cp.CdString = newCdString
	})
}

func (cp *Player) handleDefenseDuration() {
	time.AfterFunc(time.Duration(cp.Defense.DefenseDuration)*time.Millisecond, func() {
		cp.Defending = false

	})
}

func (cp *Player) defenseOnCD() bool {
	cp.CdStringMutex.RLock()
	defer cp.CdStringMutex.RUnlock()

	return string(cp.CdString[4]) == "1"
}

func (cp *Player) SetDefCd() {
	cp.CdStringMutex.Lock()
	defer cp.CdStringMutex.Unlock()
	newCdString := ut.SetNthCharTo1(cp.CdString, 4)
	cp.CdString = newCdString
}
