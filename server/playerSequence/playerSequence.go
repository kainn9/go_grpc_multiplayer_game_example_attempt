package playerSequence

import "log"

type player interface {
	GetHealth() int
}

func PrintHealth(p player) {
	health := p.GetHealth()

	log.Printf("Health %v\n", health)
}
