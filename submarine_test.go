package main

import (
	"testing"
)

func TestSubmarineScanWindowSize1(t *testing.T) {
	// 1.txt answer: 1482
	ocean, err := NewOceanFloor("input/1.txt")
	if err != nil {
		t.Fatal(err)
	}
	sub := NewSubmarine()
	sub.ScanOceanFloor(ocean, 1)
	depthDropCounter := 0
	for i, _ := range sub.oceanMap.data {
		slope, _ := sub.oceanMap.GetSlopeAtDistance(i)
		if slope < 0 {
			depthDropCounter++
		}
	}
	if depthDropCounter != 1482 {
		t.Fail()
	}
}

func TestSubmarineScanWindowSize3(t *testing.T) {
	// 1.txt answer: 1518
	ocean := OceanMap{data: []Vector2{
		{0, -199}, // A
		{1, -200}, // A B
		{2, -208}, // A B C
		{3, -210}, //   B C D
		{4, -200}, // E   C D
		{5, -207}, // E F   D
		{6, -240}, // E F G
		{7, -269}, //   F G H
		{8, -260}, //     G H
		{9, -263}, //       H
	}}
	sub := NewSubmarine()
	sub.ScanOceanFloor(ocean, 3)
	depthDropCounter := 0
	for i, _ := range sub.oceanMap.data {
		slope, _ := sub.oceanMap.GetSlopeAtDistance(i)
		if slope < 0 {
			depthDropCounter++
		}
	}
	if depthDropCounter != 5 {
		t.Fail()
	}
}
