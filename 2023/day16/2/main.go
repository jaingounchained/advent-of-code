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
				count++
			}
		}
	}
	return count
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func startBeamSpawn(tiles [][]byte, y, x int, direction Direction) int {
	energized := make([][]bool, len(tiles))
	for i := range tiles {
		energized[i] = make([]bool, len(tiles[i]))
	}

	beamSpawnMap := make(map[string]bool)

	beamSpawnMap[fmt.Sprintf("%d %d %d", y, x, direction)] = true

	spawnBeam(tiles, energized, y, x, direction, beamSpawnMap)

	return countEnergizedTiles(energized)
}

func TheFloorWillBeLava(input []string) {
	startTime := time.Now()

	tiles := parseInput(input)

	maxEnergizedTiles := 0

	for i := 0; i < len(tiles[0]); i++ {
		maxEnergizedTiles = MaxInt(
			startBeamSpawn(tiles, 0, i, south),
			maxEnergizedTiles,
		)
		maxEnergizedTiles = MaxInt(
			startBeamSpawn(tiles, len(tiles)-1, i, north),
			maxEnergizedTiles,
		)
	}

	for j := 0; j < len(tiles); j++ {
		maxEnergizedTiles = MaxInt(
			startBeamSpawn(tiles, j, 0, east),
			maxEnergizedTiles,
		)
		maxEnergizedTiles = MaxInt(
			startBeamSpawn(tiles, j, len(tiles[0])-1, west),
			maxEnergizedTiles,
		)
	}

	fmt.Println(maxEnergizedTiles)

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
