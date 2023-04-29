package worlds

import (
	"sync"
	"time"

	par "github.com/kainn9/grpc_game/server/particles"
	pl "github.com/kainn9/grpc_game/server/player"
	r "github.com/kainn9/grpc_game/server/roles"
	"github.com/solarlune/resolv"
)

/*
World receiver functions required to implement the player world interface /
intended to be used in player package
*/

func (w *World) GetHitboxMutex() *sync.RWMutex {
	return &w.hitboxMutex
}

func (w *World) GetSpace() *resolv.Space {
	return w.space
}

func (w *World) GetIndex() int {
	return w.Index
}

func (w *World) GetHeight() float64 {
	return w.height
}

func (w *World) GetParticleSystem() *par.ParticleSystem {
	return w.ParticleSystem
}

func (world *World) SpawnAtkBox(cp *pl.Player, atk *r.AttackData, index int, aid string) {
	path := atk.HitBoxSequence.HBoxPath
	inc := atk.HitBoxSequence.MovmentInc

	/*
		if currAttack is nil, cp was hit in last frame
		if len(path) == index attack sequence is over
		in both cases attack is over, enter cleanup block
	*/
	if len(path) == index || cp.CurrAttack == nil {
		cp.CurrAttack = nil
		cp.ChargeValue = 0
		cp.ChargeStart = time.Time{}

		if cp.InvincibleNoBox {
			cp.InvincibleNoBox = false
		}
		return
	}

	hBoxAgg := path[index]

	for _, hBox := range hBoxAgg {
		world.GetHitboxMutex().Lock()

		var atkObj *resolv.Object

		if !cp.FacingRight {
			atkObj = resolv.NewObject(cp.Object.X-hBox.PlayerOffX-hBox.Width+cp.Object.W, cp.Object.Y+hBox.PlayerOffY, hBox.Width, hBox.Height, "attack")
		} else {
			atkObj = resolv.NewObject(cp.Object.X+hBox.PlayerOffX, cp.Object.Y+hBox.PlayerOffY, hBox.Width, hBox.Height, "attack")
		}

		pl.InitHitboxData(atkObj, cp, cp.CurrAttack)

		hitBoxData := *pl.HBoxData(atkObj)

		hitBoxData.Player = cp
		hitBoxData.AttackData = cp.CurrAttack

		world.GetSpace().Add(atkObj)
		world.GetHitboxMutex().Unlock()

		time.AfterFunc(time.Duration(inc)*time.Millisecond, func() {
			world.RemoveHitboxFromSpace(atkObj)
		})
	}

	time.AfterFunc(time.Duration(inc)*time.Millisecond, func() {
		world.SpawnAtkBox(cp, atk, index+1, aid)
	})
}

func (world *World) RemoveHitboxFromSpace(obj *resolv.Object) {
	world.hitboxMutex.Lock()
	world.space.Remove(obj)
	world.hitboxMutex.Unlock()
}
