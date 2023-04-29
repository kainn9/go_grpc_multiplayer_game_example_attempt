package particle

import (
	"image/color"

	vector "github.com/kainn9/grpc_game/util"
)

type Particle struct {
	Position vector.Vector2
	Velocity vector.Vector2
	Size     float64
	Color    color.Color
	Id       float64 // Which sprite to render clientSide
	Lifetime float64
	Age      float64
	Damage   float64
	Active   bool
}
