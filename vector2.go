package main

import (
	"fmt"
	"math"
)

var (
	VECTOR2_UP    = Vector2{0, 1}
	VECTOR2_DOWN  = Vector2{0, -1}
	VECTOR2_LEFT  = Vector2{-1, 0}
	VECTOR2_RIGHT = Vector2{1, 0}
)

// Vector2 is the representation of 2D vectors and points
type Vector2 struct {
	X, Y float64
}

func (v Vector2) Scale(factor float64) Vector2 {
	return Vector2{v.X * factor, v.Y * factor}
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{v.X + other.X, v.Y + other.Y}
}

func (v Vector2) Subtract(other Vector2) Vector2 {
	return Vector2{v.X - other.X, v.Y - other.Y}
}

func (v Vector2) Distance(other Vector2) float64 {
	v0 := v.X - other.X
	v1 := v.Y - other.Y
	return math.Sqrt(v0*v0 + v1*v1)
}

func (v Vector2) Equal(other Vector2) bool {
	return v.X == other.X && v.Y == other.Y
}

// String returns a nicely formatted string for this object.
func (v Vector2) String() string {
	return fmt.Sprintf("(%v, %v)", v.X, v.Y)
}
