package roles

import se "github.com/kainn9/grpc_game/server/statusEffects"

var (
	Demon *Role = InitDemon()
)

func DemonAttacks() map[AtKey]*AttackData {
	atks := make(map[AtKey]*AttackData)
	atks = demonPrimaryAtk(atks)
	atks = demonSecondaryAtk(atks)

	return atks
}

func InitDemon() *Role {

	r := &Role{
		Attacks: DemonAttacks(),
		HitBoxW: 50,
		HitBoxH: 80,
		Health:  1300,

		Phys: &RolePhysStruct{
			DefaultFriction: 0.5,
			DefaultMaxSpeed: 2,
			DefaultJumpSpd:  9,
			DefaultGravity:  0.6,
		},
	}

	r.Phys.DefaultAccel = 0.5 + r.Phys.DefaultFriction

	return r
}

func demonPrimaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {
	atks[PrimaryAttackKey] = &AttackData{
		Name: PrimaryAttackKey,
		Consequence: &Consequence{
			Damage:             300,
			KnockbackX:         se.StunFloat,
			KnockbackY:         se.StunFloat,
			KnockbackXDuration: 1000,
			KnockbackYDuration: 1000,
		},
		Type: PrimaryAttackKey,
	}

	PrimaryAtkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 22),
		MovmentInc: 16.666 * 5,
	}

	path := PrimaryAtkSeq.HBoxPath

	// frame 12
	path = path.appendHboxAgg(105, 70, 15, 15, 12)
	path = path.appendHboxAgg(115, 70, 15, 15, 12)

	// frame 13
	path = path.appendHboxAgg(102, 67, 15, 15, 13)
	path = path.appendHboxAgg(115, 70, 15, 15, 13)

	// frame 14
	path = path.appendHboxAgg(102, 67, 15, 15, 14)
	path = path.appendHboxAgg(115, 70, 15, 15, 14)

	// frame 15
	path = path.appendHboxAgg(102, 67, 15, 15, 15)
	path = path.appendHboxAgg(115, 70, 15, 15, 15)

	// frame 16
	path = path.appendHboxAgg(102, 67, 15, 15, 16)
	path = path.appendHboxAgg(115, 70, 15, 15, 16)

	// frame 17
	path = path.appendHboxAgg(102, 67, 15, 15, 17)
	path.appendHboxAgg(115, 70, 15, 15, 17)

	// frame 18+ no hitbox

	atks[PrimaryAttackKey].HitBoxSequence = PrimaryAtkSeq

	return atks
}

func demonSecondaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {

	atks[SecondaryAttackKey] = &AttackData{
		Name: SecondaryAttackKey,
		Type: SecondaryAttackKey,
		Consequence: &Consequence{
			Damage:             100,
			KnockbackX:         se.HitFloat,
			KnockbackY:         se.HitFloat,
			KnockbackXDuration: se.HitDuration,
			KnockbackYDuration: se.HitDuration,
		},
	}

	atkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 21),
		MovmentInc: 16.666 * 5,
	}

	path := atkSeq.HBoxPath

	// frame 0 - 6 no hitboxes

	// frame 7
	path = path.appendHboxAgg(50, 0, 15, 15, 7)
	path = path.appendHboxAgg(57, -7, 15, 15, 7)
	path = path.appendHboxAgg(62, -12, 15, 15, 7)
	path = path.appendHboxAgg(72, -17, 15, 15, 7)
	path = path.appendHboxAgg(77, -22, 15, 15, 7)
	path = path.appendHboxAgg(82, -27, 15, 15, 7)
	path = path.appendHboxAgg(87, -32, 15, 15, 7)
	path = path.appendHboxAgg(92, -37, 15, 15, 7)

	// frame 8
	path = path.appendHboxAgg(50, 0, 15, 15, 8)
	path = path.appendHboxAgg(57, -2, 15, 15, 8)
	path = path.appendHboxAgg(62, -5, 15, 15, 8)
	path = path.appendHboxAgg(72, -8, 15, 15, 8)
	path = path.appendHboxAgg(77, -12, 15, 15, 8)
	path = path.appendHboxAgg(82, -15, 15, 15, 8)
	path = path.appendHboxAgg(87, -18, 15, 15, 8)
	path = path.appendHboxAgg(92, -21, 15, 15, 8)
	path = path.appendHboxAgg(97, -25, 15, 15, 8)
	path = path.appendHboxAgg(102, -28, 15, 15, 8)
	path = path.appendHboxAgg(107, -31, 15, 15, 8)

	// frame 9
	path = path.appendHboxAgg(50, 0, 15, 15, 9)
	path = path.appendHboxAgg(57, 3, 15, 15, 9)
	path = path.appendHboxAgg(72, 3, 8, 30, 9)
	path = path.appendHboxAgg(76, -2, 15, 15, 9)
	path = path.appendHboxAgg(82, -5, 15, 15, 9)
	path = path.appendHboxAgg(87, -8, 15, 15, 9)
	path = path.appendHboxAgg(92, -12, 15, 15, 9)
	path = path.appendHboxAgg(97, -15, 15, 15, 9)
	path = path.appendHboxAgg(102, -18, 15, 15, 9)
	path = path.appendHboxAgg(107, -21, 15, 15, 9)
	path = path.appendHboxAgg(112, -25, 15, 15, 9)
	path = path.appendHboxAgg(117, -28, 15, 15, 9)
	path = path.appendHboxAgg(122, -31, 15, 15, 9)

	// frame 10
	path = path.appendHboxAgg(50, 13, 8, 40, 10)
	path = path.appendHboxAgg(90, -5, 22, 30, 10)
	path = path.appendHboxAgg(120, -25, 30, 30, 10)

	// frame 11
	path = path.appendHboxAgg(50, 18, 8, 40, 11)
	path = path.appendHboxAgg(90, 7, 22, 30, 11)
	path = path.appendHboxAgg(120, -10, 30, 30, 11)

	// frame 12
	path = path.appendHboxAgg(50, 26, 8, 40, 12)
	path = path.appendHboxAgg(90, 15, 22, 30, 12)
	path = path.appendHboxAgg(120, 2, 30, 30, 12)

	// frame 13
	path = path.appendHboxAgg(50, 34, 8, 40, 13)
	path = path.appendHboxAgg(90, 26, 22, 30, 13)
	path = path.appendHboxAgg(120, 18, 30, 30, 13)

	// frame 14
	path = path.appendHboxAgg(50, 42, 8, 40, 14)
	path = path.appendHboxAgg(90, 38, 22, 30, 14)
	path = path.appendHboxAgg(120, 30, 30, 30, 14)

	// frame 15
	path = path.appendHboxAgg(50, 46, 8, 40, 15)
	path = path.appendHboxAgg(90, 53, 22, 30, 15)
	path = path.appendHboxAgg(120, 45, 30, 30, 15)

	// frame 16
	path = path.appendHboxAgg(50, 46, 8, 40, 16)
	path = path.appendHboxAgg(90, 53, 22, 30, 16)
	path.appendHboxAgg(120, 45, 30, 30, 16)

	// no hitboxes for frame 17+

	atks[SecondaryAttackKey].HitBoxSequence = atkSeq
	return atks
}
