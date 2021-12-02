package main

import (
	"fmt"
)

// Submarine
type Submarine struct {
	// pos is the position of the submarine
	pos Vector2

	// oceanMap is a map of the ocean floor relative to the submarine
	oceanMap OceanMap
}

// NewSubmarine
func NewSubmarine() Submarine {
	return Submarine{pos: Vector2{0, 0}}
}

func (sub *Submarine) MoveUp(distance float64) {
	sub.pos = sub.pos.Add(VECTOR2_UP.Scale(distance))
}

func (sub *Submarine) MoveDown(distance float64) {
	sub.pos = sub.pos.Add(VECTOR2_DOWN.Scale(distance))
}

func (sub *Submarine) MoveForward(distance float64) {
	sub.pos = sub.pos.Add(VECTOR2_RIGHT.Scale(distance))
}

func (sub *Submarine) MoveBack(distance float64) {
	sub.pos = sub.pos.Add(VECTOR2_LEFT.Scale(distance))
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
	for _, x := range instructions {
		sub.Move(x)
	}
}

// String returns a nicely formatted string for this object.
func (sub Submarine) String() string {
	return fmt.Sprintf("Position: %s", sub.pos.String())
}
