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
	*Windup
	*Movement
}

type Windup struct {
	Name AtKey
	Duration int
}

type Movement struct {
	Distance int
	Increment int
}

type Hitbox struct {
	Height float64
	Width  float64
	PlayerOffX  float64
	PlayerOffY  float64
}


type Consequence struct {
	Damage int
	Effects []string
}


const (
	KnightType PlayerType = "knight"
	MageType   PlayerType = "mage"
)

type AtKey string

const (
	PrimaryAttackKey AtKey = "primaryAtk"
	TestAttackKey AtKey = "test2Atk"
)
