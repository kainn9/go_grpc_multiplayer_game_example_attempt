package roles

import (
	se "github.com/kainn9/grpc_game/server/statusEffects"
)

var (
	Knight *Role = InitKnight()
)

func KnightAttacks() map[AtKey]*AttackData {
	atks := make(map[AtKey]*AttackData)
	atks = knightPrimaryAtk(atks)
	atks = knightSecondaryAtk(atks)
	atks = knightTertAtk(atks)
	atks = knightQuaternaryAtk(atks)

	return atks
}

func InitKnight() *Role {
	dm := &DefenseMovement{
		Speed:       8,
		Displacment: 100,
	}

	d := &Defense{
		Delay:           0,
		Cooldown:        0,
		DefenseType:     DefenseDodge,
		DefenseMovement: dm,
	}

	r := &Role{
		Attacks: KnightAttacks(),
		HitBoxW: 16,
		HitBoxH: 44,
		Defense: d,
		Health:  280,
		Phys: &RolePhysStruct{
			DefaultFriction: 0.5,
			DefaultMaxSpeed: 4.0,
			DefaultJumpSpd:  12.0,
			DefaultGravity:  0.75,
		},
	}

	r.Phys.DefaultAccel = 0.5 + r.Phys.DefaultFriction

	return r
}

func knightPrimaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {
	atks[PrimaryAttackKey] = &AttackData{
		Name: PrimaryAttackKey,
		Consequence: &Consequence{
			Damage:             48,
			KnockbackX:         16,
			KnockbackY:         2,
			KnockbackXDuration: 250,
			KnockbackYDuration: 50,
		},
		Type: PrimaryAttackKey,
	}

	PrimaryAtkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 4),
		MovmentInc: 16.666 * 5,
	}

	path := PrimaryAtkSeq.HBoxPath

	// frame 0
	path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 1
	path.appendHboxAgg(40, 15, 5, 10, 1)
	path.appendHboxAgg(40, 15, 5, 10, 2)
	path.appendHboxAgg(40, 15, 5, 10, 3)
	atks[PrimaryAttackKey].HitBoxSequence = PrimaryAtkSeq

	return atks
}

func knightSecondaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {

	atks[SecondaryAttackKey] = &AttackData{
		Name: SecondaryAttackKey,
		Type: SecondaryAttackKey,
		Consequence: &Consequence{
			Damage:             44,
			KnockbackX:         se.HitFloat,
			KnockbackY:         se.HitFloat,
			KnockbackXDuration: se.HitDuration,
			KnockbackYDuration: se.HitDuration,
		},
	}

	atkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 5),
		MovmentInc: 16.666 * 5,
	}

	path := atkSeq.HBoxPath

	// frame 0
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 1
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 1)

	// frame 2
	path = path.appendHboxAgg(6, 0, 8, 8, 2)
	path = path.appendHboxAgg(14, 0, 8, 8, 2)
	path = path.appendHboxAgg(22, 4, 8, 8, 2)
	path = path.appendHboxAgg(26, 4, 8, 8, 2)
	path = path.appendHboxAgg(28, 8, 8, 8, 2)
	path = path.appendHboxAgg(34, 8, 8, 8, 2)
	path = path.appendHboxAgg(36, 16, 8, 8, 2)
	path = path.appendHboxAgg(34, 24, 8, 8, 2)
	path = path.appendHboxAgg(32, 32, 8, 8, 2)
	path = path.appendHboxAgg(28, 36, 8, 8, 2)
	path = path.appendHboxAgg(22, 36, 8, 8, 2)
	path = path.appendHboxAgg(14, 41, 8, 8, 2)
	path = path.appendHboxAgg(10, 41, 8, 8, 2)

	// frame 3
	path = path.appendHboxAgg(38, 24, 8, 8, 3)
	path = path.appendHboxAgg(36, 28, 8, 8, 3)
	path = path.appendHboxAgg(32, 32, 8, 8, 3)
	path = path.appendHboxAgg(28, 36, 8, 8, 3)
	path = path.appendHboxAgg(22, 36, 8, 8, 3)
	path = path.appendHboxAgg(14, 41, 8, 8, 3)
	path = path.appendHboxAgg(10, 41, 8, 8, 3)
	path = path.appendHboxAgg(6, 41, 8, 8, 3)
	path = path.appendHboxAgg(2, 41, 8, 8, 3)
	path = path.appendHboxAgg(-2, 41, 8, 8, 3)

	// frame 4
	path.appendHboxAgg(noBox, noBox, 8, 8, 4)

	atks[SecondaryAttackKey].HitBoxSequence = atkSeq
	return atks
}

func knightTertAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {
	windup := &Windup{
		Duration: 0,
		ChargeEffect: &ChargeEffect{
			MultFactorMvDist:   200,
			MultFactorMvSpeed:  1.5,
			MultFactorDmg:      15,
			MultFactorKbxSpeed: 1.5,
			MultFactorKbxDur:   333,
			MultFactorKbyDur:   333,
			TimeLimit:          10,
		},
	}

	movement := &Movement{
		Distance:       480,
		SpeedX:         10,
		UseChargeDist:  true,
		UseChargeSpeed: true,
	}

	atks[TertAttackKey] = &AttackData{
		Name:     TertAttackKey,
		Type:     TertAttackKey,
		Windup:   windup,
		Movement: movement,
		Consequence: &Consequence{
			Damage:             15,
			KnockbackX:         6,
			KnockbackY:         3,
			KnockbackXDuration: 500,
			KnockbackYDuration: 80,

			UseChargeKbxDuration: true,
			UseChargeDmg:         true,
		},
	}

	atkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 12),
		MovmentInc: 16.666 * 5,
	}

	path := atkSeq.HBoxPath

	// frame 0
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 1
	path = path.appendHboxAgg(-10, 0, 8, 8, 1)
	path = path.appendHboxAgg(-2, 0, 8, 8, 1)
	path = path.appendHboxAgg(6, 0, 8, 8, 1)
	path = path.appendHboxAgg(14, 0, 8, 8, 1)
	path = path.appendHboxAgg(22, 4, 8, 8, 1)
	path = path.appendHboxAgg(26, 4, 8, 8, 1)
	path = path.appendHboxAgg(28, 8, 8, 8, 1)

	// frame 2
	path = path.appendHboxAgg(28, 4, 8, 8, 2)
	path = path.appendHboxAgg(32, 8, 8, 8, 2)
	path = path.appendHboxAgg(36, 12, 8, 8, 2)
	path = path.appendHboxAgg(39, 16, 8, 8, 2)
	path = path.appendHboxAgg(40, 24, 8, 8, 2)
	path = path.appendHboxAgg(36, 32, 8, 8, 2)
	path = path.appendHboxAgg(30, 36, 8, 8, 2)
	path = path.appendHboxAgg(24, 38, 8, 8, 2)
	path = path.appendHboxAgg(16, 38, 8, 8, 2)
	path = path.appendHboxAgg(8, 38, 8, 8, 2)

	// frame same as 2
	path[3] = path[2]

	// frame 4 same as 2 but slightly to right
	path = path.appendHboxAgg(36, 4, 8, 8, 4)
	path = path.appendHboxAgg(44, 8, 8, 8, 4)
	path = path.appendHboxAgg(48, 12, 8, 8, 4)
	path = path.appendHboxAgg(51, 16, 8, 8, 4)
	path = path.appendHboxAgg(50, 24, 8, 8, 4)
	path = path.appendHboxAgg(45, 32, 8, 8, 4)
	path = path.appendHboxAgg(40, 36, 8, 8, 4)
	path = path.appendHboxAgg(34, 38, 8, 8, 4)

	path[5] = path[3]
	path[6] = path[4]

	path[7] = path[5]
	path[8] = path[6]
	path[9] = path[7]
	path[10] = path[8]
	path[11] = path[9]

	atks[TertAttackKey].HitBoxSequence = atkSeq
	return atks
}

func knightQuaternaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {

	windup := &Windup{
		Duration: 0,
		ChargeEffect: &ChargeEffect{
			MultFactorDmg: 101,
			TimeLimit:     1,
		},
	}

	atks[QuaternaryAttackKey] = &AttackData{
		Name:   QuaternaryAttackKey,
		Type:   QuaternaryAttackKey,
		Windup: windup,
		Consequence: &Consequence{
			Damage:             34,
			KnockbackX:         se.StunFloat,
			KnockbackY:         se.StunFloat,
			KnockbackXDuration: 1300,
			KnockbackYDuration: 1300,

			UseChargeDmg: true,
		},
	}

	atkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 5),
		MovmentInc: 16.666 * 5,
	}

	path := atkSeq.HBoxPath

	// frame 0
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 1
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 1)

	// frame 2
	path = path.appendHboxAgg(6, 0, 8, 8, 2)
	path = path.appendHboxAgg(14, 0, 8, 8, 2)
	path = path.appendHboxAgg(22, 4, 8, 8, 2)
	path = path.appendHboxAgg(26, 4, 8, 8, 2)
	path = path.appendHboxAgg(28, 8, 8, 8, 2)
	path = path.appendHboxAgg(34, 8, 8, 8, 2)
	path = path.appendHboxAgg(36, 16, 8, 8, 2)
	path = path.appendHboxAgg(34, 24, 8, 8, 2)
	path = path.appendHboxAgg(32, 32, 8, 8, 2)
	path = path.appendHboxAgg(28, 36, 8, 8, 2)
	path = path.appendHboxAgg(22, 36, 8, 8, 2)
	path = path.appendHboxAgg(14, 41, 8, 8, 2)
	path = path.appendHboxAgg(10, 41, 8, 8, 2)

	// frame 3
	path = path.appendHboxAgg(38, 24, 8, 8, 3)
	path = path.appendHboxAgg(36, 28, 8, 8, 3)
	path = path.appendHboxAgg(32, 32, 8, 8, 3)
	path = path.appendHboxAgg(28, 36, 8, 8, 3)
	path = path.appendHboxAgg(22, 36, 8, 8, 3)
	path = path.appendHboxAgg(14, 41, 8, 8, 3)
	path = path.appendHboxAgg(10, 41, 8, 8, 3)
	path = path.appendHboxAgg(6, 41, 8, 8, 3)
	path = path.appendHboxAgg(2, 41, 8, 8, 3)
	path = path.appendHboxAgg(-2, 41, 8, 8, 3)

	// frame 4
	path.appendHboxAgg(noBox, noBox, 8, 8, 4)

	atks[QuaternaryAttackKey].HitBoxSequence = atkSeq
	return atks
}
