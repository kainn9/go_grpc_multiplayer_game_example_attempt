package statusEffects

import (
	"github.com/hajimehoshi/ebiten/v2"
	cr "github.com/kainn9/grpc_game/client/roles"
	utClient "github.com/kainn9/grpc_game/client_util"
)

var (
	stunEffectSprite *ebiten.Image
)

func loadStunEffectSprite() {
	stunEffectSprite = utClient.LoadImage("./sprites/statusEffects/stun.png")
}

func InitStunEffect() *StatusEffect {
	loadStunEffectSprite()

	animation := &cr.Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  51,
		FrameHeight: 38,
		FrameCount:  4,
		SpriteSheet: stunEffectSprite,
	}

	return &StatusEffect{
		Anim:  animation,
		Img:   stunEffectSprite,
		Width: 51,
	}
}
