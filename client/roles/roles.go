package roles

import "github.com/hajimehoshi/ebiten/v2"

/*
	Root file for roles(player classes i.e., mage, knight, etc)

	TODO: Questionable if this should be its own
	package(creates scope issues)
*/
type PlayerType string

type Role struct {
	RoleType      PlayerType
	Animations    map[string]*Animation
	HitBoxOffsetX float64
	HitBoxOffsetY float64
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

const (
	KnightType   PlayerType = "knight"
	MageType     PlayerType = "mage"
	MonkType     PlayerType = "monk"
	DemonType    PlayerType = "demon"
	WerewolfType PlayerType = "werewolf"
)

type AnimKey string

const (
	DeathRight   AnimKey = "deathRight"
	DeathLeft    AnimKey = "deathLeft"
	IdleRight    AnimKey = "idleRight"
	IdleLeft     AnimKey = "idleLeft"
	WalkRight    AnimKey = "walkRight"
	WalkLeft     AnimKey = "walkLeft"
	JumpLeft     AnimKey = "jumpLeft"
	JumpRight    AnimKey = "jumpRight"
	KbRight      AnimKey = "KbRight"
	KbLeft       AnimKey = "KbLeft"
	DefenseRight AnimKey = "defenseRight"
	DefenseLeft  AnimKey = "defenseLeft"
)
