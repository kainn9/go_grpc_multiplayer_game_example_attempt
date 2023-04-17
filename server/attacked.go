package main

import (
	"log"
	"math"
	"time"

	r "github.com/kainn9/grpc_game/server/roles"
	se "github.com/kainn9/grpc_game/server/statusEffects"
)

func (cp *player) attackedHandler() {

	if cp.defending {
		return
	}

	// Check if the player is colliding with an attack object
	if check := cp.object.Check(cp.speedX, cp.speedY, "attack"); check != nil {

		// Loop through all attack objects(being collided with)
		// only "process" attack if attack is valid
		atkObjs := check.Objects
		for _, o := range atkObjs {

			// Get the attacker of the attack object
			hitBoxData := hBoxData(o)
			atk := hitBoxData.attackData
			attacker := hitBoxData.player

			if hitIsInvalid(cp, hitBoxData) {
				continue
			}

			cp.healthHandler(attacker, atk, hitBoxData.aid)
			cp.applyAttackCC(attacker, atk)
			cp.interruptWindup()
			cp.interruptMovment()
		}
	}

	// TODO: Should this really live here(inside the attackedHandler())?
	if cp.isKnockedBackX() {
		cp.speedX = cp.kbx
	}

	if cp.isKnockedBackY() {
		cp.speedY = cp.kby
	}
}

// returns true if the attack is invalid, used to skip collision
func hitIsInvalid(cp *player, hitBoxData *hitBoxData) bool {
	return cp.hits[hitBoxData.aid] || hitBoxData.player == cp
}

func (cp *player) healthHandler(attacker *player, atk *r.AttackData, aid string) {
	serverConfig.mutex.Lock()
	cp.hits[aid] = true
	serverConfig.mutex.Unlock()

	dmg := atk.Damage

	if atk.HasChargeEffect() && atk.UseChargeDmg {
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
		delete(cp.hits, aid)
		serverConfig.mutex.Unlock()
	})
}

// TODO: Dry out the chargeEffect check and break up into
// helper funcs
// also need to cleanup naming around kbx/kby as it's confusing
// and currently refers to speed instead of distance
// allthough probably will refactor to use distance instead of speed
// down the line(like the attack movment does)
func (cp *player) applyAttackCC(attacker *player, atk *r.AttackData) {

	// speed
	kbx := atk.KnockbackX
	kby := atk.KnockbackY

	isStun := kbx == se.StunFloat && kby == se.StunFloat
	isHit := kbx == se.HitFloat && kby == se.HitFloat

	if atk.HasChargeEffect() && atk.UseChargeKbxSpeed {
		kbx += (attacker.chargeValue * atk.ChargeEffect.MultFactorKbxSpeed)
		kbx = math.Min(kbx, 16)
	}

	if attacker.object.X > cp.object.X {
		cp.kbx = -kbx
	} else {
		cp.kbx = kbx
	}

	if atk.HasChargeEffect() && atk.UseChargeKbySpeed {
		kby += (attacker.chargeValue * atk.ChargeEffect.MultFactorKbySpeed)
		kby = math.Min(kby, 16)
	}

	if (attacker.object.Y - attacker.HitBoxH) >= (cp.object.Y - cp.HitBoxH) {
		cp.kby = -kby
	} else if !atk.FixedKby {
		cp.kby = kby
	}

	if isStun {
		cp.kby = se.StunFloat
		cp.kbx = se.StunFloat
	}

	if isHit {
		cp.kby = se.HitFloat
		cp.kbx = se.HitFloat
	}

	// duration
	kbxDur := atk.KnockbackXDuration
	if atk.HasChargeEffect() && atk.UseChargeKbxDuration {
		kbxDur += int(attacker.chargeValue * atk.ChargeEffect.MultFactorKbxDur)
	}

	kbyDur := atk.KnockbackYDuration
	if atk.HasChargeEffect() && atk.UseChargeKbyDuration {
		kbyDur = int(attacker.chargeValue * atk.ChargeEffect.MultFactorKbyDur)
	}

	time.AfterFunc((time.Duration(kbxDur))*time.Millisecond, func() { cp.kbx = 0 })
	time.AfterFunc((time.Duration(kbyDur))*time.Millisecond, func() { cp.kby = 0 })
}

// simple death for now(this will just cause the player to respawn with a new PID/Full health)
func (cp *player) death() {
	cp.dead = true

	time.AfterFunc((time.Duration(1650))*time.Millisecond, func() {
		removePlayerFromGame(cp.pid, serverConfig.worldsMap[cp.worldKey])
	})

}
