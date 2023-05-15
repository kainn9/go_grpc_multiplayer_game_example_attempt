package main

import (
	"image"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	r "github.com/kainn9/grpc_game/client/roles"
	cse "github.com/kainn9/grpc_game/client/statusEffects"
	sse "github.com/kainn9/grpc_game/server/statusEffects"
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
	cooldowns        string
}

/*
Creates a player with the default Anims
*/
func NewPlayer() *Player {

	p := &Player{}

	return p
}

func animationSubImage(anim *r.Animation, s *ebiten.Image, i int, right bool, VerticalSheet bool) *ebiten.Image {
	var sx, sy int
	if !VerticalSheet {
		sx, sy = (anim.FrameOX)+i*(anim.FrameWidth), (anim.FrameOY)
	} else {
		sx, sy = (anim.FrameOX), (anim.FrameOY)+i*(anim.FrameHeight)
	}

	sub := s.SubImage(image.Rect(sx, sy, sx+(anim.FrameWidth), sy+(anim.FrameHeight))).(*ebiten.Image)

	if !right && !VerticalSheet {
		sx, sy = (anim.FrameOX)-i*(anim.FrameWidth), (anim.FrameOY)
		sub = s.SubImage(image.Rect(sx, sy, sx-(anim.FrameWidth), sy+(anim.FrameHeight))).(*ebiten.Image)
	}

	return sub
}

func standardDrawOpts(currentPlayer bool, pc *PlayerController, p *Player) ebiten.DrawImageOptions {
	drawOpts := &ebiten.DrawImageOptions{}
	if currentPlayer && p.facingRight {
		drawOpts = pc.playerCam.GetTranslation(drawOpts, (pc.playerCXpos/2)-p.HitBoxOffsetX, (pc.playerCYpos/2)-p.HitBoxOffsetY)

	} else if currentPlayer && !p.facingRight {
		drawOpts = pc.playerCam.GetTranslation(drawOpts, -p.HitBoxOffsetX+(pc.playerCXpos/2), (pc.playerCYpos/2)-p.HitBoxOffsetY)

	} else if p.facingRight {
		drawOpts = pc.playerCam.GetTranslation(drawOpts, p.x-(pc.playerCXpos/2)-p.HitBoxOffsetX, p.y-(pc.playerCYpos/2)-p.HitBoxOffsetY)
	} else {
		drawOpts = pc.playerCam.GetTranslation(drawOpts, (-p.HitBoxOffsetX)+p.x-(pc.playerCXpos/2), p.y-(pc.playerCYpos/2)-p.HitBoxOffsetY)
	}

	return *drawOpts
}

// its just stun right now
func (p *Player) drawStatusEffect(currentPlayer bool, pc *PlayerController) {
	statusOps := standardDrawOpts(currentPlayer, pc, p)
	statusOps.GeoM.Translate(p.Role.StatusEffectOffset.X, p.Role.StatusEffectOffset.Y)

	anim := cse.Stun.Anim
	i := (clientConfig.ticks / 5) % anim.FrameCount
	img := cse.Stun.Img
	sub := animationSubImage(anim, img, i, true, p.currentAnimation.VerticalSheet)

	pc.playerCam.Surface.DrawImage(sub, &statusOps)
}

func (p *Player) drawHealth(world *World, currentPlayer bool) {

	if p.dead {
		return
	}

	maxWidth := 45.0

	healthRatio := float64(p.health) / float64(p.Role.Health)

	health := healthRatio * maxWidth

	black := color.RGBA{0, 0, 0, 255}
	green := color.RGBA{0, 255, 0, 255}
	yellow := color.RGBA{255, 255, 0, 255}
	red := color.RGBA{255, 0, 0, 255}

	if int(health) <= 0 {
		health = 1
	}

	pc := world.playerController

	// using SyncMap to cache the healthbar sprite
	// to avoid creating a new image every TPS
	var healthBar *ebiten.Image
	var healthBarOutline *ebiten.Image
	outlineCacheKey := "hbOutline"

	if hbOutlineResult, ok := clientConfig.imageCache.Load(outlineCacheKey); ok {
		healthBarOutline = hbOutlineResult.(cachedImage).img
	} else {
		healthBarOutline = ebiten.NewImage(int(maxWidth+4), 9)

		cImgOL := cachedImage{
			img:       healthBarOutline,
			timestamp: time.Now(),
		}
		clientConfig.imageCache.Store(outlineCacheKey, cImgOL)

	}

	if hbResult, ok := clientConfig.imageCache.Load(healthRatio); ok {
		healthBar = hbResult.(cachedImage).img
	} else {
		healthBar = ebiten.NewImage(int(health), 5)
		cImgHB := cachedImage{
			img:       healthBar,
			timestamp: time.Now(),
		}
		clientConfig.imageCache.Store(healthRatio, cImgHB)
	}

	healthBarOpts := standardDrawOpts(currentPlayer, pc, p)

	if healthRatio < 0.6 && healthRatio > 0.3 {
		healthBar.Fill(yellow)
	} else if healthRatio <= 0.3 {
		healthBar.Fill(red)
	} else {
		healthBar.Fill(green)
	}

	healthBarOpts.GeoM.Translate(p.Role.HealthBarOffset.X, p.Role.HealthBarOffset.Y)

	healthBarOutline.Fill(black)
	outlineOpts := healthBarOpts
	outlineOpts.GeoM.Translate(-2, -2)

	pc.playerCam.Surface.DrawImage(healthBarOutline, &outlineOpts)
	pc.playerCam.Surface.DrawImage(healthBar, &healthBarOpts)
}

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

	// todo maybe make this a func
	// more move the status effect down
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
		// log.Printf("No animation for player state: %v\n", p)
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
			i = 0 % p.currentAnimation.FrameCount
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
		playerOps = pc.playerCam.GetTranslation(playerOps, -p.currentAnimation.PosOffsetX+((pc.playerCXpos/2)-p.HitBoxOffsetX), -p.currentAnimation.PosOffsetY+(pc.playerCYpos/2)-p.HitBoxOffsetY)

	} else if currentPlayer && !p.facingRight {
		playerOps = pc.playerCam.GetTranslation(playerOps, p.currentAnimation.PosOffsetX+(-p.HitBoxOffsetX)+pc.playerCXpos/2-float64(p.currentAnimation.FrameWidth-prevAnim.FrameWidth), -p.currentAnimation.PosOffsetY+(pc.playerCYpos/2)-p.HitBoxOffsetY)

	} else if p.facingRight {
		playerOps = pc.playerCam.GetTranslation(playerOps, -p.currentAnimation.PosOffsetX+(x-(pc.playerCXpos/2)-p.HitBoxOffsetX), -p.currentAnimation.PosOffsetY+y-(pc.playerCYpos/2)-p.HitBoxOffsetY)

	} else {
		playerOps = pc.playerCam.GetTranslation(playerOps, p.currentAnimation.PosOffsetX+(-p.HitBoxOffsetX)+x-(pc.playerCXpos/2)-float64(p.currentAnimation.FrameWidth-prevAnim.FrameWidth), -p.currentAnimation.PosOffsetY+y-(pc.playerCYpos/2)-p.HitBoxOffsetY)
	}

	sub := animationSubImage(p.currentAnimation, s, i, p.facingRight, p.currentAnimation.VerticalSheet)

	if hitBoxTest.on && hitBoxTest.frame >= 0 {
		sub = getAnimationFrame(p, hitBoxTest.frame, s)
	}

	// render health
	p.drawHealth(pc.world, currentPlayer)

	// render player
	pc.playerCam.Surface.DrawImage(sub, playerOps)

	// render status effect sprite
	if p.cc == string(sse.Stun) && !p.dead {
		p.drawStatusEffect(currentPlayer, world.playerController)
	}

	// render actionBar
	if currentPlayer {
		drawActionBar(pc, clientConfig.actionBar)
	}

	if clientConfig.showPlayerHitbox {
		rectImg := ebiten.NewImage(int(p.Role.HitBoxW), int(p.Role.HitBoxH))
		hitBoxOps := standardDrawOpts(currentPlayer, pc, p)
		rectImg.Fill(color.RGBA{0, 0, 255, 128})
		hitBoxOps.GeoM.Translate(p.HitBoxOffsetX, p.HitBoxOffsetY)
		pc.playerCam.Surface.DrawImage(rectImg, &hitBoxOps)
	}

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
