package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	pb "github.com/kainn9/grpc_game/proto"
)

var particleImageTest *ebiten.Image

func init() {
	particleImageTest = ebiten.NewImage(10, 10) // Replace 2, 2 with your desired particle size
	particleImageTest.Fill(color.Black)         // Fill the particle image with the color you want
}

func randomColor() color.Color {
	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func DrawParticle(world *World, particle pb.Particle) {
	if particle.Active {
		// Create a new ebiten.Image for the particle with its size
		// particleImage := ebiten.NewImage(int(particle.Size), int(particle.Size))
		randColor := randomColor().(color.RGBA)
		particleImageTest.Fill(randColor)
		// Create a new ebiten.DrawImageOptions to set the particle's position
		opts := &ebiten.DrawImageOptions{}

		pc := world.playerController

		opts = pc.playerCam.GetTranslation(opts, float64(particle.Position.X)-pc.playerCXpos/2, float64(particle.Position.Y)-pc.playerCYpos/2)

		// render particle
		pc.playerCam.Surface.DrawImage(particleImageTest, opts)
	}
}
