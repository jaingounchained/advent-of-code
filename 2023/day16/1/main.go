package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func parseInput(input []string) [][]byte {
	tiles := make([][]byte, 0)
	for _, line := range input {
		tiles = append(tiles, []byte(line))
	}
	return tiles
}

type Direction int

const (
	north = iota
	east
	south
	west
)

func spawnBeamHelper(tiles [][]byte, energized [][]bool, y, x int, direction Direction, beamSpawnMap map[string]bool) {
	if _, ok := beamSpawnMap[fmt.Sprintf("%d %d %d", y, x, direction)]; ok {
		return
	}
	beamSpawnMap[fmt.Sprintf("%d %d %d", y, x, direction)] = true
	spawnBeam(tiles, energized, y, x, direction, beamSpawnMap)
}

func spawnBeam(tiles [][]byte, energized [][]bool, y, x int, direction Direction, beamSpawnMap map[string]bool) {
	switch direction {
	case west:
		if x == -1 {
			return
		}
		for i := x; i >= 0; i-- {
			energized[y][i] = true
			switch tiles[y][i] {
			case '\\':
				spawnBeamHelper(tiles, energized, y-1, i, north, beamSpawnMap)
				return
			case '/':
				spawnBeamHelper(tiles, energized, y+1, i, south, beamSpawnMap)
				return
			case '|':
				spawnBeamHelper(tiles, energized, y-1, i, north, beamSpawnMap)
				spawnBeamHelper(tiles, energized, y+1, i, south, beamSpawnMap)
				return
			}
		}
	case east:
		if x == len(tiles[0]) {
			return
		}
		for i := x; i < len(tiles[0]); i++ {
			energized[y][i] = true
			switch tiles[y][i] {
			case '/':
				spawnBeamHelper(tiles, energized, y-1, i, north, beamSpawnMap)
				return
			case '\\':
				spawnBeamHelper(tiles, energized, y+1, i, south, beamSpawnMap)
				return
			case '|':
				spawnBeamHelper(tiles, energized, y-1, i, north, beamSpawnMap)
				spawnBeamHelper(tiles, energized, y+1, i, south, beamSpawnMap)
				return
			}
		}
	case south:
		if y == len(tiles) {
			return
		}
		for j := y; j < len(tiles); j++ {
			energized[j][x] = true
			switch tiles[j][x] {
			case '/':
				spawnBeamHelper(tiles, energized, j, x-1, west, beamSpawnMap)
				return
			case '\\':
				spawnBeamHelper(tiles, energized, j, x+1, east, beamSpawnMap)
				return
			case '-':
				spawnBeamHelper(tiles, energized, j, x-1, west, beamSpawnMap)
				spawnBeamHelper(tiles, energized, j, x+1, east, beamSpawnMap)
				return
			}
		}
	case north:
		if y == -1 {
			return
		}
		for j := y; j >= 0; j-- {
			energized[j][x] = true
			switch tiles[j][x] {
			case '\\':
				spawnBeamHelper(tiles, energized, j, x-1, west, beamSpawnMap)
				return
			case '/':
				spawnBeamHelper(tiles, energized, j, x+1, east, beamSpawnMap)
				return
			case '-':
				spawnBeamHelper(tiles, energized, j, x-1, west, beamSpawnMap)
				spawnBeamHelper(tiles, energized, j, x+1, east, beamSpawnMap)
				return
			}
		}
	}
}

func countEnergizedTiles(energized [][]bool) int {
	count := 0
	for j := 0; j < len(energized); j++ {
		for i := 0; i < len(energized[0]); i++ {
			if energized[j][i] {
				fmt.Print("#")
				count++
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	return count
}

func TheFloorWillBeLava(input []string) {
	startTime := time.Now()

	tiles := parseInput(input)

	energized := make([][]bool, len(tiles))
	for i := range tiles {
		energized[i] = make([]bool, len(tiles[i]))
	}

	beamSpawnMap := make(map[string]bool)

	beamSpawnMap[fmt.Sprintf("%d %d %d", 0, 0, east)] = true

	spawnBeam(tiles, energized, 0, 0, east, beamSpawnMap)

	fmt.Println(countEnergizedTiles(energized))

	fmt.Println("Time elapsed: ", time.Since(startTime))
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day16/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	TheFloorWillBeLava(input)
}
