package main

import (
	"fmt"
	"log"
)

func main() {
	// too high: 3959450
	diagnostics, err := LoadDiagnosticReportFromFile("input/3.txt")
	if err != nil {
		log.Fatal(err)
	}
	sub := NewSubmarine()
	sub.LoadDiagnostics(diagnostics)

	gamma := sub.CalculateGammaRate(12)
	epsilon := sub.CalculateEpsilonRate(12)

	fmt.Printf("%012b\n", gamma)
	fmt.Println(gamma)
	fmt.Printf("%012b\n", epsilon)
	fmt.Println(epsilon)

	fmt.Println(sub.CalculatePowerConsumption(12))
}
