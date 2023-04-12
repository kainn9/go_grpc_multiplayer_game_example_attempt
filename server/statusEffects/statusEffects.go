package statusEffects

type CCString string

const (
	None      CCString = ""
	KnockBack CCString = "kb"
	Stun      CCString = "stun"
	Hit       CCString = "hit"

	// hacky way of representing stun
	StunFloat float64 = 0.123
	HitFloat  float64 = 0.122

	HitDuration int = 250
)
