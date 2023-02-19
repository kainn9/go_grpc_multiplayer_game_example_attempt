package main

import ut "github.com/kainn9/grpc_game/client/util"

/*
	Putting global scope client vars and consts here to avoid confusion
*/

const (
	ScreenWidth  = 880
	ScreenHeight = 480
)

var (
	ticks     int
	worldsMap = make(map[string]WorldData)
	game      = NewGame()
	addr      = "localhost:50051"
	fullScreen = false
	// I use this for demo my client
	// addr = "ec2-54-144-156-228.compute-1.amazonaws.com:50051"
)


// dev mode stuff
var (
	rulerW         = ut.LoadImg("./sprites/rulers/wRuler.png")
	rulerH         = ut.LoadImg("./sprites/rulers/hRuler.png")
	devPreview     = false
	useHeightRuler = false
	devCamSpeed    = float64(2)
	freePlay = false
)

var (
	mainWorldBg = ut.LoadImg("./backgrounds/mapMain.png")
	altWorldBg  = ut.LoadImg("./backgrounds/mapAlt.png")
)

