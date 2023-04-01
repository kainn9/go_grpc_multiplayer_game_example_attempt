package main

import (
	"math"

	r "github.com/kainn9/grpc_game/server/roles"
)

const noMovmentStartSet int = -100

func (cp *player) movementPhase(atk *r.AttackData) {
	mv := atk.Movement

	// if no movment is set, resolve the movment/go right into the atk sequence
	if mv == nil {
		cp.resolveMovment(atk)
		return
	}

	movementSpeed := movementSpeed(mv, atk, cp)

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

		maxDist := mv.Distance

		if atk.HasChargeEffect() && mv.UseChargeDist {
			maxDist = mv.Distance + (cp.chargeValue * atk.MultFactorMvDist)
		}

		if distTraveled > maxDist {
			cp.resolveMovment(atk)
		}
	}

}

func movementSpeed(mv *r.Movement, atk *r.AttackData, cp *player) float64 {
	movmentSpeed := mv.SpeedX

	if atk.HasChargeEffect() && mv.UseChargeSpeed {
		movmentSpeed += (cp.chargeValue * atk.MultFactorMvSpeed)
		if movmentSpeed > 16 {
			movmentSpeed = 16
		}
	}

	return movmentSpeed
}

func (cp *player) resolveMovment(atk *r.AttackData) {
	cp.attackMovement = ""
	cp.maxSpeed = cp.Role.Phys.DefaultMaxSpeed
	cp.movmentStartX = noMovmentStartSet

	cp.attackSeqence(atk)
}

func (cp *player) endMovment() {
	if cp.attackMovementActive() {
		cp.attackMovement = ""
		cp.movmentStartX = noMovmentStartSet
		cp.attackSeqence(cp.currAttack)
	}
}

func (cp *player) interruptMovment() {
	cp.windup = ""
	cp.currAttack = nil
	cp.attackMovement = ""
}
