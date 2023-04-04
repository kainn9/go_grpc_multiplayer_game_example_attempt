package main

import (
	"log"
	"net"

	"net/http"
	_ "net/http/pprof"

	pb "github.com/kainn9/grpc_game/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	enablePPROF = false
)

func main() {
	// pprof for debug

	if enablePPROF {
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()

	}

	// Initialize server configuration
	initializer()

	// Listen for incoming connections
	lis, err := net.Listen("tcp", serverConfig.addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	// Log the server's listening address
	log.Printf("Listening at %s\n", serverConfig.addr)

	// Set options for the gRPC server
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(32 * 10e9),
	}

	// Check if SSL should be enabled
	tls := false

	if tls {
		// Load SSL certificate and key files
		certFile := "./ssl/server.crt"
		keyFile := "./ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed to load cert %v\n", err)
		}
		// Add TLS credentials to server options
		opts = append(opts, grpc.Creds(creds))
	}

	// Create a new gRPC server with the given options
	s := grpc.NewServer(opts...)

	// Register the PlayersServiceServer implementation with the gRPC server
	pb.RegisterPlayersServiceServer(s, &server{})

	// Stop the server when the function returns
	defer s.Stop()

	// Start serving incoming requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
