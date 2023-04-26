package roles

import (
	se "github.com/kainn9/grpc_game/server/statusEffects"
)

var (
	BirdDroid *Role = InitBirdDroid()
)

func BirdDroidAttacks() map[AtKey]*AttackData {
	atks := make(map[AtKey]*AttackData)
	atks = birdDroidPrimaryAtk(atks)
	atks = birdDroidSecondaryAtk(atks)

	return atks
}

func InitBirdDroid() *Role {

	r := &Role{
		Attacks: BirdDroidAttacks(),
		HitBoxW: 16,
		HitBoxH: 33,

		Health: 180,
		Phys: &RolePhysStruct{
			DefaultFriction: 0.5,
			DefaultMaxSpeed: 6.0,
			DefaultJumpSpd:  15.0,
			DefaultGravity:  0.75,
		},
	}

	r.Phys.DefaultAccel = 0.5 + r.Phys.DefaultFriction

	return r
}

func birdDroidPrimaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {
	atks[PrimaryAttackKey] = &AttackData{
		Name:            PrimaryAttackKey,
		InvincibleNoBox: true,
		Cooldown:        5000,
		Consequence: &Consequence{
			Damage:             80,
			KnockbackX:         12,
			KnockbackY:         2,
			KnockbackXDuration: 500,
			KnockbackYDuration: 100,
		},
		Type: PrimaryAttackKey,
	}

	atkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 9),
		MovmentInc: 16.666 * 5,
	}

	path := atkSeq.HBoxPath

	// frame 0 - 3 no hitboxes

	// frame 4
	path = path.appendHboxAgg(-3, -40, 70, 20, 4)
	path = path.appendHboxAgg(-50, 0, 30, 120, 4)

	// frame 5
	path = path.appendHboxAgg(-50, 0, 30, 120, 5)

	// frame 6
	path.appendHboxAgg(-50, 0, 30, 120, 6)

	// frame 7 + no boxes

	atks[PrimaryAttackKey].HitBoxSequence = atkSeq

	return atks
}

func birdDroidSecondaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {

	atks[SecondaryAttackKey] = &AttackData{
		Name: SecondaryAttackKey,
		Type: SecondaryAttackKey,
		Consequence: &Consequence{
			Damage:             60,
			KnockbackX:         se.HitFloat,
			KnockbackY:         se.HitFloat,
			KnockbackXDuration: se.HitDuration,
			KnockbackYDuration: se.HitDuration,
		},
	}

	atkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 16),
		MovmentInc: 16.666 * 5,
	}

	path := atkSeq.HBoxPath

	// frame 0 - 7 no hitboxes

	// frame 8
	path = path.appendHboxAgg(43, 15, 15, 140, 8)

	// frame 9
	path = path.appendHboxAgg(43, 15, 15, 140, 9)

	// frame 10
	path = path.appendHboxAgg(43, 15, 15, 140, 10)

	// frame 11
	path.appendHboxAgg(51, 18, 10, 137, 11)

	// frame 12 + no box

	atks[SecondaryAttackKey].HitBoxSequence = atkSeq
	return atks
}
