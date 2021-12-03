package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func LoadPilotInstructionsFromFile(path string) ([]Vector2, error) {
	// Load move instructions
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}
	instructions := []Vector2{}
	for _, line := range lines {
		bigram := strings.Split(line, " ")
		distance, err := strconv.ParseFloat(bigram[1], 64)
		if err != nil {
			return nil, err
		}
		switch bigram[0] {
		case "up":
			instructions = append(instructions, VECTOR2_UP.Scale(distance))
		case "down":
			instructions = append(instructions, VECTOR2_DOWN.Scale(distance))
		case "forward":
			instructions = append(instructions, VECTOR2_RIGHT.Scale(distance))
		default:
			return nil, errors.New("invalid instruction")
		}
	}
	return instructions, nil
}

func LoadDiagnosticReportFromFile(path string) ([]uint, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}
	var diagnostics []uint
	for _, line := range lines {
		x, err := strconv.ParseUint(line, 2, len(line))
		if err != nil {
			return nil, err

		}
		diagnostics = append(diagnostics, uint(x))
	}

	return diagnostics, nil
}
