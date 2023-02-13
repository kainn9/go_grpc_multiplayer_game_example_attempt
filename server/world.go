package main

import (
	"math"

	pb "github.com/kainn9/grpc_game/proto"
	"github.com/solarlune/resolv"
	"github.com/tanema/gween"
)

type World struct {
	Space  *resolv.Space
	FloatingPlatform      *resolv.Object
	FloatingPlatformTween *gween.Sequence
	State *pb.PlayerResp
	ScreenHeight float64
	ScreenWidth float64
	Stream pb.PlayersService_PlayerLocationServer
	Players map[string]*Player
	Name string
}

/*
	Creates a New World and Initializes it
*/
func NewWorld(height float64, width float64, worldBuilder BuilderFunc, name string) *World {
	w := &World{
		ScreenWidth: 	width,
		ScreenHeight:	height,
		Players: make(map[string]*Player),
		Name: name,
	}

	w.Init(worldBuilder)
	return w
}

/*
	Initializes world physics using Resolv Lib,
	*World's attrs, and a builder func
*/
func (world *World) Init(worldBuilder BuilderFunc) {
	gw := world.ScreenWidth
	gh := world.ScreenHeight


	// Define the world's Space. Here, a Space is essentially a grid (the game's width and height, or 640x360), made up of 16x16 cells. Each cell can have 0 or more Objects within it,
	// and collisions can be found by checking the Space to see if the Cells at specific positions contain (or would contain) Objects. This is a broad, simplified approach to collision
	// detection.
	world.Space = resolv.NewSpace(int(gw), int(gh), 8, 8)

	// Construct the solid level geometry. Note that the simple approach of checking cells in a Space for collision works simply when the geometry is aligned with the cells
	worldBuilder(world, gw, gh)
}


/*
	Currently where all game logic happens
 	physics is basically a rip of the resolv example:
	https://github.com/SolarLune/resolv/blob/master/examples/worldPlatformer.go
 */
func (world *World) Update(cp *Player, input string) {
	mutex.Lock()
	defer mutex.Unlock()

	cp.Object.AddTags("player")

	// World Swap Test!
	if input == "swap" {
		w, k := CurrentPlayerWorld(cp.Pid)


		if k == "alt" {
			ChangePlayersWorld(w, worldsMap["main"], cp)
		} else {
			ChangePlayersWorld(w, worldsMap["alt"], cp)
		}
		return
	}

	// Now we update the Player's movement. This is the real bread-and-butter of this example, naturally.

	friction := 0.5
	accel := 0.5 + friction
	maxSpeed := 4.0
	jumpSpd := 10.0
	gravity := 0.75

	cp.SpeedY += gravity

	if cp.WallSliding != nil && cp.SpeedY > 1 {
		cp.SpeedY = 1
	}

	// Horizontal movement is only possible when not wallsliding.
	if cp.WallSliding == nil {
		if input == "keyRight" {
			cp.SpeedX += accel
			cp.FacingRight = true

		}

		if input == "keyLeft" {
			cp.SpeedX -= accel
			cp.FacingRight = false
		}
	}


	// Apply friction and horizontal speed limiting.
	if cp.SpeedX > friction {
		cp.SpeedX -= friction
	} else if cp.SpeedX < -friction {
		cp.SpeedX += friction
	} else {
		cp.SpeedX = 0
	}

	if cp.SpeedX > maxSpeed {
		cp.SpeedX = maxSpeed
	} else if cp.SpeedX < -maxSpeed {
		cp.SpeedX = -maxSpeed
	}

	// Check for jumping.
	if input == "keySpace" {
		if input == "keyDown" && cp.OnGround != nil && cp.OnGround.HasTags("platform") {

			cp.IgnorePlatform = cp.OnGround

		} else {
			
			if cp.OnGround != nil {
				cp.SpeedY = -jumpSpd
			} else if cp.WallSliding != nil {
				// WALLJUMPING
				cp.SpeedY = -jumpSpd

				if cp.WallSliding.X > cp.Object.X {
					cp.SpeedX = -4
				} else {
					cp.SpeedX = 4
				}

				cp.WallSliding = nil
			}
		}

	}

	// We handle horizontal movement separately from vertical movement. This is, conceptually, decomposing movement into two phases / axes.
	// By decomposing movement in this manner, we can handle each case properly (i.e. stop movement horizontally separately from vertical movement, as
	// necesseary). More can be seen on this topic over on this blog post on higherorderfun.com:
	// http://higherorderfun.com/blog/2012/05/20/the-guide-to-implementing-2d-platformers/

	// dx is the horizontal delta movement variable (which is the Player's horizontal speed). If we come into contact with something, then it will
	// be that movement instead.
	dx := cp.SpeedX

	// Moving horizontally is done fairly simply; we just check to see if something solid is in front of us. If so, we move into contact with it
	// and stop horizontal movement speed. If not, then we can just move forward.

	if check := cp.Object.Check(cp.SpeedX, 0, "solid"); check != nil {

		dx = check.ContactWithCell(check.Cells[0]).X()
		cp.SpeedX = 0

		// If you're in the air, then colliding with a wall object makes you start wall sliding.
		if cp.OnGround == nil {
			cp.WallSliding = check.Objects[0]
		}

	}

	// playerOnPlayer X collision
	if check := cp.Object.Check(cp.SpeedX, 0, "player"); check != nil {
		dx = check.ContactWithCell(check.Cells[0]).X()
	}


	// Then we just apply the horizontal movement to the Player's Object. Easy-peasy.
	cp.Object.X += dx

	// Now for the vertical movement; it's the most complicated because we can land on different types of objects and need
	// to treat them all differently, but overall, it's not bad.

	// First, we set OnGround to be nil, in case we don't end up standing on anything.
	cp.OnGround = nil

	// dy is the delta movement downward, and is the vertical movement by default; similarly to dx, if we come into contact with
	// something, this will be changed to move to contact instead.

	dy := cp.SpeedY

	// We want to be sure to lock vertical movement to a maximum of the size of the Cells within the Space
	// so we don't miss any collisions by tunneling through.

	dy = math.Max(math.Min(dy, 8), -8)

	// We're going to check for collision using dy (which is vertical movement speed), but add one when moving downwards to look a bit deeper down
	// into the ground for solid objects to land on, specifically.
	checkDistance := dy
	if dy >= 0 {
		checkDistance++
	}

	// We check for any solid / stand-able objects. In actuality, there aren't any other Objects
	// with other tags in this Space, so we don't -have- to specify any tags, but it's good to be specific for clarity in this example.
	if check := cp.Object.Check(0, checkDistance, "solid", "platform", "ramp", "player"); check != nil {

		// So! Firstly, we want to see if we jumped up into something that we can slide around horizontally to avoid bumping the Player's head.

		// Sliding around a misspaced jump is a small thing that makes jumping a bit more forgiving, and is something different polished platformers
		// (like the 2D Mario games) do to make it a smidge more comfortable to play. For a visual example of this, see this excellent devlog post
		// from the extremely impressive indie game, Leilani's Island: https://forums.tigsource.com/index.php?topic=46289.msg1387138#msg1387138

		// To accomplish this sliding, we simply call Collision.SlideAgainstCell() to see if we can slide.
		// We pass the first cell, and tags that we want to avoid when sliding (i.e. we don't want to slide into cells that contain other solid objects).

		slide := check.SlideAgainstCell(check.Cells[0], "solid", "player")

		// We further ensure that we only slide if:
		// 1) We're jumping up into something (dy < 0),
		// 2) If the cell we're bumping up against contains a solid object,
		// 3) If there was, indeed, a valid slide left or right, and
		// 4) If the proposed slide is less than 8 pixels in horizontal distance. (This is a relatively arbitrary number that just so happens to be half the
		// width of a cell. This is to ensure the player doesn't slide too far horizontally.)

		if dy < 0 && check.Cells[0].ContainsTags("solid", "player") && slide != nil && math.Abs(slide.X()) <= 8 {

			// If we are able to slide here, we do so. No contact was made, and vertical speed (dy) is maintained upwards.
			cp.Object.X += slide.X()

		} else {

			// If sliding -fails-, that means the Player is jumping directly onto or into something, and we need to do more to see if we need to come into
			// contact with it. Let's press on!

			// First, we check for ramps. For ramps, we can't simply check for collision with Check(), as that's not precise enough. We need to get a bit
			// more information, and so will do so by checking its Shape (a triangular ConvexPolygon, as defined in BaseWorld.Init()) against the
			// Player's Shape (which is also a rectangular ConvexPolygon).

			// We get the ramp by simply filtering out Objects with the "ramp" tag out of the objects returned in our broad Check(), and grabbing the first one
			// if there's any at all.
			if ramps := check.ObjectsByTags("ramp"); len(ramps) > 0 {

				ramp := ramps[0]

				// For simplicity, this code assumes we can only stand on one ramp at a time as there is only one ramp in this example.
				// In actuality, if there was a possibility to have a potential collision with multiple ramps (i.e. a ramp that sits on another ramp, and the player running down
				// one onto the other), the collision testing code should probably go with the ramp with the highest confirmed intersection point out of the two.

				// Next, we see if there's been an intersection between the two Shapes using Shape.Intersection. We pass the ramp's shape, and also the movement
				// we're trying to make horizontally, as this makes Intersection return the next y-position while moving, not the one directly
				// underneath the Player. This would keep the player from getting "stuck" when walking up a ramp into the top of a solid block, if there weren't
				// a landing at the top and bottom of the ramp.

				// We use 8 here for the Y-delta so that we can easily see if you're running down the ramp (in which case you're probably in the air as you
				// move faster than you can fall in this example). This way we can maintain contact so you can always jump while running down a ramp. We only
				// continue with coming into contact with the ramp as long as you're not moving upwards (i.e. jumping).

				if contactSet := cp.Object.Shape.Intersection(dx, 8, ramp.Shape); dy >= 0 && contactSet != nil {

					// If Intersection() is successful, a ContactSet is returned. A ContactSet contains information regarding where
					// two Shapes intersect, like the individual points of contact, the center of the contacts, and the MTV, or
					// Minimum Translation Vector, to move out of contact.

					// Here, we use ContactSet.TopmostPoint() to get the top-most contact point as an indicator of where
					// we want the player's feet to be. Then we just set that position, and we're done.

					dy = contactSet.TopmostPoint()[1] - cp.Object.Bottom() + 0.1
					cp.OnGround = ramp
					cp.SpeedY = 0

				}

			}

			// Platforms are next; here, we just see if the platform is not being ignored by attempting to drop down,
			// if the player is falling on the platform (as otherwise he would be jumping through platforms), and if the platform is low enough
			// to land on. If so, we stand on it.

			// Because there's a moving floating platform, we use Collision.ContactWithObject() to ensure the player comes into contact
			// with the top of the platform object. An alternative would be to use Collision.ContactWithCell(), but that would be only if the
			// platform didn't move and were aligned with the Spatial cellular grid.

			if platforms := check.ObjectsByTags("platform"); len(platforms) > 0 {

				platform := platforms[0]

				if platform != cp.IgnorePlatform && cp.SpeedY >= 0 && cp.Object.Bottom() < platform.Y+4 {
					dy = check.ContactWithObject(platform).Y()
					cp.OnGround = platform
					cp.SpeedY = 0
				}

			}

			// Finally, we check for simple solid ground. If we haven't had any success in landing previously, or the solid ground
			// is higher than the existing ground (like if the platform passes underneath the ground, or we're walking off of solid ground
			// onto a ramp), we stand on it instead. We don't check for solid collision first because we want any ramps to override solid
			// ground (so that you can walk onto the ramp, rather than sticking to solid ground).

			// We use ContactWithObject() here because otherwise, we might come into contact with the moving platform's cells (which, naturally,
			// would be selected by a Collision.ContactWithCell() call because the cell is closest to the Player).

			if solids := check.ObjectsByTags("solid"); len(solids) > 0 && (cp.OnGround == nil || cp.OnGround.Y >= solids[0].Y) {
				dy = check.ContactWithObject(solids[0]).Y()
				cp.SpeedY = 0

				// We're only on the ground if we land on it (if the object's Y is greater than the player's).
				if solids[0].Y > cp.Object.Y {
					cp.OnGround = solids[0]
				}

			}

			// check y here
			// playerOnPlayer y collision
			if check := cp.Object.Check(0, cp.SpeedY, "player"); check != nil {
				if check.Objects[0].Y > cp.Object.Y {
					dy = check.ContactWithCell(check.Cells[0]).Y()
					cp.SpeedY = 0 // hmmm
					cp.OnGround = check.Objects[0]
				}
	
			}

			if cp.OnGround != nil {
				cp.WallSliding = nil    // Player's on the ground, so no wallsliding anymore.
				cp.IgnorePlatform = nil // Player's on the ground, so reset which platform is being ignored.
			}
		}
	}

	// Move the object on dy.
	cp.Object.Y += dy

	wallNext := 1.0
	if !cp.FacingRight {
		wallNext = -1
	}

	// If the wall next to the Player runs out, stop wall sliding.
	if c := cp.Object.Check(wallNext, 0, "solid"); cp.WallSliding != nil && c == nil {
		cp.WallSliding = nil
	}

	cp.Object.Update() // Update the player's position in the space.

}
