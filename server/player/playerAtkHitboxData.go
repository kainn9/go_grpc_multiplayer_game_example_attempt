package player

import (
	"log"

	r "github.com/kainn9/grpc_game/server/roles"
	"github.com/solarlune/resolv"
)

type hitBoxData struct {
	Player     *Player
	AttackData *r.AttackData
	Aid        string // attack id that applied universal to every hitbox per attack
}

func InitHitboxData(o *resolv.Object, p *Player, atk *r.AttackData) {
	o.Data = &hitBoxData{
		Player:     p,
		AttackData: atk,
	}
}

func assertHitboxData(o *resolv.Object) *hitBoxData {
	if data, ok := o.Data.(*hitBoxData); ok {
		return data
	}

	log.Fatalf("hitbox data is not set, this a critical error: ID:%v\nContents:%v\n", &o.Data, o.Data)
	return nil
}

func HBoxData(o *resolv.Object) *hitBoxData {
	data := assertHitboxData(o)
	return data
}
