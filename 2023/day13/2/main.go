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

func findVerticalReflections(pattern [][]byte) []int {
	columns := make([]int, 0)
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
			columns = append(columns, i)
		}
	}

	return columns
}

func findHorizontalReflections(pattern [][]byte) []int {
	rows := make([]int, 0)
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
			rows = append(rows, i)
		}
	}

	return rows
}

func switchChar(y, x int, pattern [][]byte) {
	if pattern[y][x] == '.' {
		pattern[y][x] = '#'
	} else {
		pattern[y][x] = '.'
	}
}

func printPattern(pattern [][]byte) {
	for _, line := range pattern {
		fmt.Println(string(line))
	}
}

func uniqueElements(numbers []int) []int {
	uniqueMap := make(map[int]bool)
	for _, n := range numbers {
		uniqueMap[n] = true
	}

	uniqueList := make([]int, 0)
	for k, _ := range uniqueMap {
		uniqueList = append(uniqueList, k)
	}

	return uniqueList
}

func PointOfCoincidence(input []string) {
	startTime := time.Now()

	patterns := parseInput(input)
	summary := 0

	for _, pattern := range patterns {
		originalColumns := findVerticalReflections(pattern)
		originalRows := findHorizontalReflections(pattern)

		rowsAfterFixingSmudge := make([]int, 0)
		columnsAfterFixingSmudge := make([]int, 0)
		for j := 0; j < len(pattern); j++ {
			for i := 0; i < len(pattern[0]); i++ {
				switchChar(j, i, pattern)

				columnsAfterFixingSmudge = append(columnsAfterFixingSmudge, findVerticalReflections(pattern)...)
				rowsAfterFixingSmudge = append(rowsAfterFixingSmudge, findHorizontalReflections(pattern)...)

				switchChar(j, i, pattern)
			}
		}

		uniqueRows := uniqueElements(rowsAfterFixingSmudge)
		uniqueColumns := uniqueElements(columnsAfterFixingSmudge)

		for _, n := range uniqueRows {
			summary += n * 100
		}
		for _, n := range uniqueColumns {
			summary += n
		}
		for _, n := range originalRows {
			summary -= n * 100
		}
		for _, n := range originalColumns {
			summary -= n
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
