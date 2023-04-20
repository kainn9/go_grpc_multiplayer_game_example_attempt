package player

import (
	r "github.com/kainn9/grpc_game/server/roles"
	"github.com/pborman/uuid"
)

func (cp *Player) attackSeqence(atk *r.AttackData) {
	world := cp.CurrentWorld
	aid := uuid.New()
	world.SpawnAtkBox(cp, atk, 0, aid)
}
