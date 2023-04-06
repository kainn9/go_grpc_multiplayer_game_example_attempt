package main

import (
	"context"
	_ "image/png"
	"io"
	"log"
	"math"
	_ "net/http/pprof"
	"sync"
	"time"

	pb "github.com/kainn9/grpc_game/proto"
	sr "github.com/kainn9/grpc_game/server/roles"
	u "github.com/kainn9/grpc_game/util"
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
	conn        *grpc.ClientConn
	mutex       sync.Mutex
}

var counter int
var (
	addrPort  = "localhost:50051"
	startTime = time.Now()

	bot = &yoloBot{
		inputQueue:  make([]string, 0),
		leftOrRight: 0,
		pid:         uuid.New(),
	}

	tps = 20
)

// only goes up to 10 as counter caps at 599 ticks(starts at 0)
func secondsPassed(seconds int) bool {
	mod := seconds * tps
	return math.Round(float64(counter%mod)) == float64((seconds*tps)-1)
}

func (b *yoloBot) sendCurrentInput() {

	b.mutex.Lock()
	defer b.mutex.Unlock()

	if len(b.inputQueue) == 1 {
		b.inputQueue = append(b.inputQueue, "nada")
	}

	req := pb.PlayerReq{Id: b.pid, Input: b.inputQueue[0]}
	b.inputQueue = b.inputQueue[1:]

	log.Printf("Sending Bot Req -> %v\n", &req)
	b.stream.Send(&req)
}

func enqueueRandomMovementEvent(b *yoloBot) {

	if secondsPassed(1) {
		b.leftOrRight = u.RandomInt(2)
	}

	// change req
	if b.leftOrRight > 0 {
		b.inputQueue = append(b.inputQueue, "keyLeft")
	} else {
		b.inputQueue = append(b.inputQueue, "keyRight")
	}

}

func enqueueRandomJumpEvent(b *yoloBot) {

	if !secondsPassed(1) {
		return
	}

	jumpMaybe := u.RandomInt(2)

	if jumpMaybe == 0 {
		b.inputQueue = append(b.inputQueue, "keySpace")
	}

}

func enqueueRandomAttackEvent(b *yoloBot) {

	if !secondsPassed(1) {
		return
	}

	attacks := make(map[int]sr.AtKey)

	attacks[0] = sr.PrimaryAttackKey
	attacks[1] = sr.SecondaryAttackKey
	attacks[2] = sr.TertAttackKey
	attacks[3] = sr.QuaternaryAttackKey

	attackMaybe := u.RandomInt(2)

	if attackMaybe == 0 {
		atKey := int(u.RandomInt(4))
		b.inputQueue = append(b.inputQueue, string(attacks[atKey]))
	}

}

func enqueueRandomDefenseEvent(b *yoloBot) {

	if !secondsPassed(2) {
		return
	}

	defMaybe := u.RandomInt(2)

	if defMaybe == 0 {
		b.inputQueue = append(bot.inputQueue, "defense")
	}

}

func (b *yoloBot) sendInputs() {

	enqueueRandomMovementEvent(b)
	enqueueRandomJumpEvent(b)
	enqueueRandomAttackEvent(b)
	enqueueRandomDefenseEvent(b)
	bot.sendCurrentInput()
}

func incrementCounter() {
	counter++
}

func update() error {

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

	go bot.sendInputs()

	incrementCounter()
	return nil
}

func main() {
	bot.initConnAndStream()
	ticker := time.NewTicker(time.Second / time.Duration(tps))
	defer ticker.Stop()

	for range ticker.C {
		update()
	}

}

func (b *yoloBot) initConnAndStream() {

	opts := []grpc.DialOption{}
	maxSizeOption := grpc.MaxCallRecvMsgSize(64 * 10e9)
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	opts = append(opts, creds)
	conn, err := grpc.Dial(addrPort, opts...)

	if err != nil {
		log.Printf("Did not connect: %v", err)
	}

	c := pb.NewPlayersServiceClient(conn)

	md := metadata.Pairs("pid", bot.pid)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := c.PlayerLocation(ctx, maxSizeOption)

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	b.stream = stream
	b.conn = conn

}
