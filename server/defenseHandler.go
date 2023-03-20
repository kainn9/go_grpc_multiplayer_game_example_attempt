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
		cp.movmentStartX = int(cp.object.X)
		cp.maxSpeed = cp.Defense.Speed
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
			cp.defenseCooldown = true

			time.AfterFunc(time.Duration(cp.Defense.Cooldown)*time.Millisecond, func() {
				cp.defenseCooldown = false
			})
		}
	}
}

func (cp *player) endDefenseMovement() {
	cp.defending = false
	cp.movmentStartX = noMovmentStartSet
}
