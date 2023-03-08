package main

import (
	"math"
	"time"

	r "github.com/kainn9/grpc_game/server/roles"
)

func (cp *player) windupPhase(atk *r.Attack, atKey r.AtKey) {

	if atk.Windup.HasChargeEffect() {
		cp.chargeWindupPhase(cp.Attacks[atKey])
		return
	}

	if atk.Windup != nil {
		cp.windup = atKey

		delay := atk.Windup.Duration
		time.AfterFunc(time.Duration(delay)*time.Millisecond, func() {
			resolveWindup(cp, atk)
		})

		return
	}

	resolveWindup(cp, atk)
}

func resolveWindup(cp *player, atk *r.Attack) {
	if !cp.isWindingUp() && atk.Windup != nil {
		return
	}

	cp.windup = ""
	cp.currAttack = atk
	cp.attackMovement = string(atk.Type)
}

func (cp *player) chargeWindupPhase(atk *r.Attack) {

	if cp.chargeStart.IsZero() {
		cp.chargeStart = time.Now()
		cp.windup = atk.Type
	}

	checkWindupCharge(cp, atk)
}

func checkWindupCharge(cp *player, atk *r.Attack) {
	if !cp.isWindingUp() && atk.Windup != nil {
		return
	}

	if time.Since(cp.chargeStart).Seconds() > atk.Windup.TimeLimit {
		resolveChargeWindup(cp, atk)
		return
	}

	if cp.prevEvent.Input != string(atk.Type) {
		resolveChargeWindup(cp, atk)
		return
	}

	time.AfterFunc(1*time.Second, func() {
		checkWindupCharge(cp, atk)
	})
}

func resolveChargeWindup(cp *player, atk *r.Attack) {

	chargeValue := math.Round(time.Since(cp.chargeStart).Seconds())
	cp.chargeValue = chargeValue

	resolveWindup(cp, atk)
}

func (cp *player) interruptWindup() {
	cp.chargeValue = 0
	cp.chargeStart = time.Time{}
	cp.windup = ""
	cp.attackMovement = ""
	cp.currAttack = nil
}
