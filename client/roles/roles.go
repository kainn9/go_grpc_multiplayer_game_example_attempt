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
