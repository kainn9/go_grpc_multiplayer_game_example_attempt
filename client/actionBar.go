package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	utClient "github.com/kainn9/grpc_game/client_util"
	sr "github.com/kainn9/grpc_game/server/roles"
)

// Temp implementation for demo video
// is trash tho

type actionBar struct {
	opts            *ebiten.DrawImageOptions
	bg              *ebiten.Image
	qIconSprite     *ebiten.Image
	wIconSprite     *ebiten.Image
	eIconSprite     *ebiten.Image
	rIconSprite     *ebiten.Image
	shiftIconSprite *ebiten.Image

	atkActiveSprite   *ebiten.Image
	atkIconNoCdSprite *ebiten.Image
	atkIconOnCdSprite *ebiten.Image

	defActiveSprite   *ebiten.Image
	defIconNoCdSprite *ebiten.Image
	defIconOnCdSprite *ebiten.Image

	coverSprite       *ebiten.Image
	padding           int
	iconWidth         int
	actionBarBgHeight int
	atkOrderMap       *atkOrderMap
	lettersMap        map[int]*ebiten.Image
}
type atkOrderMap map[int]sr.AtKey

func initActionBar() *actionBar {

	iconWidth := 24
	iconHeight := iconWidth
	padding := 2

	actionBarBgWidth := 1
	actionBarBgHeight := iconHeight + (padding * 2)

	drawOpts := &ebiten.DrawImageOptions{}
	drawOpts.GeoM.Translate(float64((clientConfig.screenWidth/2)-(actionBarBgWidth/2)), float64(clientConfig.screenHeight-actionBarBgHeight))

	actionBarBg := ebiten.NewImage(actionBarBgWidth, actionBarBgHeight)
	actionBarBg.Fill(color.RGBA{0, 0, 0, 255})

	coverSprite := ebiten.NewImage(iconWidth, iconHeight)
	coverSprite.Fill(color.RGBA{0, 0, 0, 255})

	atkOrderMap := &atkOrderMap{
		0: sr.PrimaryAttackKey,
		1: sr.SecondaryAttackKey,
		2: sr.TertAttackKey,
		3: sr.QuaternaryAttackKey,
	}

	ab := &actionBar{
		bg:                actionBarBg,
		opts:              drawOpts,
		padding:           padding,
		coverSprite:       coverSprite,
		atkOrderMap:       atkOrderMap,
		iconWidth:         iconWidth,
		actionBarBgHeight: actionBarBgHeight,
	}
	initIcons(ab)
	ab.lettersMap = map[int]*ebiten.Image{
		0: ab.qIconSprite,
		1: ab.wIconSprite,
		2: ab.eIconSprite,
		3: ab.rIconSprite,
	}

	return ab
}

func initIcons(ab *actionBar) {
	ab.qIconSprite = utClient.LoadImage("./sprites/actionBar/q.png")
	ab.wIconSprite = utClient.LoadImage("./sprites/actionBar/w.png")
	ab.eIconSprite = utClient.LoadImage("./sprites/actionBar/e.png")
	ab.rIconSprite = utClient.LoadImage("./sprites/actionBar/r.png")
	ab.shiftIconSprite = utClient.LoadImage("./sprites/actionBar/shift.png")

	ab.atkActiveSprite = utClient.LoadImage("./sprites/actionBar/atkIconActive.png")
	ab.atkIconNoCdSprite = utClient.LoadImage("./sprites/actionBar/atkIconNoCd.png")
	ab.atkIconOnCdSprite = utClient.LoadImage("./sprites/actionBar/atkIconOnCd.png")

	ab.defActiveSprite = utClient.LoadImage("./sprites/actionBar/defIconActive.png")

	ab.defIconNoCdSprite = utClient.LoadImage("./sprites/actionBar/defIconNoCd.png")
	ab.defIconOnCdSprite = utClient.LoadImage("./sprites/actionBar/defIconOnCd.png")
}

func drawActionBar(pc *PlayerController, ab *actionBar) {

	p := pc.world.playerMap[pc.pid]
	if p == nil {
		return
	}

	pc.playerCam.Surface.DrawImage(ab.bg, ab.opts)

	atkIconNoCdOpts := &ebiten.DrawImageOptions{}
	atkIconNoCdOpts.GeoM.Translate(2, 2)

	bgWidth, _ := ab.bg.Size()

	count := p.AttackCount
	if p.HasDefense {
		count += 1
	}

	properWidth := (count * ab.iconWidth) + ((count + 1) * ab.padding)

	if bgWidth != properWidth {
		ab.bg = ebiten.NewImage(properWidth, ab.actionBarBgHeight)

		newOpts := &ebiten.DrawImageOptions{}
		newOpts.GeoM.Translate(float64((clientConfig.screenWidth/2)-(properWidth/2)), float64(clientConfig.screenHeight-ab.actionBarBgHeight))
		ab.opts = newOpts
	}

	for i := 0; i < count; i++ {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(ab.padding+(i*24)), float64(ab.padding))

		if i == count-1 && p.HasDefense {
			ab.drawDefIcon(p, opts)
		} else if i < p.AttackCount {
			ab.drawAtkIcon(p, i, opts)
		} else {
			ab.bg.DrawImage(ab.coverSprite, opts)
		}

	}

}

func (ab *actionBar) drawAtkIcon(p *Player, i int, opts *ebiten.DrawImageOptions) {

	if p.currAttack == string((*ab.atkOrderMap)[i]) || p.windup == string((*ab.atkOrderMap)[i]) {
		ab.bg.DrawImage(ab.atkActiveSprite, opts)
	} else if string(p.cooldowns[i]) == "1" {
		ab.bg.DrawImage(ab.atkIconOnCdSprite, opts)
	} else {
		ab.bg.DrawImage(ab.atkIconNoCdSprite, opts)
		ab.bg.DrawImage(ab.lettersMap[i], opts)
	}

}

func (ab *actionBar) drawDefIcon(p *Player, opts *ebiten.DrawImageOptions) {
	if p.defending {
		ab.bg.DrawImage(ab.defActiveSprite, opts)
	} else if string(p.cooldowns[4]) == "1" {
		ab.bg.DrawImage(ab.defIconOnCdSprite, opts)
	} else {
		ab.bg.DrawImage(ab.defIconNoCdSprite, opts)

		shiftOpts := *opts
		shiftOpts.GeoM.Translate(0, 8)
		ab.bg.DrawImage(ab.shiftIconSprite, &shiftOpts)
	}
}
