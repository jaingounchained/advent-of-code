package main

import (
	"bufio"
	"fmt"
	"os"
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

func findEmptyRowsColumns(spaceGrid [][]byte) ([]int, []int) {
	emptyRows, emptyColumns := make([]int, 0), make([]int, 0)

	for y := 0; y < len(spaceGrid); y++ {
		galaxy := false
		for x := 0; x < len(spaceGrid[0]); x++ {
			if spaceGrid[y][x] == '#' {
				galaxy = true
				break
			}
		}
		if !galaxy {
			emptyRows = append(emptyRows, y)
		}
	}

	for x := 0; x < len(spaceGrid[0]); x++ {
		galaxy := false
		for y := 0; y < len(spaceGrid); y++ {
			if spaceGrid[y][x] == '#' {
				galaxy = true
				break
			}
		}
		if !galaxy {
			emptyColumns = append(emptyColumns, x)
		}
	}

	return emptyRows, emptyColumns
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

func findShortestPath(point1, point2 Point, emptyRows, emptyColumns []int, spaceExpansionRate int) int {
	largey, shorty := 0, 0
	if point1.y >= point2.y {
		largey, shorty = point1.y, point2.y
	} else {
		largey, shorty = point2.y, point1.y
	}
	emptyRowCount := 0
	for _, emptyRow := range emptyRows {
		if emptyRow > shorty && emptyRow < largey {
			emptyRowCount++
		}
	}

	largex, shortx := 0, 0
	if point1.x >= point2.x {
		largex, shortx = point1.x, point2.x
	} else {
		largex, shortx = point2.x, point1.x
	}
	emptyColumnCount := 0
	for _, emptyColumn := range emptyColumns {
		if emptyColumn > shortx && emptyColumn < largex {
			emptyColumnCount++
		}
	}

	return largey - shorty + emptyRowCount*(spaceExpansionRate-1) + largex - shortx + emptyColumnCount*(spaceExpansionRate-1)
}

func CosmicExpansion(input []string) {
	startTime := time.Now()

	spaceGrid := parseInput(input)

	emptyRows, emptyColumns := findEmptyRowsColumns(spaceGrid)
	galaxyPoints := findGalaxyPoints(spaceGrid)

	shortestPathSum := 0

	for i := 0; i < len(galaxyPoints); i++ {
		for j := i + 1; j < len(galaxyPoints); j++ {
			shortestPathSum += findShortestPath(
				galaxyPoints[i],
				galaxyPoints[j],
				emptyRows,
				emptyColumns,
				1000000,
			)
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
