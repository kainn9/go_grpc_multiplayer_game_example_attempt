package main

import "github.com/solarlune/resolv"

type builderFunc func(*world, float64, float64)

func mainWorldBuilder(world *world, gw float64, gh float64) {

	world.space.Add(
		// bounds

		// left
		resolv.NewObject(-16, 0, 16, gh, "solid"),
		// right
		resolv.NewObject(gw-16, 0, 16, gh, "solid"),

		// bottom
		resolv.NewObject(0, 0, gw, 16, "solid"),

		// top
		resolv.NewObject(0, gh-24, gw, 32, "solid"),

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

func altWorldBuilder(world *world, gw float64, gh float64) {
	world.space.Add(

		// bottom bounds
		resolv.NewObject(0, gh-16, gw, 16, "solid"),

		// Village Plat
		resolv.NewObject(1166, 3912, 6000, 10, "platform"),

		// Y-ZONE Plats
		resolv.NewObject(1008, 3931, 110, 10, "platform"),
		resolv.NewObject(909, 3973, 110, 10, "platform"),
		resolv.NewObject(902, 3818, 120, 10, "platform"),
		resolv.NewObject(996, 3687, 160, 10, "platform"),
		resolv.NewObject(673, 3717, 320, 10, "platform"),
		resolv.NewObject(512, 3844, 300, 10, "platform"),
		resolv.NewObject(676, 3632, 130, 10, "platform"),
		resolv.NewObject(906, 3588, 130, 10, "platform"),
		resolv.NewObject(1006, 3544, 130, 10, "platform"),
		resolv.NewObject(896, 3438, 130, 10, "platform"),
		resolv.NewObject(994, 3304, 160, 10, "platform"),
		resolv.NewObject(906, 3212, 130, 10, "platform"),
		resolv.NewObject(1002, 3170, 130, 10, "platform"),
		resolv.NewObject(473, 3812, 135, 10, "platform"),
		resolv.NewObject(477, 3577, 120, 10, "platform"),
		resolv.NewObject(477, 3577, 120, 10, "platform"),
		resolv.NewObject(513, 3465, 320, 10, "platform"),
		resolv.NewObject(477, 3432, 130, 10, "platform"),
		resolv.NewObject(672, 3332, 320, 10, "platform"),
		resolv.NewObject(674, 3252, 130, 10, "platform"),
		resolv.NewObject(900, 3060, 130, 10, "platform"),
		resolv.NewObject(994, 2928, 160, 10, "platform"),
		resolv.NewObject(669, 2955, 320, 10, "platform"),
		resolv.NewObject(673, 2874, 130, 10, "platform"),
		resolv.NewObject(509, 3086, 310, 10, "platform"),
		resolv.NewObject(509, 3086, 310, 10, "platform"),
		resolv.NewObject(476, 3195, 110, 10, "platform"),
		resolv.NewObject(477, 3958, 120, 10, "platform"),
		resolv.NewObject(267, 3973, 120, 10, "platform"),
		resolv.NewObject(267, 3973, 120, 10, "platform"),
		resolv.NewObject(366, 3928, 120, 10, "platform"),
		resolv.NewObject(366, 3928, 120, 10, "platform"),
		resolv.NewObject(257, 3824, 160, 10, "platform"),
		resolv.NewObject(355, 3687, 160, 10, "platform"),
		resolv.NewObject(267, 3590, 120, 10, "platform"),
		resolv.NewObject(364, 3546, 120, 10, "platform"),
		resolv.NewObject(259, 3439, 130, 10, "platform"),
		resolv.NewObject(34, 3718, 320, 10, "platform"),
		resolv.NewObject(0, 3848, 170, 10, "platform"),
		resolv.NewObject(34, 3634, 120, 10, "platform"),
		resolv.NewObject(0, 3464, 170, 10, "platform"),
		resolv.NewObject(31, 3335, 320, 10, "platform"),
		resolv.NewObject(355, 3305, 160, 10, "platform"),
		resolv.NewObject(38, 3258, 120, 10, "platform"),
		resolv.NewObject(266, 3204, 120, 10, "platform"),
		resolv.NewObject(368, 3164, 110, 10, "platform"),
		resolv.NewObject(255, 3050, 160, 10, "platform"),
		resolv.NewObject(474, 3048, 140, 10, "platform"),
		resolv.NewObject(355, 2922, 160, 10, "platform"),
		resolv.NewObject(0, 3078, 160, 10, "platform"),
		resolv.NewObject(30, 2949, 320, 10, "platform"),

		// mid section blocker left
		resolv.NewObject(206, 2584, 2030, 150, "solid"),
		resolv.NewObject(206, 2574, 2030, 10, "platform"),

		// left blocker left
		resolv.NewObject(0, 2108, 60, 540, "solid"),
		resolv.NewObject(0, 2098, 60, 10, "platform"),

		// forrest floating plats
		resolv.NewObject(64, 2639, 60, 10, "platform"),
		resolv.NewObject(128, 2549, 150, 10, "platform"),

		resolv.NewObject(305, 2500, 125, 10, "platform"),
		resolv.NewObject(452, 2450, 125, 10, "platform"),
		resolv.NewObject(615, 2392, 125, 10, "platform"),
		resolv.NewObject(797, 2359, 130, 10, "platform"),
		resolv.NewObject(797, 2359, 130, 10, "platform"),
		resolv.NewObject(956, 2316, 130, 10, "platform"),
		resolv.NewObject(1127, 2265, 130, 10, "platform"),
		resolv.NewObject(1308, 2241, 85, 10, "platform"),

		// wood forrest plat left
		resolv.NewObject(694, 2529, 1370, 10, "platform"),

		// castle floating plats
		resolv.NewObject(2093, 2484, 70, 10, "platform"),
		resolv.NewObject(2196, 2466, 30, 10, "platform"),
		resolv.NewObject(2400, 2450, 63, 10, "platform"),
		resolv.NewObject(2516, 2448, 63, 10, "platform"),
		resolv.NewObject(2611, 2428, 63, 10, "platform"),
		resolv.NewObject(2293, 2453, 63, 10, "platform"),

		// sky-town wallStalk and floaters
		resolv.NewObject(1278, 811, 54, 10, "platform"),
		resolv.NewObject(1278, 821, 54, 1275, "solid"),

		resolv.NewObject(1428, 2184, 54, 45, "solid"),
		resolv.NewObject(1428, 2174, 54, 10, "platform"),

		resolv.NewObject(1346, 2110, 54, 45, "solid"),
		resolv.NewObject(1346, 2100, 54, 10, "platform"),

		// sky-town floor left
		resolv.NewObject(0, 872, 1192, 10, "platform"),

		// sky-town floor right
		resolv.NewObject(1371, 837, 650, 10, "platform"),

		// dungeon town wall right divider
		resolv.NewObject(1970, 0, 55, 1826, "solid"),

		// rock plats
		resolv.NewObject(0, 2848, 80, 10, "platform"),
		resolv.NewObject(143, 2760, 160, 10, "platform"),
		resolv.NewObject(152, 2668, 50, 10, "platform"),
	)
}
