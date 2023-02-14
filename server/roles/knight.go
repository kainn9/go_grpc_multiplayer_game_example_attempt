package roles

var (
	Knight *Role = InitKnight()
)

func KnightAttacks() map[AtKey]*Attack {
	atks := make(map[AtKey]*Attack)

	atks[PrimaryAttackKey] = &Attack{
		Cooldown: 5,
		Width:    10,
		Height:   5,
		OffsetX:  40,
		OffsetY:  15,
		Duration: 333,
		Type:     PrimaryAttackKey,
	}

	return atks

}

func InitKnight() *Role {

	r := &Role{
		RoleType: KnightType,
		Attacks:  KnightAttacks(),
		HitBoxW:  18,
		HitBoxH:  44,
	}

	return r
}


