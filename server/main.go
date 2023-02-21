package main

import (
	"log"
	"net"

	pb "github.com/kainn9/grpc_game/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(32 * 10e9),
	}

	/*
		disable to skip ssl
		or just run make genSSL
		change value in ssl/ssl.sh
	*/
	tls := false

	if tls {
		// elasticB
		certFile := "./ssl/server.crt"
		keyFile := "./ssl/server.pem"

		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed to load cert %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	EventsHandler()

	pb.RegisterPlayersServiceServer(s, &Server{})

	defer s.Stop()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}

func EventsHandler() {
	go func() {
		for {

			if len(events) == 0 {
				continue
			}

			event := events[0]
			if event == nil {
				continue
			}

			events = events[1:]

			log.Printf("YO %v\n", event)
			EventHandler(event)
		}
	}()
}

func EventHandler(r *pb.PlayerReq) {
	var cp *Player
	pid := r.Id

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
		AddPlayerToSpace(w.Space, cp, 612, 500)
	}

	w.Update(cp, r.Input)
}
