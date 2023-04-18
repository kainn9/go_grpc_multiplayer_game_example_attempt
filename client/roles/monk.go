package roles

import (
	"github.com/hajimehoshi/ebiten/v2"
	utClient "github.com/kainn9/grpc_game/client_util"
	sr "github.com/kainn9/grpc_game/server/roles"
)

/*
File for Monk class
contains sprites/animation data
*/
var (
	Monk *Role = InitMonk()
)

var (
	monkSpriteIdleLeft  *ebiten.Image
	monkSpriteIdleRight *ebiten.Image

	monkSpriteWalkingRight *ebiten.Image
	monkSpriteWalkingLeft  *ebiten.Image

	monkSpriteJumpLeft  *ebiten.Image
	monkSpriteJumpRight *ebiten.Image

	monkDefenseLeft  *ebiten.Image
	monkDefenseRight *ebiten.Image

	monkSpriteHitRight *ebiten.Image
	monkSpriteHitLeft  *ebiten.Image

	monkSpriteKBRight *ebiten.Image
	monkSpriteKBLeft  *ebiten.Image

	monkSpriteSmashAtkRight *ebiten.Image
	monkSpriteSmashAtkLeft  *ebiten.Image

	monkSpriteEarthFistSmashRight *ebiten.Image
	monkSpriteEarthFistSmashLeft  *ebiten.Image

	monkSpriteDeathRight *ebiten.Image
	monkSpriteDeathLeft  *ebiten.Image
)

/*
Loads the default player sprites
*/
func LoadMonkSprites() {
	monkSpriteIdleLeft = utClient.LoadImage("./sprites/monk/monkIdleLeft.png")
	monkSpriteIdleRight = utClient.LoadImage("./sprites/monk/monkIdleRight.png")

	monkSpriteWalkingRight = utClient.LoadImage("./sprites/monk/monkRunningRight.png")
	monkSpriteWalkingLeft = utClient.LoadImage("./sprites/monk/monkRunningLeft.png")

	monkSpriteJumpLeft = utClient.LoadImage("./sprites/monk/monkJumpLeft.png")
	monkSpriteJumpRight = utClient.LoadImage("./sprites/monk/monkJumpRight.png")

	monkDefenseRight = utClient.LoadImage("./sprites/monk/monkDefenseRight.png")
	monkDefenseLeft = utClient.LoadImage("./sprites/monk/monkDefenseLeft.png")

	monkSpriteHitRight = utClient.LoadImage("./sprites/monk/monkHitRight.png")
	monkSpriteHitLeft = utClient.LoadImage("./sprites/monk/monkHitLeft.png")

	monkSpriteKBRight = utClient.LoadImage("./sprites/monk/monkKnockBackRight.png")
	monkSpriteKBLeft = utClient.LoadImage("./sprites/monk/monkKnockBackLeft.png")

	monkSpriteSmashAtkRight = utClient.LoadImage("./sprites/monk/monkSmashAttackRight.png")
	monkSpriteSmashAtkLeft = utClient.LoadImage("./sprites/monk/monkSmashAttackLeft.png")

	monkSpriteEarthFistSmashRight = utClient.LoadImage("./sprites/monk/monkEarthFistSmashRight.png")
	monkSpriteEarthFistSmashLeft = utClient.LoadImage("./sprites/monk/monkEarthFistSmashLeft.png")

	monkSpriteDeathRight = utClient.LoadImage("./sprites/monk/monkDeathRight.png")
	monkSpriteDeathLeft = utClient.LoadImage("./sprites/monk/monkDeathLeft.png")
}

func InitMonk() *Role {
	LoadMonkSprites()

	r := &Role{
		Animations:    MonkAnims(),
		HitBoxOffsetY: 4,
		HitBoxOffsetX: 4,
		Health:        sr.Monk.Health,
		HitBoxW:       sr.Monk.HitBoxW,
		HitBoxH:       sr.Monk.HitBoxH,
		HealthBarOffset: &Offset{
			X: -7,
			Y: -10,
		},
		StatusEffectOffset: &Offset{
			X: 0,
			Y: -6,
		},
	}

	return r
}

func MonkAnims() map[string]*Animation {
	anims := make(map[string]*Animation)

	anims[string(IdleRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  30,
		FrameHeight: 38,
		FrameCount:  6,
		SpriteSheet: monkSpriteIdleRight,
	}

	anims[string(IdleLeft)] = &Animation{
		FrameOX:     180,
		FrameOY:     0,
		FrameWidth:  30,
		FrameHeight: 38,
		FrameCount:  6,
		SpriteSheet: monkSpriteIdleLeft,
	}

	anims[string(WalkRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  33,
		FrameHeight: 62,
		FrameCount:  8,
		PosOffsetY:  14,
		SpriteSheet: monkSpriteWalkingRight,
	}

	anims[string(WalkLeft)] = &Animation{
		FrameOX:     264,
		FrameOY:     0,
		FrameWidth:  33,
		FrameHeight: 62,
		FrameCount:  8,
		PosOffsetY:  14,
		SpriteSheet: monkSpriteWalkingLeft,
	}

	anims[string(JumpLeft)] = &Animation{
		FrameOX:     105,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 49,
		FrameCount:  3,
		SpriteSheet: monkSpriteJumpLeft,
	}

	anims[string(JumpRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 49,
		FrameCount:  3,
		SpriteSheet: monkSpriteJumpRight,
	}

	anims[string(DefenseRight)] = &Animation{
		Name:        string(DefenseRight),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  45,
		FrameHeight: 48,
		FrameCount:  8,
		PosOffsetX:  10,
		PosOffsetY:  7,
		SpriteSheet: monkDefenseRight,
		Fixed:       true,
	}

	anims[string(DefenseLeft)] = &Animation{
		Name:        string(DefenseLeft),
		FrameOX:     360,
		FrameOY:     0,
		FrameWidth:  45,
		FrameHeight: 48,
		FrameCount:  8,
		PosOffsetX:  10,
		PosOffsetY:  7,
		SpriteSheet: monkDefenseLeft,
		Fixed:       true,
	}

	anims[string(HitRight)] = &Animation{
		Name:        string(HitRight),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  57,
		FrameHeight: 44,
		FrameCount:  6,
		PosOffsetX:  10,
		SpriteSheet: monkSpriteHitRight,
		Fixed:       true,
	}

	anims[string(HitLeft)] = &Animation{
		Name:        string(HitLeft),
		FrameOX:     342,
		FrameOY:     0,
		FrameWidth:  57,
		FrameHeight: 44,
		FrameCount:  6,
		PosOffsetX:  10,
		SpriteSheet: monkSpriteHitLeft,
		Fixed:       true,
	}

	stunAnimCopyRight := *anims[string(HitRight)]
	anims[string(StunRight)] = &stunAnimCopyRight

	stunAnimCopyLeft := *anims[string(HitLeft)]
	anims[string(StunLeft)] = &stunAnimCopyLeft

	anims[string(KbRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  50,
		FrameHeight: 39,
		FrameCount:  3,
		SpriteSheet: monkSpriteKBRight,
	}

	anims[string(KbLeft)] = &Animation{
		FrameOX:     150,
		FrameOY:     0,
		FrameWidth:  50,
		FrameHeight: 39,
		FrameCount:  3,
		SpriteSheet: monkSpriteKBLeft,
	}

	anims[string(DeathRight)] = &Animation{
		Name:        string(DeathRight),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  26,
		FrameHeight: 42,
		FrameCount:  24,
		SpriteSheet: monkSpriteDeathRight,
		Fixed:       true,
	}

	anims[string(DeathLeft)] = &Animation{
		Name:        string(DeathLeft),
		FrameOX:     624,
		FrameOY:     0,
		FrameWidth:  26,
		FrameHeight: 42,
		FrameCount:  24,
		SpriteSheet: monkSpriteDeathLeft,
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
		FrameWidth:  111,
		FrameHeight: 42,
		FrameCount:  13,
		SpriteSheet: monkSpriteSmashAtkRight,
		Fixed:       true,
	}

	anims[string(sr.PrimaryAttackKey)+"Left"] = &Animation{
		Name:        string(sr.PrimaryAttackKey) + "Left",
		FrameOX:     1443,
		FrameOY:     0,
		FrameWidth:  111,
		FrameHeight: 42,
		FrameCount:  13,
		SpriteSheet: monkSpriteSmashAtkLeft,
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
		FrameWidth:  89,
		FrameHeight: 55,
		FrameCount:  8,
		PosOffsetY:  15,
		SpriteSheet: monkSpriteEarthFistSmashRight,
		Fixed:       true,
	}

	a2alKey := string(sr.SecondaryAttackKey) + "Left"
	anims[a2alKey] = &Animation{
		Name:        a2alKey,
		FrameOX:     712,
		FrameOY:     0,
		FrameWidth:  89,
		FrameHeight: 55,
		FrameCount:  8,
		PosOffsetY:  15,
		SpriteSheet: monkSpriteEarthFistSmashLeft,
		Fixed:       true,
	}

	/*
		---------------------------------------------------------------------------------
		Secondary End
		---------------------------------------------------------------------------------
	*/

	return anims

}
