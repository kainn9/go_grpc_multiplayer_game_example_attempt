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
	atks = birdDroidTertAtk(atks)
	atks[QuaternaryAttackKey] = atks[TertAttackKey]

	return atks
}

func InitBirdDroid() *Role {

	r := &Role{
		Attacks: BirdDroidAttacks(),
		HitBoxW: 16,
		HitBoxH: 33,

		Health: 380,
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
		Name: PrimaryAttackKey,
		Consequence: &Consequence{
			Damage:             38,
			KnockbackX:         16,
			KnockbackY:         2,
			KnockbackXDuration: 250,
			KnockbackYDuration: 50,
		},
		Type: PrimaryAttackKey,
	}

	PrimaryAtkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 9),
		MovmentInc: 16.666 * 5,
	}

	path := PrimaryAtkSeq.HBoxPath

	// frame 0 - 2 no hitboxes

	// frame 3
	path = path.appendHboxAgg(-45, 10, 30, 30, 3)
	path = path.appendHboxAgg(-30, 20, 30, 30, 3)
	path = path.appendHboxAgg(0, 23, 30, 30, 3)
	path = path.appendHboxAgg(30, 18, 30, 30, 3)
	path = path.appendHboxAgg(45, 3, 30, 30, 3)
	path = path.appendHboxAgg(58, -13, 25, 30, 3)
	path = path.appendHboxAgg(58, -23, 15, 25, 3)
	path = path.appendHboxAgg(63, -33, 30, 15, 3)
	path = path.appendHboxAgg(48, -43, 30, 15, 3)

	// frame 4
	path = path.appendHboxAgg(-45, 10, 30, 30, 4)
	path = path.appendHboxAgg(-30, 20, 30, 30, 4)
	path.appendHboxAgg(0, 23, 30, 30, 4)

	// frame 5+ no hitbox

	atks[PrimaryAttackKey].HitBoxSequence = PrimaryAtkSeq

	return atks
}

func birdDroidSecondaryAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {

	atks[SecondaryAttackKey] = &AttackData{
		Name: SecondaryAttackKey,
		Type: SecondaryAttackKey,
		Consequence: &Consequence{
			Damage:             34,
			KnockbackX:         se.HitFloat,
			KnockbackY:         se.HitFloat,
			KnockbackXDuration: se.HitDuration,
			KnockbackYDuration: se.HitDuration,
		},
	}

	atkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 6),
		MovmentInc: 16.666 * 5,
	}

	path := atkSeq.HBoxPath

	// frame 0 - 2 no hitboxes

	// frame 3
	path = path.appendHboxAgg(-37, 10, 40, 20, 3)
	path = path.appendHboxAgg(-20, 25, 30, 110, 3)

	// frame 4
	path = path.appendHboxAgg(30, 25, 30, 60, 4)

	// frame 5
	path.appendHboxAgg(60, 25, 30, 30, 5)

	atks[SecondaryAttackKey].HitBoxSequence = atkSeq
	return atks
}

func birdDroidTertAtk(atks map[AtKey]*AttackData) map[AtKey]*AttackData {

	atks[TertAttackKey] = &AttackData{
		Name: TertAttackKey,
		Type: TertAttackKey,

		Consequence: &Consequence{
			Damage:             30,
			KnockbackX:         se.StunFloat,
			KnockbackY:         se.StunFloat,
			KnockbackXDuration: 2000,
			KnockbackYDuration: 2000,
		},
	}

	atkSeq := HitBoxSequence{
		HBoxPath:   make([]HitBoxAggregate, 9),
		MovmentInc: 16.666 * 5,
	}

	path := atkSeq.HBoxPath

	// frame 0 - 2 no hitboxes

	// frame 3
	path = path.appendHboxAgg(-50, 0, 35, 35, 3)

	// frame 4
	path = path.appendHboxAgg(-50, 0, 35, 35, 4)

	// frame 5
	path = path.appendHboxAgg(-53, -5, 35, 35, 5)

	// frame 6
	path = path.appendHboxAgg(60, 20, 35, 45, 6)
	path = path.appendHboxAgg(70, 0, 35, 45, 6)
	path = path.appendHboxAgg(65, -20, 35, 45, 6)
	path = path.appendHboxAgg(65, -35, 25, 35, 6)
	path = path.appendHboxAgg(55, -45, 25, 35, 6)
	path = path.appendHboxAgg(45, -55, 25, 35, 6)
	path = path.appendHboxAgg(25, -58, 25, 35, 6)
	path = path.appendHboxAgg(10, -58, 10, 35, 6)
	path = path.appendHboxAgg(0, -58, 10, 35, 6)
	path = path.appendHboxAgg(-10, -53, 10, 10, 6)
	path = path.appendHboxAgg(-20, -43, 10, 10, 6)
	path = path.appendHboxAgg(-30, -33, 10, 10, 6)

	// frame 7
	path = path.appendHboxAgg(60, 20, 35, 30, 7)
	path = path.appendHboxAgg(90, -10, 40, 10, 7)

	// frame 8
	path.appendHboxAgg(60, 20, 35, 30, 8)

	atks[TertAttackKey].HitBoxSequence = atkSeq
	return atks
}
