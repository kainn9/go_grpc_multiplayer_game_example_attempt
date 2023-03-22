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

const noBox = -90000

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
Monk Secondary Attack Example(only works if cp is Monk)
-----------------------------------------------------------------------------
*/

var (
	hitBoxTest = &hitboxTest{
		name:  "primaryAtk",
		on:    false,
		count: 13,
		// set to -1 to play whole anim
		frame: -1,
		left:  false,
		inc:   16.666 * 5, // 1 frame at 60fps
	}
)

func hitBoxSim(screen *ebiten.Image, cp *PlayerController) {
	inc, path := hitBoxSimSetup(hitBoxTest.inc)

	// frame 1
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 2
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 3
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 4
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 5
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 6
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 7
	path = path.appendHboxAgg(noBox, noBox, 8, 8, 0)

	// frame 8
	path = path.appendHboxAgg(45, 30, 10, 20, 8)
	path = path.appendHboxAgg(50, 25, 10, 10, 8)

	// frame 9
	path = path.appendHboxAgg(45, 30, 10, 60, 9)
	path = path.appendHboxAgg(60, 25, 10, 40, 9)
	path = path.appendHboxAgg(60, 15, 10, 35, 9)
	path = path.appendHboxAgg(77, 0, 25, 10, 9)

	// frame 10
	path = path.appendHboxAgg(45, 30, 10, 60, 10)
	path = path.appendHboxAgg(60, 25, 10, 40, 10)
	path = path.appendHboxAgg(60, 15, 10, 35, 10)
	path = path.appendHboxAgg(77, 0, 25, 10, 10)

	// frame 11
	path = path.appendHboxAgg(45, 30, 10, 60, 11)
	path = path.appendHboxAgg(60, 25, 10, 40, 11)
	path = path.appendHboxAgg(60, 15, 10, 35, 11)
	path = path.appendHboxAgg(77, 0, 25, 10, 11)

	// frame 12
	path = path.appendHboxAgg(62, 30, 10, 30, 12)
	path = path.appendHboxAgg(72, 20, 10, 15, 12)
	path = path.appendHboxAgg(76, 12, 8, 5, 12)

	startHitboxSim(screen, cp, inc, path, 0)
}

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
Tert Attack Example End
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
			ebitenutil.DrawRect(screen, cp.x-(hBox.pOffX-hBox.width/2), cp.y+hBox.pOffY, hBox.width, hBox.height, colorBox)
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
