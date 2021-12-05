package main

import (
	"fmt"
	"log"
)

func main() {
	diagnostics, err := LoadDiagnosticReportFromFile("input/3.txt")
	if err != nil {
		log.Fatal(err)
	}
	sub := NewSubmarine()
	sub.LoadDiagnostics(diagnostics)
	fmt.Println(sub.CalculateLifeSupportRating(12))
}
