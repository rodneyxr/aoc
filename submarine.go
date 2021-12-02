package main

import (
	"fmt"
	"math"
)

// Submarine
type Submarine struct {
	// pos is the position of the submarine
	pos Vector2

	// aim is the face of the submarine
	aim float64

	// oceanMap is a map of the ocean floor relative to the submarine
	oceanMap OceanMap
}

// NewSubmarine
func NewSubmarine() Submarine {
	return Submarine{pos: Vector2{0, 0}, aim: 0}
}

func (sub *Submarine) MoveUp(distance float64) {
	sub.pos = sub.pos.Add(VECTOR2_UP.Scale(distance))
}

func (sub *Submarine) MoveDown(distance float64) {
	sub.pos = sub.pos.Add(VECTOR2_DOWN.Scale(distance))
}

func (sub *Submarine) MoveBack(distance float64) {
	sub.pos = sub.pos.Add(VECTOR2_LEFT.Scale(distance))
}

func (sub *Submarine) MoveForward(distance float64) {
	move := Vector2{distance, distance * sub.aim}
	sub.pos = sub.pos.Add(move)
}

func (sub *Submarine) AimUp(amount float64) {
	sub.aim += amount
}

func (sub *Submarine) AimDown(amount float64) {
	sub.aim -= amount
}

func (sub *Submarine) Move(direction Vector2) {
	sub.pos = sub.pos.Add(direction)
}

func (sub *Submarine) ScanOceanFloor(ocean OceanMap, scanSize int) {
	sub.oceanMap = OceanMap{}
	for i := 0; i <= len(ocean.data)-scanSize; i++ {
		depthSum := 0.0
		for j := i; j < i+scanSize; j++ {
			depthSum += ocean.data[j].Y
		}
		sub.oceanMap.addData(float64(i), depthSum)
	}
	//for _, v := range ocean.data {
	//	relativePos := v.Subtract(sub.pos)
	//	sub.oceanMap.addData(relativePos.X, relativePos.Y)
	//}
}

func (sub *Submarine) Pilot(instructions []Vector2) {
	for _, move := range instructions {
		if move.Y < 0 { // Down
			sub.AimDown(math.Abs(move.Y))
		} else if move.Y > 0 { // Up
			sub.AimUp(math.Abs(move.Y))
		} else if move.X > 0 { // Forward
			sub.MoveForward(move.X)
		}
	}
}

// String returns a nicely formatted string for this object.
func (sub Submarine) String() string {
	return fmt.Sprintf("Position: %s", sub.pos.String())
}
