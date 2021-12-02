package main

import (
	"errors"
	"fmt"
	"strconv"
)

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
