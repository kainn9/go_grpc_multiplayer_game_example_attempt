package main

// TODO move these out of Global Scope and into "config" structs
import (
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
	ut "github.com/kainn9/grpc_game/util"
	"google.golang.org/grpc"
)

/*
	Putting global scope client vars and consts here to avoid confusion
*/

const (
	ScreenWidth  = 880
	ScreenHeight = 480
)

var (
	streamInit = false
	ticks      int
	worldsMap  = make(map[string]WorldData)
	game       = NewGame()
	addr       = "localhost:50051"
	fullScreen = false
	// Use This for irl network testing
	// addr        = "ec2-3-83-121-221.compute-1.amazonaws.com:50051"
	enablePPROF = false
)

// dev mode stuff
var (
	rulerW         = ut.LoadImg("./sprites/rulers/wRuler.png")
	rulerH         = ut.LoadImg("./sprites/rulers/hRuler.png")
	devPreview     = false
	useHeightRuler = false
	devCamSpeed    = float64(2)
	freePlay       = false
	ping           = 0.0
	reqT           = time.Now()
)

var (
	mainWorldBg = ut.LoadImg("./backgrounds/mapMain.png")
	altWorldBg  = ut.LoadImg("./backgrounds/mapAlt.png")
)

var audPlayer *audio.Player
var volume128 int
var connRef *grpc.ClientConn
