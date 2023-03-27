package roles

var (
	Werewolf *Role = InitWerewolf()
)

func WerewolfAttacks() map[AtKey]*AttackData {
	atks := make(map[AtKey]*AttackData)
	atks = werewolfPrimaryAtk(atks)
	atks = werewolfSecondaryAtk(atks)
	atks = werewolfTertAtk(atks)
	atks[QuaternaryAttackKey] = atks[TertAttackKey]

	return atks
}

func InitWerewolf() *Role {

	r := &Role{
		RoleType: WerewolfType,
		Attacks:  WerewolfAttacks(),
		HitBoxW:  30,
		HitBoxH:  50,
		Health:   125,

		Phys: &RolePhysStruct{
			DefaultFriction: 0.5,
			DefaultMaxSpeed: 6.6,
			DefaultJumpSpd:  15.0,
			DefaultGravity:  0.6,
		},
	}

	r.Phys.DefaultAccel = 0.5 + r.Phys.DefaultFriction

	return r
}

func werewolfPrimaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {
	atks[PrimaryAttackKey] = &AttackData{
		Name: PrimaryAttackKey,
		Consequence: &Consequence{
			Damage:             65,
			KnockbackX:         6,
			KnockbackY:         1,
			KnockbackXDuration: 1600,
			KnockbackYDuration: 500,
		},
		Type: PrimaryAttackKey,
	}

	PrimaryAtkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 8),
		MovmentInc: 16.666 * 5,
	}

	path := PrimaryAtkSeq.HBoxPath

	// frame 0-1 no hitboxes

	// frame 2
	path = path.appendHboxAgg(20, -8, 40, 56, 2)

	// frame 3
	path = path.appendHboxAgg(56, -8, 18, 20, 3)
	path = path.appendHboxAgg(46, 20, 10, 20, 3)


	// frame 4
	path = path.appendHboxAgg(56, -8, 18, 20, 4)
	path = path.appendHboxAgg(46, 20, 10, 20, 4)

	// frame 5
	path = path.appendHboxAgg(56, -8, 44, 26, 5)
	path = path.appendHboxAgg(46, 20, 20, 20, 5)

	// frame 6
	path = path.appendHboxAgg(56, 16, 20, 26, 6)
	path.appendHboxAgg(39, 20, 20, 20, 6)


	// frame 7 no hitboxes

	atks[PrimaryAttackKey].HitBoxSequence = PrimaryAtkSeq

	return atks
}

func werewolfSecondaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {

	atks[SecondaryAttackKey] = &AttackData{
		Name: SecondaryAttackKey,
		Type: SecondaryAttackKey,
		Consequence: &Consequence{
			Damage:             65,
			KnockbackX:         6,
			KnockbackY:         1,
			KnockbackXDuration: 1600,
			KnockbackYDuration: 500,
		},
	}

	atkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 10),
		MovmentInc: 16.666 * 5,
	}

	path := atkSeq.HBoxPath

	// frame 0 - 3 no hitboxes

	// frame 4
	path = path.appendHboxAgg(10, -26, 20, 23, 4)
	path = path.appendHboxAgg(76, -26, 35, 23, 4)

	// frame 5
	path = path.appendHboxAgg(14, -26, 66, 80, 5)

	// frame 6
	path = path.appendHboxAgg(18, 20, 26, 80, 6)

	// frame 7
	path.appendHboxAgg(18, 20, 26, 80, 7)


	// frame 8 - 10 no hitboxes

	atks[SecondaryAttackKey].HitBoxSequence = atkSeq
	return atks
}


func werewolfTertAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {
	windup := &Windup{
		Duration: 400,
	}

	movement := &Movement{
		Distance:       800,
		SpeedX:         16,
	}

	atks[TertAttackKey] = &AttackData{
		Name:     TertAttackKey,
		Type:     TertAttackKey,
		Windup:   windup,
		Movement: movement,
		Consequence: &Consequence{
			Damage:             50,
			KnockbackX:         6,
			KnockbackY:         1,
			KnockbackXDuration: 1600,
			KnockbackYDuration: 500,
		},
	}

	atkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 16),
		MovmentInc: 16.666 * 5,
	}

	path := atkSeq.HBoxPath

	// frame 0
	path = path.appendHboxAgg(60, 0, 20, 20, 0)
	path = path.appendHboxAgg(55, 10, 20, 20, 0)
	path = path.appendHboxAgg(45, 20, 20, 20, 0)

	// frame 1
	path = path.appendHboxAgg(45, 0, 20, 23, 1)

	// frame 2
	path = path.appendHboxAgg(25, 10, 20, 43, 2)
	path = path.appendHboxAgg(25, 20, 20, 28, 2)

	// frame 3
	path = path.appendHboxAgg(35, 0, 40, 42, 3)

	// frame 4
	path = path.appendHboxAgg(30, 10, 30, 32, 4)

	// frame 5
	path = path.appendHboxAgg(35, 0, 32, 35, 5)


	// frame 6
	path = path.appendHboxAgg(35, 0, 32, 38, 6)

	// frame 7
	path = path.appendHboxAgg(35, 0, 44, 45, 7)

	// frame 8
	path = path.appendHboxAgg(55, 0, 15, 25, 8)
	path = path.appendHboxAgg(50, 15, 15, 25, 8)
	path = path.appendHboxAgg(45, 20, 15, 25, 8)

	// frame 9
	path = path.appendHboxAgg(45, 0, 20, 24, 9)

	// frame 10
	path = path.appendHboxAgg(25, 10, 20, 43, 10)
	path = path.appendHboxAgg(25, 20, 20, 28, 10)


	// frame 11
	path = path.appendHboxAgg(45, 0, 40, 34, 12)

	// frame 12 no box

	// frame 13
	path.appendHboxAgg(10, 0, 40, 63, 13)

	// frame 14 - 16 no box

	atks[TertAttackKey].HitBoxSequence = atkSeq
	return atks
}