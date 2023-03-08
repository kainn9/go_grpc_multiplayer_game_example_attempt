package main

// TODO move these out of Global Scope and into "config" structs
import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	ut "github.com/kainn9/grpc_game/util"
	"google.golang.org/grpc"
)

/*
	Putting global scope client vars and consts here to avoid confusion
*/

type gameSettings struct {
	screenWidth  int
	screenHeight int
	streamInit   bool
	ticks        int
	worldsMap    map[string]worldData
	game         *Game
	addr         string
	fullScreen   bool
	enablePPROF  bool
	connRef      *grpc.ClientConn
	audPlayer    *audio.Player
	volume128    int
}

type devSettings struct {
	rulerW         *ebiten.Image
	rulerH         *ebiten.Image
	devPreview     bool
	useHeightRuler bool
	devCamSpeed    float64
	freePlay       bool
	ping           float64
	reqT           time.Time
}

type assets struct {
	mainWorldBg *ebiten.Image
	altWorldBg  *ebiten.Image
}

type fixedAnimTracker struct {
	pid      string
	animName string
	ticks    int
}

var clientConfig *gameSettings
var devConfig *devSettings
var assetsHelper *assets
var fixedAnims map[string]*fixedAnimTracker

func initClient() {
	clientConfig = &gameSettings{
		screenWidth:  880,
		screenHeight: 480,
		streamInit:   false,
		worldsMap:    make(map[string]worldData),
		addr:         "localhost:50051",
		fullScreen:   false,
		enablePPROF:  false,
	}

	devConfig = &devSettings{
		rulerW:         ut.LoadImg("./sprites/rulers/wRuler.png"),
		rulerH:         ut.LoadImg("./sprites/rulers/hRuler.png"),
		devPreview:     false,
		useHeightRuler: false,
		devCamSpeed:    float64(2),
		freePlay:       false,
		ping:           0.0,
		reqT:           time.Now(),
	}

	assetsHelper = &assets{
		mainWorldBg: ut.LoadImg("./backgrounds/mapMain.png"),
		altWorldBg:  ut.LoadImg("./backgrounds/mapAlt.png"),
	}

	fixedAnims = make(map[string]*fixedAnimTracker)

	clientConfig.game = NewGame()
}
