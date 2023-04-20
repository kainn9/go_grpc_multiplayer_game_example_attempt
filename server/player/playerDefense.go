package player

import (
	"math"
	"time"
)

func (cp *Player) DefenseHandler(input string) {

	if cp.Role.Defense == nil {
		return
	}

	if input == "defense" && !cp.DefenseCooldown {

		delay := cp.Defense.Delay
		time.AfterFunc(time.Duration(delay)*time.Millisecond, func() {

			cp.Defending = true

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
			cp.DefenseCooldown = true

			cp.handleDefenseCoolDown()
		}
	}
}

func (cp *Player) endDefenseMovement() {
	cp.Defending = false
	cp.MovmentStartX = noMovmentStartSet
}

func (cp *Player) handleDefenseCoolDown() {
	time.AfterFunc(time.Duration(cp.Defense.Cooldown)*time.Millisecond, func() {
		cp.DefenseCooldown = false
	})
}

func (cp *Player) handleDefenseDuration() {
	time.AfterFunc(time.Duration(cp.Defense.DefenseDuration)*time.Millisecond, func() {
		cp.Defending = false

	})
}
