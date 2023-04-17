package roles

import (
	"github.com/hajimehoshi/ebiten/v2"
	utClient "github.com/kainn9/grpc_game/client_util"
	sr "github.com/kainn9/grpc_game/server/roles"
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

	demonSpriteHitRight *ebiten.Image
	demonSpriteHitLeft  *ebiten.Image

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
	demonSpriteIdleLeft = utClient.LoadImage("./sprites/demon/demonIdleLeft.png")
	demonSpriteIdleRight = utClient.LoadImage("./sprites/demon/demonIdleRight.png")

	demonSpriteWalkingRight = utClient.LoadImage("./sprites/demon/demonRunningRight.png")
	demonSpriteWalkingLeft = utClient.LoadImage("./sprites/demon/demonRunningLeft.png")

	demonSpriteJumpLeft = utClient.LoadImage("./sprites/demon/demonJumpLeft.png")
	demonSpriteJumpRight = utClient.LoadImage("./sprites/demon/demonJumpRight.png")

	demonSpriteHitRight = utClient.LoadImage("./sprites/demon/demonHitRight.png")
	demonSpriteHitLeft = utClient.LoadImage("./sprites/demon/demonHitLeft.png")

	demonSpriteKBRight = utClient.LoadImage("./sprites/demon/demonKnockBackRight.png")
	demonSpriteKBLeft = utClient.LoadImage("./sprites/demon/demonKnockBackLeft.png")

	demonSpriteCleaveAtkRight = utClient.LoadImage("./sprites/demon/demonCleaveRight.png")
	demonSpriteCleaveAtkLeft = utClient.LoadImage("./sprites/demon/demonCleaveLeft.png")

	demonSpriteFireAtkRight = utClient.LoadImage("./sprites/demon/demonFireAtkRight.png")
	demonSpriteFireAtkLeft = utClient.LoadImage("./sprites/demon/demonFireAtkLeft.png")

	demonSpriteDeathRight = utClient.LoadImage("./sprites/demon/demonDeathRight.png")
	demonSpriteDeathLeft = utClient.LoadImage("./sprites/demon/demonDeathLeft.png")
}

func InitDemon() *Role {
	LoadDemonSprites()

	r := &Role{
		RoleType:      sr.DemonType,
		Animations:    DemonAnims(),
		HitBoxOffsetY: 30,
		HitBoxOffsetX: 30,
		Health:        sr.Demon.Health,
		HitBoxW:       sr.Demon.HitBoxW,
		HitBoxH:       sr.Demon.HitBoxH,
		HealthBarOffset: &Offset{
			X: 30,
			Y: -10,
		},
		StatusEffectOffset: &Offset{
			X: 30,
			Y: -6,
		},
	}

	return r
}

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
		PosOffsetX:  -10,
		SpriteSheet: demonSpriteWalkingRight,
	}

	anims[string(WalkLeft)] = &Animation{
		FrameOX:     1068,
		FrameOY:     0,
		FrameWidth:  89,
		FrameHeight: 135,
		FrameCount:  12,
		PosOffsetX:  -10,
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

	anims[string(HitRight)] = &Animation{
		Name:        string(HitRight),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  126,
		FrameHeight: 126,
		FrameCount:  6,
		PosOffsetX:  28,
		PosOffsetY:  16,
		SpriteSheet: demonSpriteHitRight,
		Fixed:       true,
	}

	anims[string(HitLeft)] = &Animation{
		Name:        string(HitLeft),
		FrameOX:     756,
		FrameOY:     0,
		FrameWidth:  126,
		FrameHeight: 126,
		FrameCount:  6,
		PosOffsetX:  28,
		PosOffsetY:  16,
		SpriteSheet: demonSpriteHitLeft,
		Fixed:       true,
	}

	stunAnimCopyRight := *anims[string(HitRight)]
	anims[string(StunRight)] = &stunAnimCopyRight

	stunAnimCopyLeft := *anims[string(HitLeft)]
	anims[string(StunLeft)] = &stunAnimCopyLeft

	anims[string(KbRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  126,
		FrameHeight: 126,
		FrameCount:  6,
		PosOffsetX:  28,
		PosOffsetY:  16,
		SpriteSheet: demonSpriteKBRight,
	}

	anims[string(KbLeft)] = &Animation{
		FrameOX:     756,
		FrameOY:     0,
		FrameWidth:  126,
		FrameHeight: 126,
		FrameCount:  6,
		PosOffsetX:  28,
		PosOffsetY:  16,
		SpriteSheet: demonSpriteKBLeft,
	}

	anims[string(DeathRight)] = &Animation{
		Name:        string(DeathRight),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  147,
		FrameHeight: 138,
		FrameCount:  24,
		PosOffsetX:  28,
		PosOffsetY:  26,
		SpriteSheet: demonSpriteDeathRight,
		Fixed:       true,
	}

	anims[string(DeathLeft)] = &Animation{
		Name:        string(DeathLeft),
		FrameOX:     3528,
		FrameOY:     0,
		FrameWidth:  147,
		FrameHeight: 138,
		FrameCount:  24,
		PosOffsetX:  28,
		PosOffsetY:  26,
		SpriteSheet: demonSpriteDeathLeft,
		Fixed:       true,
	}

	/*
		---------------------------------------------------------------------------------
		Primary Attack
		---------------------------------------------------------------------------------
	*/
	anims[string(sr.PrimaryAttackKey)+"Right"] = &Animation{
		Name:        string(sr.PrimaryAttackKey) + "Right",
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  176,
		FrameHeight: 121,
		FrameCount:  22,
		PosOffsetX:  10,
		PosOffsetY:  5,
		SpriteSheet: demonSpriteCleaveAtkRight,
		Fixed:       true,
	}

	anims[string(sr.PrimaryAttackKey)+"Left"] = &Animation{
		Name:        string(sr.PrimaryAttackKey) + "Left",
		FrameOX:     3872,
		FrameOY:     0,
		FrameWidth:  176,
		FrameHeight: 121,
		FrameCount:  22,
		PosOffsetX:  10,
		PosOffsetY:  5,
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
		FrameOY:     0,
		FrameWidth:  288,
		FrameHeight: 160,
		FrameCount:  21,
		PosOffsetX:  80,
		PosOffsetY:  50,
		SpriteSheet: demonSpriteFireAtkRight,
		Fixed:       true,
	}

	a2alKey := string(sr.SecondaryAttackKey) + "Left"
	anims[a2alKey] = &Animation{
		Name:        a2alKey,
		FrameOX:     6048,
		FrameOY:     0,
		FrameWidth:  288,
		FrameHeight: 160,
		FrameCount:  21,
		PosOffsetX:  80,
		PosOffsetY:  50,
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
