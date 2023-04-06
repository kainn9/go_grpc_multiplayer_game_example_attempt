package util

import (
	"math/rand"
	"time"
)

// returns random number from 0 -> n - 1
func RandomInt(n int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(n)
}
