package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func parseInput(input []string) [][]byte {
	spaceGrid := make([][]byte, 0)

	for _, line := range input {
		points := []byte(line)
		spaceGrid = append(spaceGrid, points)
	}

	return spaceGrid
}

func printSpaceGrid(spaceGrid [][]byte) {
	for _, bytes := range spaceGrid {
		fmt.Println(string(bytes))
	}
}

func transposeMatrix(matrix [][]byte) [][]byte {
	transpose := make([][]byte, 0)

	for x := 0; x < len(matrix[0]); x++ {
		row := make([]byte, 0)
		for y := 0; y < len(matrix); y++ {
			row = append(row, matrix[y][x])
		}
		transpose = append(transpose, row)
	}

	return transpose
}

func expandSpaceGridRowWise(spaceGrid [][]byte) [][]byte {
	expandedSpaceGrid := make([][]byte, 0)

	for y := 0; y < len(spaceGrid); y++ {
		galaxy := false
		for x := 0; x < len(spaceGrid[0]); x++ {
			if spaceGrid[y][x] == '#' {
				galaxy = true
				break
			}
		}
		expandedSpaceGrid = append(expandedSpaceGrid, spaceGrid[y])
		if !galaxy {
			expandedSpaceGrid = append(expandedSpaceGrid, []byte(strings.Repeat(".", len(spaceGrid[0]))))
		}
	}

	return expandedSpaceGrid
}

func expandSpaceGrid(spaceGrid [][]byte) [][]byte {
	spaceGrid2 := expandSpaceGridRowWise(spaceGrid)

	spaceGrid3 := transposeMatrix(spaceGrid2)

	spaceGrid4 := expandSpaceGridRowWise(spaceGrid3)

	return spaceGrid4
}

type Point struct {
	y int
	x int
}

func findGalaxyPoints(spaceGrid [][]byte) []Point {
	points := make([]Point, 0)

	for y, grid := range spaceGrid {
		for x, point := range grid {
			if point == '#' {
				points = append(points, Point{
					y,
					x,
				})
			}
		}
	}

	return points
}

func mod(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func findShortestPath(point1, point2 Point) int {
	return mod(point1.y-point2.y) + mod(point1.x-point2.x)
}

func CosmicExpansion(input []string) {
	startTime := time.Now()

	spaceGrid := parseInput(input)

	galaxyPoints := findGalaxyPoints(expandSpaceGrid(spaceGrid))

	shortestPathSum := 0

	for i := 0; i < len(galaxyPoints); i++ {
		for j := i + 1; j < len(galaxyPoints); j++ {
			shortestPathSum += findShortestPath(galaxyPoints[i], galaxyPoints[j])
		}
	}

	fmt.Println(shortestPathSum)

	fmt.Println("Time elapsed: ", time.Since(startTime))
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day11/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	CosmicExpansion(input)
}
