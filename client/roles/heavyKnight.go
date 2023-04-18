package roles

import (
	"github.com/hajimehoshi/ebiten/v2"
	utClient "github.com/kainn9/grpc_game/client_util"
	sr "github.com/kainn9/grpc_game/server/roles"
)

/*
File for HeavyKnight class
contains sprites/animation data
*/

var (
	HeavyKnight *Role = InitHeavyKnight()
)

var (
	heavyKnightSpriteIdleLeft  *ebiten.Image
	heavyKnightSpriteIdleRight *ebiten.Image

	heavyKnightSpriteWalkingRight *ebiten.Image
	heavyKnightSpriteWalkingLeft  *ebiten.Image

	heavyKnightSpriteJumpLeft  *ebiten.Image
	heavyKnightSpriteJumpRight *ebiten.Image

	heavyKnightSpriteHitRight *ebiten.Image
	heavyKnightSpriteHitLeft  *ebiten.Image

	heavyKnightSpriteKBRight *ebiten.Image
	heavyKnightSpriteKBLeft  *ebiten.Image

	heavyKnightSpriteDeathRight *ebiten.Image
	heavyKnightSpriteDeathLeft  *ebiten.Image

	heavyKnightSpriteSweepSlashRight *ebiten.Image
	heavyKnightSpriteSweepSlashLeft  *ebiten.Image

	heavyKnightSpriteLowSweepRight *ebiten.Image
	heavyKnightSpriteLowSweepLeft  *ebiten.Image

	heavyKnightSpriteSmashSlashRight *ebiten.Image
	heavyKnightSpriteSmashSlashLeft  *ebiten.Image
)

/*
Loads the default player sprites
*/
func LoadHeavyKnightSprites() {
	heavyKnightSpriteIdleLeft = utClient.LoadImage("./sprites/heavyKnight/heavyKnightIdleLeft.png")
	heavyKnightSpriteIdleRight = utClient.LoadImage("./sprites/heavyKnight/heavyKnightIdleRight.png")

	heavyKnightSpriteWalkingRight = utClient.LoadImage("./sprites/heavyKnight/heavyKnightRunningRight.png")
	heavyKnightSpriteWalkingLeft = utClient.LoadImage("./sprites/heavyKnight/heavyKnightRunningLeft.png")

	heavyKnightSpriteJumpLeft = utClient.LoadImage("./sprites/heavyKnight/heavyKnightJumpLeft.png")
	heavyKnightSpriteJumpRight = utClient.LoadImage("./sprites/heavyKnight/heavyKnightJumpRight.png")

	heavyKnightSpriteHitRight = utClient.LoadImage("./sprites/heavyKnight/heavyKnightHitRight.png")
	heavyKnightSpriteHitLeft = utClient.LoadImage("./sprites/heavyKnight/heavyKnightHitLeft.png")

	heavyKnightSpriteKBRight = utClient.LoadImage("./sprites/heavyKnight/heavyKnightKnockBackRight.png")
	heavyKnightSpriteKBLeft = utClient.LoadImage("./sprites/heavyKnight/heavyKnightKnockBackLeft.png")

	heavyKnightSpriteDeathRight = utClient.LoadImage("./sprites/heavyKnight/heavyKnightDeathLeft.png")
	heavyKnightSpriteDeathLeft = utClient.LoadImage("./sprites/heavyKnight/heavyKnightDeathLeft.png")

	heavyKnightSpriteSweepSlashRight = utClient.LoadImage("./sprites/heavyKnight/heavyKnightSweepSlashRight.png")
	heavyKnightSpriteSweepSlashLeft = utClient.LoadImage("./sprites/heavyKnight/heavyKnightSweepSlashLeft.png")

	heavyKnightSpriteLowSweepRight = utClient.LoadImage("./sprites/heavyKnight/heavyKnightLowSweepRight.png")
	heavyKnightSpriteLowSweepLeft = utClient.LoadImage("./sprites/heavyKnight/heavyKnightLowSweepLeft.png")

	heavyKnightSpriteSmashSlashRight = utClient.LoadImage("./sprites/heavyKnight/heavyKnightSmashSlashRight.png")
	heavyKnightSpriteSmashSlashLeft = utClient.LoadImage("./sprites/heavyKnight/heavyKnightSmashSlashLeft.png")
}

func InitHeavyKnight() *Role {
	LoadHeavyKnightSprites()

	r := &Role{
		Animations:    HeavyKnightAnims(),
		HitBoxOffsetY: 72,
		HitBoxOffsetX: 85,
		Health:        sr.HeavyKnight.Health,
		HitBoxW:       sr.HeavyKnight.HitBoxW,
		HitBoxH:       sr.HeavyKnight.HitBoxH,
		HealthBarOffset: &Offset{
			X: 80,
			Y: 50,
		},
		StatusEffectOffset: &Offset{
			X: 80,
			Y: 50,
		},
	}

	return r
}

// TODO MAKE ANIM KEYS CONSTS
func HeavyKnightAnims() map[string]*Animation {
	anims := make(map[string]*Animation)

	anims[string(IdleRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  10,
		SpriteSheet: heavyKnightSpriteIdleRight,
	}

	anims[string(IdleLeft)] = &Animation{
		FrameOX:     2000,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  10,
		SpriteSheet: heavyKnightSpriteIdleLeft,
	}

	anims[string(WalkRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  8,
		SpriteSheet: heavyKnightSpriteWalkingRight,
	}

	anims[string(WalkLeft)] = &Animation{
		FrameOX:     1600,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  8,
		SpriteSheet: heavyKnightSpriteWalkingLeft,
	}

	anims[string(JumpLeft)] = &Animation{
		FrameOX:     142,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 162,
		FrameCount:  1,
		PosOffsetX:  -40,
		PosOffsetY:  -20,
		SpriteSheet: heavyKnightSpriteJumpLeft,
	}

	anims[string(JumpRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 162,
		FrameCount:  1,
		PosOffsetX:  -40,
		PosOffsetY:  -20,
		SpriteSheet: heavyKnightSpriteJumpRight,
	}

	anims[string(HitRight)] = &Animation{
		Name:        string(HitRight),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  3,
		SpriteSheet: heavyKnightSpriteHitRight,
		Fixed:       true,
	}

	anims[string(HitLeft)] = &Animation{
		Name:        string(HitLeft),
		FrameOX:     600,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  3,
		SpriteSheet: heavyKnightSpriteHitLeft,
		Fixed:       true,
	}

	stunAnimCopyRight := *anims[string(HitRight)]
	anims[string(StunRight)] = &stunAnimCopyRight

	stunAnimCopyLeft := *anims[string(HitLeft)]
	anims[string(StunLeft)] = &stunAnimCopyLeft

	anims[string(KbRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  4,
		SpriteSheet: heavyKnightSpriteKBRight,
	}

	anims[string(KbLeft)] = &Animation{
		FrameOX:     800,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  4,
		SpriteSheet: heavyKnightSpriteKBLeft,
	}

	anims[string(DeathRight)] = &Animation{
		Name:        string(DeathLeft),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  24,
		SpriteSheet: heavyKnightSpriteDeathRight,
		Fixed:       true,
	}

	anims[string(DeathLeft)] = &Animation{
		Name:        string(DeathLeft),
		FrameOX:     2000,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  24,
		SpriteSheet: heavyKnightSpriteDeathLeft,
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
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  6,
		SpriteSheet: heavyKnightSpriteSweepSlashRight,
		Fixed:       true,
	}

	anims[string(sr.PrimaryAttackKey)+"Left"] = &Animation{
		Name:        string(sr.PrimaryAttackKey) + "Left",
		FrameOX:     1200,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  6,
		SpriteSheet: heavyKnightSpriteSweepSlashLeft,
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
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  6,

		SpriteSheet: heavyKnightSpriteLowSweepRight,
		Fixed:       true,
	}

	a2alKey := string(sr.SecondaryAttackKey) + "Left"
	anims[a2alKey] = &Animation{
		Name:        a2alKey,
		FrameOX:     1200,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  6,

		SpriteSheet: heavyKnightSpriteLowSweepLeft,
		Fixed:       true,
	}

	/*
		---------------------------------------------------------------------------------
		Secondary End
		---------------------------------------------------------------------------------
	*/

	/*
		---------------------------------------------------------------------------------
		Tert Attack
		---------------------------------------------------------------------------------
	*/

	a3arKey := "tertAtkRight"
	anims[a3arKey] = &Animation{
		Name:        a3arKey,
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  9,
		SpriteSheet: heavyKnightSpriteSmashSlashRight,
		Fixed:       true,
	}

	a3alKey := "tertAtkLeft"
	anims[a3alKey] = &Animation{
		Name:        a3alKey,
		FrameOX:     1800,
		FrameOY:     0,
		FrameWidth:  200,
		FrameHeight: 200,
		FrameCount:  9,
		SpriteSheet: heavyKnightSpriteSmashSlashLeft,
		Fixed:       true,
	}

	/*
		---------------------------------------------------------------------------------
		Tert Attack END
		---------------------------------------------------------------------------------
	*/

	return anims

}
