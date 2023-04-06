package main

import (
	"context"
	_ "image/png"
	"io"
	"log"
	"math"
	_ "net/http/pprof"
	"time"

	pb "github.com/kainn9/grpc_game/proto"
	sr "github.com/kainn9/grpc_game/server/roles"
	ut "github.com/kainn9/grpc_game/util"
	"github.com/pborman/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

/*
Simple Script to load "BOTS"(they're very dumb) into the game
since loading multiple clients gets to heavy for proper testing

Note: they do crash eventually, but gets the job done for now.
*/
type yoloBot struct {
	leftOrRight int64
	inputQueue  []string
	pid         string
	stream      pb.PlayersService_PlayerLocationClient
}

var counter int

var (
	startTime = time.Now()

	bot = &yoloBot{
		inputQueue:  make([]string, 0),
		leftOrRight: 0,
		pid:         uuid.New(),
	}
)

// only goes up to 10 as counter caps at 599 ticks(starts at 0)
func secondsPassed(seconds int) bool {
	mod := seconds * 60
	return math.Round(float64(counter%mod)) == float64((seconds*60)-1)
}

func (*yoloBot) sendCurrentInput() {

	if len(bot.inputQueue) == 1 {

		bot.inputQueue = append(bot.inputQueue, "keyRight")
	}

	req := pb.PlayerReq{Id: bot.pid, Input: bot.inputQueue[0]}
	bot.inputQueue = bot.inputQueue[1:]

	log.Printf("Sending Bot Req -> %v: %v\n", bot.pid, bot.inputQueue[0])
	bot.stream.Send(&req)
}

func sendRandomMovementEvent() {

	if secondsPassed(3) {
		bot.leftOrRight = ut.RandomInt(2)
	}

	// change req
	if bot.leftOrRight > 0 {
		bot.inputQueue = append(bot.inputQueue, "keyLeft")
	} else {
		bot.inputQueue = append(bot.inputQueue, "keyRight")
	}

}

func sendRandomJumpEvent() {

	if !secondsPassed(2) {
		return
	}

	jumpMaybe := ut.RandomInt(4)

	if jumpMaybe == 0 {
		bot.inputQueue = append(bot.inputQueue, "keySpace")
	}

}

func sendRandomAttackEvent() {

	if !secondsPassed(5) {
		return
	}

	attacks := make(map[int]sr.AtKey)

	attacks[0] = sr.PrimaryAttackKey
	attacks[1] = sr.SecondaryAttackKey
	attacks[2] = sr.TertAttackKey
	attacks[3] = sr.QuaternaryAttackKey

	attackMaybe := ut.RandomInt(3)

	if attackMaybe == 0 {
		atKey := int(ut.RandomInt(4))
		bot.inputQueue = append(bot.inputQueue, string(attacks[atKey]))
	}

}

func sendRandomDefenseEvent() {

	if !secondsPassed(7) {
		return
	}

	defMaybe := ut.RandomInt(2)

	if defMaybe == 0 {
		bot.inputQueue = append(bot.inputQueue, "defense")
	}

}

func (*yoloBot) sendInputs() {
	sendRandomMovementEvent()
	sendRandomJumpEvent()
	sendRandomAttackEvent()
	sendRandomDefenseEvent()
	bot.sendCurrentInput()
}

func incrementCounter() {
	counter++
}

func update() error {

	bot.sendInputs()

	go func() {

		for {
			r, err := bot.stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {

				duration := time.Since(startTime)
				log.Fatalf("Bot Stream crashed! %v - %v\n", err, duration.Seconds())
				break
			}

			// toggle true if you want this log
			if false {
				log.Printf("State %v\n", r.Players)
			}

		}
	}()

	incrementCounter()
	return nil
}

func main() {
	bot.stream = botStream()

	ticker := time.NewTicker(time.Second / 60)
	defer ticker.Stop()

	for range ticker.C {
		update()
	}

}

func botStream() pb.PlayersService_PlayerLocationClient {

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

	md := metadata.Pairs("pid", bot.pid)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := c.PlayerLocation(ctx, maxSizeOption)

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	return stream

}
