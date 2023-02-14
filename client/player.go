package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Player Sprites(for now)
var (
	playerSpriteIdleLeft  *ebiten.Image
	playerSpriteIdleRight *ebiten.Image

	playerSpriteWalkingRight *ebiten.Image
	playerSpriteWalkingLeft  *ebiten.Image
	playerSpriteJumpLeft     *ebiten.Image
	playerSpriteJumpRight    *ebiten.Image
)

type Player struct {
	Animations  []Animation
	SpeedX      float64
	SpeedY      float64
	X           float64
	Y           float64
	FacingRight bool
	Jumping     bool
}

/*
Animation "system"(lol) will need to be redone better...
Will probably start to tackle when adding NPC's
or having various player sprites
*/
type Animation struct {
	FrameOX     int
	FrameOY     int
	FrameWidth  int
	FrameHeight int
	FrameCount  int
	Name        string
	SpriteSheet *ebiten.Image
}

/*
Creates a player with the default Anims
*/
func NewPlayer() *Player {

	idleRight := Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  32,
		FrameHeight: 48,
		FrameCount:  8,
		SpriteSheet: playerSpriteIdleRight,
	}
	idleLeft := Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  32,
		FrameHeight: 48,
		FrameCount:  8,
		SpriteSheet: playerSpriteIdleLeft,
	}

	walkRight := Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 48,
		FrameCount:  8,
		SpriteSheet: playerSpriteWalkingRight,
	}

	walkLeft := Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 48,
		FrameCount:  8,
		SpriteSheet: playerSpriteWalkingLeft,
	}

	jumpLeft := Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 48,
		FrameCount:  1,
		SpriteSheet: playerSpriteJumpLeft,
	}

	jumpRight := Animation{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  35,
		FrameHeight: 48,
		FrameCount:  1,
		SpriteSheet: playerSpriteJumpRight,
	}

	p := &Player{
		Animations: []Animation{
			idleRight,
			idleLeft,
			walkRight,
			walkLeft,
			jumpLeft,
			jumpRight,
		},
	}
	return p
}

/*
Renders a Player using their Anim's
and game tick count
*/
func DrawPlayer(world *World, p *Player, currentPlayer bool) {
	currentAnimation := p.Animations[0]

	/* TODO: gotta dry this up...
	no need to keep calling:
	i := (ticks / 5) % currentAnimation.FrameCount
	over and over...
	*/
	i := (ticks / 5) % currentAnimation.FrameCount

	s := playerSpriteIdleRight

	if !p.FacingRight {
		s = playerSpriteIdleLeft
		currentAnimation = p.Animations[1]
		i = (ticks / 5) % currentAnimation.FrameCount
	}

	if p.SpeedX > 0 && p.FacingRight {
		s = playerSpriteWalkingRight
		currentAnimation = p.Animations[2]
		i = (ticks / 5) % currentAnimation.FrameCount
	}

	if p.SpeedX < 0 && !p.FacingRight {
		s = playerSpriteWalkingLeft
		currentAnimation = p.Animations[3]
		i = (ticks / 5) % currentAnimation.FrameCount
	}

	if p.Jumping && !p.FacingRight {
		s = playerSpriteJumpLeft
		currentAnimation = p.Animations[4]
		i = (ticks / 5) % currentAnimation.FrameCount
	}

	if p.Jumping && p.FacingRight {
		s = playerSpriteJumpRight
		currentAnimation = p.Animations[5]
		i = (ticks / 5) % currentAnimation.FrameCount
	}

	/*
		Logic for rendering current and other players
		inside the game camera
	*/
	x := p.X
	y := p.Y

	pc := world.PlayerController
	pcX := pc.X
	pcY := pc.Y
	playerOps := &ebiten.DrawImageOptions{}

	if currentPlayer {
		playerOps = pc.Cam.GetTranslation(playerOps, x/2, y/2)
	} else {
		playerOps = pc.Cam.GetTranslation(playerOps, x-(pcX/2), y-(pcY/2))
	}

	sx, sy := (currentAnimation.FrameOX)+i*(currentAnimation.FrameWidth), (currentAnimation.FrameOY)
	sub := s.SubImage(image.Rect(sx, sy, sx+(currentAnimation.FrameWidth), sy+(currentAnimation.FrameHeight))).(*ebiten.Image)
	pc.Cam.Surface.DrawImage(sub, playerOps)
}
