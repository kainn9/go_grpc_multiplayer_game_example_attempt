package roles

import (
	"github.com/hajimehoshi/ebiten/v2"
)

/*
	Root file for roles(player classes i.e., mage, knight, etc)

	TODO: Questionable if this should be its own
	package(creates scope issues)
*/

type Role struct {
	Animations         map[string]*Animation
	HitBoxOffsetX      float64
	HitBoxOffsetY      float64
	Health             int
	HitBoxW            float64
	HitBoxH            float64
	HealthBarOffset    *Offset
	StatusEffectOffset *Offset
	AttackCount        int
	HasDefense         bool
}

type Animation struct {
	Fixed       bool
	FrameOX     int
	FrameOY     int
	FrameWidth  int
	FrameHeight int
	FrameCount  int
	Name        string
	SpriteSheet *ebiten.Image
	PosOffsetX  float64
	PosOffsetY  float64
}

type AnimKey string

type Offset struct {
	X float64
	Y float64
}

const (
	DeathRight   AnimKey = "deathRight"
	DeathLeft    AnimKey = "deathLeft"
	IdleRight    AnimKey = "idleRight"
	IdleLeft     AnimKey = "idleLeft"
	WalkRight    AnimKey = "walkRight"
	WalkLeft     AnimKey = "walkLeft"
	JumpLeft     AnimKey = "jumpLeft"
	JumpRight    AnimKey = "jumpRight"
	KbRight      AnimKey = "kbRight"
	KbLeft       AnimKey = "kbLeft"
	HitRight     AnimKey = "hitRight"
	HitLeft      AnimKey = "hitLeft"
	StunRight    AnimKey = "stunRight"
	StunLeft     AnimKey = "stunLeft"
	DefenseRight AnimKey = "defenseRight"
	DefenseLeft  AnimKey = "defenseLeft"
)
