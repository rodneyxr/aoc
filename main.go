package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
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

func (sub *Submarine) ScanOceanFloor(ocean OceanMap) {
	for _, v := range ocean.data {
		relativePos := v.Subtract(sub.pos)
		sub.oceanMap.addData(relativePos.X, relativePos.Y)
	}
}

// String returns a nicely formatted string for this object.
func (sub Submarine) String() string {
	return fmt.Sprintf("Position: %s", sub.pos.String())
}

// OceanMap represents a map of the ocean floor
type OceanMap struct {
	data []Vector2
}

// NewOceanFloor
func NewOceanFloor(path string) (OceanMap, error) {
	ocean := OceanMap{}

	// Read the input for day1
	lines, err := readLines(path)
	if err != nil {
		return ocean, err
	}

	// Convert the lines to Vector2 coordinates
	for x, y := range lines {
		if depth, err := strconv.ParseFloat(y, 64); err != nil {
			return ocean, err
		} else {
			ocean.addData(float64(x), -depth)
		}
	}

	return ocean, nil
}

func (o *OceanMap) addData(distance, depth float64) {
	o.data = append(o.data, Vector2{X: distance, Y: depth})
}

func (o *OceanMap) GetSlopeAtDistance(distance int) (float64, error) {
	// Check if distance exists
	if distance < 0 || distance > len(o.data) {
		return 0, errors.New("no data for given distance")
	}

	// Calculate the slope between the previous distance
	if distance == 0 {
		return 0, nil
	}

	v1 := o.data[distance-1]
	v2 := o.data[distance]
	slope := (v2.Y - v1.Y) / (v2.X - v1.X)
	return slope, nil
}

func (o OceanMap) Print() {
	for _, v := range o.data {
		fmt.Println(v)
	}
}

func main() {
	ocean, err := NewOceanFloor("input/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	sub := NewSubmarine()
	sub.ScanOceanFloor(ocean)
	//sub.oceanMap.Print()
	depthDropCounter := 0
	for i, v := range sub.oceanMap.data {
		slope, _ := sub.oceanMap.GetSlopeAtDistance(i)
		fmt.Println(v, slope)
		if slope < 0 {
			depthDropCounter++
		}
	}
	fmt.Println(depthDropCounter)
}
