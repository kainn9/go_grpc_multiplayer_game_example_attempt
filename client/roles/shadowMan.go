package roles

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	utClient "github.com/kainn9/grpc_game/client_util"
	sr "github.com/kainn9/grpc_game/server/roles"
)

/*
File for shadowMan class
contains sprites/animation data
*/
var (
	shadowMan *Role = InitShadowMan()
)

var (
	shadowManSpriteIdleLeft  *ebiten.Image
	shadowManSpriteIdleRight *ebiten.Image

	shadowManSpriteWalkingRight *ebiten.Image
	shadowManSpriteWalkingLeft  *ebiten.Image

	shadowManSpriteJumpLeft  *ebiten.Image
	shadowManSpriteJumpRight *ebiten.Image

	shadowManSpriteHitRight *ebiten.Image
	shadowManSpriteHitLeft  *ebiten.Image

	shadowManSpriteKBRight *ebiten.Image
	shadowManSpriteKBLeft  *ebiten.Image

	shadowManSpriteCleaveAtkRight *ebiten.Image
	shadowManSpriteCleaveAtkLeft  *ebiten.Image

	shadowManSpriteFireAtkRight *ebiten.Image
	shadowManSpriteFireAtkLeft  *ebiten.Image

	shadowManSpriteDeathRight *ebiten.Image
	shadowManSpriteDeathLeft  *ebiten.Image
)

/*
Loads the default player sprites
*/
func LoadshadowManSprites() {
	shadowManSpriteIdleLeft = utClient.LoadImage("./sprites/shadowOfStorms/idleLeft.png")
	shadowManSpriteIdleRight = utClient.LoadImage("./sprites/shadowOfStorms/idle.png")

	shadowManSpriteWalkingRight = utClient.LoadImage("./sprites/shadowOfStorms/run.png")
	shadowManSpriteWalkingLeft = utClient.LoadImage("./sprites/shadowOfStorms/runLeftTest.png")

	shadowManSpriteJumpLeft = utClient.LoadImage("./sprites/shadowOfStorms/runLeft.png")
	shadowManSpriteJumpRight = utClient.LoadImage("./sprites/shadowOfStorms/run.png")

	shadowManSpriteHitRight = utClient.LoadImage("./sprites/shadowOfStorms/damageAndDeath.png")
	shadowManSpriteHitLeft = utClient.LoadImage("./sprites/shadowOfStorms/damageAndDeath.png")

	shadowManSpriteKBRight = utClient.LoadImage("./sprites/shadowOfStorms/damageAndDeath.png")
	shadowManSpriteKBLeft = utClient.LoadImage("./sprites/shadowOfStorms/damageAndDeath.png")

	shadowManSpriteCleaveAtkRight = utClient.LoadImage("./sprites/shadowOfStorms/attack2.png")
	shadowManSpriteCleaveAtkLeft = utClient.LoadImage("./sprites/shadowOfStorms/attack2Left.png")

	shadowManSpriteFireAtkRight = utClient.LoadImage("./sprites/shadowOfStorms/attack1.png")
	shadowManSpriteFireAtkLeft = utClient.LoadImage("./sprites/shadowOfStorms/attack1Left.png")

	shadowManSpriteDeathRight = utClient.LoadImage("./sprites/shadowOfStorms/attack1.png")
	shadowManSpriteDeathLeft = utClient.LoadImage("./sprites/shadowOfStorms/attack1.png")
}

func InitShadowMan() *Role {
	LoadshadowManSprites()
	fmt.Print("Hello World")
	r := &Role{
		Animations:    shadowManAnims(),
		HitBoxOffsetY: 60,
		HitBoxOffsetX: 30,
		Health:        sr.ShadowMan.Health,
		HitBoxW:       sr.ShadowMan.HitBoxW,
		HitBoxH:       sr.ShadowMan.HitBoxH,
		HealthBarOffset: &Offset{
			X: 20,
			Y: 40,
		},
		StatusEffectOffset: &Offset{
			X: 0,
			Y: -10,
		},
	}

	return r
}

func shadowManAnims() map[string]*Animation {
	fmt.Print("Hello World2")
	anims := make(map[string]*Animation)

	anims[string(IdleRight)] = &Animation{
		FrameOX:       0,
		FrameOY:       0,
		FrameWidth:    75,
		FrameHeight:   100,
		FrameCount:    1,
		SpriteSheet:   shadowManSpriteIdleRight,
		VerticalSheet: true,
	}

	anims[string(IdleLeft)] = &Animation{
		FrameOX:       55,
		FrameOY:       0,
		FrameWidth:    75,
		FrameHeight:   100,
		FrameCount:    1,
		SpriteSheet:   shadowManSpriteIdleLeft,
		VerticalSheet: true,
	}

	anims[string(WalkRight)] = &Animation{
		FrameOX:       0,
		FrameOY:       0,
		FrameWidth:    75,
		FrameHeight:   90,
		FrameCount:    6,
		PosOffsetX:    0,
		SpriteSheet:   shadowManSpriteWalkingRight,
		VerticalSheet: true,
	}

	anims[string(WalkLeft)] = &Animation{
		FrameOX:       0,
		FrameOY:       0,
		FrameWidth:    75,
		FrameHeight:   90,
		FrameCount:    6,
		PosOffsetX:    0,
		SpriteSheet:   shadowManSpriteWalkingLeft,
		VerticalSheet: true,
	}

	anims[string(JumpLeft)] = &Animation{
		FrameOX:       0,
		FrameOY:       0,
		FrameWidth:    75,
		FrameHeight:   90,
		FrameCount:    6,
		SpriteSheet:   shadowManSpriteWalkingLeft,
		VerticalSheet: true,
	}

	anims[string(JumpRight)] = &Animation{
		FrameOX:       0,
		FrameOY:       0,
		FrameWidth:    75,
		FrameHeight:   90,
		FrameCount:    6,
		SpriteSheet:   shadowManSpriteWalkingRight,
		VerticalSheet: true,
	}

	anims[string(HitRight)] = &Animation{
		Name:          string(HitRight),
		FrameOX:       0,
		FrameOY:       0,
		FrameWidth:    126,
		FrameHeight:   126,
		FrameCount:    1,
		PosOffsetX:    28,
		PosOffsetY:    16,
		SpriteSheet:   shadowManSpriteHitRight,
		Fixed:         true,
		VerticalSheet: true,
	}

	anims[string(HitLeft)] = &Animation{
		Name:          string(HitLeft),
		FrameOX:       756,
		FrameOY:       0,
		FrameWidth:    126,
		FrameHeight:   126,
		FrameCount:    1,
		PosOffsetX:    28,
		PosOffsetY:    16,
		SpriteSheet:   shadowManSpriteHitLeft,
		Fixed:         true,
		VerticalSheet: true,
	}

	stunAnimCopyRight := *anims[string(HitRight)]
	anims[string(StunRight)] = &stunAnimCopyRight

	stunAnimCopyLeft := *anims[string(HitLeft)]
	anims[string(StunLeft)] = &stunAnimCopyLeft

	anims[string(KbRight)] = &Animation{
		FrameOX:       0,
		FrameOY:       0,
		FrameWidth:    126,
		FrameHeight:   126,
		FrameCount:    1,
		PosOffsetX:    28,
		PosOffsetY:    16,
		SpriteSheet:   shadowManSpriteKBRight,
		VerticalSheet: true,
	}

	anims[string(KbLeft)] = &Animation{
		FrameOX:       756,
		FrameOY:       0,
		FrameWidth:    126,
		FrameHeight:   126,
		FrameCount:    1,
		PosOffsetX:    28,
		PosOffsetY:    16,
		SpriteSheet:   shadowManSpriteKBLeft,
		VerticalSheet: true,
	}

	anims[string(DeathRight)] = &Animation{
		Name:          string(DeathRight),
		FrameOX:       0,
		FrameOY:       0,
		FrameWidth:    147,
		FrameHeight:   138,
		FrameCount:    1,
		PosOffsetX:    28,
		PosOffsetY:    26,
		SpriteSheet:   shadowManSpriteDeathRight,
		Fixed:         true,
		VerticalSheet: true,
	}

	anims[string(DeathLeft)] = &Animation{
		Name:          string(DeathLeft),
		FrameOX:       3528,
		FrameOY:       0,
		FrameWidth:    147,
		FrameHeight:   138,
		FrameCount:    1,
		PosOffsetX:    28,
		PosOffsetY:    26,
		SpriteSheet:   shadowManSpriteDeathLeft,
		Fixed:         true,
		VerticalSheet: true,
	}

	/*
		---------------------------------------------------------------------------------
		Primary Attack
		---------------------------------------------------------------------------------
	*/
	anims[string(sr.PrimaryAttackKey)+"Right"] = &Animation{
		Name:          string(sr.PrimaryAttackKey) + "Right",
		FrameOX:       0,
		FrameOY:       0,
		FrameWidth:    200,
		FrameHeight:   90,
		FrameCount:    10,
		PosOffsetX:    10,
		PosOffsetY:    5,
		SpriteSheet:   shadowManSpriteCleaveAtkRight,
		Fixed:         true,
		VerticalSheet: true,
	}

	anims[string(sr.PrimaryAttackKey)+"Left"] = &Animation{
		Name:          string(sr.PrimaryAttackKey) + "Left",
		FrameOX:       0,
		FrameOY:       0,
		FrameWidth:    150,
		FrameHeight:   90,
		FrameCount:    10,
		PosOffsetX:    10,
		PosOffsetY:    5,
		SpriteSheet:   shadowManSpriteCleaveAtkLeft,
		Fixed:         true,
		VerticalSheet: true,
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
		Name:          a2arKey,
		FrameOX:       0,
		FrameOY:       0,
		FrameWidth:    200,
		FrameHeight:   90,
		FrameCount:    20,
		PosOffsetX:    0,
		PosOffsetY:    5,
		SpriteSheet:   shadowManSpriteFireAtkRight,
		Fixed:         true,
		VerticalSheet: true,
	}

	a2alKey := string(sr.SecondaryAttackKey) + "Left"
	anims[a2alKey] = &Animation{
		Name:          a2alKey,
		FrameOX:       0,
		FrameOY:       0,
		FrameWidth:    150,
		FrameHeight:   90,
		FrameCount:    20,
		PosOffsetX:    20,
		PosOffsetY:    5,
		SpriteSheet:   shadowManSpriteFireAtkLeft,
		Fixed:         true,
		VerticalSheet: true,
	}

	/*
		---------------------------------------------------------------------------------
		Secondary End
		---------------------------------------------------------------------------------
	*/
	fmt.Print("Hello World3")
	return anims

}
