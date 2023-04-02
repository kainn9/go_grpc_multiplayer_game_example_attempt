package main

import "github.com/solarlune/resolv"

type builderFunc func(*world, float64, float64)

func introWorldBuilder(world *world, gw float64, gh float64) {
	portal := resolv.NewObject(1541, 150, 40, 40, "portal")
	initPortal(portal, 3, 16, 250)

	world.space.Add(

		// left bounds
		resolv.NewObject(0, 0, 16, gh, "bounds"),
		// right bounds
		resolv.NewObject(gw-16, 0, 16, gh, "bounds"),

		// floor 1 left
		resolv.NewObject(0, 752, 468, 10, "solid"),

		// floor 2 mid
		resolv.NewObject(546, 732, 470, 10, "solid"),

		// rock
		resolv.NewObject(318, 548, 127, 200, "solid"),

		// woodPlat long
		resolv.NewObject(484, 506, 737, 5, "platform"),

		// grassMid floatingPlat
		resolv.NewObject(361, 276, 800, 5, "platform"),

		// grassTop floatingPlat left
		resolv.NewObject(0, 245, 278, 5, "platform"),

		// floatingBLock
		resolv.NewObject(1101, 753, 55, 50, "solid"),

		// castle floor
		resolv.NewObject(1237, 728, 390, 10, "solid"),

		// castle midPlat
		resolv.NewObject(1362, 419, 250, 5, "platform"),

		// castle topPlat
		resolv.NewObject(1232, 198, 400, 5, "platform"),

		// castle left side plats p1
		resolv.NewObject(1237, 55, 30, 5, "platform"),
		resolv.NewObject(1237, 94, 30, 5, "platform"),
		resolv.NewObject(1237, 134, 30, 5, "platform"),
		resolv.NewObject(1237, 165, 30, 5, "platform"),

		// castle left side plats p2
		resolv.NewObject(1239, 242, 30, 5, "platform"),
		resolv.NewObject(1239, 293, 30, 5, "platform"),
		resolv.NewObject(1239, 339, 30, 5, "platform"),
		resolv.NewObject(1239, 384, 30, 5, "platform"),
		resolv.NewObject(1239, 437, 30, 5, "platform"),
		resolv.NewObject(1239, 490, 30, 5, "platform"),

		// inner castle plats small
		resolv.NewObject(1307, 533, 30, 5, "platform"),
		resolv.NewObject(1368, 553, 30, 5, "platform"),
		resolv.NewObject(1436, 590, 30, 5, "platform"),

		// inner castle plats med
		resolv.NewObject(1463, 640, 60, 5, "platform"),
		resolv.NewObject(1355, 659, 60, 5, "platform"),

		// ladder1
		resolv.NewObject(696, 310, 40, 5, "platform"),
		resolv.NewObject(696, 343, 40, 5, "platform"),
		resolv.NewObject(696, 376, 40, 5, "platform"),
		resolv.NewObject(696, 415, 40, 5, "platform"),
		resolv.NewObject(696, 455, 40, 5, "platform"),

		// ladder2
		resolv.NewObject(1021, 310, 40, 5, "platform"),
		resolv.NewObject(1021, 343, 40, 5, "platform"),
		resolv.NewObject(1021, 376, 40, 5, "platform"),
		resolv.NewObject(1021, 415, 40, 5, "platform"),
		resolv.NewObject(1021, 455, 40, 5, "platform"),

		// portal
		portal,
	)
}

func landOfYohoPassageOneBuilder(world *world, gw float64, gh float64) {

	portal := resolv.NewObject(gw-16, 0, 18, gh, "portal")
	initPortal(portal, 4, 40, 500)

	world.space.Add(

		// left bounds
		resolv.NewObject(0, 0, 16, gh, "bounds"),

		// floor
		resolv.NewObject(0, 328, 1000, 40, "solid"),

		portal,
	)
}

func landOfYohoPassageTwoBuilder(world *world, gw float64, gh float64) {
	portalBack := resolv.NewObject(0, 408, 16, 500, "portal")
	initPortal(portalBack, 3, 880, 250)

	portalForwards := resolv.NewObject(0, 0, 16, 340, "portal")
	initPortal(portalForwards, 5, 3160, 402)

	world.space.Add(

		// right bounds jumpable
		resolv.NewObject(gw-16, 0, 16, gh, "solid"),

		// floor left bottom
		resolv.NewObject(0, 581, 825, 200, "solid"),

		// floor right bottom
		resolv.NewObject(914, 581, 200, 200, "solid"),

		// plats
		resolv.NewObject(949, 514, 125, 5, "platform"),
		resolv.NewObject(949, 459, 125, 5, "platform"),
		resolv.NewObject(949, 408, 125, 5, "platform"),
		resolv.NewObject(949, 354, 125, 5, "platform"),

		// top floor
		resolv.NewObject(0, 336, 938, 80, "solid"),
		portalBack,
		portalForwards,
	)
}

func landOfYohoVillageBuilder(world *world, gw float64, gh float64) {
	portalBack := resolv.NewObject(gw-16, 0, 18, gh, "portal")
	initPortal(portalBack, 4, 104, 292)

	world.space.Add(
		// left
		resolv.NewObject(0, 0, 16, gh, "bounds"),

		// floor
		resolv.NewObject(0, 502, 3400, 300, "solid"),

		// castleZone Sub Solid
		resolv.NewObject(0, 395, 421, 300, "solid"),

		// castleZone Plat
		resolv.NewObject(0, 350, 1160, 5, "platform"),
		portalBack,
	)
}

func mainWorldBuilder(world *world, gw float64, gh float64) {

	world.space.Add(

		// left
		resolv.NewObject(0, 0, 16, gh, "bounds"),
		// right
		resolv.NewObject(gw-16, 0, 16, gh, "bounds"),

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

		// left bounds
		resolv.NewObject(0, 0, 16, gh, "bounds"),
		// right bounds
		resolv.NewObject(gw-16, 0, 16, gh, "bounds"),

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
