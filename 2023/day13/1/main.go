package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func parseInput(input []string) [][][]byte {
	input = append(input, "")
	patterns := make([][][]byte, 0)

	pattern := make([][]byte, 0)
	for _, line := range input {
		if line == "" {
			patterns = append(patterns, pattern)
			pattern = [][]byte{}
			continue
		}
		pattern = append(pattern, []byte(line))
	}

	return patterns
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findVerticalReflection(pattern [][]byte) int {
	for i := 1; i < len(pattern[0]); i++ {
		isPattern := true
		bound := MinInt(i, len(pattern[0])-i)
		for j, l := i, i-1; j < i+bound; j, l = j+1, l-1 {
			for k := 0; k < len(pattern); k++ {
				if isPattern && (pattern[k][j] != pattern[k][l]) {
					isPattern = false
				}
			}
		}
		if isPattern {
			return i
		}
	}

	return -1
}

func findHorizontalReflection(pattern [][]byte) int {
	for i := 1; i < len(pattern); i++ {
		isPattern := true
		bound := MinInt(i, len(pattern)-i)
		for j, l := i, i-1; j < i+bound; j, l = j+1, l-1 {
			for k := 0; k < len(pattern[0]); k++ {
				if isPattern && (pattern[j][k] != pattern[l][k]) {
					isPattern = false
				}
			}
		}
		if isPattern {
			return i
		}
	}

	return -1
}

func PointOfCoincidence(input []string) {
	startTime := time.Now()

	patterns := parseInput(input)
	summary := 0

	for _, pattern := range patterns {
		if columns := findVerticalReflection(pattern); columns > 0 {
			summary += columns
			continue
		}

		if rows := findHorizontalReflection(pattern); rows > 0 {
			summary += rows * 100
			continue
		}
	}

	fmt.Println(summary)

	fmt.Println("Time elapsed: ", time.Since(startTime))
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day13/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	PointOfCoincidence(input)
}
