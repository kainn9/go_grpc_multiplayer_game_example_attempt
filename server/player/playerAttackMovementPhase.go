package player

import (
	"math"

	r "github.com/kainn9/grpc_game/server/roles"
)

const noMovmentStartSet int = -100

func (cp *Player) MovementPhase(atk *r.AttackData) {
	mv := atk.Movement

	// if no movment is set, resolve the movment/go right into the atk sequence
	if mv == nil {
		cp.resolveMovment(atk)
		return
	}

	movementSpeed := movementSpeed(mv, atk, cp)

	if cp.MovmentStartX == noMovmentStartSet {
		cp.MaxSpeed = float64(movementSpeed)
		cp.MovmentStartX = int(cp.Object.X)
	} else {

		if cp.FacingRight {
			cp.SpeedX = float64(movementSpeed)
		} else {
			cp.SpeedX = float64(-movementSpeed)
		}

		distTraveled := math.Abs(float64(cp.MovmentStartX) - cp.Object.X)

		maxDist := mv.Distance

		if atk.HasChargeEffect() && mv.UseChargeDist {
			maxDist = mv.Distance + (cp.ChargeValue * atk.MultFactorMvDist)
		}

		if distTraveled > maxDist {
			cp.resolveMovment(atk)
		}
	}

}

func movementSpeed(mv *r.Movement, atk *r.AttackData, cp *Player) float64 {
	movmentSpeed := mv.SpeedX

	if atk.HasChargeEffect() && mv.UseChargeSpeed {
		movmentSpeed += (cp.ChargeValue * atk.MultFactorMvSpeed)
		if movmentSpeed > 16 {
			movmentSpeed = 16
		}
	}

	return movmentSpeed
}

func (cp *Player) resolveMovment(atk *r.AttackData) {
	cp.AttackMovement = ""
	cp.MaxSpeed = cp.Role.Phys.DefaultMaxSpeed
	cp.MovmentStartX = noMovmentStartSet

	cp.attackSeqence(atk)
}

func (cp *Player) endMovment() {
	if cp.AttackMovementActive() {
		cp.AttackMovement = ""
		cp.MovmentStartX = noMovmentStartSet
		cp.attackSeqence(cp.CurrAttack)
	}
}

func (cp *Player) interruptMovment() {
	cp.Windup = ""
	cp.CurrAttack = nil
	cp.AttackMovement = ""
}
