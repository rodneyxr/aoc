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

	// diagnostics is a list of binary numbers representing diagnostic readings
	diagnostics []uint
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

func (sub *Submarine) LoadDiagnostics(diagnostics []uint) {
	sub.diagnostics = diagnostics
}

func (sub *Submarine) CalculatePowerConsumption(bits int) uint {
	return sub.CalculateGammaRate(bits) * sub.CalculateEpsilonRate(bits)
}

// CalculateGammaRate reads diagnostics to find the MOST common bit for each position in the bitstring
func (sub *Submarine) CalculateGammaRate(bits int) uint {
	return CalculateGammaRateAbstract(sub.diagnostics, bits)
}

func CalculateGammaRateAbstract(data []uint, bits int) uint {
	gamma := uint(0b0)
	for i := bits - 1; i >= 0; i-- {
		count := uint(0)
		for _, diag := range data {
			// Get the bit
			bit := (diag >> i) & 0b1

			// Add the bit to the counts
			count += bit
		}
		// If double the count is bigger than the total, that means 1 is
		// the most common
		if count*2 >= uint(len(data)) {
			gamma += uint(1) << uint(i)
		}
	}
	return gamma
}

// CalculateEpsilonRate reads diagnostics to find the LEAST common bit for each position in the bitstring
func (sub *Submarine) CalculateEpsilonRate(bits int) uint {
	return CalculateEpsilonRateAbstract(sub.diagnostics, bits)
}

func CalculateEpsilonRateAbstract(data []uint, bits int) uint {
	mask := uint((math.Pow(2, float64(bits))) - 1)
	return ^CalculateGammaRateAbstract(data, bits) & mask
}

// CalculateOxygenGeneratorRating
func (sub Submarine) CalculateOxygenGeneratorRating(bits int) uint {
	// Copy the diagnostics
	diagCopy := make([]uint, len(sub.diagnostics))
	copy(diagCopy, sub.diagnostics)

	// Calculate the initial oxygen rate
	oxy := sub.CalculateGammaRate(bits)

	// lambda
	nbit := func(x uint, n int) uint { return (x >> (n)) & 0b1}

	// Filter by each bit
	for i := bits; i >= 0; i-- {
		var filter []uint
		for _, diag := range diagCopy {
			if nbit(diag, i) == nbit(oxy, i) {
				filter = append(filter, diag)
			}
		}
		if len(filter) == 0 {
			return diagCopy[0]
		}
		diagCopy = filter
		oxy = CalculateGammaRateAbstract(diagCopy, bits)
	}
	return oxy
}

func (sub Submarine) CalculateC02ScrubberRating(bits int) uint {
	// Copy the diagnostics
	diagCopy := make([]uint, len(sub.diagnostics))
	copy(diagCopy, sub.diagnostics)

	// Calculate the initial oxygen rate
	c02 := sub.CalculateEpsilonRate(bits)

	// lambda
	nbit := func(x uint, n int) uint { return (x >> (n)) & 0b1}

	// Filter by each bit
	for i := bits; i >= 0; i-- {
		var filter []uint
		for _, diag := range diagCopy {
			if nbit(diag, i) == nbit(c02, i) {
				filter = append(filter, diag)
			}
		}
		if len(filter) == 0 {
			return diagCopy[0]
		}
		diagCopy = filter
		c02 = CalculateEpsilonRateAbstract(diagCopy, bits)
	}
	return c02
}

func (sub Submarine) CalculateLifeSupportRating(bits int) uint {
	return sub.CalculateOxygenGeneratorRating(bits) * sub.CalculateC02ScrubberRating(bits)
}

// String returns a nicely formatted string for this object.
func (sub Submarine) String() string {
	return fmt.Sprintf("Position: %s", sub.pos.String())
}
