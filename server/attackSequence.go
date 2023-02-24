package main

import (
	"time"

	r "github.com/kainn9/grpc_game/server/roles"
	"github.com/solarlune/resolv"
)

func (cp *player) attackSeqence(atk *r.Attack) {
	world := serverConfig.worldsMap[cp.worldKey]
	spawnAtkBox(world.space, cp, atk.HitBoxSequence.MovmentInc, atk.HitBoxSequence.HBoxPath, 0)
}


func spawnAtkBox(space *resolv.Space, cp *player, inc float64, path r.HBoxPath, index int) {
	if len(path) == index {
		cp.currAttack = nil
		cp.chargeValue = 0
		cp.chargeStart = time.Time{}
		return
	}	

	hBoxAgg := path[index]

	hitBoxToClear := make([]*resolv.Object, len(hBoxAgg))
	serverConfig.mutex.Lock()
	for _, hBox := range hBoxAgg {
		
		var atkObj *resolv.Object

		if !cp.facingRight {
			atkObj = resolv.NewObject(cp.object.X- (hBox.PlayerOffX -hBox.Width/2), cp.object.Y+hBox.PlayerOffY, hBox.Width, hBox.Height, "attack")
		} else {
			atkObj = resolv.NewObject(cp.object.X + hBox.PlayerOffX, cp.object.Y + hBox.PlayerOffY, hBox.Width, hBox.Height, "attack")
		}


		serverConfig.AOTP[atkObj] = cp
		serverConfig.OTA[atkObj] = cp.currAttack

	

		hitBoxToClear = append(hitBoxToClear, atkObj)
		space.Add(atkObj)
	}		
	serverConfig.mutex.Unlock()

	
	time.AfterFunc(time.Duration(inc)*time.Millisecond, func() {
		removeHitBoxAggFromAOTPAndOTA(space, hitBoxToClear)
	})
	
	time.AfterFunc(time.Duration(inc)*time.Millisecond, func() {
		spawnAtkBox(space, cp, inc, path, index + 1)
	})
}


func removeHitBoxAggFromAOTPAndOTA(space *resolv.Space, objects []*resolv.Object) {
	serverConfig.mutex.Lock()
	defer serverConfig.mutex.Unlock()

	for _, obj := range objects {

		if serverConfig.AOTP[obj] != nil {
			delete(serverConfig.AOTP, obj)
			space.Remove(obj)
		}

		if serverConfig.OTA[obj] != nil {
			delete(serverConfig.OTA, obj)
		}
		
	}
}

