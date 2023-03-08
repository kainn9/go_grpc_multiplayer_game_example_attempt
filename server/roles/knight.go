package roles

var (
	Knight *Role = InitKnight()
)

func KnightAttacks() map[AtKey]*Attack {
	atks := make(map[AtKey]*Attack)
	atks = primaryAtk(atks)
	atks = secondaryAtk(atks)
	atks = tertAtk(atks)

	return atks
}

func InitKnight() *Role {
	d := &Defense{
		Speed:       16,
		Delay:       0,
		Displacment: 80,
		Cooldown:    1000,
	}

	r := &Role{
		RoleType: KnightType,
		Attacks:  KnightAttacks(),
		HitBoxW:  16,
		HitBoxH:  44,
		Defense:  d,
	}

	return r
}

func primaryAtk(atks map[AtKey]*Attack) map[AtKey]*Attack {
	atks[PrimaryAttackKey] = &Attack{
		Name:     PrimaryAttackKey,
		Cooldown: 5,
		Consequence: &Consequence{
			Damage:             25,
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

func secondaryAtk(atks map[AtKey]*Attack) map[AtKey]*Attack {
	windup := &Windup{
		Duration: 750,
	}

	movement := &Movement{
		Distance: 480,
		SpeedX:   10,
	}

	atks[SecondaryAttackKey] = &Attack{
		Name:     SecondaryAttackKey,
		Cooldown: 5,
		Type:     SecondaryAttackKey,
		Windup:   windup,
		Movement: movement,
		Consequence: &Consequence{
			Damage:             10,
			KnockbackX:         6,
			KnockbackY:         6,
			KnockbackXDuration: 500,
			KnockbackYDuration: 150,
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

	atks[SecondaryAttackKey].HitBoxSequence = atkSeq
	return atks
}

// TODO: make this more dry
// maybe doing something like tertAtk = secondaryAtk and only changing the charge effect
// but I'm too lazy to do that rn
// also make sure its a clone and not a reference/pointer

// only difference from secondary is the charge effect
// cause I'm too lazy to setup new attack/animation rn
// but ironically this proves its not very dry...
// at least for attacks that are almost identical
func tertAtk(atks map[AtKey]*Attack) map[AtKey]*Attack {
	windup := &Windup{
		Duration: 750,
		ChargeEffect: &ChargeEffect{
			MultFactorMvDist:   200,
			MultFactorMvSpeed:  1.5,
			MultFactorDmg:      3,
			MultFactorKbxSpeed: 1.5,
			MultFactorKbxDur:   333,
			TimeLimit:          10,
		},
	}

	movement := &Movement{
		Distance:       480,
		SpeedX:         10,
		UseChargeDist:  true,
		UseChargeSpeed: true,
	}

	atks[TertAttackKey] = &Attack{
		Name:     TertAttackKey,
		Cooldown: 5,
		Type:     TertAttackKey,
		Windup:   windup,
		Movement: movement,
		Consequence: &Consequence{
			Damage:             10,
			KnockbackX:         6,
			KnockbackY:         6,
			KnockbackXDuration: 500,
			KnockbackYDuration: 150,

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
