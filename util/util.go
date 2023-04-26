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

func SetNthCharTo1(str string, n int) string {

	// Convert the string to a byte slice
	bytes := []byte(str)

	// Set the nth byte to 49 ('1' in ASCII)
	bytes[n] = 49

	// Convert the byte slice back to a string
	return string(bytes)
}

func SetNthCharTo0(str string, n int) string {
	// Convert the string to a byte slice
	bytes := []byte(str)

	// Set the nth byte to 48 ('0' in ASCII)
	bytes[n] = 48

	// Convert the byte slice back to a string
	return string(bytes)
}
