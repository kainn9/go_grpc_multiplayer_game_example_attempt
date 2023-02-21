package roles

import (
	"github.com/hajimehoshi/ebiten/v2"
	sr "github.com/kainn9/grpc_game/server/roles"
	ut "github.com/kainn9/grpc_game/util"
)

/*
File for Knight class
contains sprites/animation data
*/
var (
	Knight *Role = InitKnight()
)

// Player Sprites(for now)
var (
	playerSpriteIdleLeft  *ebiten.Image
	playerSpriteIdleRight *ebiten.Image

	playerSpriteWalkingRight *ebiten.Image
	playerSpriteWalkingLeft  *ebiten.Image
	playerSpriteJumpLeft     *ebiten.Image
	playerSpriteJumpRight    *ebiten.Image
	playerSpriteAttackLeft   *ebiten.Image
	playerSpriteAttackRight  *ebiten.Image
	playerSpriteKBRight      *ebiten.Image
	playerSpriteKBLeft       *ebiten.Image
)

/*
Loads the default player sprites
*/
func LoadKnightSprites() {
	playerSpriteIdleLeft = ut.LoadImg("./sprites/knight/knightIdleLeft.png")
	playerSpriteIdleRight = ut.LoadImg("./sprites/knight/knightIdleRight.png")

	playerSpriteWalkingRight = ut.LoadImg("./sprites/knight/knightRunningRight.png")
	playerSpriteWalkingLeft = ut.LoadImg("./sprites/knight/knightRunningLeft.png")
	playerSpriteJumpLeft = ut.LoadImg("./sprites/knight/knightJumpLeft.png")
	playerSpriteJumpRight = ut.LoadImg("./sprites/knight/knightJumpRight.png")
	playerSpriteAttackRight = ut.LoadImg("./sprites/knight/knightAttackRight.png")
	playerSpriteAttackLeft = ut.LoadImg("./sprites/knight/knightAttackLeft.png")
	playerSpriteKBRight = ut.LoadImg("./sprites/knight/knightKnockBackRight.png")
	playerSpriteKBLeft = ut.LoadImg("./sprites/knight/knightKnockBackLeft.png")
}

func InitKnight() *Role {
	LoadKnightSprites()

	r := &Role{
		RoleType:      KnightType,
		Animations:    KnightAnims(),
		HitBoxOffsetY: 4,
		HitBoxOffsetX: 8,
	}

	return r
}

func KnightAnims() map[string]*Animation {
	anims := make(map[string]*Animation)

	anims["idleRight"] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  32,
		FrameHeight: 48,
		FrameCount:  8,
		SpriteSheet: playerSpriteIdleRight,
	}

	anims["idleLeft"] = &Animation{
		FrameOX:     256,
		FrameOY:     0,
		FrameWidth:  32,
		FrameHeight: 48,
		FrameCount:  8,
		SpriteSheet: playerSpriteIdleLeft,
	}

	anims["walkRight"] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 48,
		FrameCount:  8,
		SpriteSheet: playerSpriteWalkingRight,
	}

	anims["walkLeft"] = &Animation{
		FrameOX:     280,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 48,
		FrameCount:  8,
		SpriteSheet: playerSpriteWalkingLeft,
	}

	anims["jumpLeft"] = &Animation{
		FrameOX:     35,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 48,
		FrameCount:  1,
		SpriteSheet: playerSpriteJumpLeft,
	}

	anims["jumpRight"] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 48,
		FrameCount:  1,
		SpriteSheet: playerSpriteJumpRight,
	}

	anims[string(sr.PrimaryAttackKey)+"Right"] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  65,
		FrameHeight: 48,
		FrameCount:  4,
		SpriteSheet: playerSpriteAttackRight,
	}

	anims[string(sr.PrimaryAttackKey)+"Left"] = &Animation{
		FrameOX:     260,
		FrameOY:     0,
		FrameWidth:  65,
		FrameHeight: 48,
		FrameCount:  4,
		SpriteSheet: playerSpriteAttackLeft,
	}

	anims["KbRight"] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  30,
		FrameHeight: 48,
		FrameCount:  4,
		SpriteSheet: playerSpriteKBRight,
	}

	anims["KbLeft"] = &Animation{
		FrameOX:     120,
		FrameOY:     0,
		FrameWidth:  30,
		FrameHeight: 32,
		FrameCount:  4,
		SpriteSheet: playerSpriteKBLeft,
	}

	return anims

}
