package main

import "github.com/solarlune/resolv"

type BuilderFunc func(*World, float64, float64) 


func MainWorldBuilder (world *World, gw float64, gh float64) {

	world.Space.Add(
		// bounds
		resolv.NewObject(0, 0, 16, gh, "solid"),
		resolv.NewObject(gw-16, 0, 16, gh, "solid"),
		resolv.NewObject(0, 0, gw, 16, "solid"),
		resolv.NewObject(0, gh-24, gw, 32, "solid"),
		resolv.NewObject(160, gh-56, 160, 32, "solid"),
		resolv.NewObject(320, 64, 32, 160, "solid"),
		resolv.NewObject(64, 128, 16, 160, "solid"),
		resolv.NewObject(gw-128, 64, 128, 16, "solid"),
		resolv.NewObject(gw-128, gh-88, 128, 16, "solid"),



		// big rock
		resolv.NewObject(468, 590, 260, 300, "solid"),
		resolv.NewObject(468, 580, 260, 30, "platform"),


		// plat2right of big rock
		resolv.NewObject(774, 676, 158, 10, "platform"),


		// Big platform thing
		resolv.NewObject(959, 633, 2500, 10, "platform"),

		// wooden plats
		resolv.NewObject(1059, 570, 127, 5, "platform"),
		resolv.NewObject(1156, 505, 127, 5, "platform"),
		resolv.NewObject(1281, 448, 127, 5, "platform"),
		resolv.NewObject(1418, 398, 125, 5, "platform"),
		resolv.NewObject(1546, 350, 125, 5, "platform"),
		resolv.NewObject(1716, 305, 125, 5, "platform"), 
		resolv.NewObject(1904, 256, 125, 5, "platform"),

		resolv.NewObject(2086, 198, 125, 5, "platform"),
		resolv.NewObject(2258, 146, 125, 5, "platform"),


		// floating rock top right
		resolv.NewObject(2496, 100, 158, 5, "platform"),
	)
}

func AltWorldBuilder(world *World, gw float64, gh float64) {
	world.Space.Add(

		// bounds 
		resolv.NewObject(0, 0, 16, gh, "solid"),
		resolv.NewObject(gw-16, 0, 16, gh, "solid"),
		resolv.NewObject(0, 0, gw, 16, "solid"),
		resolv.NewObject(0, gh-24, gw, 32, "solid"),
		resolv.NewObject(160, gh-56, 160, 32, "solid"),
		resolv.NewObject(320, 64, 32, 160, "solid"),
		resolv.NewObject(64, 128, 16, 160, "solid"),
		resolv.NewObject(gw-128, 64, 128, 16, "solid"),
		resolv.NewObject(gw-128, gh-88, 128, 16, "solid"),


		// floor
		resolv.NewObject(0, 660, gw, 10, "platform"),
	)
}
