package roles

var (
	Mage *Role = InitMage()
)

func MageAttacks() map[AtKey]*AttackData {
	atks := make(map[AtKey]*AttackData)
	atks = magePrimaryAtk(atks)
	atks = mageSecondaryAtk(atks)
	atks = mageTertAtk(atks)
	atks[QuaternaryAttackKey] = atks[TertAttackKey]

	return atks
}

func InitMage() *Role {

	dm := &DefenseMovement{
		Speed:       16,
		Displacment: 500,
	}

	d := &Defense{
		Delay:           0,
		Cooldown:        100,
		DefenseType:     DefenseDodge,
		DefenseMovement: dm,
	}

	r := &Role{
		RoleType: MageType,
		Attacks:  MageAttacks(),
		HitBoxW:  16,
		HitBoxH:  44,
		Defense:  d,
		Health:   85,
		Phys: &RolePhysStruct{
			DefaultFriction: 0.5,
			DefaultMaxSpeed: 5.0,
			DefaultJumpSpd:  14.0,
			DefaultGravity:  0.6,
		},
	}

	r.Phys.DefaultAccel = 0.5 + r.Phys.DefaultFriction

	return r
}

func magePrimaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {
	atks[PrimaryAttackKey] = &AttackData{
		Name: PrimaryAttackKey,
		Consequence: &Consequence{
			Damage:             70,
			KnockbackX:         0.1,
			KnockbackY:         0.1,
			KnockbackXDuration: 1250,
			KnockbackYDuration: 1250,
		},
		Type: PrimaryAttackKey,
	}

	PrimaryAtkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 21),
		MovmentInc: 16.666 * 5,
	}

	path := PrimaryAtkSeq.HBoxPath

	// frame 0 - 10 no hitboxess

	// frame 11
	path = path.appendHboxAgg(-33, -30, 47, 8, 11)
	path = path.appendHboxAgg(-63, -30, 47, 8, 11)
	path = path.appendHboxAgg(43, -30, 47, 8, 11)
	path = path.appendHboxAgg(73, -30, 47, 8, 11)

	// frame 12
	path = path.appendHboxAgg(-33, -30, 47, 8, 12)
	path = path.appendHboxAgg(-63, -30, 47, 8, 12)
	path = path.appendHboxAgg(43, -30, 47, 8, 12)
	path = path.appendHboxAgg(73, -30, 47, 8, 12)

	// frame 13
	path = path.appendHboxAgg(-33, -24, 47, 8, 13)
	path = path.appendHboxAgg(-63, -24, 47, 8, 13)
	path = path.appendHboxAgg(43, -24, 47, 8, 13)
	path = path.appendHboxAgg(73, -24, 47, 8, 13)

	// frame 14
	path = path.appendHboxAgg(-33, -16, 47, 8, 14)
	path = path.appendHboxAgg(-63, -16, 47, 8, 14)
	path = path.appendHboxAgg(43, -16, 47, 8, 14)
	path = path.appendHboxAgg(73, -16, 47, 8, 14)

	// frame 15
	path = path.appendHboxAgg(-33, -6, 47, 8, 15)
	path = path.appendHboxAgg(-63, -6, 47, 8, 15)
	path = path.appendHboxAgg(43, -6, 47, 8, 15)
	path = path.appendHboxAgg(73, -6, 47, 8, 15)

	// frame 16
	path = path.appendHboxAgg(-33, 10, 35, 8, 16)
	path = path.appendHboxAgg(-63, 10, 35, 8, 16)
	path = path.appendHboxAgg(43, 10, 35, 8, 16)
	path = path.appendHboxAgg(73, 10, 35, 8, 16)

	// frame 17
	path = path.appendHboxAgg(-33, 10, 35, 8, 17)
	path = path.appendHboxAgg(-63, 10, 35, 8, 17)
	path = path.appendHboxAgg(43, 10, 35, 8, 17)
	path.appendHboxAgg(73, 10, 35, 8, 17)

	atks[PrimaryAttackKey].HitBoxSequence = PrimaryAtkSeq

	return atks
}

func mageSecondaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {

	atks[SecondaryAttackKey] = &AttackData{
		Name: SecondaryAttackKey,
		Type: SecondaryAttackKey,
		Consequence: &Consequence{
			Damage:             80,
			KnockbackX:         4,
			KnockbackY:         2,
			KnockbackXDuration: 650,
			KnockbackYDuration: 140,
		},
	}

	atkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 11),
		MovmentInc: 16.666 * 5,
	}

	path := atkSeq.HBoxPath

	// frame 1 - 3 no hitbox

	// frame 4
	path = path.appendHboxAgg(33, -10, 55, 40, 4)
	path = path.appendHboxAgg(23, 0, 35, 10, 4)
	path = path.appendHboxAgg(73, 0, 35, 10, 4)

	// frame 5
	path = path.appendHboxAgg(33, -10, 55, 40, 5)
	path = path.appendHboxAgg(23, 0, 35, 10, 5)
	path = path.appendHboxAgg(73, 0, 35, 10, 5)

	// frame 6
	path = path.appendHboxAgg(30, -25, 25, 50, 6)
	path = path.appendHboxAgg(15, -10, 55, 80, 6)

	// frame 7
	path = path.appendHboxAgg(33, -10, 55, 50, 7)
	path = path.appendHboxAgg(23, 0, 45, 10, 7)
	path = path.appendHboxAgg(82, 0, 45, 12, 7)

	// frame 8
	path = path.appendHboxAgg(33, -10, 55, 50, 8)
	path = path.appendHboxAgg(23, 0, 45, 10, 8)
	path = path.appendHboxAgg(82, 0, 45, 12, 8)

	// frame 9
	path = path.appendHboxAgg(23, 5, 40, 15, 9)
	path = path.appendHboxAgg(30, -5, 15, 60, 9)
	path.appendHboxAgg(83, 5, 40, 15, 9)

	// frame 10 + no hitbox

	atks[SecondaryAttackKey].HitBoxSequence = atkSeq
	return atks
}

func mageTertAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {

	movement := &Movement{
		Distance: 100,
		SpeedX:   13,
	}

	atks[TertAttackKey] = &AttackData{
		Name: TertAttackKey,
		Type: TertAttackKey,

		Movement: movement,
		Consequence: &Consequence{
			Damage:             1,
			KnockbackX:         0.1,
			KnockbackY:         0.1,
			KnockbackXDuration: 4000,
			KnockbackYDuration: 4000,
		},
	}

	atkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 12),
		MovmentInc: 16.666 * 5,
	}

	path := atkSeq.HBoxPath

	// frame 0 - 2 no hitboxes

	// frame 3
	path = path.appendHboxAgg(-50, 0, 50, 110, 3)

	// frame 4
	path = path.appendHboxAgg(-50, 0, 50, 110, 4)

	// frame 5
	path = path.appendHboxAgg(-50, 0, 50, 110, 5)

	// frame 6
	path = path.appendHboxAgg(-70, -10, 60, 150, 6)

	// frame 7
	path = path.appendHboxAgg(-70, -10, 60, 150, 7)

	// frame 8
	path = path.appendHboxAgg(-40, 10, 40, 90, 8)

	// frame 9
	path = path.appendHboxAgg(-40, 10, 40, 90, 9)

	// frame 10
	path.appendHboxAgg(-30, 30, 20, 80, 10)

	// frame 11 + no box

	atks[TertAttackKey].HitBoxSequence = atkSeq
	return atks
}
