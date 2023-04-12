package statusEffects

import (
	"github.com/hajimehoshi/ebiten/v2"
	cr "github.com/kainn9/grpc_game/client/roles"
)

type StatusEffect struct {
	Img   *ebiten.Image
	Anim  *cr.Animation
	Width float64
}

var (
	Stun = InitStunEffect()
)
