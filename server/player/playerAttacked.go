package player

import (
	"log"
	"math"
	"time"

	r "github.com/kainn9/grpc_game/server/roles"
	se "github.com/kainn9/grpc_game/server/statusEffects"
)

func (cp *Player) AttackedHandler() {

	if cp.Defending {
		return
	}

	// Check if the player is colliding with an attack object
	if check := cp.Object.Check(cp.SpeedX, cp.SpeedY, "attack"); check != nil {

		// Loop through all attack objects(being collided with)
		// only "process" attack if attack is valid
		atkObjs := check.Objects
		for _, o := range atkObjs {

			// Get the attacker of the attack object
			hitBoxData := HBoxData(o)
			atk := hitBoxData.AttackData
			attacker := hitBoxData.Player

			if hitIsInvalid(cp, hitBoxData) {
				continue
			}

			cp.HealthHandler(attacker, atk, hitBoxData.Aid)
			cp.initAttackCC(attacker, atk)
			cp.interruptWindup()
			cp.interruptMovment()
		}
	}

	cp.applyCC()
}

// returns true if the attack is invalid, used to skip collision
func hitIsInvalid(cp *Player, hitBoxData *hitBoxData) bool {
	cp.HitsMutex.RLock()
	hit := cp.Hits[hitBoxData.Aid]
	cp.HitsMutex.RUnlock()
	return hit || hitBoxData.Player == cp
}

func (cp *Player) HealthHandler(attacker *Player, atk *r.AttackData, aid string) {
	cp.HitsMutex.Lock()
	cp.Hits[aid] = true
	cp.HitsMutex.Unlock()

	dmg := atk.Damage
	log.Println("Yo Charge")

	if atk.HasChargeEffect() && atk.UseChargeDmg {
		dmg += int(math.Round(attacker.ChargeValue * atk.ChargeEffect.MultFactorDmg))
	}

	cp.Health -= dmg
	log.Printf("Player %s was hit by %s for %d damage\n", cp.Pid, attacker.Pid, atk.Damage)

	if cp.Health <= 0 {
		cp.death()
		log.Printf("Player %s has died\n", cp.Pid)
	}

	time.AfterFunc((time.Duration(500))*time.Millisecond, func() {
		cp.HitsMutex.Lock()
		delete(cp.Hits, aid)
		cp.HitsMutex.Unlock()
	})
}

// NOTE:
// KBY/KBX refers to speed, not distance
// may want to fix/change this down the line
func (cp *Player) initAttackCC(attacker *Player, atk *r.AttackData) {

	// speed
	kbx := atk.KnockbackX
	kby := atk.KnockbackY

	isStun := kbx == se.StunFloat && kby == se.StunFloat
	isHit := kbx == se.HitFloat && kby == se.HitFloat

	if atk.HasChargeEffect() && atk.UseChargeKbxSpeed {
		kbx += (attacker.ChargeValue * atk.ChargeEffect.MultFactorKbxSpeed)
		kbx = math.Min(kbx, 16)
	}

	if attacker.Object.X > cp.Object.X {
		cp.Kbx = -kbx
	} else {
		cp.Kbx = kbx
	}

	if atk.HasChargeEffect() && atk.UseChargeKbySpeed {
		kby += (attacker.ChargeValue * atk.ChargeEffect.MultFactorKbySpeed)
		kby = math.Min(kby, 16)
	}

	if (attacker.Object.Y-attacker.HitBoxH) >= (cp.Object.Y-cp.HitBoxH) || atk.FixedKby {
		cp.Kby = -kby
	} else {
		cp.Kby = kby
	}

	if isStun {
		cp.Kby = se.StunFloat
		cp.Kbx = se.StunFloat
	}

	if isHit {
		cp.Kby = se.HitFloat
		cp.Kbx = se.HitFloat
	}

	// duration
	kbxDur := atk.KnockbackXDuration
	if atk.HasChargeEffect() && atk.UseChargeKbxDuration {
		kbxDur += int(attacker.ChargeValue * atk.ChargeEffect.MultFactorKbxDur)
	}

	kbyDur := atk.KnockbackYDuration
	if atk.HasChargeEffect() && atk.UseChargeKbyDuration {
		kbyDur = int(attacker.ChargeValue * atk.ChargeEffect.MultFactorKbyDur)
	}

	stamp := time.Now()

	cp.KbStampMutex.Lock()
	cp.KbStamp = stamp
	cp.KbStampMutex.Unlock()

	time.AfterFunc((time.Duration(kbxDur))*time.Millisecond, func() {
		cp.KbStampMutex.RLock()
		defer cp.KbStampMutex.RUnlock()

		if cp.KbStamp == stamp {
			cp.Kbx = 0
		}

	})

	time.AfterFunc((time.Duration(kbyDur))*time.Millisecond, func() {
		cp.KbStampMutex.RLock()
		defer cp.KbStampMutex.RUnlock()

		if cp.KbStamp == stamp {
			cp.Kby = 0
		}
	})
}

func (cp *Player) applyCC() {
	if cp.IsKnockedBackX() {
		cp.SpeedX = cp.Kbx
	}

	if cp.IsKnockedBackY() {
		cp.SpeedY = cp.Kby
	}
}

// simple death for now(this will just cause the player to respawn with a new PID/Full health)
func (cp *Player) death() {
	cp.Dying = true
}
