package util

import "math"

type Vector2 struct {
	X, Y float64
}

// Add returns the result of adding another vector to the current vector.
func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{X: v.X + other.X, Y: v.Y + other.Y}
}

// Sub returns the result of subtracting another vector from the current vector.
func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{X: v.X - other.X, Y: v.Y - other.Y}
}

// Mul returns the result of scaling the current vector by a scalar value.
func (v Vector2) Mul(scalar float64) Vector2 {
	return Vector2{X: v.X * scalar, Y: v.Y * scalar}
}

// Length returns the length (magnitude) of the current vector.
func (v Vector2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize returns the normalized (unit) vector of the current vector.
func (v Vector2) Normalize() Vector2 {
	length := v.Length()
	if length == 0 {
		return v
	}
	return Vector2{X: v.X / length, Y: v.Y / length}
}

func (v Vector2) Scaled(scale float64) Vector2 {
	return Vector2{
		X: v.X * scale,
		Y: v.Y * scale,
	}
}
