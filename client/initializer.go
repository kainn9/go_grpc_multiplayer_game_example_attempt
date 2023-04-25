package main

// TODO move these out of Global Scope and into "config" structs
import (
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	r "github.com/kainn9/grpc_game/client/roles"
	utClient "github.com/kainn9/grpc_game/client_util"
	"google.golang.org/grpc"
)

/*
	Putting global scope client vars and consts here to avoid confusion
*/

type gameSettings struct {
	screenWidth      int
	screenHeight     int
	streamInit       bool
	ticks            int
	worldsMap        map[int]worldData
	game             *Game
	addr             string
	fullScreen       bool
	enablePPROF      bool
	connRef          *grpc.ClientConn
	audPlayer        *audio.Player
	volume128        int
	showHelp         bool
	roles            map[int32]*r.Role
	showPlayerHitbox bool
	imageCache       sync.Map
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

type worldBackgrounds struct {
	worldOne             *ebiten.Image
	worldTwo             *ebiten.Image
	worldThree           *ebiten.Image
	landOfYohoPassageOne *ebiten.Image
	landOfYohoPassageTwo *ebiten.Image
	landOfYohoVillage    *ebiten.Image
}

type fixedAnimTracker struct {
	pid      string
	animName string
	ticks    int
}

type cachedImage struct {
	img       *ebiten.Image
	timestamp time.Time // TODO create a go routine to clear cache based on timestamp
}

var clientConfig *gameSettings
var devConfig *devSettings
var wBgHelper *worldBackgrounds
var fixedAnims map[string]*fixedAnimTracker
var ADDR string

func initClient() {
	addr := ADDR

	if addr == "" {
		addr = "localhost"
	}

	clientConfig = &gameSettings{
		screenWidth:      880,
		screenHeight:     480,
		streamInit:       false,
		worldsMap:        make(map[int]worldData),
		addr:             addr + ":50051",
		fullScreen:       false,
		enablePPROF:      false,
		showHelp:         true,
		roles:            make(map[int32]*r.Role),
		showPlayerHitbox: true,
		imageCache:       sync.Map{},
	}

	clientConfig.roles[0] = r.InitKnight()
	clientConfig.roles[1] = r.InitMonk()
	clientConfig.roles[2] = r.InitDemon()
	clientConfig.roles[3] = r.InitWerewolf()
	clientConfig.roles[4] = r.InitMage()
	clientConfig.roles[5] = r.InitHeavyKnight()
	clientConfig.roles[6] = r.InitBirdDroid()

	devConfig = &devSettings{
		rulerW:         utClient.LoadImage("./sprites/rulers/wRuler.png"),
		rulerH:         utClient.LoadImage("./sprites/rulers/hRuler.png"),
		devPreview:     false,
		useHeightRuler: false,
		devCamSpeed:    float64(2),
		freePlay:       false,
		ping:           0.0,
		reqT:           time.Now(),
	}

	wBgHelper = &worldBackgrounds{
		worldOne:             utClient.LoadImage("./backgrounds/worldOne.png"),
		worldTwo:             utClient.LoadImage("./backgrounds/worldTwo.png"),
		worldThree:           utClient.LoadImage("./backgrounds/worldThree.png"),
		landOfYohoPassageOne: utClient.LoadImage("./backgrounds/landOfYoho/landOfYohoPassageOne.png"),
		landOfYohoPassageTwo: utClient.LoadImage("./backgrounds/landOfYoho/landOfYohoPassageTwo.png"),
		landOfYohoVillage:    utClient.LoadImage("./backgrounds/landOfYoho/landOfYohoVillage.png"),
	}

	clientConfig.worldsMap[0] = *NewWorldData(848, 1600, wBgHelper.worldOne)
	clientConfig.worldsMap[1] = *NewWorldData(848, 3200, wBgHelper.worldTwo)
	clientConfig.worldsMap[2] = *NewWorldData(4000, 6000, wBgHelper.worldThree)
	clientConfig.worldsMap[3] = *NewWorldData(480, 960, wBgHelper.landOfYohoPassageOne)
	clientConfig.worldsMap[4] = *NewWorldData(756, 1100, wBgHelper.landOfYohoPassageTwo)
	clientConfig.worldsMap[5] = *NewWorldData(600, 3278, wBgHelper.landOfYohoVillage)

	fixedAnims = make(map[string]*fixedAnimTracker)

	clientConfig.game = NewGame()
}
