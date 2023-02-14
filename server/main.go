package main

import (
	"log"
	"net"
	"sync"

	pb "github.com/kainn9/grpc_game/proto"
	"github.com/solarlune/resolv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	mutex         sync.RWMutex
	addr          string = ":50051"
	mainW                = NewWorld(4000, 6000, MainWorldBuilder, "main")
	altW                 = NewWorld(848, 3200, AltWorldBuilder, "alt")
	worldsMap            = make(map[string]*World)
	activePlayers        = make(map[string]*Player)
	AOTP                 = make(map[*resolv.Object]*Player) // map of Attack resolv objects to Player struct
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

	pb.RegisterPlayersServiceServer(s, &Server{})

	defer s.Stop()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
