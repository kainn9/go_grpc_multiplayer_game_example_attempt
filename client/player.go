package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	r "github.com/kainn9/grpc_game/client/roles"
)

type Player struct {
	SpeedX      float64
	SpeedY      float64
	X           float64
	Y           float64
	FacingRight bool
	Jumping     bool
	CC          string
	CurrAttack  string
	r.Role
	currentAnimation *r.Animation
}

/*
Creates a player with the default Anims
*/
func NewPlayer() *Player {

	p := &Player{
		Role: *r.InitKnight(),
	}
	return p
}

/*
Renders a Player using their Anim's
and game tick count
*/
func DrawPlayer(world *World, p *Player, currentPlayer bool) {

	defaultAnim := p.Animations["idleRight"]

	if p.currentAnimation == nil {
		p.currentAnimation = defaultAnim
	}

	prevAnim := defaultAnim

	if !p.FacingRight {

		p.currentAnimation = p.Animations["idleLeft"]

	}

	if p.SpeedX > 0 && p.FacingRight {

		p.currentAnimation = p.Animations["walkRight"]

	}

	if p.SpeedX < 0 && !p.FacingRight {
		p.currentAnimation = p.Animations["walkLeft"]

	}

	if p.Jumping && !p.FacingRight {
		p.currentAnimation = p.Animations["jumpLeft"]

	}

	if p.Jumping && p.FacingRight {
		p.currentAnimation = p.Animations["jumpRight"]

	}

	if p.CurrAttack != "" {
		if p.FacingRight {
			p.currentAnimation = p.Animations[p.CurrAttack+"Right"]
		} else {
			p.currentAnimation = p.Animations[p.CurrAttack+"Left"]
		}

	}

	if p.CC != "" {

		if p.FacingRight {
			p.currentAnimation = p.Animations[p.CC+"Right"]
		} else {
			p.currentAnimation = p.Animations[p.CC+"Left"]
		}

	}
	i := (ticks / 5) % p.currentAnimation.FrameCount
	s := p.currentAnimation.SpriteSheet

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

	if currentPlayer && p.FacingRight {
		playerOps = pc.Cam.GetTranslation(playerOps, (x/2)-p.HitBoxOffsetX, (y/2)-p.HitBoxOffsetY)

	} else if currentPlayer && !p.FacingRight {
		playerOps = pc.Cam.GetTranslation(playerOps, (-p.HitBoxOffsetX)+x/2-float64(p.currentAnimation.FrameWidth-prevAnim.FrameWidth), (y/2)-p.HitBoxOffsetY)

	} else if p.FacingRight {
		playerOps = pc.Cam.GetTranslation(playerOps, x-(pcX/2)-p.HitBoxOffsetX, y-(pcY/2)-p.HitBoxOffsetY)

	} else {
		playerOps = pc.Cam.GetTranslation(playerOps, (-p.HitBoxOffsetX)+x-(pcX/2)-float64(p.currentAnimation.FrameWidth-prevAnim.FrameWidth), y-(pcY/2)-p.HitBoxOffsetY)
	}


	// Render the Anims
	sx, sy := (p.currentAnimation.FrameOX)+i*(p.currentAnimation.FrameWidth), (p.currentAnimation.FrameOY)
	sub := s.SubImage(image.Rect(sx, sy, sx+(p.currentAnimation.FrameWidth), sy+(p.currentAnimation.FrameHeight))).(*ebiten.Image)

	if !p.FacingRight {
		sx, sy = (p.currentAnimation.FrameOX)-i*(p.currentAnimation.FrameWidth), (p.currentAnimation.FrameOY)
		sub = s.SubImage(image.Rect(sx, sy, sx-(p.currentAnimation.FrameWidth), sy+(p.currentAnimation.FrameHeight))).(*ebiten.Image)
	}
	pc.Cam.Surface.DrawImage(sub, playerOps)
}
