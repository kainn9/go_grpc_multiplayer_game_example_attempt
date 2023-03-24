package main

import (
	"log"

	r "github.com/kainn9/grpc_game/server/roles"
	"github.com/solarlune/resolv"
)

type hitBoxData struct {
	player     *player
	attackData *r.AttackData
	aid        string // attack id that applied universal to every hitbox per attack
}

func initHitboxData(o *resolv.Object, p *player, atk *r.AttackData) {
	o.Data = &hitBoxData{
		player:     p,
		attackData: atk,
	}
}

func assertHitboxData(o *resolv.Object) *hitBoxData {
	if data, ok := o.Data.(*hitBoxData); ok {
		return data
	}

	log.Fatalf("hitbox data is not set, this a critical error: ID:%v\nContents:%v\n", &o.Data, o.Data)
	return nil
}

func hBoxData(o *resolv.Object) *hitBoxData {
	data := assertHitboxData(o)
	return data
}
