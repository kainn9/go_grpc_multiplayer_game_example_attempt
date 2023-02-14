package main

import (
	"math"
	"time"

	r "github.com/kainn9/grpc_game/server/roles"
	"github.com/solarlune/resolv"
)

type Player struct {
	Object         *resolv.Object
	SpeedX         float64
	SpeedY         float64
	OnGround       *resolv.Object
	WallSliding    *resolv.Object
	FacingRight    bool
	IgnorePlatform *resolv.Object
	Pid            string
	WorldKey       string
	r.Role
	CurrAttack  *r.Attack
	KnockedBack float64
	PlayerPh
	GravBoost bool
}

type PlayerPh struct {
	friction float64
	accel    float64
	maxSpeed float64
	jumpSpd  float64
	gravity  float64
}

func NewPlayer(pid string, worldKey string) *Player {

	ph := &PlayerPh{
		friction: defaultFriction,
		accel:    defaultAccel,
		maxSpeed: defaultMaxSpeed,
		jumpSpd:  defaultJumpSpd,
		gravity:  defaultGravity,
	}

	p := &Player{
		Pid:      pid,
		WorldKey: worldKey,
		Role:     *r.Knight,
		PlayerPh: *ph,
	}

	return p
}

/*
Creates a resolve object from a player
and attaches it to a resolv space(usually one per world)
*/
func AddPlayerToSpace(space *resolv.Space, p *Player, x float64, y float64) *Player {
	p.Object = resolv.NewObject(x, y, p.HitBoxW, p.HitBoxH)
	p.Object.SetShape(resolv.NewRectangle(0, 0, p.Object.W, p.Object.H))

	space.Add(p.Object)
	return p
}

/*
Helper to disconnect player
*/
func DisconnectPlayer(pid string, w *World) {
	w.Space.Remove(w.Players[pid].Object)
	delete(w.Players, pid)
	delete(activePlayers, pid)

}

/*
Helper to change a players current world
*/
func ChangePlayersWorld(oldWorld *World, newWorld *World, cp *Player, x float64, y float64) {
	delete(oldWorld.Players, cp.Pid)
	oldWorld.Space.Remove(cp.Object)
	newWorld.Players[cp.Pid] = cp
	AddPlayerToSpace(newWorld.Space, cp, x, y)
	cp.WorldKey = newWorld.Name
}

func (cp *Player) WorldTransferHandler(input string) {
	// World Swap Test!
	if input == "swap" {
		w, k := CurrentPlayerWorld(cp.Pid)

		if k == "alt" {
			ChangePlayersWorld(w, worldsMap["main"], cp, 1250, 3700)
		} else {
			ChangePlayersWorld(w, worldsMap["alt"], cp, 612, 500)
		}
		return
	}

}

func (cp *Player) CanMove() bool {
	return cp.CurrAttack == nil && cp.KnockedBack == 0
}

func (cp *Player) Gravity() {
	cp.SpeedY += cp.gravity

	if cp.WallSliding != nil && cp.SpeedY > 1 {
		cp.SpeedY = 1
	}

}

func (cp *Player) JumpHandler(input string) {
	if !cp.CanMove() {
		return
	}

	if input == "keySpace" {
		if input == "keyDown" && cp.OnGround != nil && cp.OnGround.HasTags("platform") {

			cp.IgnorePlatform = cp.OnGround

		} else {

			if cp.OnGround != nil {
				cp.SpeedY = -cp.jumpSpd
			} else if cp.WallSliding != nil {
				// WALLJUMPING
				cp.SpeedY = -cp.jumpSpd

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

func (cp *Player) AttackHandler(input string, world *World) {
	
	if input == "primaryAttack" && cp.CurrAttack == nil {
		cp.PrimaryAttack(world)
	}


}

func (cp *Player) AttackedHandler() {
	if check := cp.Object.Check(cp.SpeedX, cp.SpeedY, "attack"); check != nil {

		atkObjs := check.Objects

		for _, o := range atkObjs {
			// use map to check who attacks belongs to don't effect owner
			if o != nil {

				attacker := AOTP[o]

				if attacker == cp || attacker == nil {
					continue
				}

				if attacker.Object.X > cp.Object.X {
					cp.SpeedY = 0
					cp.KnockedBack = -100

				} else {
					cp.SpeedY = 0
					cp.KnockedBack = 100
				}

				time.AfterFunc(1*time.Second, func() { cp.KnockedBack = 0 })
			}

		}

	}

	if cp.KnockedBack != 0 {
		cp.SpeedX += cp.KnockedBack
	}
}

func (cp *Player) HorizontalMovementHandler(input string, worldWidth float64) {

	// Can't move while attacking(for now)
	if cp.CurrAttack != nil{
		cp.SpeedX = 0
		return

	}

	if cp.KnockedBack != 0 {
		cp.maxSpeed = 30
	} else {
		cp.maxSpeed = defaultMaxSpeed
		HorizontalMovementListener(input, cp)
	}

	// Apply friction and horizontal speed limiting.
	if cp.SpeedX > cp.friction {
		cp.SpeedX -= cp.friction
	} else if cp.SpeedX < -cp.friction {
		cp.SpeedX += cp.friction
	} else {
		cp.SpeedX = 0
	}

	if cp.SpeedX > cp.maxSpeed {
		cp.SpeedX = cp.maxSpeed
	} else if cp.SpeedX < -cp.maxSpeed {
		cp.SpeedX = -cp.maxSpeed
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

	// Then we just apply the horizontal movement to the Player's Object.
	newXPos := cp.Object.X + dx

	if newXPos > 30 && newXPos < worldWidth-30 {
		cp.Object.X = newXPos
	}

}

func HorizontalMovementListener(input string, cp *Player) {
	// Horizontal movement is only possible when not wallsliding.
	if cp.WallSliding == nil {
		if input == "keyRight" {
			cp.SpeedX += cp.accel
			cp.FacingRight = true

		}

		if input == "keyLeft" {
			cp.SpeedX -= cp.accel
			cp.FacingRight = false
		}
	}
}

func (cp *Player) VerticalMovmentHandler(input string, world *World) {
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
	newYPos := cp.Object.Y + dy

	if newYPos < world.Height - 10 && newYPos > 10 {
			cp.Object.Y += dy
	}

}

// TODO: Move this somewhere btr
type CCString string

const (
	None      CCString = ""
	KnockBack CCString = "Kb"
)

func (cp *Player) isCC() CCString {
	if cp.KnockedBack != 0 {
		return KnockBack
	}
	return None
}


func(cp *Player) PrimaryAttack(world *World) {

		cs := cp.Object.Space
		atk := cp.Attacks[r.PrimaryAttackKey]
		cp.CurrAttack = atk

		atkObj := resolv.NewObject(cp.Object.X+atk.OffsetX, cp.Object.Y+atk.OffsetY, atk.Width, atk.Height, "attack")

		if !cp.FacingRight {
			// modify to calc player width as origin is top left corner I think
			atkObj = resolv.NewObject(cp.Object.X-(atk.OffsetX-atk.Width/2), cp.Object.Y+atk.OffsetY, atk.Width, atk.Height, "attack")
		}

		AOTP[atkObj] = cp

		cp.CurrAttack = atk

		cs.Add(
			atkObj,
		)

		time.AfterFunc(time.Duration(atk.Duration)*time.Millisecond, func() {
			world.removeAtk(atkObj)
			cp.CurrAttack = nil
		})
}
