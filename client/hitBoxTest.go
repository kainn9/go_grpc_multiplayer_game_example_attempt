package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type hitboxTest struct {
	on    bool
	name  string
	count int
	frame int
	left  bool
	inc   float64
	// sadly pWidth needs to be hard coded for accurate sim when left is set to true, since client doesn't actually know player
	// player width. Width of player hitboxes can be found in roles folder in server
	pWidth float64
}

type hitBox struct {
	pOffX  float64
	pOffY  float64
	height float64
	width  float64
}

type hitBoxAggregate []hitBox

type hitBoxSequence struct {
	hBoxPath   hBoxPath
	movmentInc float64
}

type hBoxPath []hitBoxAggregate

const noBox = -90000 // TODO: deprecate noBox(old examples still use), as you can just omit path values

// use this to preview for full anim len when freezing a frame
func previewPathAllFrames(x float64, y float64, h float64, w float64, path hBoxPath, count int) {
	for i := 0; i < count; i++ {
		path = path.appendHboxAgg(x, y, h, w, i)
	}
}

/*
	------------------------------------------------
		SETUP TEST HITBOX TEST HERE
	------------------------------------------------
*/

/*
-----------------------------------------------------------------------------
Knight Tert Attack Example(only works if CP is Knight)
-----------------------------------------------------------------------------
*/
// var (
// 	hitBoxTest = &hitboxTest{
// 		name:  string(sr.TertAttackKey),
// 		on:    false,
// 		count: 12,
// 		// set to -1 to play whole anim
// 		frame: -1,
// 		left:  false,
// 		inc:   16.666 * 5, // 1 frame at 60fps
// 	}
// )

// func hitBoxSim(screen *ebiten.Image, cp *PlayerController) {
// 	inc, path := hitBoxSimSetup(hitBoxTest.inc)

// 	// frame 0
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

// 	// frame 1
// 	path = path.appendHboxAgg(-10, 0, 8, 8, 1)
// 	path = path.appendHboxAgg(-2, 0, 8, 8, 1)
// 	path = path.appendHboxAgg(6, 0, 8, 8, 1)
// 	path = path.appendHboxAgg(14, 0, 8, 8, 1)
// 	path = path.appendHboxAgg(22, 4, 8, 8, 1)
// 	path = path.appendHboxAgg(26, 4, 8, 8, 1)
// 	path = path.appendHboxAgg(28, 8, 8, 8, 1)

// 	// frame 2
// 	path = path.appendHboxAgg(28, 4, 8, 8, 2)
// 	path = path.appendHboxAgg(32, 8, 8, 8, 2)
// 	path = path.appendHboxAgg(36, 12, 8, 8, 2)
// 	path = path.appendHboxAgg(39, 16, 8, 8, 2)
// 	path = path.appendHboxAgg(40, 24, 8, 8, 2)
// 	path = path.appendHboxAgg(36, 32, 8, 8, 2)
// 	path = path.appendHboxAgg(30, 36, 8, 8, 2)
// 	path = path.appendHboxAgg(24, 38, 8, 8, 2)
// 	path = path.appendHboxAgg(16, 38, 8, 8, 2)
// 	path = path.appendHboxAgg(8, 38, 8, 8, 2)

// 	// frame same as 2
// 	path[3] = path[2]

// 	// frame 4 same as 2 but slightly to right
// 	path = path.appendHboxAgg(36, 4, 8, 8, 4)
// 	path = path.appendHboxAgg(44, 8, 8, 8, 4)
// 	path = path.appendHboxAgg(48, 12, 8, 8, 4)
// 	path = path.appendHboxAgg(51, 16, 8, 8, 4)
// 	path = path.appendHboxAgg(50, 24, 8, 8, 4)
// 	path = path.appendHboxAgg(45, 32, 8, 8, 4)
// 	path = path.appendHboxAgg(40, 36, 8, 8, 4)
// 	path = path.appendHboxAgg(34, 38, 8, 8, 4)

// 	path[5] = path[3]
// 	path[6] = path[4]

// 	path[7] = path[5]
// 	path[8] = path[6]
// 	path[9] = path[7]
// 	path[10] = path[8]
// 	path[11] = path[9]

// 	startHitboxSim(screen, cp, inc, path, 0)
// }

/*
-----------------------------------------------------------------------------
Knight Tert Attack Example End
-----------------------------------------------------------------------------
*/

/*
-----------------------------------------------------------------------------
Mage Tert Attack Example(only works if CP is Mage)
-----------------------------------------------------------------------------
*/
// var (
// 	hitBoxTest = &hitboxTest{
// 		name:  string(sr.TertAttackKey),
// 		on:    false,
// 		count: 12,
// 		// set to -1 to play whole anim
// 		frame:  -1,
// 		left:   true,
// 		inc:    16.666 * 5, // 1 frame at 60fps
// 		pWidth: 16,
// 	}
// )

// func hitBoxSim(screen *ebiten.Image, cp *PlayerController) {
// 	inc, path := hitBoxSimSetup(hitBoxTest.inc)

// 	// frame 0 - 2 no hitboxes

// 	// frame 3
// 	path = path.appendHboxAgg(-50, 0, 50, 110, 3)

// 	// frame 4
// 	path = path.appendHboxAgg(-50, 0, 50, 110, 4)

// 	// frame 5
// 	path = path.appendHboxAgg(-50, 0, 50, 110, 5)

// 	// frame 6
// 	path = path.appendHboxAgg(-70, -10, 60, 150, 6)

// 	// frame 7
// 	path = path.appendHboxAgg(-70, -10, 60, 150, 7)

// 	// frame 8
// 	path = path.appendHboxAgg(-40, 10, 40, 90, 8)

// 	// frame 9
// 	path = path.appendHboxAgg(-40, 10, 40, 90, 9)

// 	// frame 10
// 	path = path.appendHboxAgg(-30, 30, 20, 80, 10)

// 	// frame 11 + no box

// 	previewPathAllFrames(-30, 30, 20, 80, path, 12)
// 	startHitboxSim(screen, cp, inc, path, 0)
// }

/*
-----------------------------------------------------------------------------
Mage Tert Attack Example End
-----------------------------------------------------------------------------
*/

/*
-----------------------------------------------------------------------------
Mage Secondary Attack Example(only works if CP is Mage)
-----------------------------------------------------------------------------
*/
// var (
// 	hitBoxTest = &hitboxTest{
// 		name:  string(sr.SecondaryAttackKey),
// 		on:    false,
// 		count: 12,
// 		// set to -1 to play whole anim
// 		frame:  -1,
// 		left:   true,
// 		inc:    16.666 * 5, // 1 frame at 60fps
// 		pWidth: 16,
// 	}
// )

// func hitBoxSim(screen *ebiten.Image, cp *PlayerController) {
// 	inc, path := hitBoxSimSetup(hitBoxTest.inc)

// 	// frame 1 - 3 no hitbox

// 	// frame 4
// 	path = path.appendHboxAgg(33, -10, 55, 40, 4)
// 	path = path.appendHboxAgg(23, 0, 35, 10, 4)
// 	path = path.appendHboxAgg(73, 0, 35, 10, 4)

// 	// frame 5
// 	path = path.appendHboxAgg(33, -10, 55, 40, 5)
// 	path = path.appendHboxAgg(23, 0, 35, 10, 5)
// 	path = path.appendHboxAgg(73, 0, 35, 10, 5)

// 	// frame 6
// 	path = path.appendHboxAgg(30, -25, 25, 50, 6)
// 	path = path.appendHboxAgg(15, -10, 55, 80, 6)

// 	// frame 7
// 	path = path.appendHboxAgg(33, -10, 55, 50, 7)
// 	path = path.appendHboxAgg(23, 0, 45, 10, 7)
// 	path = path.appendHboxAgg(82, 0, 45, 12, 7)

// 	// frame 8
// 	path = path.appendHboxAgg(33, -10, 55, 50, 8)
// 	path = path.appendHboxAgg(23, 0, 45, 10, 8)
// 	path = path.appendHboxAgg(82, 0, 45, 12, 8)

// 	// frame 9
// 	path = path.appendHboxAgg(23, 5, 40, 15, 9)
// 	path = path.appendHboxAgg(30, -5, 15, 60, 9)
// 	path = path.appendHboxAgg(83, 5, 40, 15, 9)

// 	// frame 10 + no hitbox

// 	startHitboxSim(screen, cp, inc, path, 0)
// }

/*
-----------------------------------------------------------------------------
Mage Secondary Attack Example End
-----------------------------------------------------------------------------
*/

/*
-----------------------------------------------------------------------------
Mage Primary Attack Example(only works if CP is Mage)
-----------------------------------------------------------------------------
*/
// var (
// 	hitBoxTest = &hitboxTest{
// 		name:  string(sr.PrimaryAttackKey),
// 		on:    false,
// 		count: 21,
// 		// set to -1 to play whole anim
// 		frame: -1,
// 		left:  false,
// 		inc:   16.666 * 5, // 1 frame at 60fps
// 	}
// )

// func hitBoxSim(screen *ebiten.Image, cp *PlayerController) {
// 	inc, path := hitBoxSimSetup(hitBoxTest.inc)

// 	// frame 0 - 10 no hitboxess

// 	// frame 11
// 	path = path.appendHboxAgg(-33, -30, 47, 8, 11)
// 	path = path.appendHboxAgg(-63, -30, 47, 8, 11)
// 	path = path.appendHboxAgg(43, -30, 47, 8, 11)
// 	path = path.appendHboxAgg(73, -30, 47, 8, 11)

// 	// frame 12
// 	path = path.appendHboxAgg(-33, -30, 47, 8, 12)
// 	path = path.appendHboxAgg(-63, -30, 47, 8, 12)
// 	path = path.appendHboxAgg(43, -30, 47, 8, 12)
// 	path = path.appendHboxAgg(73, -30, 47, 8, 12)

// 	// frame 13
// 	path = path.appendHboxAgg(-33, -24, 47, 8, 13)
// 	path = path.appendHboxAgg(-63, -24, 47, 8, 13)
// 	path = path.appendHboxAgg(43, -24, 47, 8, 13)
// 	path = path.appendHboxAgg(73, -24, 47, 8, 13)

// 	// frame 14
// 	path = path.appendHboxAgg(-33, -16, 47, 8, 14)
// 	path = path.appendHboxAgg(-63, -16, 47, 8, 14)
// 	path = path.appendHboxAgg(43, -16, 47, 8, 14)
// 	path = path.appendHboxAgg(73, -16, 47, 8, 14)

// 	// frame 15
// 	path = path.appendHboxAgg(-33, -6, 47, 8, 15)
// 	path = path.appendHboxAgg(-63, -6, 47, 8, 15)
// 	path = path.appendHboxAgg(43, -6, 47, 8, 15)
// 	path = path.appendHboxAgg(73, -6, 47, 8, 15)

// 	// frame 16
// 	path = path.appendHboxAgg(-33, 10, 35, 8, 16)
// 	path = path.appendHboxAgg(-63, 10, 35, 8, 16)
// 	path = path.appendHboxAgg(43, 10, 35, 8, 16)
// 	path = path.appendHboxAgg(73, 10, 35, 8, 16)

// 	// frame 17
// 	path = path.appendHboxAgg(-33, 10, 35, 8, 17)
// 	path = path.appendHboxAgg(-63, 10, 35, 8, 17)
// 	path = path.appendHboxAgg(43, 10, 35, 8, 17)
// 	path = path.appendHboxAgg(73, 10, 35, 8, 17)

// 	// frame 17+ no box

// 	startHitboxSim(screen, cp, inc, path, 0)
// }

/*
-----------------------------------------------------------------------------
Mage Primary Attack Example End
-----------------------------------------------------------------------------
*/

/*
-----------------------------------------------------------------------------
Werewolf Tert Attack Example(only works if cp is Werewolf, was lazy with these hitboxes)
-----------------------------------------------------------------------------
*/

// var (
// 	hitBoxTest = &hitboxTest{
// 		name:  "tertAtk",
// 		on:    false,
// 		count: 17,
// 		// set to -1 to play whole anim
// 		frame: -1,
// 		left:  false,
// 		inc:   16.666 * 5, // 1 frame at 60fps
// 	}
// )

// func hitBoxSim(screen *ebiten.Image, cp *PlayerController) {
// 	inc, path := hitBoxSimSetup(hitBoxTest.inc)

// 	// frame 0
// 	path = path.appendHboxAgg(60, 0, 20, 20, 0)
// 	path = path.appendHboxAgg(55, 10, 20, 20, 0)
// 	path = path.appendHboxAgg(45, 20, 20, 20, 0)

// 	// frame 1
// 	path = path.appendHboxAgg(45, 0, 20, 23, 1)

// 	// frame 2
// 	path = path.appendHboxAgg(25, 10, 20, 43, 2)
// 	path = path.appendHboxAgg(25, 20, 20, 28, 2)

// 	// frame 3
// 	path = path.appendHboxAgg(35, 0, 40, 42, 3)

// 	// frame 4
// 	path = path.appendHboxAgg(30, 10, 30, 32, 4)

// 	// frame 5
// 	path = path.appendHboxAgg(35, 0, 32, 35, 5)

// 	// frame 6
// 	path = path.appendHboxAgg(35, 0, 32, 38, 6)

// 	// frame 7
// 	path = path.appendHboxAgg(35, 0, 44, 45, 7)

// 	// frame 8
// 	path = path.appendHboxAgg(55, 0, 15, 25, 8)
// 	path = path.appendHboxAgg(50, 15, 15, 25, 8)
// 	path = path.appendHboxAgg(45, 20, 15, 25, 8)

// 	// frame 9
// 	path = path.appendHboxAgg(45, 0, 20, 24, 9)

// 	// frame 10
// 	path = path.appendHboxAgg(25, 10, 20, 43, 10)
// 	path = path.appendHboxAgg(25, 20, 20, 28, 10)

// 	// frame 11
// 	path = path.appendHboxAgg(45, 0, 40, 34, 12)

// 	// frame 12 no box

// 	// frame 13
// 	path = path.appendHboxAgg(10, 0, 40, 63, 13)

// 	// frame 14 - 16 no box

// 	startHitboxSim(screen, cp, inc, path, 0)
// }

/*
-----------------------------------------------------------------------------
Werewolf Tert Attack Example END
-----------------------------------------------------------------------------
*/

/*
-----------------------------------------------------------------------------
Werewolf Secondary Attack Example(only works if cp is Werewolf, was lazy with these hitboxes)
-----------------------------------------------------------------------------
*/

// var (
// 	hitBoxTest = &hitboxTest{
// 		name:  "secondaryAtk",
// 		on:    false,
// 		count: 10,
// 		// set to -1 to play whole anim
// 		frame: -1,
// 		left:  false,
// 		inc:   16.666 * 5, // 1 frame at 60fps
// 	}
// )

// func hitBoxSim(screen *ebiten.Image, cp *PlayerController) {
// 	inc, path := hitBoxSimSetup(hitBoxTest.inc)

// 	// frame 0 - 3 no hitboxes

// 	// frame 4
// 	path = path.appendHboxAgg(10, -26, 20, 23, 4)
// 	path = path.appendHboxAgg(76, -26, 35, 23, 4)

// 	// frame 5
// 	path = path.appendHboxAgg(14, -26, 66, 80, 5)

// 	// frame 6
// 	path = path.appendHboxAgg(18, 20, 26, 80, 6)

// 	// frame 7
// 	path = path.appendHboxAgg(18, 20, 26, 80, 7)

// 	// frame 8 - 10 no hitboxes

// 	startHitboxSim(screen, cp, inc, path, 0)
// }

/*
-----------------------------------------------------------------------------
Werewolf Secondary Attack Example END
-----------------------------------------------------------------------------
*/

/*
-----------------------------------------------------------------------------
Werewolf Primary Attack Example(only works if cp is Werewolf, was lazy with these hitboxes)
-----------------------------------------------------------------------------
*/

// var (
// 	hitBoxTest = &hitboxTest{
// 		name:  "primaryAtk",
// 		on:    false,
// 		count: 8,
// 		// set to -1 to play whole anim
// 		frame: -1,
// 		left:  false,
// 		inc:   16.666 * 5, // 1 frame at 60fps
// 	}
// )

// func hitBoxSim(screen *ebiten.Image, cp *PlayerController) {
// 	inc, path := hitBoxSimSetup(hitBoxTest.inc)

// 	// frame 0-1 no hitboxes

// 	// frame 2
// 	path = path.appendHboxAgg(20, -8, 40, 56, 2)

// 	// frame 3
// 	path = path.appendHboxAgg(56, -8, 18, 20, 3)
// 	path = path.appendHboxAgg(46, 20, 10, 20, 3)

// 	// frame 4
// 	path = path.appendHboxAgg(56, -8, 18, 20, 4)
// 	path = path.appendHboxAgg(46, 20, 10, 20, 4)

// 	// frame 5
// 	path = path.appendHboxAgg(56, -8, 44, 26, 5)
// 	path = path.appendHboxAgg(46, 20, 20, 20, 5)

// 	// frame 6
// 	path = path.appendHboxAgg(56, 16, 20, 26, 6)
// 	path = path.appendHboxAgg(39, 20, 20, 20, 6)

// 	// frame 7 no hitboxes

// 	startHitboxSim(screen, cp, inc, path, 0)
// }

/*
-----------------------------------------------------------------------------
Werewolf Primary Attack Example End
-----------------------------------------------------------------------------
*/

/*
-----------------------------------------------------------------------------
Demon Secondary Attack Example(only works if cp is Demon)
-----------------------------------------------------------------------------
*/

// var (
// 	hitBoxTest = &hitboxTest{
// 		name:  "secondaryAtk",
// 		on:    false,
// 		count: 21,
// 		// set to -1 to play whole anim
// 		frame: -1,
// 		left:  false,
// 		inc:   16.666 * 5, // 1 frame at 60fps
// 	}
// )

// func hitBoxSim(screen *ebiten.Image, cp *PlayerController) {
// 	inc, path := hitBoxSimSetup(hitBoxTest.inc)

// 	// frame 0 - 6 no hitboxes

// 	// frame 7
// 	path = path.appendHboxAgg(50, 0, 15, 15, 7)
// 	path = path.appendHboxAgg(57, -7, 15, 15, 7)
// 	path = path.appendHboxAgg(62, -12, 15, 15, 7)
// 	path = path.appendHboxAgg(72, -17, 15, 15, 7)
// 	path = path.appendHboxAgg(77, -22, 15, 15, 7)
// 	path = path.appendHboxAgg(82, -27, 15, 15, 7)
// 	path = path.appendHboxAgg(87, -32, 15, 15, 7)
// 	path = path.appendHboxAgg(92, -37, 15, 15, 7)

// 	// frame 8
// 	path = path.appendHboxAgg(50, 0, 15, 15, 8)
// 	path = path.appendHboxAgg(57, -2, 15, 15, 8)
// 	path = path.appendHboxAgg(62, -5, 15, 15, 8)
// 	path = path.appendHboxAgg(72, -8, 15, 15, 8)
// 	path = path.appendHboxAgg(77, -12, 15, 15, 8)
// 	path = path.appendHboxAgg(82, -15, 15, 15, 8)
// 	path = path.appendHboxAgg(87, -18, 15, 15, 8)
// 	path = path.appendHboxAgg(92, -21, 15, 15, 8)
// 	path = path.appendHboxAgg(97, -25, 15, 15, 8)
// 	path = path.appendHboxAgg(102, -28, 15, 15, 8)
// 	path = path.appendHboxAgg(107, -31, 15, 15, 8)

// 	// frame 9
// 	path = path.appendHboxAgg(50, 0, 15, 15, 9)
// 	path = path.appendHboxAgg(57, 3, 15, 15, 9)
// 	path = path.appendHboxAgg(72, 3, 8, 30, 9)
// 	path = path.appendHboxAgg(76, -2, 15, 15, 9)
// 	path = path.appendHboxAgg(82, -5, 15, 15, 9)
// 	path = path.appendHboxAgg(87, -8, 15, 15, 9)
// 	path = path.appendHboxAgg(92, -12, 15, 15, 9)
// 	path = path.appendHboxAgg(97, -15, 15, 15, 9)
// 	path = path.appendHboxAgg(102, -18, 15, 15, 9)
// 	path = path.appendHboxAgg(107, -21, 15, 15, 9)
// 	path = path.appendHboxAgg(112, -25, 15, 15, 9)
// 	path = path.appendHboxAgg(117, -28, 15, 15, 9)
// 	path = path.appendHboxAgg(122, -31, 15, 15, 9)

// 	// frame 10
// 	path = path.appendHboxAgg(50, 13, 8, 40, 10)
// 	path = path.appendHboxAgg(90, -5, 22, 30, 10)
// 	path = path.appendHboxAgg(120, -25, 30, 30, 10)

// 	// frame 11
// 	path = path.appendHboxAgg(50, 18, 8, 40, 11)
// 	path = path.appendHboxAgg(90, 7, 22, 30, 11)
// 	path = path.appendHboxAgg(120, -10, 30, 30, 11)

// 	// frame 12
// 	path = path.appendHboxAgg(50, 26, 8, 40, 12)
// 	path = path.appendHboxAgg(90, 15, 22, 30, 12)
// 	path = path.appendHboxAgg(120, 2, 30, 30, 12)

// 	// frame 13
// 	path = path.appendHboxAgg(50, 34, 8, 40, 13)
// 	path = path.appendHboxAgg(90, 26, 22, 30, 13)
// 	path = path.appendHboxAgg(120, 18, 30, 30, 13)

// 	// frame 14
// 	path = path.appendHboxAgg(50, 42, 8, 40, 14)
// 	path = path.appendHboxAgg(90, 38, 22, 30, 14)
// 	path = path.appendHboxAgg(120, 30, 30, 30, 14)

// 	// frame 15
// 	path = path.appendHboxAgg(50, 46, 8, 40, 15)
// 	path = path.appendHboxAgg(90, 53, 22, 30, 15)
// 	path = path.appendHboxAgg(120, 45, 30, 30, 15)

// 	// frame 16
// 	path = path.appendHboxAgg(50, 46, 8, 40, 16)
// 	path = path.appendHboxAgg(90, 53, 22, 30, 16)
// 	path = path.appendHboxAgg(120, 45, 30, 30, 16)

// 	// no hitboxes for frame 17+

// 	startHitboxSim(screen, cp, inc, path, 0)
// }

/*
-----------------------------------------------------------------------------
Demon Secondary Attack Example End
-----------------------------------------------------------------------------
*/

// /*
// -----------------------------------------------------------------------------
// Demon Primary Attack Example(only works if cp is Demon)
// -----------------------------------------------------------------------------
// */

var (
	hitBoxTest = &hitboxTest{
		name:  "primaryAtk",
		on:    false,
		count: 22,
		// set to -1 to play whole anim
		frame:  -1,
		left:   true,
		inc:    16.666 * 5, // 1 frame at 60fps,
		pWidth: 50.0,
	}
)

func hitBoxSim(screen *ebiten.Image, cp *PlayerController) {
	inc, path := hitBoxSimSetup(hitBoxTest.inc)

	// frame 0 - 11 have no hitboxes

	// frame 12
	path = path.appendHboxAgg(105, 70, 15, 15, 12)
	path = path.appendHboxAgg(115, 70, 15, 15, 12)

	// frame 13
	path = path.appendHboxAgg(102, 67, 15, 15, 13)
	path = path.appendHboxAgg(115, 70, 15, 15, 13)

	// frame 14
	path = path.appendHboxAgg(102, 67, 15, 15, 14)
	path = path.appendHboxAgg(115, 70, 15, 15, 14)

	// frame 15
	path = path.appendHboxAgg(102, 67, 15, 15, 15)
	path = path.appendHboxAgg(115, 70, 15, 15, 15)

	// frame 16
	path = path.appendHboxAgg(102, 67, 15, 15, 16)
	path = path.appendHboxAgg(115, 70, 15, 15, 16)

	// frame 17
	path = path.appendHboxAgg(102, 67, 15, 15, 17)
	path = path.appendHboxAgg(115, 70, 15, 15, 17)

	// frame 18+ no hitbox

	startHitboxSim(screen, cp, inc, path, 0)
}

/*
-----------------------------------------------------------------------------
Demon Primary Attack Example End
-----------------------------------------------------------------------------
*/

/*
-----------------------------------------------------------------------------
Monk Secondary Attack Example(only works if cp is Monk)
-----------------------------------------------------------------------------
*/

// var (
// 	hitBoxTest = &hitboxTest{
// 		name:  "primaryAtk",
// 		on:    false,
// 		count: 13,
// 		// set to -1 to play whole anim
// 		frame: -1,
// 		left:  false,
// 		inc:   16.666 * 5, // 1 frame at 60fps
// 	}
// )

// func hitBoxSim(screen *ebiten.Image, cp *PlayerController) {
// 	inc, path := hitBoxSimSetup(hitBoxTest.inc)

// 	// frame 1
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

// 	// frame 2
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

// 	// frame 3
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

// 	// frame 4
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

// 	// frame 5
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

// 	// frame 6
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

// 	// frame 7
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

// 	// frame 8
// 	path = path.appendHboxAgg(45, 30, 10, 20, 8)
// 	path = path.appendHboxAgg(50, 25, 10, 10, 8)

// 	// frame 9
// 	path = path.appendHboxAgg(45, 30, 10, 60, 9)
// 	path = path.appendHboxAgg(60, 25, 10, 40, 9)
// 	path = path.appendHboxAgg(60, 15, 10, 35, 9)
// 	path = path.appendHboxAgg(77, 0, 25, 10, 9)

// 	// frame 10
// 	path = path.appendHboxAgg(45, 30, 10, 60, 10)
// 	path = path.appendHboxAgg(60, 25, 10, 40, 10)
// 	path = path.appendHboxAgg(60, 15, 10, 35, 10)
// 	path = path.appendHboxAgg(77, 0, 25, 10, 10)

// 	// frame 11
// 	path = path.appendHboxAgg(45, 30, 10, 60, 11)
// 	path = path.appendHboxAgg(60, 25, 10, 40, 11)
// 	path = path.appendHboxAgg(60, 15, 10, 35, 11)
// 	path = path.appendHboxAgg(77, 0, 25, 10, 11)

// 	// frame 12
// 	path = path.appendHboxAgg(62, 30, 10, 30, 12)
// 	path = path.appendHboxAgg(72, 20, 10, 15, 12)
// 	path = path.appendHboxAgg(76, 12, 8, 5, 12)

// 	startHitboxSim(screen, cp, inc, path, 0)
// }

/*
-----------------------------------------------------------------------------
Monk Secondary Attack Example(only works if cp is Monk)
-----------------------------------------------------------------------------
*/

// var (
// 	hitBoxTest = &hitboxTest{
// 		name:  "secondaryAtk",
// 		on:    false,
// 		count: 8,
// 		// set to -1 to play whole anim
// 		frame: -1,
// 		left:  false,
// 		inc:   16.666 * 5, // 1 frame at 60fps
// 	}
// )

// func hitBoxSim(screen *ebiten.Image, cp *PlayerController) {
// 	inc, path := hitBoxSimSetup(hitBoxTest.inc)

// 	// frame 0
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

// 	// frame 1
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 1)

// 	// frame 2
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 2)

// 	// frame 3
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 3)

// 	// frame 4
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 4)

// 	// frame 5
// 	path = path.appendHboxAgg(33, -5, 20, 25, 5)

// 	// frame 6
// 	path = path.appendHboxAgg(35, -4, 25, 27, 6)

// 	// frame 7
// 	path = path.appendHboxAgg(38, 12, 20, 20, 7)

// 	startHitboxSim(screen, cp, inc, path, 0)
// }

/*
-----------------------------------------------------------------------------
Monk Secondary Attack Example End
-----------------------------------------------------------------------------
*/

/*
-----------------------------------------------------------------------------
Knight Secondary Attack Example(only works if cp is Knight)
-----------------------------------------------------------------------------
*/

// var (
// 	hitBoxTest = &hitboxTest{
// 		name:  "secondaryAtk",
// 		on:    false,
// 		count: 5,
// 		// set to -1 to play whole anim
// 		frame: -1,
// 		left:  false,
// 		inc:   16.666 * 5, // 1 frame at 60fps
// 	}
// )

// func hitBoxSim(screen *ebiten.Image, cp *PlayerController) {
// 	inc, path := hitBoxSimSetup(hitBoxTest.inc)

// 	// frame 0
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

// 	// frame 1
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 1)

// 	// frame 2
// 	path = path.appendHboxAgg(6, 0, 8, 8, 2)
// 	path = path.appendHboxAgg(14, 0, 8, 8, 2)
// 	path = path.appendHboxAgg(22, 4, 8, 8, 2)
// 	path = path.appendHboxAgg(26, 4, 8, 8, 2)
// 	path = path.appendHboxAgg(28, 8, 8, 8, 2)
// 	path = path.appendHboxAgg(34, 8, 8, 8, 2)
// 	path = path.appendHboxAgg(36, 16, 8, 8, 2)
// 	path = path.appendHboxAgg(34, 24, 8, 8, 2)
// 	path = path.appendHboxAgg(32, 32, 8, 8, 2)
// 	path = path.appendHboxAgg(28, 36, 8, 8, 2)
// 	path = path.appendHboxAgg(22, 36, 8, 8, 2)
// 	path = path.appendHboxAgg(14, 41, 8, 8, 2)
// 	path = path.appendHboxAgg(10, 41, 8, 8, 2)

// 	// frame 3
// 	path = path.appendHboxAgg(38, 24, 8, 8, 3)
// 	path = path.appendHboxAgg(36, 28, 8, 8, 3)
// 	path = path.appendHboxAgg(32, 32, 8, 8, 3)
// 	path = path.appendHboxAgg(28, 36, 8, 8, 3)
// 	path = path.appendHboxAgg(22, 36, 8, 8, 3)
// 	path = path.appendHboxAgg(14, 41, 8, 8, 3)
// 	path = path.appendHboxAgg(10, 41, 8, 8, 3)
// 	path = path.appendHboxAgg(6, 41, 8, 8, 3)
// 	path = path.appendHboxAgg(2, 41, 8, 8, 3)
// 	path = path.appendHboxAgg(-2, 41, 8, 8, 3)

// 	// frame 4
// 	path = path.appendHboxAgg(noBox, noBox, 8, 8, 4)
// 	startHitboxSim(screen, cp, inc, path, 0)
// }

/*
-----------------------------------------------------------------------------
Knight Secondary Attack Example End
-----------------------------------------------------------------------------
*/

/*
	------------------------------------------------
		NO NEED TO TOUCH ANYTHING BELOW HERE
		UNLESS CHANGING HITBOX TEST "SYSTEM"
	------------------------------------------------
*/

func (path hBoxPath) appendHboxAgg(x float64, y float64, h float64, w float64, index int) hBoxPath {

	path[index] = append(path[index], hitBox{
		pOffX:  x,
		pOffY:  y,
		height: h,
		width:  w,
	})

	return path
}

func spawnBox(screen *ebiten.Image, cp *PlayerController, inc float64, path hBoxPath, index int) {

	if len(path) == index {
		hitBoxTest.on = false
		return
	}

	hBoxAgg := path[index]

	colorBox := color.RGBA{R: 128, G: 0, B: 128, A: 255}
	cp.world.bg.Clear()

	for _, hBox := range hBoxAgg {

		if hitBoxTest.left {

			ebitenutil.DrawRect(screen, cp.x-hBox.pOffX-hBox.width+hitBoxTest.pWidth, cp.y+hBox.pOffY, hBox.width, hBox.height, colorBox)
		} else {
			ebitenutil.DrawRect(screen, cp.x+hBox.pOffX, cp.y+hBox.pOffY, hBox.width, hBox.height, colorBox)

		}
	}

	time.AfterFunc(time.Duration(inc)*time.Millisecond, func() {
		spawnBox(screen, cp, inc, path, index+1)
	})
}

func hitBoxSimSetup(inc float64) (float64, hBoxPath) {
	simPath := &hitBoxSequence{
		movmentInc: inc,
	}

	simPath.hBoxPath = make(hBoxPath, hitBoxTest.count)

	return simPath.movmentInc, simPath.hBoxPath
}

func startHitboxSim(screen *ebiten.Image, cp *PlayerController, inc float64, path hBoxPath, index int) {
	hitBoxTest.on = true

	spawnBox(screen, cp, inc, path, 0)
}
