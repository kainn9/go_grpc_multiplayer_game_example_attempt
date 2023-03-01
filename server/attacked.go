package main

import (
	"log"
	"math"
	"time"

	r "github.com/kainn9/grpc_game/server/roles"
	"github.com/solarlune/resolv"
)

func (cp *player) attackedHandler() {

	// Check if the player is colliding with an attack object
	if check := cp.object.Check(cp.speedX, cp.speedY, "attack"); check != nil {

		// Loop through all attack objects(being collided with) and check if they belong to another player.
		atkObjs := check.Objects
		for _, o := range atkObjs {
			if o == nil {
				continue
			}

			// Get the attacker of the attack object
			atk, attacker := getAttackAndAttacker(o)

			// If the attacker is the same as the player being attacked,
			// or if the attacker is nil,
			// or if the player has already been hit by this attack, skip this collision
			if checkForValidAttackHit(attacker, cp, atk) {
				continue
			}


			cp.healthHandler(attacker, atk)
			cp.knockBackHandler(attacker, atk)			
		}
	}

	// TODO: Should this really live here(inside the attackedHandler())?
	if cp.isKnockedBackX() {
		cp.speedX += cp.kbx
	}

	if cp.isKnockedBackY() {
			cp.speedY += cp.kby
	}	
}


func getAttackAndAttacker(o *resolv.Object) (*r.Attack, *player) {
	serverConfig.mutex.RLock()
	attacker := serverConfig.AOTP[o]
	atk := serverConfig.OTA[o]
	serverConfig.mutex.RUnlock()

	return atk, attacker
}

func checkForValidAttackHit(attacker *player, cp *player, atk *r.Attack) bool {
	return attacker == cp || attacker == nil || (serverConfig.HTAP[attacker.pid + string(atk.Name)])
}


func (cp *player) healthHandler(attacker *player, atk *r.Attack) {
	serverConfig.mutex.Lock()
	serverConfig.HTAP[attacker.pid + string(atk.Name)] = true
	serverConfig.mutex.Unlock()


	dmg := atk.Damage

	if  atk.Windup != nil && atk.ChargeEffect != nil && atk.UseChargeDmg {
		dmg += int(math.Round(attacker.chargeValue * atk.ChargeEffect.MultFactorDmg))
	}

	cp.health -= dmg
	log.Printf("Player %s was hit by %s for %d damage\n", cp.pid, attacker.pid, atk.Damage)

	if cp.health <= 0 {
		cp.death()
		log.Printf("Player %s has died\n", cp.pid)
	}


	time.AfterFunc((time.Duration(500))*time.Millisecond, func() { 
		serverConfig.mutex.Lock()
		delete(serverConfig.HTAP, attacker.pid + string(atk.Name))
		serverConfig.mutex.Unlock() 
	})
}

// TODO: Dry out the chargeEffect check and break up into
// helper funcs
// also need to cleanup naming around kbx/kby as it's confusing
// and currently refers to speed instead of distance
// allthough probably will refactor to use distance instead of speed
// down the line(like the attack movment does)
func (cp *player) knockBackHandler(attacker *player, atk *r.Attack) {

	kbx := atk.KnockbackX
	if atk.Windup != nil && atk.ChargeEffect != nil && atk.UseChargeKbxSpeed {
		kbx += (attacker.chargeValue * atk.ChargeEffect.MultFactorKbxSpeed)
		kbx = math.Min(kbx, 16)
	}
			
	if attacker.object.X > cp.object.X {
		cp.kbx = -kbx
	} else {
		cp.kbx = kbx
	}

	kby := atk.KnockbackY
	if atk.Windup != nil && atk.ChargeEffect != nil && atk.UseChargeKbySpeed {
		kby += (attacker.chargeValue * atk.ChargeEffect.MultFactorKbySpeed)
		kby = math.Min(kby, 16)
	}

	if attacker.object.Y >= cp.object.Y {
		cp.kby = -kby
	} else {
		cp.kby = kby
	}

	kbxDur := atk.KnockbackXDuration
	if atk.Windup != nil && atk.ChargeEffect != nil && atk.UseChargeKbxDuration {
		kbxDur += int(attacker.chargeValue * atk.ChargeEffect.MultFactorKbxDur)
	}

	kbyDur := atk.KnockbackYDuration
	if atk.Windup != nil && atk.ChargeEffect != nil && atk.UseChargeKbyDuration {
		kbyDur = int(attacker.chargeValue * atk.ChargeEffect.MultFactorKbyDur)
	}

	time.AfterFunc((time.Duration(kbxDur))*time.Millisecond, func() { cp.kbx = 0 })
	time.AfterFunc((time.Duration(kbyDur))*time.Millisecond, func() { cp.kby = 0 })
}

// simple death for now(this will just cause the player to respawn with a new PID/Full health)
func  (cp *player) death() {
	removePlayerFromGame(cp.pid, serverConfig.worldsMap[cp.worldKey])
}
