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

func TestSubmarineMove(t *testing.T) {
	sub := NewSubmarine()
	sub.MoveForward(5)
	sub.AimDown(5)
	sub.MoveForward(8)
	sub.AimUp(3)
	sub.AimDown(8)
	sub.MoveForward(2)

	if !sub.pos.Equal(Vector2{15, -60}) {
		t.Fail()
	}
}

func TestSubmarineDiagnostics(t *testing.T) {
	diagnostics, err := LoadDiagnosticReportFromFile("input/test/3.txt")
	if err != nil {
		t.Fatal(err)
	}
	sub := NewSubmarine()
	sub.LoadDiagnostics(diagnostics)
	bits := 5
	gammaRate := sub.CalculateGammaRate(bits)
	//fmt.Printf("%05b\n", gammaRate)
	if gammaRate != 0b10110 {
		t.Fail()
	}

	epsilonRate := sub.CalculateEpsilonRate(bits)
	//fmt.Printf("%05b\n", epsilonRate)
	if epsilonRate != 0b01001 {
		t.Fail()
	}

	powerConsumption := sub.CalculatePowerConsumption(bits)
	if powerConsumption != 198 {
		t.Fail()
	}
}
