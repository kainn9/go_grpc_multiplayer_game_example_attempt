package roles

import (
	"github.com/hajimehoshi/ebiten/v2"
	utClient "github.com/kainn9/grpc_game/client_util"
	sr "github.com/kainn9/grpc_game/server/roles"
)

/*
File for Knight class
contains sprites/animation data
*/

var (
	Knight *Role = InitKnight()
)

var (
	knightSpriteIdleLeft  *ebiten.Image
	knightSpriteIdleRight *ebiten.Image

	knightSpriteWalkingRight *ebiten.Image
	knightSpriteWalkingLeft  *ebiten.Image

	knightSpriteJumpLeft  *ebiten.Image
	knightSpriteJumpRight *ebiten.Image

	knightSpriteStabLeft  *ebiten.Image
	knightSpriteStabRight *ebiten.Image

	knightSpriteKBRight *ebiten.Image
	knightSpriteKBLeft  *ebiten.Image

	knightSpriteDashSlashWULeft  *ebiten.Image
	knightSpriteDashSlashWURight *ebiten.Image

	knightSpriteDashSlashMVRight *ebiten.Image
	knightSpriteDashSlashMVLeft  *ebiten.Image

	knightSpriteDashSlashLeft  *ebiten.Image
	knightSpriteDashSlashRight *ebiten.Image

	knightSpriteQuickSlashLeft  *ebiten.Image
	knightSpriteQuickSlashRight *ebiten.Image

	knightSpriteQuickSlashWindupLeft  *ebiten.Image
	knightSpriteQuickSlashWindupRight *ebiten.Image

	knightSpriteSlideRight *ebiten.Image
	knightSpriteSlideLeft  *ebiten.Image

	knightSpriteDeathRight *ebiten.Image
	knightSpriteDeathLeft  *ebiten.Image
)

/*
Loads the default player sprites
*/
func LoadKnightSprites() {
	knightSpriteIdleLeft = utClient.LoadImage("./sprites/knight/knightIdleLeft.png")
	knightSpriteIdleRight = utClient.LoadImage("./sprites/knight/knightIdleRight.png")

	knightSpriteWalkingRight = utClient.LoadImage("./sprites/knight/knightRunningRight.png")
	knightSpriteWalkingLeft = utClient.LoadImage("./sprites/knight/knightRunningLeft.png")

	knightSpriteJumpLeft = utClient.LoadImage("./sprites/knight/knightJumpLeft.png")
	knightSpriteJumpRight = utClient.LoadImage("./sprites/knight/knightJumpRight.png")

	knightSpriteStabRight = utClient.LoadImage("./sprites/knight/knightStabRight.png")
	knightSpriteStabLeft = utClient.LoadImage("./sprites/knight/knightStabLeft.png")

	knightSpriteKBRight = utClient.LoadImage("./sprites/knight/knightKnockBackRight.png")
	knightSpriteKBLeft = utClient.LoadImage("./sprites/knight/knightKnockBackLeft.png")

	knightSpriteDashSlashMVLeft = utClient.LoadImage("./sprites/knight/knightDashSlashMovementLeft.png")
	knightSpriteDashSlashMVRight = utClient.LoadImage("./sprites/knight/knightDashSlashMovementRight.png")

	knightSpriteDashSlashWULeft = utClient.LoadImage("./sprites/knight/knightDashSlashWindupLeft.png")
	knightSpriteDashSlashWURight = utClient.LoadImage("./sprites/knight/knightDashSlashWindupRight.png")

	knightSpriteDashSlashLeft = utClient.LoadImage("./sprites/knight/knightDashSlashLeft.png")
	knightSpriteDashSlashRight = utClient.LoadImage("./sprites/knight/knightDashSlashRight.png")

	knightSpriteQuickSlashRight = utClient.LoadImage("./sprites/knight/knightQuickSlashRight.png")
	knightSpriteQuickSlashLeft = utClient.LoadImage("./sprites/knight/knightQuickSlashLeft.png")

	knightSpriteSlideRight = utClient.LoadImage("./sprites/knight/knightSlideRight.png")
	knightSpriteSlideLeft = utClient.LoadImage("./sprites/knight/knightSlideLeft.png")

	knightSpriteQuickSlashWindupLeft = utClient.LoadImage("./sprites/knight/knightQuickSlashWindupLeft.png")
	knightSpriteQuickSlashWindupRight = utClient.LoadImage("./sprites/knight/knightQuickSlashWindupRight.png")

	knightSpriteDeathRight = utClient.LoadImage("./sprites/knight/knightDeathRight.png")
	knightSpriteDeathLeft = utClient.LoadImage("./sprites/knight/knightDeathLeft.png")
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

// TODO MAKE ANIM KEYS CONSTS
func KnightAnims() map[string]*Animation {
	anims := make(map[string]*Animation)

	anims["idleRight"] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  32,
		FrameHeight: 48,
		FrameCount:  8,
		SpriteSheet: knightSpriteIdleRight,
	}

	anims["idleLeft"] = &Animation{
		FrameOX:     256,
		FrameOY:     0,
		FrameWidth:  32,
		FrameHeight: 48,
		FrameCount:  8,
		SpriteSheet: knightSpriteIdleLeft,
	}

	anims["walkRight"] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 48,
		FrameCount:  8,
		SpriteSheet: knightSpriteWalkingRight,
	}

	anims["walkLeft"] = &Animation{
		FrameOX:     280,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 48,
		FrameCount:  8,
		SpriteSheet: knightSpriteWalkingLeft,
	}

	anims["jumpLeft"] = &Animation{
		FrameOX:     35,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 48,
		FrameCount:  1,
		SpriteSheet: knightSpriteJumpLeft,
	}

	anims["jumpRight"] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 48,
		FrameCount:  1,
		SpriteSheet: knightSpriteJumpRight,
	}

	anims["defenseRight"] = &Animation{
		Name:        "defenseRight",
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  50,
		FrameHeight: 48,
		FrameCount:  20,
		SpriteSheet: knightSpriteSlideRight,
		Fixed:       true,
	}

	anims["defenseLeft"] = &Animation{
		Name:        "defenseLeft",
		FrameOX:     1000,
		FrameOY:     0,
		FrameWidth:  50,
		FrameHeight: 48,
		FrameCount:  20,
		SpriteSheet: knightSpriteSlideLeft,
		Fixed:       true,
	}

	anims["KbRight"] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  30,
		FrameHeight: 48,
		FrameCount:  4,
		SpriteSheet: knightSpriteKBRight,
	}

	anims["KbLeft"] = &Animation{
		FrameOX:     120,
		FrameOY:     0,
		FrameWidth:  30,
		FrameHeight: 32,
		FrameCount:  4,
		SpriteSheet: knightSpriteKBLeft,
	}

	anims[string(DeathRight)] = &Animation{
		Name:        string(DeathLeft),
		FrameOX:     0,
		FrameOY:     14,
		FrameWidth:  70,
		FrameHeight: 70,
		FrameCount:  24,
		SpriteSheet: knightSpriteDeathRight,
		Fixed:       true,
	}

	anims[string(DeathLeft)] = &Animation{
		Name:        string(DeathLeft),
		FrameOX:     1680,
		FrameOY:     14,
		FrameWidth:  70,
		FrameHeight: 70,
		FrameCount:  24,
		SpriteSheet: knightSpriteDeathLeft,
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
		FrameWidth:  65,
		FrameHeight: 48,
		FrameCount:  4,
		SpriteSheet: knightSpriteStabRight,
		Fixed:       true,
	}

	anims[string(sr.PrimaryAttackKey)+"Left"] = &Animation{
		Name:        "primaryAtkleft",
		FrameOX:     260,
		FrameOY:     0,
		FrameWidth:  65,
		FrameHeight: 48,
		FrameCount:  4,
		SpriteSheet: knightSpriteStabLeft,
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
		FrameOY:     20,
		FrameWidth:  70,
		FrameHeight: 70,
		FrameCount:  5,
		PosOffsetX:  14,
		SpriteSheet: knightSpriteQuickSlashRight,
		Fixed:       true,
	}

	a2alKey := string(sr.SecondaryAttackKey) + "Left"
	anims[a2alKey] = &Animation{
		Name:        a2alKey,
		FrameOX:     350,
		FrameOY:     20,
		FrameWidth:  70,
		FrameHeight: 70,
		FrameCount:  5,
		PosOffsetX:  14,
		SpriteSheet: knightSpriteQuickSlashLeft,
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
		FrameOX:     440,
		FrameOY:     0,
		FrameWidth:  40,
		FrameHeight: 48,
		FrameCount:  11,
		SpriteSheet: knightSpriteDashSlashMVLeft,
		Fixed:       true,
	}

	a3mrKey := string(sr.TertAttackKey) + "MovementRight"
	anims[a3mrKey] = &Animation{
		Name:        a3mrKey,
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  40,
		FrameHeight: 48,
		FrameCount:  11,
		SpriteSheet: knightSpriteDashSlashMVRight,
		Fixed:       true,
	}

	a3wurKey := string(sr.TertAttackKey) + "WindupRight"
	anims[a3wurKey] = &Animation{
		Name:        a3wurKey,
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  70,
		FrameHeight: 48,
		FrameCount:  11,
		PosOffsetX:  14,
		SpriteSheet: knightSpriteDashSlashWURight,
		Fixed:       true,
	}

	a3wulKey := string(sr.TertAttackKey) + "WindupLeft"
	anims[a3wulKey] = &Animation{
		Name:        a3wulKey,
		FrameOX:     770,
		FrameOY:     0,
		FrameWidth:  70,
		FrameHeight: 48,
		FrameCount:  11,
		PosOffsetX:  14,
		SpriteSheet: knightSpriteDashSlashWULeft,
		Fixed:       true,
	}

	a3arKey := "tertAtkRight"
	anims[a3arKey] = &Animation{
		Name:        a3arKey,
		FrameOX:     0,
		FrameOY:     10,
		FrameWidth:  75,
		FrameHeight: 60,
		FrameCount:  12,
		SpriteSheet: knightSpriteDashSlashRight,
		Fixed:       true,
	}

	a3alKey := "tertAtkLeft"
	anims[a3alKey] = &Animation{
		Name:        a3alKey,
		FrameOX:     900,
		FrameOY:     10,
		FrameWidth:  75,
		FrameHeight: 60,
		FrameCount:  12,
		SpriteSheet: knightSpriteDashSlashLeft,
		Fixed:       true,
	}

	/*
		---------------------------------------------------------------------------------
		Tert Attack END
		---------------------------------------------------------------------------------
	*/

	/*
		---------------------------------------------------------------------------------
		Quaternary Attack
		---------------------------------------------------------------------------------
	*/

	a4wurKey := string(sr.QuaternaryAttackKey) + "WindupRight"
	anims[a4wurKey] = &Animation{
		Name:        a4wurKey,
		FrameOX:     0,
		FrameOY:     20,
		FrameWidth:  70,
		FrameHeight: 70,
		FrameCount:  12,
		PosOffsetX:  14,
		SpriteSheet: knightSpriteQuickSlashWindupRight,
		Fixed:       true,
	}

	a4wulKey := string(sr.QuaternaryAttackKey) + "WindupLeft"
	anims[a4wulKey] = &Animation{
		Name:        a4wulKey,
		FrameOX:     840,
		FrameOY:     20,
		FrameWidth:  70,
		FrameHeight: 70,
		FrameCount:  12,
		PosOffsetX:  14,
		SpriteSheet: knightSpriteQuickSlashWindupLeft,
		Fixed:       true,
	}

	a4arKey := string(sr.QuaternaryAttackKey) + "Right"
	anims[a4arKey] = &Animation{
		Name:        a4arKey,
		FrameOX:     0,
		FrameOY:     20,
		FrameWidth:  70,
		FrameHeight: 70,
		FrameCount:  5,
		PosOffsetX:  14,
		SpriteSheet: knightSpriteQuickSlashRight,
		Fixed:       true,
	}

	a4alKey := string(sr.QuaternaryAttackKey) + "Left"
	anims[a4alKey] = &Animation{
		Name:        a4alKey,
		FrameOX:     350,
		FrameOY:     20,
		FrameWidth:  70,
		FrameHeight: 70,
		FrameCount:  5,
		PosOffsetX:  14,
		SpriteSheet: knightSpriteQuickSlashLeft,
		Fixed:       true,
	}

	/*
		---------------------------------------------------------------------------------
		Quaternary End
		---------------------------------------------------------------------------------
	*/

	return anims

}
