package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func parseInput(input []string) [][]byte {
	grid := make([][]byte, 0)

	for _, line := range input {
		grid = append(grid, []byte(line))
	}

	return grid
}

func printGrid(grid [][]byte) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}

func rotateClockwise(grid [][]byte) [][]byte {
	rotatedGrid := make([][]byte, 0)

	for i := 0; i < len(grid[0]); i++ {
		line := make([]byte, 0)
		for j := len(grid) - 1; j >= 0; j-- {
			line = append(line, grid[j][i])
		}
		rotatedGrid = append(rotatedGrid, line)
	}

	return rotatedGrid
}

func tiltNorth(grid [][]byte) {
	for i := 0; i < len(grid[0]); i++ {
		for j := 1; j < len(grid); j++ {
			if grid[j][i] == 'O' {
				shiftRockPosition := -1
				for k := j - 1; k >= 0; k-- {
					if grid[k][i] == 'O' || grid[k][i] == '#' {
						shiftRockPosition = k
						break
					}
				}
				grid[j][i], grid[shiftRockPosition+1][i] = grid[shiftRockPosition+1][i], grid[j][i]
			}
		}
	}
}

func calculateLoad(grid [][]byte) int {
	load := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'O' {
				load += len(grid) - i
			}
		}
	}

	return load
}

func stringFromByteArray(byteArray [][]byte) string {
	str := ""
	for _, e := range byteArray {
		str += string(e)
	}
	return str
}

func ParabolicReflectorDish(input []string) {
	startTime := time.Now()

	originalGrid := parseInput(input)

	gridCopy := append([][]byte(nil), originalGrid...)

	seenGrids := make(map[string]int)

	cycleStart, cycleLength := 0, 0
	i := 0
	for {
		i++
		for i := 0; i < 4; i++ {
			tiltNorth(gridCopy)
			gridCopy = rotateClockwise(gridCopy)
		}
		gridStr := stringFromByteArray(gridCopy)
		if v, ok := seenGrids[gridStr]; ok {
			cycleStart = v
			cycleLength = i - v
			break
		}
		seenGrids[gridStr] = i
	}

	identicalGridIndex := (1000000000-cycleStart)%cycleLength + cycleStart

	for i := 1; i <= identicalGridIndex; i++ {
		for i := 0; i < 4; i++ {
			tiltNorth(originalGrid)
			originalGrid = rotateClockwise(originalGrid)
		}
	}

	fmt.Println(calculateLoad(originalGrid))

	fmt.Println("Time elapsed: ", time.Since(startTime))
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day14/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	ParabolicReflectorDish(input)
}
