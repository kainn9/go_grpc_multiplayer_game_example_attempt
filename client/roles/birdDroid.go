package roles

import (
	"github.com/hajimehoshi/ebiten/v2"
	utClient "github.com/kainn9/grpc_game/client_util"
	sr "github.com/kainn9/grpc_game/server/roles"
)

/*
File for BirdDroid class
contains sprites/animation data
*/

var (
	BirdDroid *Role = InitBirdDroid()
)

var (
	birdDroidSpriteIdleLeft  *ebiten.Image
	birdDroidSpriteIdleRight *ebiten.Image

	birdDroidSpriteWalkingRight *ebiten.Image
	birdDroidSpriteWalkingLeft  *ebiten.Image

	birdDroidSpriteJumpLeft  *ebiten.Image
	birdDroidSpriteJumpRight *ebiten.Image

	birdDroidSpriteHitRight *ebiten.Image
	birdDroidSpriteHitLeft  *ebiten.Image

	birdDroidSpriteKBRight *ebiten.Image
	birdDroidSpriteKBLeft  *ebiten.Image

	birdDroidSpriteDeathRight *ebiten.Image
	birdDroidSpriteDeathLeft  *ebiten.Image

	birdDroidSpriteTeleSlamRight *ebiten.Image
	birdDroidSpriteTeleSlamLeft  *ebiten.Image

	birdDroidSpriteLazerShotRight *ebiten.Image
	birdDroidSpriteLazerShotLeft  *ebiten.Image
)

/*
Loads the default player sprites
*/
func LoadBirdDroidSprites() {
	birdDroidSpriteIdleLeft = utClient.LoadImage("./sprites/birdDroid/birdDroidIdleLeft.png")
	birdDroidSpriteIdleRight = utClient.LoadImage("./sprites/birdDroid/birdDroidIdleRight.png")

	birdDroidSpriteWalkingRight = utClient.LoadImage("./sprites/birdDroid/birdDroidRunningRight.png")
	birdDroidSpriteWalkingLeft = utClient.LoadImage("./sprites/birdDroid/birdDroidRunningLeft.png")

	birdDroidSpriteJumpLeft = utClient.LoadImage("./sprites/birdDroid/birdDroidJumpLeft.png")
	birdDroidSpriteJumpRight = utClient.LoadImage("./sprites/birdDroid/birdDroidJumpRight.png")

	birdDroidSpriteHitRight = utClient.LoadImage("./sprites/birdDroid/birdDroidHitRight.png")
	birdDroidSpriteHitLeft = utClient.LoadImage("./sprites/birdDroid/birdDroidHitLeft.png")

	birdDroidSpriteKBRight = utClient.LoadImage("./sprites/birdDroid/birdDroidKnockBackRight.png")
	birdDroidSpriteKBLeft = utClient.LoadImage("./sprites/birdDroid/birdDroidKnockBackLeft.png")

	birdDroidSpriteDeathRight = utClient.LoadImage("./sprites/birdDroid/birdDroidDeathRight.png")
	birdDroidSpriteDeathLeft = utClient.LoadImage("./sprites/birdDroid/birdDroidDeathLeft.png")

	birdDroidSpriteTeleSlamRight = utClient.LoadImage("./sprites/birdDroid/birdDroidTeleSlamRight.png")
	birdDroidSpriteTeleSlamLeft = utClient.LoadImage("./sprites/birdDroid/birdDroidTeleSlamLeft.png")

	birdDroidSpriteLazerShotRight = utClient.LoadImage("./sprites/birdDroid/birdDroidLazerShotRight.png")
	birdDroidSpriteLazerShotLeft = utClient.LoadImage("./sprites/birdDroid/birdDroidLazerShotLeft.png")

}

func InitBirdDroid() *Role {
	LoadBirdDroidSprites()

	r := &Role{
		Animations:    BirdDroidAnims(),
		HitBoxOffsetY: 73,
		HitBoxOffsetX: 40,
		Health:        sr.BirdDroid.Health,
		HitBoxW:       sr.BirdDroid.HitBoxW,
		HitBoxH:       sr.BirdDroid.HitBoxH,
		AttackCount:   len(sr.BirdDroid.Attacks),
		HasDefense:    sr.BirdDroid.Defense != nil,
		HealthBarOffset: &Offset{
			X: 30,
			Y: 50,
		},
		StatusEffectOffset: &Offset{
			X: 30,
			Y: 50,
		},
	}

	return r
}

// TODO MAKE ANIM KEYS CONSTS
func BirdDroidAnims() map[string]*Animation {
	anims := make(map[string]*Animation)

	anims[string(IdleRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  7,
		SpriteSheet: birdDroidSpriteIdleRight,
	}

	anims[string(IdleLeft)] = &Animation{
		FrameOX:     994,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  7,
		PosOffsetX:  -44,
		SpriteSheet: birdDroidSpriteIdleLeft,
	}

	anims[string(WalkRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  9,
		SpriteSheet: birdDroidSpriteWalkingRight,
	}

	anims[string(WalkLeft)] = &Animation{
		FrameOX:     1278,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  9,
		PosOffsetX:  -44,
		SpriteSheet: birdDroidSpriteWalkingLeft,
	}

	anims[string(JumpLeft)] = &Animation{
		FrameOX:     142,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  1,
		PosOffsetX:  -44,
		SpriteSheet: birdDroidSpriteJumpLeft,
	}

	anims[string(JumpRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  1,
		SpriteSheet: birdDroidSpriteJumpRight,
	}

	anims[string(HitRight)] = &Animation{
		Name:        string(HitRight),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  12,
		SpriteSheet: birdDroidSpriteHitRight,
		Fixed:       true,
	}

	anims[string(HitLeft)] = &Animation{
		Name:        string(HitLeft),
		FrameOX:     1704,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  12,
		PosOffsetX:  -44,
		SpriteSheet: birdDroidSpriteHitLeft,
		Fixed:       true,
	}

	stunAnimCopyRight := *anims[string(HitRight)]
	anims[string(StunRight)] = &stunAnimCopyRight

	stunAnimCopyLeft := *anims[string(HitLeft)]
	anims[string(StunLeft)] = &stunAnimCopyLeft

	anims[string(KbRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  4,
		SpriteSheet: birdDroidSpriteKBRight,
	}

	anims[string(KbLeft)] = &Animation{
		FrameOX:     568,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  4,
		PosOffsetX:  -44,
		SpriteSheet: birdDroidSpriteKBLeft,
	}

	anims[string(DeathRight)] = &Animation{
		Name:        string(DeathRight),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  24,
		SpriteSheet: birdDroidSpriteDeathRight,
		Fixed:       true,
	}

	anims[string(DeathLeft)] = &Animation{
		Name:        string(DeathLeft),
		FrameOX:     3408,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  24,
		PosOffsetX:  -44,
		SpriteSheet: birdDroidSpriteDeathLeft,
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
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  9,
		PosOffsetX:  25,
		SpriteSheet: birdDroidSpriteTeleSlamRight,
		Fixed:       true,
	}

	anims[string(sr.PrimaryAttackKey)+"Left"] = &Animation{
		Name:        string(sr.PrimaryAttackKey) + "Left",
		FrameOX:     1278,
		FrameOY:     0,
		FrameWidth:  142,
		FrameHeight: 107,
		FrameCount:  9,
		PosOffsetX:  -19,
		SpriteSheet: birdDroidSpriteTeleSlamLeft,
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
		FrameWidth:  229,
		FrameHeight: 107,
		FrameCount:  16,

		SpriteSheet: birdDroidSpriteLazerShotRight,
		Fixed:       true,
	}

	a2alKey := string(sr.SecondaryAttackKey) + "Left"
	anims[a2alKey] = &Animation{
		Name:        a2alKey,
		FrameOX:     3664,
		FrameOY:     0,
		FrameWidth:  229,
		FrameHeight: 107,
		FrameCount:  16,
		PosOffsetX:  -44,
		SpriteSheet: birdDroidSpriteLazerShotLeft,
		Fixed:       true,
	}

	/*
		---------------------------------------------------------------------------------
		Secondary End
		---------------------------------------------------------------------------------
	*/

	return anims

}
