package main

import (
	"math"
	"math/rand"
	"time"

	pb "github.com/kainn9/grpc_game/proto"
	r "github.com/kainn9/grpc_game/server/roles"
	"github.com/solarlune/resolv"
)

// player represents a player in the game
type player struct {
	object          *resolv.Object // The Resolv physics object representing the player
	speedX          float64        // The player's horizontal speed
	speedY          float64        // The player's vertical speed
	onGround        *resolv.Object // The Resolv physics object representing the ground the player is standing on
	wallSliding     *resolv.Object // The Resolv physics object representing the wall the player is sliding on
	facingRight     bool           // Whether the player is facing right or left
	ignorePlatform  *resolv.Object // The Resolv physics object representing a platform that the player can ignore collision with
	pid             string         // The player's unique identifier
	worldKey        string         // The key of the world the player is currently in
	r.Role                         // The player's role
	currAttack      *r.AttackData  // The player's current attack
	playerPh                       // The player's physics parameters
	gravBoost       bool           // Whether the player is receiving a gravity boost
	windup          r.AtKey
	chargeStart     time.Time
	chargeValue     float64
	attackMovement  string
	movmentStartX   int
	prevEvent       *pb.PlayerReq
	health          int
	defending       bool
	defenseCooldown bool
	hits            map[string]bool
}

// playerPh represents the physics parameters of a player
type playerPh struct {
	friction float64 // The player's friction
	accel    float64 // The player's acceleration
	maxSpeed float64 // The player's maximum speed
	jumpSpd  float64 // The player's jump speed
	gravity  float64 // The player's gravity
	kbx      float64 // The force of the knockback on the player
	kby      float64 // The force of the knockback on the player
}

type ccString string

const (
	None      ccString = ""
	KnockBack ccString = "Kb"
)

func (cp *player) isCC() ccString {
	if cp.isKnockedBack() {
		return KnockBack
	}

	return None
}

func (cp *player) isKnockedBack() bool {
	return cp.isKnockedBackX() || cp.isKnockedBackY()
}

func (cp *player) isKnockedBackX() bool {
	return cp.kbx != 0
}

func (cp *player) isKnockedBackY() bool {
	return cp.kby != 0
}

func (cp *player) isAttacking() bool {
	return cp.currAttack != nil
}

func (cp *player) isWindingUp() bool {
	return cp.windup != ""
}

func (cp *player) attackMovementActive() bool {
	return cp.attackMovement != ""
}

// canAcceptInputs returns if a player is in a "controllable" state
func (cp *player) canAcceptInputs() bool {
	// returns true if player
	// is not in knockback, windup or attack state
	return !cp.isKnockedBack() && !cp.isWindingUp() && !cp.isAttacking() && !cp.defending
}

// newPlayer creates a new player with the given unique identifier and world key
func newPlayer(pid string, worldKey string) *player {

	// some tempcode for getting a random role each time a player is spawned
	randomRole := make(map[int32]r.Role)
	randomRole[0] = *r.Knight
	randomRole[1] = *r.Monk

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random key
	keys := make([]int32, 0, len(randomRole))

	for k := range randomRole {
		keys = append(keys, k)
	}

	randomKey := keys[rand.Intn(len(keys))]

	role := randomRole[randomKey]
	// end of the temp code

	p := &player{
		pid:           pid,
		worldKey:      worldKey,
		Role:          role,
		movmentStartX: -100,
		health:        role.Health,
		hits:          make(map[string]bool),
	}

	ph := &playerPh{
		friction: role.Phys.DefaultFriction,
		accel:    role.Phys.DefaultAccel,
		maxSpeed: role.Phys.DefaultMaxSpeed,
		jumpSpd:  role.Phys.DefaultJumpSpd,
		gravity:  role.Phys.DefaultGravity,
	}

	p.playerPh = *ph

	return p
}

// addPlayerToSpace adds a player to a Resolv space with the given coordinates
func addPlayerToSpace(space *resolv.Space, p *player, x float64, y float64) *player {

	p.object = resolv.NewObject(x, y, p.HitBoxW, p.HitBoxH)

	p.object.SetShape(resolv.NewRectangle(0, 0, p.object.W, p.object.H))

	initHitboxData(p.object, p, nil)
	space.Add(p.object)

	return p
}

// removePlayerFromGame removes a player from the game with the given unique identifier and world
func removePlayerFromGame(pid string, w *world) {
	// Stop server crash if client disconnects before fully loading/creating a player
	serverConfig.mutex.RLock()
	if w.players[pid] == nil {
		return
	}

	obj := w.players[pid].object
	serverConfig.mutex.RUnlock()

	w.space.Remove(obj)

	serverConfig.mutex.Lock()
	delete(w.players, pid)
	delete(serverConfig.activePlayers, pid)
	serverConfig.mutex.Unlock()
}

// changePlayersWorld swaps a player from their old world to a new world,
// updating their position and worldKey in the process.
func changePlayersWorld(oldWorld *world, newWorld *world, cp *player, x float64, y float64) {
	serverConfig.mutex.Lock()
	delete(oldWorld.players, cp.pid)
	oldWorld.space.Remove(cp.object)
	newWorld.players[cp.pid] = cp
	serverConfig.mutex.Unlock()
	addPlayerToSpace(newWorld.space, cp, x, y)
	cp.worldKey = newWorld.name
}

// worldTransferHandler handles a player's request to switch worlds.
// If the input string is "swap", the player's current world is identified,
// and the player is moved to the alternative world if they are in the main world
// or the main world if they are in the alternative world.
func (cp *player) worldTransferHandler(input string) {
	// World Swap Test!
	if input == "swap" {
		w, k := currentPlayerWorld(cp.pid)

		if k == "alt" {
			changePlayersWorld(w, serverConfig.worldsMap["main"], cp, 612, 500)
		} else {
			changePlayersWorld(w, serverConfig.worldsMap["alt"], cp, 1250, 3700)
		}
		return
	}
}

// gravityHandler applies gravity to the player, adjusting their speedY accordingly.
func (cp *player) gravityHandler() {
	cp.speedY += cp.gravity

	if cp.wallSliding != nil && cp.speedY > 1 {
		cp.speedY = 1
	}
}

// jumpHandler handles a player's jump input.
// If the input is "keySpace", the player will jump if they are not performing an attack,
// and if they are either on the ground or wall sliding.
// If the player is wall sliding, a wall jump will be executed.
func (cp *player) jumpHandler(input string) {

	if input == "keySpace" {
		if input == "keyDown" && cp.onGround != nil && cp.onGround.HasTags("platform") {
			cp.ignorePlatform = cp.onGround
		} else {
			if cp.onGround != nil {
				cp.speedY = -cp.jumpSpd
			} else if cp.wallSliding != nil {
				// WALLJUMPING
				cp.speedY = -cp.jumpSpd

				if cp.wallSliding.X > cp.object.X {
					cp.speedX = -4
				} else {
					cp.speedX = 4
				}

				cp.wallSliding = nil
			}
		}
	}
}

// horizontalMovementHandler handles the horizontal movement of the player based on user input and collision detection.
func (cp *player) horizontalMovementHandler(input string, worldWidth float64) {

	// TODO: Clean this jank up, and make a better way to handle speed modz
	if cp.isKnockedBack() {
		cp.maxSpeed = math.Abs(math.Max(math.Abs(cp.kbx), math.Abs(cp.kby)))
	} else if !cp.attackMovementActive() && !cp.defending {
		cp.maxSpeed = cp.Role.Phys.DefaultMaxSpeed
	}
	// end of TODO above -----------------------------

	if cp.speedX > cp.friction {
		cp.speedX -= cp.friction // decrease speed by friction value if speed is greater than friction
	} else if cp.speedX < -cp.friction {
		cp.speedX += cp.friction // increase speed by friction value if speed is smaller than negative of friction
	} else {
		cp.speedX = 0 // if speed is between negative and positive friction value, set speed to 0
	}

	if cp.speedX > cp.maxSpeed {
		cp.speedX = cp.maxSpeed // limit speed to maxSpeed if it's greater than maxSpeed
	} else if cp.speedX < -cp.maxSpeed {
		cp.speedX = -cp.maxSpeed // limit speed to negative maxSpeed if it's smaller than negative maxSpeed
	}

	// We handle horizontal movement separately from vertical movement. This is, conceptually, decomposing movement into two phases / axes.
	// By decomposing movement in this manner, we can handle each case properly (i.e. stop movement horizontally separately from vertical movement, as
	// necessary). More can be seen on this topic over on this blog post on higherorderfun.com:
	// http://higherorderfun.com/blog/2012/05/20/the-guide-to-implementing-2d-platformers/

	// dx is the horizontal delta movement variable (which is the Player's horizontal speed). If we come into contact with something, then it will
	// be that movement instead.
	dx := cp.speedX

	// Moving horizontally is done fairly simply; we just check to see if something solid is in front of us. If so, we move into contact with it
	// and stop horizontal movement speed. If not, then we can just move forward.
	if check := cp.object.Check(cp.speedX, 0, "solid"); check != nil {
		dx = check.ContactWithCell(check.Cells[0]).X() // set delta movement to the distance to the object we collide with
		cp.speedX = 0                                  // stop horizontal movement
		if cp.onGround == nil {
			cp.wallSliding = check.Objects[0] // set wallSliding to the object we collide with if player is in the air
		}

		cp.endMovment()
		cp.endDefenseMovement()
	}

	// playerOnPlayer X collision
	if check := cp.object.Check(cp.speedX, 0, "player"); check != nil {

		obj := check.Objects[0]

		data := hBoxData(obj)
		otherPlayer := data.player

		if !otherPlayer.defending && !cp.defending {
			dx = check.ContactWithCell(check.Cells[0]).X() // set delta movement to the distance to the player we collide with
			cp.endMovment()

			if (math.Abs(cp.object.X-otherPlayer.object.X) < 3) && (math.Abs(cp.object.Y-otherPlayer.object.Y) < 10) {
				dx += cp.object.W
			}
		}

	}

	// Then we just apply the horizontal movement to the Player's object.
	newXPos := cp.object.X + dx // calculate new x position

	// temp hack to prevent player from moving out of world
	// eventually might just kill player if they go out of bounds
	// or actually add real walls lol
	if newXPos > 30 && newXPos < worldWidth-30 {
		cp.object.X = newXPos // update player's x position if it's within the world's width limits(rando chose 30 for now)

	} else if cp.attackMovementActive() { // don't want players to get stuck in movment(normally will end on collision or distance traveled, but this hack isn't a real collision so...)
		cp.resolveMovment(cp.currAttack)
	} else if cp.defending {
		cp.endDefenseMovement()
	}

}

// This function is responsible for handling the horizontal movement of a player object based on input.
// If the player is not currently wall-sliding, then they can move horizontally by accelerating in the input direction.
func (cp *player) horizontalMovementListener(input string) {
	if cp.wallSliding == nil {
		if input == "keyRight" {
			cp.speedX += cp.accel
			cp.facingRight = true
		}

		if input == "keyLeft" {
			cp.speedX -= cp.accel
			cp.facingRight = false
		}
	}
}

// Handler for vertical movement/collisions where it sets onGround to nil, applies gravity and checks for collisions with
// different objects such as platforms, solid ground, or other players. If a collision occurs, the player is moved to contact
// the object, and any special actions (such as sliding or landing on a platform) are taken.

func (cp *player) verticalMovmentHandler(input string, world *world) {
	// Now for the vertical movement; it's the most complicated because we can land on different types of objects and need
	// to treat them all differently, but overall, it's not bad.

	// First, we set onGround to be nil, in case we don't end up standing on anything.
	cp.onGround = nil

	// dy is the delta movement downward, and is the vertical movement by default; similarly to dx, if we come into contact with
	// something, this will be changed to move to contact instead.

	dy := cp.speedY

	// We want to be sure to lock vertical movement to a maximum of the size of the Cells within the Space
	// so we don't miss any collisions by tunneling through.

	dy = math.Max(math.Min(dy, float64(cellY)), float64(-cellY))

	// We're going to check for collision using dy (which is vertical movement speed), but add one when moving downwards to look a bit deeper down
	// into the ground for solid objects to land on, specifically.
	checkDistance := dy
	if dy >= 0 {
		checkDistance++
	}

	// We check for any solid / stand-able objects.
	if check := cp.object.Check(0, checkDistance, "solid", "platform", "player"); check != nil {

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
			cp.object.X += slide.X()

		} else {

			// Platforms are next; here, we just see if the platform is not being ignored by attempting to drop down,
			// if the player is falling on the platform (as otherwise he would be jumping through platforms), and if the platform is low enough
			// to land on. If so, we stand on it.

			if platforms := check.ObjectsByTags("platform"); len(platforms) > 0 {

				platform := platforms[0]

				if platform != cp.ignorePlatform && cp.speedY >= 0 && cp.object.Bottom() < platform.Y+4 {
					dy = check.ContactWithObject(platform).Y()
					cp.onGround = platform
					cp.speedY = 0
				}

			}
			// basic solids collision
			if check := cp.object.Check(0, cp.speedY, "solid"); check != nil {
				if check.Objects[0].Y > cp.object.Y {
					dy = check.ContactWithCell(check.Cells[0]).Y()

					cp.onGround = check.Objects[0]
				}

				cp.speedY = 0
			}

			// playerOnPlayer y collision
			if check := cp.object.Check(0, cp.speedY, "player"); check != nil {
				if check.Objects[0].Y > cp.object.Y {
					dy = check.ContactWithCell(check.Cells[0]).Y()
					cp.speedY = 0
					cp.onGround = check.Objects[0]
				} else {
					check.Objects[0].Y += dy
					check.Objects[0].Update()
				}

			}

			if cp.onGround != nil {
				cp.wallSliding = nil    // Player's on the ground, so no wallSliding anymore.
				cp.ignorePlatform = nil // Player's on the ground, so reset which platform is being ignored.
			}
		}
	}

	// Move the object on dy.
	newYPos := cp.object.Y + dy

	// temp hack to prevent player from moving out of world
	// eventually might just kill player if they go out of bounds
	// or actually add real walls lol
	if newYPos < world.height-10 && newYPos > 10 {
		cp.object.Y += dy
	}

}
