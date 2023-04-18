package roles

import (
	"github.com/hajimehoshi/ebiten/v2"
	utClient "github.com/kainn9/grpc_game/client_util"
	sr "github.com/kainn9/grpc_game/server/roles"
)

/*
File for Werewolf class
contains sprites/animation data
*/
var (
	Werewolf *Role = InitWerewolf()
)

var (
	werewolfSpriteIdleLeft  *ebiten.Image
	werewolfSpriteIdleRight *ebiten.Image

	werewolfSpriteWalkingRight *ebiten.Image
	werewolfSpriteWalkingLeft  *ebiten.Image

	werewolfSpriteJumpLeft  *ebiten.Image
	werewolfSpriteJumpRight *ebiten.Image

	werewolfSpriteHitRight *ebiten.Image
	werewolfSpriteHitLeft  *ebiten.Image

	werewolfSpriteKBRight *ebiten.Image
	werewolfSpriteKBLeft  *ebiten.Image

	werewolfSpriteDeathRight *ebiten.Image
	werewolfSpriteDeathLeft  *ebiten.Image

	werewolfSpriteSPWindupRight *ebiten.Image
	werewolfSpriteSPWindupLeft  *ebiten.Image

	werewolfSpriteSPMovementRight *ebiten.Image
	werewolfSpriteSPMovementLeft  *ebiten.Image

	werewolfSpriteSPAtkRight *ebiten.Image
	werewolfSpriteSPAtkLeft  *ebiten.Image

	werewolfSpriteDoubleSlashRight *ebiten.Image
	werewolfSpriteDoubleSlashLeft  *ebiten.Image

	werewolfSpriteFlipSlashLeft  *ebiten.Image
	werewolfSpriteFlipSlashRight *ebiten.Image
)

/*
Loads the default player sprites
*/
func LoadWerewolfSprites() {
	werewolfSpriteIdleLeft = utClient.LoadImage("./sprites/werewolf/werewolfIdleLeft.png")
	werewolfSpriteIdleRight = utClient.LoadImage("./sprites/werewolf/werewolfIdleRight.png")

	werewolfSpriteWalkingRight = utClient.LoadImage("./sprites/werewolf/werewolfRunningRight.png")
	werewolfSpriteWalkingLeft = utClient.LoadImage("./sprites/werewolf/werewolfRunningLeft.png")

	werewolfSpriteJumpLeft = utClient.LoadImage("./sprites/werewolf/werewolfJumpLeft.png")
	werewolfSpriteJumpRight = utClient.LoadImage("./sprites/werewolf/werewolfJumpRight.png")

	werewolfSpriteHitRight = utClient.LoadImage("./sprites/werewolf/werewolfHitRight.png")
	werewolfSpriteHitLeft = utClient.LoadImage("./sprites/werewolf/werewolfHitLeft.png")

	werewolfSpriteKBRight = utClient.LoadImage("./sprites/werewolf/werewolfKnockBackRight.png")
	werewolfSpriteKBLeft = utClient.LoadImage("./sprites/werewolf/werewolfKnockBackLeft.png")

	werewolfSpriteDeathRight = utClient.LoadImage("./sprites/werewolf/werewolfDeathRight.png")
	werewolfSpriteDeathLeft = utClient.LoadImage("./sprites/werewolf/werewolfDeathLeft.png")

	werewolfSpriteSPWindupRight = utClient.LoadImage("./sprites/werewolf/werewolfSPWindupRight.png")
	werewolfSpriteSPWindupLeft = utClient.LoadImage("./sprites/werewolf/werewolfSPWindupLeft.png")

	werewolfSpriteSPMovementRight = utClient.LoadImage("./sprites/werewolf/werewolfSPMovementRight.png")
	werewolfSpriteSPMovementLeft = utClient.LoadImage("./sprites/werewolf/werewolfSPMovementLeft.png")

	werewolfSpriteSPAtkRight = utClient.LoadImage("./sprites/werewolf/werewolfSPAtkRight.png")
	werewolfSpriteSPAtkLeft = utClient.LoadImage("./sprites/werewolf/werewolfSPAtkLeft.png")

	werewolfSpriteDoubleSlashRight = utClient.LoadImage("./sprites/werewolf/werewolfDoubleSlashRight.png")
	werewolfSpriteDoubleSlashLeft = utClient.LoadImage("./sprites/werewolf/werewolfDoubleSlashLeft.png")

	werewolfSpriteFlipSlashLeft = utClient.LoadImage("./sprites/werewolf/werewolfFlipSlashLeft.png")
	werewolfSpriteFlipSlashRight = utClient.LoadImage("./sprites/werewolf/werewolfFlipSlashRight.png")
}

func InitWerewolf() *Role {
	LoadWerewolfSprites()

	r := &Role{
		Animations:    WerewolfAnims(),
		HitBoxOffsetY: 0,
		HitBoxOffsetX: 20,
		Health:        sr.Werewolf.Health,
		HitBoxW:       sr.Werewolf.HitBoxW,
		HitBoxH:       sr.Werewolf.HitBoxH,
		HealthBarOffset: &Offset{
			X: 15,
			Y: -12,
		},
		StatusEffectOffset: &Offset{
			X: 15,
			Y: -8,
		},
	}

	return r
}

func WerewolfAnims() map[string]*Animation {
	anims := make(map[string]*Animation)

	anims[string(IdleRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  69,
		FrameHeight: 82,
		FrameCount:  10,
		PosOffsetY:  30,
		SpriteSheet: werewolfSpriteIdleRight,
	}

	anims[string(IdleLeft)] = &Animation{
		FrameOX:     690,
		FrameOY:     0,
		FrameWidth:  69,
		FrameHeight: 82,
		FrameCount:  10,
		PosOffsetY:  30,
		SpriteSheet: werewolfSpriteIdleLeft,
	}

	anims[string(WalkRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  87,
		FrameHeight: 82,
		FrameCount:  6,
		PosOffsetY:  15,
		SpriteSheet: werewolfSpriteWalkingRight,
	}

	anims[string(WalkLeft)] = &Animation{
		FrameOX:     522,
		FrameOY:     0,
		FrameWidth:  87,
		FrameHeight: 82,
		FrameCount:  6,
		PosOffsetY:  15,
		SpriteSheet: werewolfSpriteWalkingLeft,
	}

	anims[string(JumpLeft)] = &Animation{
		FrameOX:     162,
		FrameOY:     0,
		FrameWidth:  54,
		FrameHeight: 82,
		FrameCount:  3,
		PosOffsetY:  10,
		SpriteSheet: werewolfSpriteJumpLeft,
	}

	anims[string(JumpRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  54,
		FrameHeight: 82,
		FrameCount:  3,
		PosOffsetY:  10,
		SpriteSheet: werewolfSpriteJumpRight,
	}

	anims[string(HitRight)] = &Animation{
		Name:        string(HitRight),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  160,
		FrameHeight: 96,
		FrameCount:  6,
		SpriteSheet: werewolfSpriteHitRight,
		PosOffsetX:  40,
		PosOffsetY:  30,
		Fixed:       true,
	}

	anims[string(HitLeft)] = &Animation{
		Name:        string(HitLeft),
		FrameOX:     960,
		FrameOY:     0,
		FrameWidth:  160,
		FrameHeight: 96,
		FrameCount:  6,
		SpriteSheet: werewolfSpriteHitLeft,
		PosOffsetX:  40,
		PosOffsetY:  30,
		Fixed:       true,
	}

	stunAnimCopyRight := *anims[string(HitRight)]
	anims[string(StunRight)] = &stunAnimCopyRight

	stunAnimCopyLeft := *anims[string(HitLeft)]
	anims[string(StunLeft)] = &stunAnimCopyLeft

	anims[string(KbRight)] = &Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  76,
		FrameHeight: 82,
		FrameCount:  6,
		PosOffsetY:  30,
		SpriteSheet: werewolfSpriteKBRight,
	}

	anims[string(KbLeft)] = &Animation{
		FrameOX:     456,
		FrameOY:     0,
		FrameWidth:  76,
		FrameHeight: 82,
		FrameCount:  6,
		PosOffsetY:  30,
		SpriteSheet: werewolfSpriteKBLeft,
	}

	anims[string(DeathRight)] = &Animation{
		Name:        string(DeathRight),
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  81,
		FrameHeight: 82,
		FrameCount:  24,
		PosOffsetY:  20,
		SpriteSheet: werewolfSpriteDeathRight,
		Fixed:       true,
	}

	anims[string(DeathLeft)] = &Animation{
		Name:        string(DeathLeft),
		FrameOX:     1944,
		FrameOY:     0,
		FrameWidth:  81,
		FrameHeight: 82,
		FrameCount:  24,
		PosOffsetY:  20,
		SpriteSheet: werewolfSpriteDeathLeft,
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
		FrameWidth:  160,
		FrameHeight: 96,
		FrameCount:  8,
		PosOffsetX:  50,
		PosOffsetY:  45,
		SpriteSheet: werewolfSpriteDoubleSlashRight,
		Fixed:       true,
	}

	anims[string(sr.PrimaryAttackKey)+"Left"] = &Animation{
		Name:        string(sr.PrimaryAttackKey) + "Left",
		FrameOX:     1280,
		FrameOY:     0,
		FrameWidth:  160,
		FrameHeight: 96,
		FrameCount:  8,
		PosOffsetX:  50,
		PosOffsetY:  45,
		SpriteSheet: werewolfSpriteDoubleSlashLeft,
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
		FrameWidth:  160,
		FrameHeight: 96,
		FrameCount:  10,
		PosOffsetX:  30,
		PosOffsetY:  50,
		SpriteSheet: werewolfSpriteFlipSlashRight,
		Fixed:       true,
	}

	a2alKey := string(sr.SecondaryAttackKey) + "Left"
	anims[a2alKey] = &Animation{
		Name:        a2alKey,
		FrameOX:     1600,
		FrameOY:     0,
		FrameWidth:  160,
		FrameHeight: 96,
		FrameCount:  10,
		PosOffsetX:  30,
		PosOffsetY:  50,
		SpriteSheet: werewolfSpriteFlipSlashLeft,
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
		FrameOX:     960,
		FrameOY:     0,
		FrameWidth:  160,
		FrameHeight: 96,
		FrameCount:  6,
		PosOffsetX:  30,
		PosOffsetY:  30,
		SpriteSheet: werewolfSpriteSPMovementLeft,
		Fixed:       true,
	}

	a3mrKey := string(sr.TertAttackKey) + "MovementRight"
	anims[a3mrKey] = &Animation{
		Name:        a3mrKey,
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  160,
		FrameHeight: 96,
		FrameCount:  6,
		PosOffsetX:  30,
		PosOffsetY:  30,
		SpriteSheet: werewolfSpriteSPMovementRight,
		Fixed:       true,
	}

	a3wurKey := string(sr.TertAttackKey) + "WindupRight"
	anims[a3wurKey] = &Animation{
		Name:        a3wurKey,
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  160,
		FrameHeight: 96,
		FrameCount:  4,
		PosOffsetX:  30,
		PosOffsetY:  30,
		SpriteSheet: werewolfSpriteSPWindupRight,
		Fixed:       true,
	}

	a3wulKey := string(sr.TertAttackKey) + "WindupLeft"
	anims[a3wulKey] = &Animation{
		Name:        a3wulKey,
		FrameOX:     960,
		FrameOY:     0,
		FrameWidth:  160,
		FrameHeight: 96,
		FrameCount:  4,
		PosOffsetX:  30,
		PosOffsetY:  30,
		SpriteSheet: werewolfSpriteSPWindupLeft,
		Fixed:       true,
	}

	a3arKey := "tertAtkRight"
	anims[a3arKey] = &Animation{
		Name:        a3arKey,
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  160,
		FrameHeight: 96,
		FrameCount:  17,
		PosOffsetX:  50,
		PosOffsetY:  30,
		SpriteSheet: werewolfSpriteSPAtkRight,
		Fixed:       true,
	}

	a3alKey := "tertAtkLeft"
	anims[a3alKey] = &Animation{
		Name:        a3alKey,
		FrameOX:     2720,
		FrameOY:     0,
		FrameWidth:  160,
		FrameHeight: 96,
		FrameCount:  17,
		PosOffsetX:  50,
		PosOffsetY:  30,
		SpriteSheet: werewolfSpriteSPAtkLeft,
		Fixed:       true,
	}

	/*
		---------------------------------------------------------------------------------
		Tert Attack END
		---------------------------------------------------------------------------------
	*/

	return anims

}
