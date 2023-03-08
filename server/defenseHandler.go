package main

import (
	"math"
	"time"
)

func (cp *player) defenseHandler(input string) {

	if input == "defense" && !cp.defenseCooldown {

		delay := cp.Defense.Delay
		time.AfterFunc(time.Duration(delay)*time.Millisecond, func() {

			cp.defending = true

		})
	}

}

func (cp *player) handleDefenseMovement() {
	movementSpeed := cp.Defense.Speed

	if cp.movmentStartX == noMovmentStartSet {
		cp.maxSpeed = float64(movementSpeed)
		cp.movmentStartX = int(cp.object.X)
	} else {

		if cp.facingRight {
			cp.speedX = float64(movementSpeed)
		} else {
			cp.speedX = float64(-movementSpeed)
		}

		distTraveled := math.Abs(float64(cp.movmentStartX) - cp.object.X)

		maxDist := cp.Defense.Displacment

		if distTraveled > maxDist {
			cp.defending = false
			cp.movmentStartX = noMovmentStartSet
			cp.maxSpeed = float64(gamePhys.defaultMaxSpeed)
			cp.defenseCooldown = true

			time.AfterFunc(time.Duration(cp.Defense.Cooldown)*time.Millisecond, func() {
				cp.defenseCooldown = false
			})
		}
	}
}
