package roles

type PlayerType string

type Role struct {
	RoleType PlayerType
	Attacks  map[AtKey]*Attack
	HitBoxW  float64
	HitBoxH  float64
}

type Attack struct {
	Cooldown int
	Width    float64
	Height   float64
	OffsetX  float64
	OffsetY  float64
	Duration int
	Type     AtKey
}

const (
	KnightType PlayerType = "knight"
	MageType   PlayerType = "mage"
)

type AtKey string

const (
	PrimaryAttackKey AtKey = "primaryAtk"
)
