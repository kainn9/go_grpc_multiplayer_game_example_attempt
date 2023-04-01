package roles

import (
	"github.com/hajimehoshi/ebiten/v2"
	sr "github.com/kainn9/grpc_game/server/roles"
	ut "github.com/kainn9/grpc_game/util"
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
	mageSpriteIdleLeft = ut.LoadImg("./sprites/mage/mageIdleLeft.png")
	mageSpriteIdleRight = ut.LoadImg("./sprites/mage/mageIdleRight.png")

	mageSpriteWalkingRight = ut.LoadImg("./sprites/mage/mageRunningRight.png")
	mageSpriteWalkingLeft = ut.LoadImg("./sprites/mage/mageRunningLeft.png")

	mageSpriteJumpLeft = ut.LoadImg("./sprites/mage/mageJumpLeft.png")
	mageSpriteJumpRight = ut.LoadImg("./sprites/mage/mageJumpRight.png")

	mageSpriteKBRight = ut.LoadImg("./sprites/mage/mageKnockBackRight.png")
	mageSpriteKBLeft = ut.LoadImg("./sprites/mage/mageKnockBackLeft.png")

	mageSpriteDeathRight = ut.LoadImg("./sprites/mage/mageDeathRight.png")
	mageSpriteDeathLeft = ut.LoadImg("./sprites/mage/mageDeathLeft.png")

	mageSpriteLightSwordsRight = ut.LoadImg("./sprites/mage/mageLightSwordsRight.png")
	mageSpriteLightSwordsLeft = ut.LoadImg("./sprites/mage/mageLightSwordsLeft.png")

	mageSpriteDefenseRight = ut.LoadImg("./sprites/mage/mageDefenseRight.png")
	mageSpriteDefenseLeft = ut.LoadImg("./sprites/mage/mageDefenseLeft.png")

	mageSpriteFireBlastRight = ut.LoadImg("./sprites/mage/mageFireBlastRight.png")
	mageSpriteFireBlastLeft = ut.LoadImg("./sprites/mage/mageFireBlastLeft.png")

	mageSpriteIceBlastRight = ut.LoadImg("./sprites/mage/mageIceBlastRight.png")
	mageSpriteIceBlastLeft = ut.LoadImg("./sprites/mage/mageIceBlastLeft.png")

}

func InitMage() *Role {
	LoadMageSprites()

	r := &Role{
		RoleType:      MageType,
		Animations:    MageAnims(),
		HitBoxOffsetY: 52,
		HitBoxOffsetX: 88,
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
		Name:        "primaryAtkRight",
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  192,
		FrameHeight: 112,
		FrameCount:  21,
		SpriteSheet: mageSpriteLightSwordsRight,
		Fixed:       true,
	}

	anims[string(sr.PrimaryAttackKey)+"Left"] = &Animation{
		Name:        "primaryAtkleft",
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
