package main

import (
	"errors"
	"io"
	"log"

	pb "github.com/kainn9/grpc_game/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.PlayersServiceServer
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

func LocateFromPID(pid string) (world *World, worldKey string, err error) {

	player := activePlayers[pid]

	if player != nil {
		w := worldsMap[player.WorldKey]

		return w, player.WorldKey, nil
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
func CurrentPlayerWorld(pid string) (*World, string) {

	// adding new player to default world
	// or setting current world to where
	// the active player is located
	w, k, err := LocateFromPID(pid)

	if err != nil {
		return worldsMap["main"], "main"
	}
	return w, k
}

func (s *Server) PlayerLocation(stream pb.PlayersService_PlayerLocationServer) error {
	// Available Maps
	worldsMap["main"] = mainW
	worldsMap["alt"] = altW

	md, _ := metadata.FromIncomingContext(stream.Context())
	pid := md["pid"][0]

	log.Printf("Player Connection Recieved %v\n", pid)

	for {
		w, _ := CurrentPlayerWorld(pid)
		req, err := stream.Recv()

		if err == io.EOF {
			log.Printf("EOF")
			DisconnectPlayer(pid, w)
			return nil
		}

		if err != nil {
			switch status.Code(err) {
			case codes.Canceled:

				log.Println("connection has been closed")
				log.Printf("Removing player: %v\n", pid)
				DisconnectPlayer(pid, w)

			default:

				log.Printf("Error while reading client stream %v\n", err)
				log.Printf("Removing player: %v\n", pid)
				DisconnectPlayer(pid, w)
			}

			return nil
		}

		requestHandler(req, pid)
		responseHandler(stream, pid)

	}
}

/*
Handle incoming stream from client
*/
func requestHandler(r *pb.PlayerReq, pid string) {
	var cp *Player
	w, k := CurrentPlayerWorld(pid)

	if activePlayers[pid] == nil {

		cp = NewPlayer(pid, k)

		mutex.Lock()
		activePlayers[pid] = cp
		w.Players[pid] = cp
		mutex.Unlock()
	} else {
		cp = activePlayers[pid]
	}

	if cp.Object == nil {
		AddPlayerToSpace(w.Space, cp, 1250, 3700)
	}

	w.Update(cp, r.Input)
}

/*
Handles sending back the game state to client.
Currently only sends the data for the world
that the currentPlayer(client on the other
side of stream) resides inside of
*/
func responseHandler(stream pb.PlayersService_PlayerLocationServer, pid string) {

	mutex.Lock()
	defer mutex.Unlock()

	w, wk := CurrentPlayerWorld(pid)

	res := &pb.PlayerResp{}

	for k := range w.Players {
		curr := w.Players[k]

		jumping := curr.OnGround == nil && curr.WallSliding == nil

		currAtk := ""
		if curr.CurrAttack != nil && curr.CurrAttack.Type != "" {
			currAtk = string(curr.CurrAttack.Type)
		}

		p := &pb.Player{
			Id:          curr.Pid,
			Lx:          curr.Object.X,
			Ly:          curr.Object.Y,
			FacingRight: curr.FacingRight,
			SpeedX:      curr.SpeedX,
			SpeedY:      curr.SpeedY,
			World:       wk,
			Jumping:     jumping,
			CurrAttack:  currAtk,
			CC:          string(curr.isCC()),
		}

		res.Players = append(res.Players, p)
	}

	err := stream.Send(res)
	if err != nil {
		log.Fatalf("Error while sending data to client: %v\n", err)
	}
}
