package main

import (
	"errors"
	"io"
	"log"
	"time"

	pb "github.com/kainn9/grpc_game/proto"
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

/*
	There is no authN or persistance currently.
	Identity boils down to a randomly generated
	"Player ID"(pid) that is created when the
	client first connects to the server and the
	bidirectional stream is setup

	Anwyay, we use the following two functions
	to track which world the player is using these PIDs
*/

func locateFromPID(pid string) (world *world, worldKey string, err error) {

	player := serverConfig.activePlayers[pid]

	if player != nil {
		w := serverConfig.worldsMap[player.worldKey]

		if w != nil {
			return w, player.worldKey, nil
		}

	}

	return nil, "", errors.New("no match found")
}

/*
Helper that uses LocateFromPID
to return a world that a player
is currently attached to. Unlike
LocateFromPID this defaults to returning
the main/starting world
*/
func currentPlayerWorld(pid string) (*world, string) {

	// adding new player to default world
	// or setting current world to where
	// the active player is located
	w, k, err := locateFromPID(pid)

	if err != nil {
		return serverConfig.worldsMap["main"], "main"
	}
	return w, k
}

func (s *server) PlayerLocation(stream pb.PlayersService_PlayerLocationServer) error {

	md, _ := metadata.FromIncomingContext(stream.Context())
	pid := md["pid"][0]
	var prevReq *pb.PlayerReq

	log.Printf("Player Connection Recieved %v\n", pid)

	for {

		w, _ := currentPlayerWorld(pid)

		stalledWrapperInstance := stalledWrapper{stalled: true}

		stalledWrapperInstance.stalledHandler(pid, prevReq)

		req, err := stream.Recv()

		if err == io.EOF {
			log.Printf("EOF")
			removePlayerFromGame(pid, w)
			return nil
		}

		if err != nil {
			switch status.Code(err) {
			case codes.Canceled:

				log.Println("connection has been closed")
				log.Printf("Removing player: %v\n", pid)
				removePlayerFromGame(pid, w)

			default:

				log.Printf("Error while reading client stream %v\n", err)
				log.Printf("Removing player: %v\n", pid)
				removePlayerFromGame(pid, w)
			}

			return nil
		}

		prevReq = req
		stalledWrapperInstance.stalled = false

		initPlayer(req)

		// this can be nil when a player is transfering worlds
		if w.players[pid] != nil {
			w.players[pid].prevEvent = prevReq
		}

		newEvent := newEvent(req, false)
		newEvent.enqueue(w)

		responseHandler(stream, pid)

	}

}

func initPlayer(r *pb.PlayerReq) (*player, *world) {
	var cp *player
	pid := r.Id

	w, k := currentPlayerWorld(pid)

	if serverConfig.activePlayers[pid] == nil {

		cp = newPlayer(pid, k)

		serverConfig.mutex.Lock()
		serverConfig.activePlayers[pid] = cp
		w.players[pid] = cp
		serverConfig.mutex.Unlock()
	} else {
		cp = serverConfig.activePlayers[pid]
	}

	if cp.object == nil {
		addPlayerToSpace(w.space, cp, 612, 500)
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

	serverConfig.mutex.Lock()
	defer serverConfig.mutex.Unlock()

	w, wk := currentPlayerWorld(pid)

	res := &pb.PlayerResp{}

	for k := range w.players {
		curr := w.players[k]

		jumping := curr.onGround == nil && curr.wallSliding == nil

		currAtk := ""
		if curr.currAttack != nil && curr.currAttack.Type != "" {
			currAtk = string(curr.currAttack.Type)
		}

		p := &pb.Player{
			Id:             curr.pid,
			Lx:             curr.object.X,
			Ly:             curr.object.Y,
			FacingRight:    curr.facingRight,
			SpeedX:         curr.speedX,
			SpeedY:         curr.speedY,
			World:          wk,
			Jumping:        jumping,
			CurrAttack:     currAtk,
			CC:             string(curr.isCC()),
			Windup:         string(curr.windup),
			AttackMovement: string(curr.attackMovement),
			Health:         int32(curr.health),
			// Defending:      curr.defending,
		}

		res.Players = append(res.Players, p)
	}

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

					w, _, err := locateFromPID(pid)
					if err != nil {
						return
					}

					prevEvent := newEvent(prevReq, true)
					prevEvent.enqueue(w)
				} else {
					return
				}
			}

		})
	}()
}
