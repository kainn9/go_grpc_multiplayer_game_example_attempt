package roles

import (
	"github.com/hajimehoshi/ebiten/v2"
	sr "github.com/kainn9/grpc_game/server/roles"
	ut "github.com/kainn9/grpc_game/util"
)

/*
File for Demon class
contains sprites/animation data
*/
var (
	Demon *Role = InitMonk()
)

var (
	demonSpriteIdleLeft  *ebiten.Image
	demonSpriteIdleRight *ebiten.Image

	demonSpriteWalkingRight *ebiten.Image
	demonSpriteWalkingLeft  *ebiten.Image

	demonSpriteJumpLeft  *ebiten.Image
	demonSpriteJumpRight *ebiten.Image


	demonSpriteKBRight *ebiten.Image
	demonSpriteKBLeft  *ebiten.Image

	demonSpriteCleaveAtkRight *ebiten.Image
	demonSpriteCleaveAtkLeft  *ebiten.Image

	demonSpriteFireAtkRight *ebiten.Image
	demonSpriteFireAtkLeft  *ebiten.Image

	demonSpriteDeathRight *ebiten.Image
	demonSpriteDeathLeft  *ebiten.Image
)

/*
Loads the default player sprites
*/
func LoadDemonSprites() {
	demonSpriteIdleLeft = ut.LoadImg("./sprites/demon/demonIdleLeft.png")
	demonSpriteIdleRight = ut.LoadImg("./sprites/demon/demonIdleRight.png")

	demonSpriteWalkingRight = ut.LoadImg("./sprites/demon/demonRunningRight.png")
	demonSpriteWalkingLeft = ut.LoadImg("./sprites/demon/demonRunningLeft.png")

	demonSpriteJumpLeft = ut.LoadImg("./sprites/demon/demonJumpLeft.png")
	demonSpriteJumpRight = ut.LoadImg("./sprites/demon/demonJumpRight.png")


	demonSpriteKBRight = ut.LoadImg("./sprites/demon/demonKnockBackRight.png")
	demonSpriteKBLeft = ut.LoadImg("./sprites/demon/demonKnockBackLeft.png")

	demonSpriteCleaveAtkRight = ut.LoadImg("./sprites/demon/demonCleaveRight.png")
	demonSpriteCleaveAtkLeft = ut.LoadImg("./sprites/demon/demonCleaveLeft.png")

	demonSpriteFireAtkRight = ut.LoadImg("./sprites/demon/demonFireAtkRight.png")
	demonSpriteFireAtkLeft = ut.LoadImg("./sprites/demon/demonFireAtkLeft.png")

	demonSpriteDeathRight = ut.LoadImg("./sprites/demon/demonDeathRight.png")
	demonSpriteDeathLeft = ut.LoadImg("./sprites/demon/demonDeathLeft.png")
}

func InitDemon() *Role {
	LoadDemonSprites()

	r := &Role{
		RoleType:      DemonType,
		Animations:    DemonAnims(),
		HitBoxOffsetY: 30,
		HitBoxOffsetX: 30,
	}

	return r
}

// TODO MAKE ANIM KEYS CONSTS
func DemonAnims() map[string]*Animation {
	anims := make(map[string]*Animation)

	anims[string(IdleRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  107,
		FrameHeight: 111,
		FrameCount:  6,
		SpriteSheet: demonSpriteIdleRight,
	}

	anims[string(IdleLeft)] = &Animation{
		FrameOX:     642,
		FrameOY:     0,
		FrameWidth:  107,
		FrameHeight: 111,
		FrameCount:  6,
		SpriteSheet: demonSpriteIdleLeft,
	}

	anims[string(WalkRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  89,
		FrameHeight: 135,
		FrameCount:  12,
		PosOffsetX: -10,
		SpriteSheet: demonSpriteWalkingRight,
	}

	anims[string(WalkLeft)] = &Animation{
		FrameOX:     1068,
		FrameOY:     0,
		FrameWidth:  89,
		FrameHeight: 135,
		FrameCount:  12,
		PosOffsetX: -10,
		SpriteSheet: demonSpriteWalkingLeft,
	}

	anims[string(JumpLeft)] = &Animation{
		FrameOX:     110,
		FrameOY:     0,
		FrameWidth:  111,
		FrameHeight: 112,
		FrameCount:  1,
		SpriteSheet: demonSpriteJumpLeft,
	}

	anims[string(JumpRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  111,
		FrameHeight: 112,
		FrameCount:  1,
		SpriteSheet: demonSpriteJumpRight,
	}

	anims[string(KbRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  116,
		FrameHeight: 116,
		FrameCount:  1,
		SpriteSheet: demonSpriteKBRight,
	}

	anims[string(KbLeft)] = &Animation{
		FrameOX:     116,
		FrameOY:     0,
		FrameWidth:  116,
		FrameHeight: 116,
		FrameCount:  1,
		SpriteSheet: demonSpriteKBLeft,
	}

	anims[string(DeathRight)] = &Animation{
		Name:        string(DeathRight),
		FrameOX:     0,
		FrameOY:     26,
		FrameWidth:  147,
		FrameHeight: 138,
		FrameCount:  24,
		PosOffsetX: 28,
		SpriteSheet: demonSpriteDeathRight,
		Fixed:       true,
	}

	anims[string(DeathLeft)] = &Animation{
		Name:        string(DeathLeft),
		FrameOX:     3528,
		FrameOY:     26,
		FrameWidth:  147,
		FrameHeight: 138,
		FrameCount:  24,
		PosOffsetX: 28,
		SpriteSheet: demonSpriteDeathLeft,
		Fixed:       true,
	}

	/*
		---------------------------------------------------------------------------------
		Primary Attack
		---------------------------------------------------------------------------------
	*/
	anims[string(sr.PrimaryAttackKey)+"Right"] = &Animation{
		Name:        string(sr.PrimaryAttackKey)+"Right",
		FrameOX:     0,
		FrameOY:     5,
		FrameWidth:  176,
		FrameHeight: 121,
		FrameCount:  22,
		PosOffsetX: 10,
		SpriteSheet: demonSpriteCleaveAtkRight,
		Fixed:       true,
	}

	anims[string(sr.PrimaryAttackKey)+"Left"] = &Animation{
		Name:        string(sr.PrimaryAttackKey)+"Left",
		FrameOX:     3872,
		FrameOY:     5,
		FrameWidth:  176,
		FrameHeight: 121,
		FrameCount:  22,
		PosOffsetX: 10,
		SpriteSheet: demonSpriteCleaveAtkLeft,
		Fixed:       true,
	}
	/*
		---------------------------------------------------------------------------------
		Primary Attack END
		---------------------------------------------------------------------------------
	*/

	/*
		---------------------------------------------------------------------------------
		Secondary Attack
		---------------------------------------------------------------------------------
	*/

	a2arKey := string(sr.SecondaryAttackKey) + "Right"
	anims[a2arKey] = &Animation{
		Name:        a2arKey,
		FrameOX:     0,
		FrameOY:     50,
		FrameWidth:  288,
		FrameHeight: 160,
		FrameCount:  21,
		PosOffsetX: 80,
		SpriteSheet: demonSpriteFireAtkRight,
		Fixed:       true,
	}

	a2alKey := string(sr.SecondaryAttackKey) + "Left"
	anims[a2alKey] = &Animation{
		Name:        a2alKey,
		FrameOX:     6048,
		FrameOY:     50,
		FrameWidth:  288,
		FrameHeight: 160,
		FrameCount:  21,
		PosOffsetX: 80,
		SpriteSheet: demonSpriteFireAtkLeft,
		Fixed:       true,
	}

	/*
		---------------------------------------------------------------------------------
		Secondary End
		---------------------------------------------------------------------------------
	*/

	return anims

}
