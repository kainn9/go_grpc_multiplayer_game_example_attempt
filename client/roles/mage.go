package roles

import (
	"github.com/hajimehoshi/ebiten/v2"
	utClient "github.com/kainn9/grpc_game/client_util"
	sr "github.com/kainn9/grpc_game/server/roles"
)

/*
File for Mage class
contains sprites/animation data
*/
var (
	Mage *Role = InitMage()
)

var (
	mageSpriteIdleLeft  *ebiten.Image
	mageSpriteIdleRight *ebiten.Image

	mageSpriteWalkingRight *ebiten.Image
	mageSpriteWalkingLeft  *ebiten.Image

	mageSpriteJumpLeft  *ebiten.Image
	mageSpriteJumpRight *ebiten.Image

	mageSpriteHitRight *ebiten.Image
	mageSpriteHitLeft  *ebiten.Image

	mageSpriteKBRight *ebiten.Image
	mageSpriteKBLeft  *ebiten.Image

	mageSpriteDeathRight *ebiten.Image
	mageSpriteDeathLeft  *ebiten.Image

	mageSpriteLightSwordsRight *ebiten.Image
	mageSpriteLightSwordsLeft  *ebiten.Image

	mageSpriteFireBlastRight *ebiten.Image
	mageSpriteFireBlastLeft  *ebiten.Image

	mageSpriteIceBlastRight *ebiten.Image
	mageSpriteIceBlastLeft  *ebiten.Image

	mageSpriteDefenseRight *ebiten.Image
	mageSpriteDefenseLeft  *ebiten.Image
)

/*
Loads the default player sprites
*/
func LoadMageSprites() {
	mageSpriteIdleLeft = utClient.LoadImage("./sprites/mage/mageIdleLeft.png")
	mageSpriteIdleRight = utClient.LoadImage("./sprites/mage/mageIdleRight.png")

	mageSpriteWalkingRight = utClient.LoadImage("./sprites/mage/mageRunningRight.png")
	mageSpriteWalkingLeft = utClient.LoadImage("./sprites/mage/mageRunningLeft.png")

	mageSpriteJumpLeft = utClient.LoadImage("./sprites/mage/mageJumpLeft.png")
	mageSpriteJumpRight = utClient.LoadImage("./sprites/mage/mageJumpRight.png")

	mageSpriteHitRight = utClient.LoadImage("./sprites/mage/mageHitRight.png")
	mageSpriteHitLeft = utClient.LoadImage("./sprites/mage/mageHitRight.png")

	mageSpriteKBRight = utClient.LoadImage("./sprites/mage/mageKnockBackRight.png")
	mageSpriteKBLeft = utClient.LoadImage("./sprites/mage/mageKnockBackLeft.png")

	mageSpriteDeathRight = utClient.LoadImage("./sprites/mage/mageDeathRight.png")
	mageSpriteDeathLeft = utClient.LoadImage("./sprites/mage/mageDeathLeft.png")

	mageSpriteLightSwordsRight = utClient.LoadImage("./sprites/mage/mageLightSwordsRight.png")
	mageSpriteLightSwordsLeft = utClient.LoadImage("./sprites/mage/mageLightSwordsLeft.png")

	mageSpriteDefenseRight = utClient.LoadImage("./sprites/mage/mageDefenseRight.png")
	mageSpriteDefenseLeft = utClient.LoadImage("./sprites/mage/mageDefenseLeft.png")

	mageSpriteFireBlastRight = utClient.LoadImage("./sprites/mage/mageFireBlastRight.png")
	mageSpriteFireBlastLeft = utClient.LoadImage("./sprites/mage/mageFireBlastLeft.png")

	mageSpriteIceBlastRight = utClient.LoadImage("./sprites/mage/mageIceBlastRight.png")
	mageSpriteIceBlastLeft = utClient.LoadImage("./sprites/mage/mageIceBlastLeft.png")

}

func InitMage() *Role {
	LoadMageSprites()

	r := &Role{
		Animations:    MageAnims(),
		HitBoxOffsetY: 52,
		HitBoxOffsetX: 88,
		Health:        sr.Mage.Health,
		HitBoxW:       sr.Mage.HitBoxW,
		HitBoxH:       sr.Mage.HitBoxH,
		AttackCount:   len(sr.Mage.Attacks),
		HasDefense:    sr.Mage.Defense != nil,
		HealthBarOffset: &Offset{
			X: 75,
			Y: 28,
		},
		StatusEffectOffset: &Offset{
			X: 75,
			Y: 34,
		},
	}

	return r
}

func MageAnims() map[string]*Animation {
	anims := make(map[string]*Animation)

	anims[string(IdleRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  8,
		SpriteSheet: mageSpriteIdleRight,
	}

	anims[string(IdleLeft)] = &Animation{
		FrameOX:     1536,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  8,
		SpriteSheet: mageSpriteIdleLeft,
	}

	anims[string(WalkRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  10,
		SpriteSheet: mageSpriteWalkingRight,
	}

	anims[string(WalkLeft)] = &Animation{
		FrameOX:     1920,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  10,
		SpriteSheet: mageSpriteWalkingLeft,
	}

	anims[string(JumpLeft)] = &Animation{
		FrameOX:     576,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  3,
		SpriteSheet: mageSpriteJumpLeft,
	}

	anims[string(JumpRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  3,
		SpriteSheet: mageSpriteJumpRight,
	}

	anims[string(HitRight)] = &Animation{
		Name:        string(HitRight),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  6,
		SpriteSheet: mageSpriteHitRight,
		Fixed:       true,
	}

	anims[string(HitLeft)] = &Animation{
		Name:        string(HitLeft),
		FrameOX:     1152,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  6,
		SpriteSheet: mageSpriteHitLeft,
		Fixed:       true,
	}

	stunAnimCopyRight := *anims[string(HitRight)]
	anims[string(StunRight)] = &stunAnimCopyRight

	stunAnimCopyLeft := *anims[string(HitLeft)]
	anims[string(StunLeft)] = &stunAnimCopyLeft

	anims[string(KbRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  3,
		SpriteSheet: mageSpriteKBRight,
	}

	anims[string(KbLeft)] = &Animation{
		FrameOX:     576,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  3,
		SpriteSheet: mageSpriteKBLeft,
	}

	anims[string(DeathRight)] = &Animation{
		Name:        string(DeathRight),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  24,
		SpriteSheet: mageSpriteDeathRight,
		Fixed:       true,
	}

	anims[string(DeathLeft)] = &Animation{
		Name:        string(DeathLeft),
		FrameOX:     4608,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  24,
		SpriteSheet: mageSpriteDeathLeft,
		Fixed:       true,
	}

	anims[string(DefenseRight)] = &Animation{
		Name:        string(DefenseRight),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  7,
		SpriteSheet: mageSpriteDefenseRight,
		Fixed:       true,
	}

	anims[string(DefenseLeft)] = &Animation{
		Name:        string(DefenseLeft),
		FrameOX:     1344,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  7,
		SpriteSheet: mageSpriteDefenseLeft,
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
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  21,
		SpriteSheet: mageSpriteLightSwordsRight,
		Fixed:       true,
	}

	anims[string(sr.PrimaryAttackKey)+"Left"] = &Animation{
		Name:        string(sr.PrimaryAttackKey) + "Left",
		FrameOX:     4032,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  21,
		SpriteSheet: mageSpriteLightSwordsLeft,
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
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  12,
		SpriteSheet: mageSpriteFireBlastRight,
		Fixed:       true,
	}

	a2alKey := string(sr.SecondaryAttackKey) + "Left"
	anims[a2alKey] = &Animation{
		Name:        a2alKey,
		FrameOX:     2304,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  12,
		SpriteSheet: mageSpriteFireBlastLeft,
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

	a3mlKey := string(sr.TertAttackKey) + "MovementLeft"
	anims[a3mlKey] = &Animation{
		Name:        a3mlKey,
		FrameOX:     192,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  1,
		SpriteSheet: mageSpriteIceBlastLeft,
		Fixed:       true,
	}

	a3mrKey := string(sr.TertAttackKey) + "MovementRight"
	anims[a3mrKey] = &Animation{
		Name:        a3mrKey,
		FrameOX:     2112,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  1,
		SpriteSheet: mageSpriteIceBlastRight,
		Fixed:       true,
	}

	a3arKey := "tertAtkRight"
	anims[a3arKey] = &Animation{
		Name:        a3arKey,
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  12,
		SpriteSheet: mageSpriteIceBlastRight,
		Fixed:       true,
	}

	a3alKey := "tertAtkLeft"
	anims[a3alKey] = &Animation{
		Name:        a3alKey,
		FrameOX:     2304,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  12,
		SpriteSheet: mageSpriteIceBlastLeft,
		Fixed:       true,
	}

	/*
		---------------------------------------------------------------------------------
		Tert Attack END
		---------------------------------------------------------------------------------
	*/

	return anims

}
