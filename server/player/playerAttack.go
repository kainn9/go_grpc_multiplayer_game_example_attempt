package player

import (
	"time"

	"image/color"
	"math/rand"

	particle "github.com/kainn9/grpc_game/server/particles"
	r "github.com/kainn9/grpc_game/server/roles"
	"github.com/kainn9/grpc_game/util"
	ut "github.com/kainn9/grpc_game/util"
)

// attackHandler handles player's attack inputs
func (cp *Player) AttackHandler(input string, world World) {

	// can't attack while attacking yo
	if cp.CurrAttack != nil {
		return
	}

	if input == string(r.PrimaryAttackKey) {
		// lol testing
		if cp.Role.Attacks[r.PrimaryAttackKey].Damage == 420 {
			ps := world.GetParticleSystem()
			createRandomParticles(ps, 50, cp)
		}
		cp.attack(world, r.PrimaryAttackKey)
	}

	if input == string(r.SecondaryAttackKey) {
		cp.attack(world, r.SecondaryAttackKey)
	}

	if input == string(r.TertAttackKey) {
		cp.attack(world, r.TertAttackKey)
	}

	if input == string(r.QuaternaryAttackKey) {
		cp.attack(world, r.QuaternaryAttackKey)
	}
}

func (cp *Player) attack(world World, atKey r.AtKey) {
	atk := cp.Attacks[atKey]

	if atk == nil {
		return
	}

	cp.CdStringMutex.RLock()
	onCd := string(cp.CdString[r.AtkOrderMap[atKey]]) == "1"
	cp.CdStringMutex.RUnlock()

	if onCd {
		return
	}

	if atk.Cooldown != 0 {

		cp.CdStringMutex.Lock()
		newCdString := ut.SetNthCharTo1(cp.CdString, r.AtkOrderMap[atKey])
		cp.CdString = newCdString
		cp.CdStringMutex.Unlock()

		time.AfterFunc((time.Duration(atk.Cooldown))*time.Millisecond, func() {
			cp.CdStringMutex.Lock()
			defer cp.CdStringMutex.Unlock()

			newCdString := ut.SetNthCharTo0(cp.CdString, r.AtkOrderMap[atKey])
			cp.CdString = newCdString
		})
	}

	if atk.InvincibleNoBox {
		cp.InvincibleNoBox = true
	}

	cp.windupPhase(atk, atKey)

}

func createRandomParticles(particleSystem *particle.ParticleSystem, count int, cp *Player) {
	for i := 0; i < count; i++ {
		position := util.Vector2{
			X: cp.Object.X,
			Y: cp.Object.Y,
		}
		velocity := util.Vector2{
			X: rand.Float64()*100 + 3, // random value between 1 and 3
			Y: rand.Float64() * 20,    // random value between 0 and 2
		}
		size := rand.Float64()*15 + 1       // random value between 1 and 5
		color := color.RGBA{R: 255, A: 255} // fixed color: red
		lifetime := rand.Float64()*3 + 2    // random value between 2 and 5 seconds
		particleSystem.AddParticle(position, velocity, size, color, lifetime)
	}
}
