package roles

import (
	"github.com/hajimehoshi/ebiten/v2"
	sr "github.com/kainn9/grpc_game/server/roles"
	ut "github.com/kainn9/grpc_game/util"
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
func LoadmonkSprites() {
	monkSpriteIdleLeft = ut.LoadImg("./sprites/monk/monkIdleLeft.png")
	monkSpriteIdleRight = ut.LoadImg("./sprites/monk/monkIdleRight.png")

	monkSpriteWalkingRight = ut.LoadImg("./sprites/monk/monkRunningRight.png")
	monkSpriteWalkingLeft = ut.LoadImg("./sprites/monk/monkRunningLeft.png")

	monkSpriteJumpLeft = ut.LoadImg("./sprites/monk/monkJumpLeft.png")
	monkSpriteJumpRight = ut.LoadImg("./sprites/monk/monkJumpRight.png")

	monkDefenseRight = ut.LoadImg("./sprites/monk/monkDefenseRight.png")
	monkDefenseLeft = ut.LoadImg("./sprites/monk/monkDefenseLeft.png")

	monkSpriteKBRight = ut.LoadImg("./sprites/monk/monkKnockBackRight.png")
	monkSpriteKBLeft = ut.LoadImg("./sprites/monk/monkKnockBackLeft.png")

	monkSpriteSmashAtkRight = ut.LoadImg("./sprites/monk/monkSmashAttackRight.png")
	monkSpriteSmashAtkLeft = ut.LoadImg("./sprites/monk/monkSmashAttackLeft.png")

	monkSpriteEarthFistSmashRight = ut.LoadImg("./sprites/monk/monkEarthFistSmashRight.png")
	monkSpriteEarthFistSmashLeft = ut.LoadImg("./sprites/monk/monkEarthFistSmashLeft.png")

	monkSpriteDeathRight = ut.LoadImg("./sprites/monk/monkDeathRight.png")
	monkSpriteDeathLeft = ut.LoadImg("./sprites/monk/monkDeathLeft.png")
}

func InitMonk() *Role {
	LoadmonkSprites()

	r := &Role{
		RoleType:      MonkType,
		Animations:    MonkAnims(),
		HitBoxOffsetY: 4,
		HitBoxOffsetX: 4,
	}

	return r
}

// TODO MAKE ANIM KEYS CONSTS
func MonkAnims() map[string]*Animation {
	anims := make(map[string]*Animation)

	anims["idleRight"] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  30,
		FrameHeight: 38,
		FrameCount:  6,
		SpriteSheet: monkSpriteIdleRight,
	}

	anims["idleLeft"] = &Animation{
		FrameOX:     180,
		FrameOY:     0,
		FrameWidth:  30,
		FrameHeight: 38,
		FrameCount:  6,
		SpriteSheet: monkSpriteIdleLeft,
	}

	anims["walkRight"] = &Animation{
		FrameOX:     0,
		FrameOY:     14,
		FrameWidth:  33,
		FrameHeight: 62,
		FrameCount:  8,
		SpriteSheet: monkSpriteWalkingRight,
	}

	anims["walkLeft"] = &Animation{
		FrameOX:     264,
		FrameOY:     14,
		FrameWidth:  33,
		FrameHeight: 62,
		FrameCount:  8,
		SpriteSheet: monkSpriteWalkingLeft,
	}

	anims["jumpLeft"] = &Animation{
		FrameOX:     105,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 49,
		FrameCount:  3,
		SpriteSheet: monkSpriteJumpLeft,
	}

	anims["jumpRight"] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 49,
		FrameCount:  3,
		SpriteSheet: monkSpriteJumpRight,
	}

	anims["defenseRight"] = &Animation{
		Name:        "defenseRight",
		FrameOX:     0,
		FrameOY:     7,
		FrameWidth:  45,
		FrameHeight: 48,
		FrameCount:  8,
		PosOffsetX:  10,
		SpriteSheet: monkDefenseRight,
		Fixed:       true,
	}

	anims["defenseLeft"] = &Animation{
		Name:        "defenseLeft",
		FrameOX:     360,
		FrameOY:     7,
		FrameWidth:  45,
		FrameHeight: 48,
		FrameCount:  8,
		PosOffsetX:  10,
		SpriteSheet: monkDefenseLeft,
		Fixed:       true,
	}

	anims["KbRight"] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  50,
		FrameHeight: 39,
		FrameCount:  3,
		SpriteSheet: monkSpriteKBRight,
	}

	anims["KbLeft"] = &Animation{
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
		Name:        "primaryAtkRight",
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  111,
		FrameHeight: 42,
		FrameCount:  13,
		SpriteSheet: monkSpriteSmashAtkRight,
		Fixed:       true,
	}

	anims[string(sr.PrimaryAttackKey)+"Left"] = &Animation{
		Name:        "primaryAtkleft",
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
		FrameOY:     15,
		FrameWidth:  89,
		FrameHeight: 55,
		FrameCount:  8,
		SpriteSheet: monkSpriteEarthFistSmashRight,
		Fixed:       true,
	}

	a2alKey := string(sr.SecondaryAttackKey) + "Left"
	anims[a2alKey] = &Animation{
		Name:        a2alKey,
		FrameOX:     712,
		FrameOY:     15,
		FrameWidth:  89,
		FrameHeight: 55,
		FrameCount:  8,
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
