package main

import (
	"time"

	r "github.com/kainn9/grpc_game/server/roles"
	"github.com/pborman/uuid"
	"github.com/solarlune/resolv"
)

func (cp *player) attackSeqence(atk *r.AttackData) {
	world := serverConfig.worldsMap[cp.worldKey]
	aid := uuid.New()
	spawnAtkBox(world, cp, atk, 0, aid)
}

func spawnAtkBox(world *world, cp *player, atk *r.AttackData, index int, aid string) {
	path := atk.HitBoxSequence.HBoxPath
	inc := atk.HitBoxSequence.MovmentInc

	/*
		if currAttack is nil, cp was hit in last frame
		if len(path) == index attack sequence is over
		in both cases attack is over, enter cleanup block
	*/
	if len(path) == index || cp.currAttack == nil {
		cp.currAttack = nil
		cp.chargeValue = 0
		cp.chargeStart = time.Time{}
		return
	}

	hBoxAgg := path[index]

	for _, hBox := range hBoxAgg {

		var atkObj *resolv.Object

		if !cp.facingRight {
			atkObj = resolv.NewObject(cp.object.X-(hBox.PlayerOffX-hBox.Width/2), cp.object.Y+hBox.PlayerOffY, hBox.Width, hBox.Height, "attack")
		} else {
			atkObj = resolv.NewObject(cp.object.X+hBox.PlayerOffX, cp.object.Y+hBox.PlayerOffY, hBox.Width, hBox.Height, "attack")
		}

		initHitboxData(atkObj, cp, cp.currAttack)

		hitBoxData := *hBoxData(atkObj)

		hitBoxData.player = cp
		hitBoxData.attackData = cp.currAttack


		world.hitboxMutex.Lock()
		world.space.Add(atkObj)
		world.hitboxMutex.Unlock()


		time.AfterFunc(time.Duration(inc)*time.Millisecond, func() {
			removeHitboxFromSpace(world, atkObj)
		})
	}

	time.AfterFunc(time.Duration(inc)*time.Millisecond, func() {
		spawnAtkBox(world, cp, atk, index+1, aid)
	})
}

func removeHitboxFromSpace(world *world, obj *resolv.Object) {
	world.hitboxMutex.Lock()
	defer world.hitboxMutex.Unlock()

	world.space.Remove(obj)
}
