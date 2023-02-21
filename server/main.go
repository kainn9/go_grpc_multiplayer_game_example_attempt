package main

import (
	"log"
	"net"
	"sync"
	"time"

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
		certFile := "./ssl/server.crt"
		keyFile := "./ssl/server.pem"

		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed to load cert %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	centralTickLoop()

	s := grpc.NewServer(opts...)

	pb.RegisterPlayersServiceServer(s, &Server{})

	defer s.Stop()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}

// TODO: Scope these to World/Arena
var (
	events = make([]*pb.PlayerReq, 0)
	kotex  sync.RWMutex
)

func centralTickLoop() {
	go func() {
		ticker := time.NewTicker(time.Second / 60)
		defer ticker.Stop()

		for range ticker.C {
			// should this be concurrent?
			processEventsPerTick()

		}
	}()
}

func processEventsPerTick() {
	if len(events) > 25 {
		log.Printf("LEN! 25 %v\n", len(events) > 25)
		log.Printf("LEN! 50 %v\n", len(events) > 50)
		log.Printf("LEN! 100 %v\n", len(events) > 100)
	}

	// iterate over the events while removing the current element
	for i := 0; i < 100; i++ {

		if len(events) == 0 || i > len(events)-1 {
			break
		}

		// process event here
		ev := events[i]

		w, _, err := LocateFromPID(ev.Id)
		if err == nil {

			cp := w.Players[ev.Id]

			if cp != nil {
				requestHandler(cp, w, ev)
			}

		}

		kotex.RLock()
		events = append(events[:i], events[i+1:]...)
		defer kotex.RUnlock()
		i--
	}

}
