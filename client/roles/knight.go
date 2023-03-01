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
	knightSpriteIdleLeft  *ebiten.Image
	knightSpriteIdleRight *ebiten.Image

	knightSpriteWalkingRight *ebiten.Image
	knightSpriteWalkingLeft  *ebiten.Image
	knightSpriteJumpLeft     *ebiten.Image
	knightSpriteJumpRight    *ebiten.Image
	knightSpriteAttackLeft   *ebiten.Image
	knightSpriteAttackRight  *ebiten.Image
	knightSpriteKBRight      *ebiten.Image
	knightSpriteKBLeft       *ebiten.Image
	knightSpriteWULeft   	*ebiten.Image
	knightSpriteWURight   	*ebiten.Image
	knightSpriteMVRight      	*ebiten.Image
	knightSpriteMVLeft   	*ebiten.Image
	knightSecondaryAttackLeft       *ebiten.Image
	knightSecondaryAttackRight       *ebiten.Image
)

/*
Loads the default player sprites
*/
func LoadKnightSprites() {
	knightSpriteIdleLeft = ut.LoadImg("./sprites/knight/knightIdleLeft.png")
	knightSpriteIdleRight = ut.LoadImg("./sprites/knight/knightIdleRight.png")

	knightSpriteWalkingRight = ut.LoadImg("./sprites/knight/knightRunningRight.png")
	knightSpriteWalkingLeft = ut.LoadImg("./sprites/knight/knightRunningLeft.png")
	knightSpriteJumpLeft = ut.LoadImg("./sprites/knight/knightJumpLeft.png")
	knightSpriteJumpRight = ut.LoadImg("./sprites/knight/knightJumpRight.png")
	knightSpriteAttackRight = ut.LoadImg("./sprites/knight/knightAttackRight.png")
	knightSpriteAttackLeft = ut.LoadImg("./sprites/knight/knightAttackLeft.png")
	knightSpriteKBRight = ut.LoadImg("./sprites/knight/knightKnockBackRight.png")
	knightSpriteKBLeft = ut.LoadImg("./sprites/knight/knightKnockBackLeft.png")
	knightSpriteMVLeft = ut.LoadImg("./sprites/knight/knightSecondaryMovementLeft.png")
	knightSpriteMVRight = ut.LoadImg("./sprites/knight/knightSecondaryMovementRight.png")
	knightSpriteWULeft = ut.LoadImg("./sprites/knight/knightSecondaryWindupLeft.png")
	knightSpriteWURight = ut.LoadImg("./sprites/knight/knightSecondaryWindupRight.png")
	knightSecondaryAttackLeft =  ut.LoadImg("./sprites/knight/knightSecondarySlashAttackLeft.png")
	knightSecondaryAttackRight =  ut.LoadImg("./sprites/knight/knightSecondarySlashAttackRight.png")
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

	anims[string(sr.PrimaryAttackKey)+"Right"] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  65,
		FrameHeight: 48,
		FrameCount:  4,
		SpriteSheet: knightSpriteAttackRight,
	}

	anims[string(sr.PrimaryAttackKey)+"Left"] = &Animation{
		FrameOX:     260,
		FrameOY:     0,
		FrameWidth:  65,
		FrameHeight: 48,
		FrameCount:  4,
		SpriteSheet: knightSpriteAttackLeft,
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


	
	a2mlKey := string(sr.SecondaryAttackKey) + "MovementLeft"
	anims[a2mlKey] = &Animation{
		Name: a2mlKey,
		FrameOX:     440,
		FrameOY:     0,
		FrameWidth:  40,
		FrameHeight: 48,
		FrameCount:  11,
		SpriteSheet: knightSpriteMVLeft,
		Fixed: true,
		
	}
	a2mrKey := string(sr.SecondaryAttackKey) + "MovementRight"
	anims[a2mrKey] = &Animation{
		Name: a2mrKey,
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  40,
		FrameHeight: 48,
		FrameCount:  11,
		SpriteSheet: knightSpriteMVRight,
		Fixed: true,
		
	}

	a2wurKey := string(sr.SecondaryAttackKey) + "WindupRight"
	anims[a2wurKey] = &Animation{
		Name: a2wurKey,
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  70,
		FrameHeight: 48,
		FrameCount:  11,
		PosOffsetX: 14,
		SpriteSheet: knightSpriteWURight,
		Fixed: true,
		
	}
	a2wulKey := string(sr.SecondaryAttackKey) + "WindupLeft"
	anims[a2wulKey] = &Animation{
		Name: a2wulKey,
		FrameOX:     770,
		FrameOY:     0,
		FrameWidth:  70,
		FrameHeight: 48,
		FrameCount:  11,
		PosOffsetX: 14,
		SpriteSheet: knightSpriteWULeft,
		Fixed: true,
		
	}

	a2arKey := "secondaryAtkRight"
	anims[a2arKey] = &Animation{
		Name: a2arKey,
		FrameOX:     0,
		FrameOY:     10,
		FrameWidth:  75,
		FrameHeight: 60,
		FrameCount:  12,
		SpriteSheet: knightSecondaryAttackRight,
		Fixed: true,
		
	}

	a2alKey := "secondaryAtkLeft"
	anims[a2alKey] = &Animation{
		Name: a2alKey,
		FrameOX:     900,
		FrameOY:     10,
		FrameWidth:  75,
		FrameHeight: 60,
		FrameCount:  12,
		SpriteSheet: knightSecondaryAttackLeft,
		Fixed: true,
	}



	a3alKey := "tertAtkLeft"
	a3arKey := "tertAtkRight"
	a3mlKey := string(sr.TertAttackKey) + "MovementLeft"
	a3mrKey := string(sr.TertAttackKey) + "MovementRight"
	a3wulKey := string(sr.TertAttackKey) + "WindupLeft"
	a3wurKey := string(sr.TertAttackKey) + "WindupRight"
	
	anims[a3arKey] = anims[a2arKey]
	anims[a3arKey].Name = a3arKey

	anims[a3alKey] = anims[a2alKey]
	anims[a3alKey].Name = a3alKey
	
	anims[a3wulKey] = anims[a2wulKey]
	anims[a3wulKey].Name = a3wulKey

	anims[a3wurKey] = anims[a2wurKey]
	anims[a3wurKey].Name = a3wurKey

		
	anims[a3mrKey] = anims[a2mrKey]
	anims[a3mrKey].Name = a3mrKey

	anims[a3mlKey] = anims[a2mlKey]
	anims[a3mlKey].Name = a3mlKey


	return anims

}
