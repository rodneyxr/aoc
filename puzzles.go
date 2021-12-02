package main

import (
	"fmt"
	"log"
)

func day1Part1() {
	ocean, err := NewOceanFloor("input/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	sub := NewSubmarine()
	sub.ScanOceanFloor(ocean, 1)
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

func day1Part2() {
	ocean, err := NewOceanFloor("input/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	sub := NewSubmarine()
	sub.ScanOceanFloor(ocean, 3)
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

// breaks day2Part1
func day2Part2() {
	sub := NewSubmarine()
	instructions, err := LoadPilotInstructionsFromFile("input/2.txt")
	if err != nil {
		log.Fatal(err)
	}
	sub.Pilot(instructions)
	fmt.Println(sub)
	fmt.Print(int(sub.pos.X * -sub.pos.Y))
}