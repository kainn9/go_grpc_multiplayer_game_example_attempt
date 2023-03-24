package roles

type PlayerType string

type Role struct {
	RoleType PlayerType
	Attacks  map[AtKey]*AttackData
	HitBoxW  float64
	HitBoxH  float64
	Defense  *Defense
	Health   int

	Phys *RolePhysStruct
}

type AttackData struct {
	Name     AtKey
	Duration int
	Type     AtKey
	*Windup
	*Movement
	*Consequence
	HitBoxSequence HitBoxSequence
}

type Windup struct {
	Duration int
	*ChargeEffect
}

type ChargeEffect struct {
	MultFactorDmg      float64
	MultFactorMvSpeed  float64
	MultFactorMvDist   float64
	MultFactorKbxSpeed float64
	MultFactorKbySpeed float64
	MultFactorKbxDur   float64
	MultFactorKbyDur   float64
	TimeLimit          float64
}

type Movement struct {
	Distance       float64
	SpeedX         float64
	UseChargeDist  bool
	UseChargeSpeed bool
}

type HitBox struct {
	Height     float64
	Width      float64
	PlayerOffX float64
	PlayerOffY float64
}

type HitBoxAggregate []HitBox
type HBoxPath []HitBoxAggregate

type HitBoxSequence struct {
	HBoxPath   HBoxPath
	MovmentInc float64
}

type Consequence struct {
	Damage               int
	KnockbackX           float64
	KnockbackY           float64
	KnockbackXDuration   int
	KnockbackYDuration   int
	UseChargeKbyDuration bool
	UseChargeKbxDuration bool
	UseChargeKbxSpeed    bool
	UseChargeKbySpeed    bool
	UseChargeDmg         bool
}

type Defense struct {
	Displacment float64
	Delay       float64
	Speed       float64
	Cooldown    float64
}

// custom phys per role basis
type RolePhysStruct struct {
	DefaultFriction float64
	DefaultAccel    float64
	DefaultMaxSpeed float64
	DefaultJumpSpd  float64
	DefaultGravity  float64
}

const (
	KnightType PlayerType = "knight"
	MageType   PlayerType = "mage"
	MonkType   PlayerType = "monk"
)

type AtKey string

const (
	PrimaryAttackKey    AtKey = "primaryAtk"
	SecondaryAttackKey  AtKey = "secondaryAtk"
	TertAttackKey       AtKey = "tertAtk"
	QuaternaryAttackKey AtKey = "quaAtk"
)

const noBox = -10000

func (path HBoxPath) appendHboxAgg(x float64, y float64, h float64, w float64, index int) HBoxPath {

	path[index] = append(path[index], HitBox{
		PlayerOffX: x,
		PlayerOffY: y,
		Height:     h,
		Width:      w,
	})

	return path
}

func (wu *Windup) HasChargeEffect() bool {
	if wu == nil {
		return false
	}
	return wu.ChargeEffect != nil
}
