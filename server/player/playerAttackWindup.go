package player

import (
	"math"
	"time"

	r "github.com/kainn9/grpc_game/server/roles"
)

func (cp *Player) windupPhase(atk *r.AttackData, atKey r.AtKey) {

	if atk.Windup.HasChargeEffect() {
		cp.chargeWindupPhase(cp.Attacks[atKey])
		return
	}

	if atk.Windup != nil {
		cp.Windup = atKey

		delay := atk.Windup.Duration
		time.AfterFunc(time.Duration(delay)*time.Millisecond, func() {
			resolveWindup(cp, atk)
		})

		return
	}

	resolveWindup(cp, atk)
}

func resolveWindup(cp *Player, atk *r.AttackData) {
	if !cp.IsWindingUp() && atk.Windup != nil {
		return
	}

	cp.Windup = ""
	cp.CurrAttack = atk
	cp.AttackMovement = string(atk.Type)
}

func (cp *Player) chargeWindupPhase(atk *r.AttackData) {

	if cp.ChargeStart.IsZero() {
		cp.ChargeStart = time.Now()
		cp.Windup = atk.Type
	}

	checkWindupCharge(cp, atk)
}

func checkWindupCharge(cp *Player, atk *r.AttackData) {

	if !cp.IsWindingUp() && atk.Windup != nil {
		return
	}

	if time.Since(cp.ChargeStart).Seconds() > atk.Windup.TimeLimit {
		resolveChargeWindup(cp, atk)
		return
	}

	if cp.PrevEvent.Input != string(atk.Type) {
		resolveChargeWindup(cp, atk)
		return
	}

	time.AfterFunc(1*time.Second, func() {
		checkWindupCharge(cp, atk)
	})
}

func resolveChargeWindup(cp *Player, atk *r.AttackData) {

	chargeValue := math.Round(time.Since(cp.ChargeStart).Seconds())
	cp.ChargeValue = chargeValue

	resolveWindup(cp, atk)
}

func (cp *Player) interruptWindup() {
	cp.ChargeValue = 0
	cp.ChargeStart = time.Time{}
	cp.Windup = ""
	cp.AttackMovement = ""
	cp.CurrAttack = nil
}
