package player

import (
	"math/rand"
	"sync"
	"time"

	pb "github.com/kainn9/grpc_game/proto"

	r "github.com/kainn9/grpc_game/server/roles"
	se "github.com/kainn9/grpc_game/server/statusEffects"
	"github.com/solarlune/resolv"
)

// player represents a player in the game
type Player struct {
	Object               *resolv.Object // The Resolv physics object representing the player
	SpeedX               float64        // The player's horizontal speed
	SpeedY               float64        // The player's vertical speed
	OnGround             *resolv.Object // The Resolv physics object representing the ground the player is standing on
	WallSliding          *resolv.Object // The Resolv physics object representing the wall the player is sliding on
	FacingRight          bool           // Whether the player is facing right or left
	IgnorePlatform       *resolv.Object // The Resolv physics object representing a platform that the player can ignore collision with
	Pid                  string         // The player's unique identifier
	*r.Role                             // The player's role
	CurrAttack           *r.AttackData  // The player's current attack
	PlayerPh                            // The player's physics parameters
	Windup               r.AtKey
	ChargeStart          time.Time
	ChargeValue          float64
	AttackMovement       string
	MovmentStartX        int
	PrevEvent            *pb.PlayerReq
	Health               int
	Defending            bool
	Hits                 map[string]bool
	HitsMutex            sync.RWMutex
	Dying                bool
	DeathCallBackPending bool
	KbStamp              time.Time
	KbStampMutex         sync.RWMutex
	roleMutex            sync.RWMutex
	CurrentWorld         World
	CdString             string
	CdStringMutex        sync.RWMutex
	InvincibleNoBox      bool
}

// playerPh represents the physics parameters of a player
type PlayerPh struct {
	Friction float64 // The player's friction
	Accel    float64 // The player's acceleration
	MaxSpeed float64 // The player's maximum speed
	JumpSpd  float64 // The player's jump speed
	Gravity  float64 // The player's gravity
	Kbx      float64 // The force of the knockback on the player
	Kby      float64 // The force of the knockback on the player
}

type World interface {
	GetHitboxMutex() *sync.RWMutex
	GetSpace() *resolv.Space
	RemoveHitboxFromSpace(*resolv.Object)
	GetIndex() int
	GetHeight() float64
	SpawnAtkBox(*Player, *r.AttackData, int, string)
}

func (cp *Player) IsCC() se.CCString {
	if cp.IsHit() {
		return se.Hit
	}

	if cp.IsStunned() {
		return se.Stun
	}

	if cp.IsKnockedBack() {
		return se.KnockBack
	}
	return se.None
}

func (cp *Player) IsHit() bool {
	return cp.Kbx == se.HitFloat && cp.Kby == se.HitFloat
}

func (cp *Player) IsStunned() bool {
	return cp.Kbx == se.StunFloat && cp.Kby == se.StunFloat
}

func (cp *Player) IsKnockedBack() bool {
	return cp.IsKnockedBackX() || cp.IsKnockedBackY()
}

func (cp *Player) IsKnockedBackX() bool {
	return cp.Kbx != 0
}

func (cp *Player) IsKnockedBackY() bool {
	return cp.Kby != 0
}

func (cp *Player) IsAttacking() bool {
	return cp.CurrAttack != nil
}

func (cp *Player) IsWindingUp() bool {
	return cp.Windup != ""
}

func (cp *Player) AttackMovementActive() bool {
	return cp.AttackMovement != ""
}

func (cp *Player) CanAcceptInputs() bool {

	return !cp.Dying && !cp.IsKnockedBack() && !cp.IsWindingUp() && !cp.IsAttacking() && !cp.Defending
}

// newPlayer creates a new player with the given unique identifier and world key
func NewPlayer(pid string, world World) *Player {

	// some tempcode for getting a random role each time a player is spawned
	randomRole := make(map[int32]*r.Role)
	randomRole[0] = r.Knight
	randomRole[1] = r.Monk
	randomRole[2] = r.Demon
	randomRole[3] = r.Werewolf
	randomRole[4] = r.Mage
	randomRole[5] = r.HeavyKnight
	randomRole[6] = r.BirdDroid

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random key
	keys := make([]int32, 0, len(randomRole))

	for k := range randomRole {
		keys = append(keys, k)
	}

	randomKey := keys[rand.Intn(len(keys))]

	role := randomRole[randomKey] // change to lock role
	// end of the temp code

	p := &Player{
		Pid:           pid,
		CurrentWorld:  world,
		Role:          role,
		MovmentStartX: -100,
		Health:        role.Health,
		Hits:          make(map[string]bool),
		KbStampMutex:  sync.RWMutex{},
		roleMutex:     sync.RWMutex{},
		CdString:      "00000",
		CdStringMutex: sync.RWMutex{},
		HitsMutex:     sync.RWMutex{},
	}

	ph := &PlayerPh{
		Friction: role.Phys.DefaultFriction,
		Accel:    role.Phys.DefaultAccel,
		MaxSpeed: role.Phys.DefaultMaxSpeed,
		JumpSpd:  role.Phys.DefaultJumpSpd,
		Gravity:  role.Phys.DefaultGravity,
	}

	p.PlayerPh = *ph

	return p
}

func (cp *Player) RotateRoleData(newRole *r.Role) {
	cp.roleMutex.Lock()
	defer cp.roleMutex.Unlock()

	cp.Role = newRole
	cp.Health = cp.Role.Health

	cp.PlayerPh = PlayerPh{
		Friction: cp.Role.Phys.DefaultFriction,
		Accel:    cp.Role.Phys.DefaultAccel,
		MaxSpeed: cp.Role.Phys.DefaultMaxSpeed,
		JumpSpd:  cp.Role.Phys.DefaultJumpSpd,
		Gravity:  cp.Role.Phys.DefaultGravity,
	}
}
