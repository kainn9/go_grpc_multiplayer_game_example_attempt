package main

import (
	"io"
	"log"
	"time"

	pb "github.com/kainn9/grpc_game/proto"
	evt "github.com/kainn9/grpc_game/server/event"
	particle "github.com/kainn9/grpc_game/server/particles"
	pl "github.com/kainn9/grpc_game/server/player"
	wr "github.com/kainn9/grpc_game/server/worlds"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.PlayersServiceServer
}

type stalledWrapper struct {
	stalled bool
}

func (s *server) PlayerLocation(stream pb.PlayersService_PlayerLocationServer) error {

	md, _ := metadata.FromIncomingContext(stream.Context())
	pid := md["pid"][0]
	var prevReq *pb.PlayerReq

	log.Printf("Player Connection Recieved %v\n", pid)

	for {

		w, _ := wr.CurrentPlayerWorldOrDefault(pid)

		stalledWrapperInstance := stalledWrapper{stalled: true}

		stalledWrapperInstance.stalledHandler(pid, prevReq)

		req, err := stream.Recv()

		if err == io.EOF {
			log.Printf("EOF")
			wr.RemovePlayerFromWorld(pid, w)
			return nil
		}

		if err != nil {
			switch status.Code(err) {
			case codes.Canceled:

				log.Println("connection has been closed")
				log.Printf("Removing player: %v\n", pid)
				wr.RemovePlayerFromWorld(pid, w)

			default:

				log.Printf("Error while reading client stream %v\n", err)
				log.Printf("Removing player: %v\n", pid)
				wr.RemovePlayerFromWorld(pid, w)
			}

			return nil
		}

		prevReq = req
		stalledWrapperInstance.stalled = false

		initPlayer(req)

		// this can be nil when a player is transfering worlds
		w.WPlayersMutex.Lock()
		p := w.Players[pid]

		if p != nil {
			w.Players[pid].PrevEvent = prevReq
		}
		w.WPlayersMutex.Unlock()

		newEvent := evt.NewEvent(req, false)
		wr.Enqueue(w, newEvent)

		responseHandler(stream, pid)

	}

}

func initPlayer(r *pb.PlayerReq) (*pl.Player, *wr.World) {
	var cp *pl.Player
	pid := r.Id

	w, _ := wr.CurrentPlayerWorldOrDefault(pid)

	wr.WorldsConfig.Mutex.RLock()
	activePlayer := wr.WorldsConfig.ActivePlayers[pid]
	wr.WorldsConfig.Mutex.RUnlock()

	if activePlayer == nil {

		cp = pl.NewPlayer(pid, w)

		wr.WorldsConfig.Mutex.Lock()
		wr.WorldsConfig.ActivePlayers[pid] = cp
		wr.WorldsConfig.Mutex.Unlock()

		w.WPlayersMutex.Lock()
		w.Players[pid] = cp
		w.WPlayersMutex.Unlock()

	} else {
		cp = activePlayer
	}

	if cp.Object == nil {
		wr.AddPlayerToSpace(w, cp, float64(w.WorldSpawnCords.X), float64(w.WorldSpawnCords.Y))
	}

	return cp, w
}

/*
Handles sending back the game state to client.
Currently only sends the data for the world
that the currentPlayer(client on the other
side of stream) resides inside of
*/
func responseHandler(stream pb.PlayersService_PlayerLocationServer, pid string) {

	w, wk := wr.CurrentPlayerWorldOrDefault(pid)
	res := &pb.PlayerResp{}

	w.WPlayersMutex.RLock()
	for k := range w.Players {
		curr := w.Players[k]

		if curr == nil || curr.Object == nil {
			continue
		}
		jumping := curr.OnGround == nil && curr.WallSliding == nil

		currAtk := ""
		if curr.CurrAttack != nil && curr.CurrAttack.Type != "" {
			currAtk = string(curr.CurrAttack.Type)
		}

		p := &pb.Player{
			Id:             curr.Pid,
			Lx:             curr.Object.X,
			Ly:             curr.Object.Y,
			FacingRight:    curr.FacingRight,
			SpeedX:         curr.SpeedX,
			SpeedY:         curr.SpeedY,
			World:          int32(wk),
			Jumping:        jumping,
			CurrAttack:     currAtk,
			CC:             string(curr.IsCC()),
			Windup:         string(curr.Windup),
			AttackMovement: string(curr.AttackMovement),
			Health:         int32(curr.Health),
			Defending:      curr.Defending,
			Role:           wr.WorldsConfig.Roles[curr.Role],
			Dead:           curr.Dying,
			Cooldowns:      curr.CdString,
		}

		res.Players = append(res.Players, p)
	}
	w.WPlayersMutex.RUnlock()

	// Convert ParticleSystem to proto version (look for better way)
	testParticleSystem := w.ParticleSystem
	res.ParticleSystem = particle.ConvertToProtoParticleSystem(testParticleSystem)

	err := stream.Send(res)

	if err != nil {
		log.Fatalf("Error while sending data to client: %v\n", err)
	}
}

func (s *stalledWrapper) stalledHandler(pid string, prevReq *pb.PlayerReq) {

	go func() {
		// Note:
		// best on something local testing looks like
		// we don't need the timeAFterFunc since the ticker handles the 16.666ms
		// delay, but leaving for now in caseI need to change it back at some point
		time.AfterFunc(0*time.Millisecond, func() {

			ticker := time.NewTicker(time.Second / 60)
			defer ticker.Stop()

			for range ticker.C {
				// Questionable if prevRequest should be used.
				if prevReq != nil && s.stalled {

					w, _, err := wr.LocateFromPID(pid)
					if err != nil {
						return
					}

					prevEvent := evt.NewEvent(prevReq, true)
					wr.Enqueue(w, prevEvent)
				} else {
					return
				}
			}

		})
	}()
}
