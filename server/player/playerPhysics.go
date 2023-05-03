package player

import (
	"math"

	gc "github.com/kainn9/grpc_game/server/globalConstants"
	r "github.com/kainn9/grpc_game/server/roles"
)

func validXCollision(cp *Player, otherPlayer *Player) bool {
	return (!otherPlayer.Defending || otherPlayer.Defense.DefenseType == r.DefenseBlock) && (!cp.Defending || cp.Defense.DefenseType == r.DefenseBlock) && (!cp.Dying && !otherPlayer.Dying) && (!cp.InvincibleNoBox && !otherPlayer.InvincibleNoBox)
}

// gravityHandler applies gravity to the player, adjusting their speedY accordingly.
func (cp *Player) GravityHandler() {
	cp.SpeedY += cp.Gravity

	if cp.WallSliding != nil && cp.SpeedY > 1 {
		cp.SpeedY = 1
	}
}

// jumpHandler handles a player's jump input.
// If the input is "keySpace", the player will jump if they are not performing an attack,
// and if they are either on the ground or wall sliding.
// If the player is wall sliding, a wall jump will be executed.
func (cp *Player) JumpHandler(input string) {

	if input == "keySpace" {
		if input == "keyDown" && cp.OnGround != nil && cp.OnGround.HasTags("platform") {
			// cp.IgnorePlatform = cp.OnGround
		} else {
			if cp.OnGround != nil {
				cp.SpeedY = -cp.JumpSpd
			} else if cp.WallSliding != nil {
				// WALLJUMPING
				cp.SpeedY = -cp.JumpSpd

				if cp.WallSliding.X > cp.Object.X {
					cp.SpeedX = -4
				} else {
					cp.SpeedX = 4
				}

				cp.WallSliding = nil
			}
		}
	}
}

// Handler for vertical movement/collisions where it sets onGround to nil, applies gravity and checks for collisions with
// different objects such as platforms, solid ground, or other players. If a collision occurs, the player is moved to contact
// the object, and any special actions (such as sliding or landing on a platform) are taken.
func (cp *Player) VerticalMovmentHandler(input string, world World) {
	// Now for the vertical movement; it's the most complicated because we can land on different types of objects and need
	// to treat them all differently, but overall, it's not bad.

	// First, we set onGround to be nil, in case we don't end up standing on anything.
	cp.OnGround = nil

	// dy is the delta movement downward, and is the vertical movement by default; similarly to dx, if we come into contact with
	// something, this will be changed to move to contact instead.

	dy := cp.SpeedY

	// We want to be sure to lock vertical movement to a maximum of the size of the Cells within the Space
	// so we don't miss any collisions by tunneling through.

	dy = math.Max(math.Min(dy, float64(gc.CELL_Y)), float64(-gc.CELL_Y))

	// We're going to check for collision using dy (which is vertical movement speed), but add one when moving downwards to look a bit deeper down
	// into the ground for solid objects to land on, specifically.
	checkDistance := dy
	if dy >= 0 {
		checkDistance++
	}

	// We check for any solid / stand-able objects.
	if check := cp.Object.Check(0, checkDistance, "solid", "platform", "player"); check != nil {

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

			// Platforms are next; here, we just see if the platform is not being ignored by attempting to drop down,
			// if the player is falling on the platform (as otherwise he would be jumping through platforms), and if the platform is low enough
			// to land on. If so, we stand on it.

			if platforms := check.ObjectsByTags("platform"); len(platforms) > 0 {

				minY := platforms[0].Y
				minP := platforms[0]
				for i, p := range platforms {
					minY = math.Max(minY, p.Y) // lower y actually means lower pos
					if minY == p.Y {
						minP = platforms[i]
					}
				}

				if cp.SpeedY >= 0 && cp.Object.Bottom() < minP.Y+4 {
					cp.OnGround = minP
				}

				if input == "keyDown" && cp.OnGround != nil && cp.OnGround.HasTags("platform") {
					cp.IgnorePlatform = cp.OnGround
				}

				if minP != cp.IgnorePlatform && cp.SpeedY >= 0 && cp.Object.Bottom() < minP.Y+4 {
					dy = check.ContactWithObject(minP).Y()
					cp.SpeedY = 0
				}
			}

			// basic solids collision
			if check := cp.Object.Check(0, cp.SpeedY, "solid"); check != nil {
				if check.Objects[0].Y > cp.Object.Y {
					dy = check.ContactWithCell(check.Cells[0]).Y()

					cp.OnGround = check.Objects[0]
				}

				cp.SpeedY = 0
			}

			// playerOnPlayer Y collision
			if check := cp.Object.Check(0, cp.SpeedY, "player"); check != nil {

				invalidPlayerYCollide := (cp.Dying || HBoxData(check.Objects[0]).Player.Dying)

				if (check.Objects[0].Y > cp.Object.Y) && !invalidPlayerYCollide {
					dy = check.ContactWithCell(check.Cells[0]).Y()
					cp.SpeedY = 0
					cp.OnGround = check.Objects[0]
				} else {
					check.Objects[0].Y += dy
				}

			}

			if cp.OnGround != nil {
				cp.WallSliding = nil    // Player's on the ground, so no wallSliding anymore.
				cp.IgnorePlatform = nil // Player's on the ground, so reset which platform is being ignored.
			}
		}
	}

	// Move the object on dy.
	newYPos := cp.Object.Y + dy

	// top bounds(top left is 0,0)
	if newYPos > 0 {
		cp.Object.Y += dy
	}

	// player fell too far
	if newYPos > cp.CurrentWorld.GetHeight()-cp.Object.H {
		cp.death()
	}

}

// horizontalMovementHandler handles the horizontal movement of the player based on user input and collision detection.
func (cp *Player) HorizontalMovementHandler(input string) {

	// TODO: Clean this jank up, and make a better way to handle speed modz
	if cp.IsKnockedBack() {
		cp.MaxSpeed = math.Abs(math.Max(math.Abs(cp.Kbx), math.Abs(cp.Kby)))
	} else if !cp.AttackMovementActive() && !cp.Defending {
		cp.MaxSpeed = cp.Role.Phys.DefaultMaxSpeed
	}
	// end of TODO above -----------------------------

	if cp.SpeedX > cp.Friction {
		cp.SpeedX -= cp.Friction // decrease speed by friction value if speed is greater than friction
	} else if cp.SpeedX < -cp.Friction {
		cp.SpeedX += cp.Friction // increase speed by friction value if speed is smaller than negative of friction
	} else {
		cp.SpeedX = 0 // if speed is between negative and positive friction value, set speed to 0
	}

	if cp.SpeedX > cp.MaxSpeed {
		cp.SpeedX = cp.MaxSpeed // limit speed to maxSpeed if it's greater than maxSpeed
	} else if cp.SpeedX < -cp.MaxSpeed {
		cp.SpeedX = -cp.MaxSpeed // limit speed to negative maxSpeed if it's smaller than negative maxSpeed
	}

	// We handle horizontal movement separately from vertical movement. This is, conceptually, decomposing movement into two phases / axes.
	// By decomposing movement in this manner, we can handle each case properly (i.e. stop movement horizontally separately from vertical movement, as
	// necessary). More can be seen on this topic over on this blog post on higherorderfun.com:
	// http://higherorderfun.com/blog/2012/05/20/the-guide-to-implementing-2d-platformers/

	// dx is the horizontal delta movement variable (which is the Player's horizontal speed). If we come into contact with something, then it will
	// be that movement instead.
	dx := cp.SpeedX

	// Moving horizontally is done fairly simply; we just check to see if something solid is in front of us. If so, we move into contact with it
	// and stop horizontal movement speed. If not, then we can just move forward.
	if check := cp.Object.Check(cp.SpeedX, 0, "solid", "bounds"); check != nil {
		dx = check.ContactWithCell(check.Cells[0]).X() // set delta movement to the distance to the object we collide with
		cp.SpeedX = 0                                  // stop horizontal movement
		if cp.OnGround == nil && check.Objects[0].HasTags("solid") {
			cp.WallSliding = check.Objects[0] // set wallSliding to the object we collide with if player is in the air
		}

		cp.endMovment()
		cp.endDefenseMovement()
	}

	// playerOnPlayer X collision
	if check := cp.Object.Check(cp.SpeedX, 0, "player"); check != nil {

		obj := check.Objects[0]

		data := HBoxData(obj)
		otherPlayer := data.Player

		if validXCollision(cp, otherPlayer) {
			dx = check.ContactWithCell(check.Cells[0]).X() // set delta movement to the distance to the player we collide with
			cp.endMovment()

			if (math.Abs(cp.Object.X-otherPlayer.Object.X) < 3) && (math.Abs(cp.Object.Y-otherPlayer.Object.Y) < 10) {
				dx += cp.Object.W
			}
		}

	}

	// Then we just apply the horizontal movement to the Player's object.
	newXPos := cp.Object.X + dx // calculate new x position
	cp.Object.X = newXPos
}

// This function is responsible for handling the horizontal movement of a player object based on input.
// If the player is not currently wall-sliding, then they can move horizontally by accelerating in the input direction.
func (cp *Player) HorizontalMovementListener(input string) {
	if cp.WallSliding == nil {
		if input == "keyRight" {
			cp.SpeedX += cp.Accel
			cp.FacingRight = true
		}

		if input == "keyLeft" {
			cp.SpeedX -= cp.Accel
			cp.FacingRight = false
		}
	}
}
