package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	r "github.com/kainn9/grpc_game/client/roles"
)

type Player struct {
	id             string
	speedX         float64
	speedY         float64
	x              float64
	y              float64
	facingRight    bool
	jumping        bool
	cc             string
	currAttack     string
	windup         string
	attackMovement string
	r.Role
	currentAnimation *r.Animation
	health           int
	defending        bool
	dead             bool
}

/*
Creates a player with the default Anims
*/
func NewPlayer() *Player {

	p := &Player{}

	return p
}

/*
TODO: DOC AND CLEAN UP
*/
func DrawPlayer(world *World, p *Player, currentPlayer bool) {

	defaultAnim := p.Animations["idleRight"]

	if p.currentAnimation == nil {
		p.currentAnimation = defaultAnim
	}

	prevAnim := defaultAnim

	if !p.facingRight {

		p.currentAnimation = p.Animations["idleLeft"]

	}

	if p.speedX > 0 && p.facingRight {

		p.currentAnimation = p.Animations["walkRight"]

	}

	if p.speedX < 0 && !p.facingRight {
		p.currentAnimation = p.Animations["walkLeft"]

	}

	if p.jumping && !p.facingRight {
		p.currentAnimation = p.Animations["jumpLeft"]

	}

	if p.jumping && p.facingRight {
		p.currentAnimation = p.Animations["jumpRight"]

	}

	if p.currAttack != "" {
		if p.facingRight {
			p.currentAnimation = p.Animations[p.currAttack+"Right"]
		} else {
			p.currentAnimation = p.Animations[p.currAttack+"Left"]
		}
	}

	if p.windup != "" {
		if p.facingRight {
			p.currentAnimation = p.Animations[p.windup+"WindupRight"]
		} else {
			p.currentAnimation = p.Animations[p.windup+"WindupLeft"]
		}
	}

	if p.attackMovement != "" {
		if p.facingRight {
			p.currentAnimation = p.Animations[p.attackMovement+"MovementRight"]
		} else {
			p.currentAnimation = p.Animations[p.attackMovement+"MovementLeft"]
		}
	}

	if p.cc != "" {

		if p.facingRight {
			p.currentAnimation = p.Animations[p.cc+"Right"]
		} else {
			p.currentAnimation = p.Animations[p.cc+"Left"]
		}

	}

	if p.defending {

		if p.facingRight {
			p.currentAnimation = p.Animations["defenseRight"]
		} else {
			p.currentAnimation = p.Animations["defenseLeft"]
		}

	}

	if p.dead {

		if p.facingRight {
			p.currentAnimation = p.Animations[string(r.DeathRight)]
		} else {
			p.currentAnimation = p.Animations[string(r.DeathLeft)]
		}

	}

	if currentPlayer && hitBoxTest.on {
		if p.facingRight {
			p.currentAnimation = p.Animations[hitBoxTest.name+"Right"]
		} else {
			p.currentAnimation = p.Animations[hitBoxTest.name+"Left"]
		}
	}

	if p.currentAnimation == nil {
		log.Printf("No animation for player state: %v\n", p)
		return
	}

	i := (clientConfig.ticks / 5) % p.currentAnimation.FrameCount
	s := p.currentAnimation.SpriteSheet

	if p.currentAnimation.Fixed {
		fixedAnimKey := p.id + p.currentAnimation.Name
		fixedAnimationCheck := fixedAnims[fixedAnimKey]

		if fixedAnimationCheck == nil {
			fixedAnims[fixedAnimKey] = &fixedAnimTracker{
				pid:      p.id,
				animName: p.currentAnimation.Name,
				ticks:    0,
			}
		} else {
			fixedTicks := fixedAnims[fixedAnimKey].ticks
			i = (fixedTicks / 5) % p.currentAnimation.FrameCount
		}
	}

	/*
		Logic for rendering current and other players
		inside the game camera
	*/
	x := p.x
	y := p.y

	pc := world.playerController

	playerOps := &ebiten.DrawImageOptions{}

	if currentPlayer && p.facingRight {
		playerOps = pc.playerCam.GetTranslation(playerOps, -p.currentAnimation.PosOffsetX+((pc.playerCXpos/2)-p.HitBoxOffsetX), (pc.playerCYpos/2)-p.HitBoxOffsetY)

	} else if currentPlayer && !p.facingRight {
		playerOps = pc.playerCam.GetTranslation(playerOps, p.currentAnimation.PosOffsetX+(-p.HitBoxOffsetX)+pc.playerCXpos/2-float64(p.currentAnimation.FrameWidth-prevAnim.FrameWidth), (pc.playerCYpos/2)-p.HitBoxOffsetY)

	} else if p.facingRight {
		playerOps = pc.playerCam.GetTranslation(playerOps, -p.currentAnimation.PosOffsetX+(x-(pc.playerCXpos/2)-p.HitBoxOffsetX), y-(pc.playerCYpos/2)-p.HitBoxOffsetY)

	} else {
		playerOps = pc.playerCam.GetTranslation(playerOps, p.currentAnimation.PosOffsetX+(-p.HitBoxOffsetX)+x-(pc.playerCXpos/2)-float64(p.currentAnimation.FrameWidth-prevAnim.FrameWidth), y-(pc.playerCYpos/2)-p.HitBoxOffsetY)
	}

	// Render the Anims
	sx, sy := (p.currentAnimation.FrameOX)+i*(p.currentAnimation.FrameWidth), (p.currentAnimation.FrameOY)
	sub := s.SubImage(image.Rect(sx, sy, sx+(p.currentAnimation.FrameWidth), sy+(p.currentAnimation.FrameHeight))).(*ebiten.Image)

	if !p.facingRight {
		sx, sy = (p.currentAnimation.FrameOX)-i*(p.currentAnimation.FrameWidth), (p.currentAnimation.FrameOY)
		sub = s.SubImage(image.Rect(sx, sy, sx-(p.currentAnimation.FrameWidth), sy+(p.currentAnimation.FrameHeight))).(*ebiten.Image)
	}

	if hitBoxTest.on && hitBoxTest.frame >= 0 {
		sub = getAnimationFrame(p, hitBoxTest.frame, s)
		
	}

	// render player
	pc.playerCam.Surface.DrawImage(sub, playerOps)

	/*
	-------------------------------------------------------
		Uncomment this and place values in NewImage to preview player hitbox— expand this to hitboxTest
	-------------------------------------------------------
	*/
	// rectImg := ebiten.NewImage(50, 98)
	// rectImg.Fill(color.RGBA{0, 0, 255, 128})
	// playerOps.GeoM.Translate(p.HitBoxOffsetX, p.HitBoxOffsetY)
	// pc.playerCam.Surface.DrawImage(rectImg, playerOps)
	/*
	-------------------------------------------------------
	 end
	-------------------------------------------------------
	*/		
	
}

func getAnimationFrame(p *Player, i int, s *ebiten.Image) *ebiten.Image {
	sx, sy := (p.currentAnimation.FrameOX)+i*(p.currentAnimation.FrameWidth), (p.currentAnimation.FrameOY)
	sub := s.SubImage(image.Rect(sx, sy, sx+(p.currentAnimation.FrameWidth), sy+(p.currentAnimation.FrameHeight))).(*ebiten.Image)

	if !p.facingRight {
		sx, sy = (p.currentAnimation.FrameOX)-i*(p.currentAnimation.FrameWidth), (p.currentAnimation.FrameOY)
		sub = s.SubImage(image.Rect(sx, sy, sx-(p.currentAnimation.FrameWidth), sy+(p.currentAnimation.FrameHeight))).(*ebiten.Image)
	}

	alphaValue := uint8(10)
	bounds := sub.Bounds()

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := sub.At(x, y).RGBA()
			sub.Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), alphaValue})
		}
	}

	return sub
}
