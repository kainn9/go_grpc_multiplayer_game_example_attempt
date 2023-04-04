package main

import (
	"context"
	_ "image/png"
	"io"
	"log"

	_ "net/http/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	pb "github.com/kainn9/grpc_game/proto"
	"github.com/pborman/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var s pb.PlayersService_PlayerLocationClient
var pid string

type Game struct{}

func (g *Game) Draw(screen *ebiten.Image) {}

func (g *Game) Update() error {

	req := pb.PlayerReq{Id: pid, Input: "keySpace"}
	s.Send(&req)

	go func() {

		for {
			_, err := s.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Bot Stream crashed! %v\n", err)
				break
			}
		}
	}()

	return nil
}

func (g *Game) Layout(int, int) (int, int) {
	return 1, 1
}

func main() {
	s = botStream()
	ebiten.RunGame(&Game{})

}

func botStream() pb.PlayersService_PlayerLocationClient {
	pid := uuid.New()
	addr := "localhost:50051"
	opts := []grpc.DialOption{}
	maxSizeOption := grpc.MaxCallRecvMsgSize(32 * 10e9)
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	opts = append(opts, creds)
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	c := pb.NewPlayersServiceClient(conn)

	md := metadata.Pairs("pid", pid)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := c.PlayerLocation(ctx, maxSizeOption)

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	return stream

}
