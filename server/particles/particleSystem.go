package particle

import (
	"image/color"

	proto "github.com/kainn9/grpc_game/proto"
	utils "github.com/kainn9/grpc_game/util"
)

type ParticleSystem struct {
	Particles []Particle
}

func (ps *ParticleSystem) AddParticle(position utils.Vector2, velocity utils.Vector2, size float64, color color.Color, lifetime float64) {
	particle := NewParticle(position, velocity, size, color, lifetime)
	ps.Particles = append(ps.Particles, particle)
}

func (ps *ParticleSystem) Update(dt float64) {
	for i := 0; i < len(ps.Particles); i++ {
		particle := &ps.Particles[i]

		// Skip the update for inactive particles
		if !particle.Active {
			continue
		}

		// Update the particle's position based on its velocity and the time delta (dt)
		particle.Position = particle.Position.Add(particle.Velocity.Scaled(dt))

		// Update the particle's age
		particle.Age += dt

		// Deactivate particles that have exceeded their lifetime
		if particle.Age >= particle.Lifetime {
			particle.Active = false
		}
	}
}

func NewParticle(position utils.Vector2, velocity utils.Vector2, size float64, color color.Color, lifetime float64) Particle {
	return Particle{
		Position: position,
		Velocity: velocity,
		Size:     size,
		Color:    color,
		Lifetime: lifetime,
		Age:      0,
		Active:   true,
		Damage:   387,
	}
}

func ConvertToProtoParticleSystem(ps *ParticleSystem) *proto.ParticleSystem {
	// Create a new instance of the protobuf ParticleSystem
	protoPS := proto.ParticleSystem{}

	// Copy over each particle
	for i := 0; i < len(ps.Particles); i++ {
		currentParticle := ps.Particles[i]
		protoParticle := proto.Particle{
			Position: &proto.Vector2{
				X: float32(currentParticle.Position.X),
				Y: float32(currentParticle.Position.Y),
			},
			Velocity: &proto.Vector2{
				X: float32(currentParticle.Velocity.X),
				Y: float32(currentParticle.Velocity.Y),
			},
			Size:     currentParticle.Size,
			Id:       currentParticle.Id,
			Lifetime: currentParticle.Lifetime,
			Age:      currentParticle.Age,
			Damage:   currentParticle.Damage,
			Active:   currentParticle.Active,
		}
		protoPS.Particles = append(protoPS.Particles, &protoParticle)
	}

	return &protoPS
}
