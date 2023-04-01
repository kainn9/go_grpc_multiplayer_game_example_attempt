package roles

var (
	Monk *Role = InitMonk()
)

func MonkAttacks() map[AtKey]*AttackData {
	atks := make(map[AtKey]*AttackData)
	atks = monkPrimaryAtk(atks)
	atks = monkSecondaryAtk(atks)

	// temp
	atks[TertAttackKey] = atks[SecondaryAttackKey]
	atks[QuaternaryAttackKey] = atks[SecondaryAttackKey]

	return atks
}

func InitMonk() *Role {

	d := &Defense{
		Delay:           0,
		Cooldown:        300,
		DefenseDuration: 1000,
		DefenseType:     DefenseBlock,
	}

	r := &Role{
		RoleType: MonkType,
		Attacks:  MonkAttacks(),
		HitBoxW:  12,
		HitBoxH:  35,
		Defense:  d,
		Health:   380,

		Phys: &RolePhysStruct{
			DefaultFriction: 0.5,
			DefaultMaxSpeed: 3.5,
			DefaultJumpSpd:  13.0,
			DefaultGravity:  0.6,
		},
	}

	r.Phys.DefaultAccel = 0.5 + r.Phys.DefaultFriction

	return r
}

func monkPrimaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {
	atks[PrimaryAttackKey] = &AttackData{
		Name: PrimaryAttackKey,
		Consequence: &Consequence{
			Damage:             15,
			KnockbackX:         0.5,
			KnockbackY:         0.5,
			KnockbackXDuration: 1500,
			KnockbackYDuration: 1500,
		},
		Type: PrimaryAttackKey,
	}

	PrimaryAtkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 13),
		MovmentInc: 16.666 * 5,
	}

	path := PrimaryAtkSeq.HBoxPath
	// frame 0
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 1
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 2
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 3
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 4
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 5
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 6
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 7
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 8
	path = path.appendHboxAgg(45, 30, 10, 20, 8)
	path = path.appendHboxAgg(50, 25, 10, 10, 8)

	// frame 9
	path = path.appendHboxAgg(45, 30, 10, 60, 9)
	path = path.appendHboxAgg(60, 25, 10, 40, 9)
	path = path.appendHboxAgg(60, 15, 10, 35, 9)
	path = path.appendHboxAgg(77, 0, 25, 10, 9)

	// frame 10
	path = path.appendHboxAgg(45, 30, 10, 60, 10)
	path = path.appendHboxAgg(60, 25, 10, 40, 10)
	path = path.appendHboxAgg(60, 15, 10, 35, 10)
	path = path.appendHboxAgg(77, 0, 25, 10, 10)

	// frame 11
	path = path.appendHboxAgg(45, 30, 10, 60, 11)
	path = path.appendHboxAgg(60, 25, 10, 40, 11)
	path = path.appendHboxAgg(60, 15, 10, 35, 11)
	path = path.appendHboxAgg(77, 0, 25, 10, 11)

	// frame 12
	path = path.appendHboxAgg(62, 30, 10, 30, 12)
	path = path.appendHboxAgg(72, 20, 10, 15, 12)
	path.appendHboxAgg(76, 12, 8, 5, 12)

	atks[PrimaryAttackKey].HitBoxSequence = PrimaryAtkSeq

	return atks
}

func monkSecondaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {

	atks[SecondaryAttackKey] = &AttackData{
		Name: SecondaryAttackKey,
		Type: SecondaryAttackKey,
		Consequence: &Consequence{
			Damage:             15,
			KnockbackX:         0,
			KnockbackY:         -0.1,
			KnockbackXDuration: 0,
			KnockbackYDuration: 1500,
		},
	}

	atkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 8),
		MovmentInc: 16.666 * 5,
	}

	path := atkSeq.HBoxPath

	// frame 0
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 1
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 1)

	// frame 2
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 2)

	// frame 3
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 3)

	// frame 4
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 4)

	// frame 5
	path = path.appendHboxAgg(33, -5, 20, 25, 5)

	// frame 6
	path = path.appendHboxAgg(35, -4, 25, 27, 6)

	// frame 7
	path.appendHboxAgg(38, 12, 20, 20, 7)

	atks[SecondaryAttackKey].HitBoxSequence = atkSeq
	return atks
}
