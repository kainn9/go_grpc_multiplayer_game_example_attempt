package event

import (
	pb "github.com/kainn9/grpc_game/proto"
	sr "github.com/kainn9/grpc_game/server/roles"
)

type Event struct {
	*pb.PlayerReq
	Stalled       bool
	EventCategory EventCategory
}

type EventCategory string

const (
	EventRestricted EventCategory = "rest"
	EventNormal     EventCategory = "norm"
	EventAdmin      EventCategory = "admin"
)

type EventInput string

const (
	KeyLeft      EventInput = "keyLeft"
	KeyRight     EventInput = "keyRight"
	KeySpace     EventInput = "keySpace"
	KeyDown      EventInput = "keyDown"
	PrimaryAtk   EventInput = EventInput(sr.PrimaryAttackKey)
	SecondaryAtk EventInput = EventInput(sr.SecondaryAttackKey)
	TertAtk      EventInput = EventInput(sr.TertAttackKey)
	QuaAtk       EventInput = EventInput(sr.QuaternaryAttackKey)
	Defense      EventInput = "defense"
	Nada         EventInput = "nada"
	Swap         EventInput = "swap"
	RoleSwap     EventInput = "roleSwap"
)

var ValidEvents = map[EventInput]EventInput{
	KeyLeft:      KeyLeft,
	KeyRight:     KeyRight,
	KeySpace:     KeySpace,
	KeyDown:      KeyDown,
	PrimaryAtk:   PrimaryAtk,
	SecondaryAtk: SecondaryAtk,
	TertAtk:      TertAtk,
	QuaAtk:       QuaAtk,
	Defense:      Defense,
	Nada:         Nada,
	Swap:         Swap,
	RoleSwap:     RoleSwap,
}

var AttackEvents = map[EventInput]EventInput{
	PrimaryAtk:   PrimaryAtk,
	SecondaryAtk: SecondaryAtk,
	TertAtk:      TertAtk,
	QuaAtk:       QuaAtk,
}

// note we don't include keySpace/Jump
// to allow "double up" for this event
var PhysEvents = map[EventInput]EventInput{
	KeyLeft:  KeyLeft,
	KeyRight: KeyRight,
	KeyDown:  KeyDown,
}

var AdminEvents = map[EventInput]EventInput{
	Swap:     Swap,
	RoleSwap: RoleSwap,
}

func NewEvent(req *pb.PlayerReq, stalled bool) *Event {

	e := &Event{
		PlayerReq: req,
		Stalled:   stalled,
	}

	if PhysEvents[EventInput(req.Input)] != "" || AttackEvents[EventInput(req.Input)] != "" {
		e.EventCategory = EventRestricted
	} else if AdminEvents[EventInput(req.Input)] != "" {
		e.EventCategory = EventAdmin
	} else {
		e.EventCategory = EventNormal
	}

	return e
}

func (e *Event) Valid() bool {
	return ValidEvents[EventInput(e.Input)] != ""
}
