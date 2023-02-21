package main

import (
	"math"
	"time"

	r "github.com/kainn9/grpc_game/server/roles"
	"github.com/solarlune/resolv"
)

// attackHandler handles player's attack inputs
func (cp *player) attackHandler(input string, world *world) {


	// can't attack while attacking
	if cp.currAttack != nil {
		return
	}


	// If the player pressed the primary attack button and the player is not currently attacking, initiate a primary attack
	if input == "primaryAttack" {
		cp.primaryAttack(world)
	}

	if input == string(r.TestAttackKey) {
		cp.test2Attack(world, r.TestAttackKey)
	}

}


func (cp *player) test2Attack(world *world, atKey r.AtKey) {

	// cs := cp.object.Space
	atk := cp.Attacks[atKey]

	if atk.Windup != nil {
		cp.windup = atKey

		delay := atk.Windup.Duration
		time.AfterFunc(time.Duration(delay)*time.Millisecond, func() { 
			
			cp.windup = ""
			cp.currAttack = atk
			cp.attackMovement = true
		})
	}
}

func (cp *player) movementPhase(atk *r.Attack) {
		mv := atk.Movement

		// if attack has no movment, return
		if mv == nil {
			return
		} // TODO: Fire atk immediately if no movement

		

		if cp.atkMovmentStartX == -100 {
			cp.maxSpeed = float64(mv.Increment)
			cp.atkMovmentStartX = int(cp.object.X)
		} else {

				

			if cp.facingRight {
				cp.speedX = float64(mv.Increment)
			} else {
				cp.speedX = float64(-mv.Increment)
			}

			distTraveled := int(math.Abs(float64(cp.atkMovmentStartX) - cp.object.X))

			if distTraveled > mv.Distance {		

				cp.attackMovement = false
				cp.maxSpeed = gamePhys.defaultMaxSpeed
				cp.atkMovmentStartX = -100

				// TODO: FIRE DASH ATTACK HERE
				cp.currAttack = nil // just nilling it for now, until todo
			}
		}
}



////--------------------------------


func (cp *player) primaryAttack(world *world) {

	cs := cp.object.Space
	atk := cp.Attacks[r.PrimaryAttackKey]
	cp.currAttack = atk

	atkObj := resolv.NewObject(cp.object.X+atk.OffsetX, cp.object.Y+atk.OffsetY, atk.Width, atk.Height, "attack")

	if !cp.facingRight {
		// modify to calc player width as origin is top left corner I think
		atkObj = resolv.NewObject(cp.object.X-(atk.OffsetX-atk.Width/2), cp.object.Y+atk.OffsetY, atk.Width, atk.Height, "attack")
	}

	serverConfig.AOTP[atkObj] = cp

	cp.currAttack = atk

	cs.Add(
		atkObj,
	)

	time.AfterFunc(time.Duration(atk.Duration)*time.Millisecond, func() {
		world.removeAtk(atkObj)
		cp.currAttack = nil
	})
}
