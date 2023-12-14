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

func tiltGridNorth(grid [][]byte) {
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

func ParabolicReflectorDish(input []string) {
	startTime := time.Now()

	grid := parseInput(input)

	tiltGridNorth(grid)

	load := calculateLoad(grid)

	fmt.Println(load)

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
