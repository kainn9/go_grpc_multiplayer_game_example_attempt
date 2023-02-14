package roles

import "github.com/hajimehoshi/ebiten/v2"

type PlayerType string

type Role struct {
	RoleType      PlayerType
	Animations    map[string]*Animation
	HitBoxOffsetX float64
	HitBoxOffsetY float64
}

type Animation struct {
	FrameOX     int
	FrameOY     int
	FrameWidth  int
	FrameHeight int
	FrameCount  int
	Name        string
	SpriteSheet *ebiten.Image
}

const (
	KnightType PlayerType = "knight"
	MageType   PlayerType = "mage"
)
