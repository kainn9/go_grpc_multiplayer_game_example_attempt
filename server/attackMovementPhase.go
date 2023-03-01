package main

import (
	"math"

	r "github.com/kainn9/grpc_game/server/roles"
)

const noMovmentStartSet int = -100

func (cp *player) movementPhase(atk *r.Attack) {
		mv := atk.Movement


		// if no movment is set, resolve the movment/go right into the atk sequence
		if mv == nil {
			cp.resolveMovment(atk)
			return
		}
		
	
		movementSpeed := movementSpeed(mv, atk, cp)

		if cp.atkMovmentStartX == noMovmentStartSet {
			cp.maxSpeed = float64(movementSpeed)
			cp.atkMovmentStartX = int(cp.object.X)
		} else {

			if cp.facingRight {
				cp.speedX = float64(movementSpeed)
			} else {
				cp.speedX = float64(-movementSpeed)
			}

			distTraveled := math.Abs(float64(cp.atkMovmentStartX) - cp.object.X)


			maxDist := mv.Distance

			if atk.ChargeEffect != nil && mv.UseChargeDist {
				maxDist = mv.Distance + (cp.chargeValue * atk.MultFactorMvDist)
			}
			
			if distTraveled > maxDist {		
				cp.resolveMovment(atk)
			}
		}
		
}


func movementSpeed(mv *r.Movement, atk *r.Attack, cp *player) float64 {
	movmentSpeed := mv.SpeedX

	if atk.ChargeEffect != nil && mv.UseChargeSpeed {
		movmentSpeed += (cp.chargeValue * atk.MultFactorMvSpeed)
		if movmentSpeed > 16 {
			movmentSpeed = 16
		}
	}

	return movmentSpeed
}

func (cp *player) resolveMovment(atk *r.Attack) {
	cp.attackMovement = ""
	cp.maxSpeed = gamePhys.defaultMaxSpeed
	cp.atkMovmentStartX = noMovmentStartSet
			
	cp.attackSeqence(atk)
}

func (cp *player) endMovment(){
	if cp.attackMovementActive() {
		cp.attackMovement = ""
		cp.atkMovmentStartX = noMovmentStartSet
		cp.attackSeqence(cp.currAttack)
	}
}

