package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func parseInput(input []string) [][]byte {
	pipeMaze := make([][]byte, 0)
	for _, line := range input {
		pipeMaze = append(pipeMaze, []byte(line))
	}

	return pipeMaze
}

type Point struct {
	y int
	x int
}

func findStartingPoint(pipeMaze [][]byte) Point {
	for y, line := range pipeMaze {
		for x, point := range line {
			if point == 'S' {
				return Point{
					y,
					x,
				}
			}
		}
	}
	return Point{}
}

type Direction int

const (
	north Direction = iota
	east
	south
	west
)

type DirectionInfo struct {
	direction             Direction
	pointAddition         Point
	sourceValidPipes      []byte
	destinationValidPipes []byte
}

func (directionInfo DirectionInfo) validPipe(sourcePipe, destinationPipe byte) bool {
	validSourcePipe := false
	for _, sourceValidPipe := range directionInfo.sourceValidPipes {
		if sourcePipe == sourceValidPipe {
			validSourcePipe = true
			break
		}
	}

	validDestinationPipe := false
	for _, destinationValidPipe := range directionInfo.destinationValidPipes {
		if destinationPipe == destinationValidPipe {
			validDestinationPipe = true
			break
		}
	}

	return validSourcePipe && validDestinationPipe
}

var loopPoints []Point
var seenGrid [][]bool
var directionInfo []DirectionInfo

func validDestination(destinationY, destinationX int, sourcePipeByte byte, pipeMaze [][]byte, directioninfo DirectionInfo) bool {
	// Out of bounds
	if destinationY < 0 || destinationY > len(pipeMaze)-1 {
		return false
	}
	if destinationX < 0 || destinationX > len(pipeMaze[0])-1 {
		return false
	}

	// Point already seen
	// No connecting pipes possible
	return !seenGrid[destinationY][destinationX] &&
		directioninfo.validPipe(sourcePipeByte, pipeMaze[destinationY][destinationX])
}

func floodFillDFS(pipeMaze [][]byte, currentPoint Point) bool {
	// pre recurse
	seenGrid[currentPoint.y][currentPoint.x] = true
	loopPoints = append(loopPoints, currentPoint)

	// recurse
	for _, d := range directionInfo {
		nextPointY := currentPoint.y + d.pointAddition.y
		nextPointX := currentPoint.x + d.pointAddition.x
		if validDestination(
			nextPointY,
			nextPointX,
			pipeMaze[currentPoint.y][currentPoint.x],
			pipeMaze,
			d,
		) {
			floodFillDFS(pipeMaze, Point{
				nextPointY,
				nextPointX,
			})
			return true
		}
	}

	return true
}

func floodFillBFS(pipeMaze [][]byte, startingPoint Point) {
	loopPoints = append(loopPoints, startingPoint)

	// Initialize queue
	queue := []Point{startingPoint}
	seenGrid[startingPoint.y][startingPoint.x] = true

	for len(queue) != 0 {
		currentPoint := queue[0]
		queue = queue[1:]

		for _, d := range directionInfo {
			nextPointY := currentPoint.y + d.pointAddition.y
			nextPointX := currentPoint.x + d.pointAddition.x
			if validDestination(
				nextPointY,
				nextPointX,
				pipeMaze[currentPoint.y][currentPoint.x],
				pipeMaze,
				d,
			) {
				seenGrid[nextPointY][nextPointX] = true
				queue = append(queue, Point{
					nextPointY,
					nextPointX,
				})
				loopPoints = append(loopPoints, Point{
					nextPointY,
					nextPointX,
				})
			}
		}
	}
}

func mod(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func enclosedPoints(loopPoints []Point) int {
	// Shoelace theoram
	enclosedArea := 0
	for i := 0; i < len(loopPoints)-1; i++ {
		enclosedArea += (loopPoints[i].x - loopPoints[i+1].x) * loopPoints[i].y
	}
	enclosedArea += (loopPoints[len(loopPoints)-1].x - loopPoints[0].x) * loopPoints[len(loopPoints)-1].y

	enclosedArea = mod(enclosedArea)

	// Picks theoram: A = i + b/2 -1
	enclosedPoints := enclosedArea - len(loopPoints)/2 + 1

	return enclosedPoints
}

func PipeMaze(input []string) {
	startTime := time.Now()

	pipeMaze := parseInput(input)

	startingPoint := findStartingPoint(pipeMaze)

	directionInfo = []DirectionInfo{
		{
			north,
			Point{-1, 0},
			[]byte{'S', '|', 'J', 'L'},
			[]byte{'|', '7', 'F'},
		},
		{
			east,
			Point{0, 1},
			[]byte{'S', '-', 'L', 'F'},
			[]byte{'-', 'J', '7'},
		},
		{
			south,
			Point{1, 0},
			[]byte{'S', '|', '7', 'F'},
			[]byte{'|', 'L', 'J'},
		},
		{
			west,
			Point{0, -1},
			[]byte{'S', '-', 'J', '7'},
			[]byte{'-', 'L', 'F'},
		},
	}

	loopPoints = make([]Point, 0)

	seenGrid = make([][]bool, len(pipeMaze))
	for i := 0; i < len(pipeMaze); i++ {
		seenGrid[i] = make([]bool, len(pipeMaze[0]))
	}

	floodFillDFS(pipeMaze, startingPoint)
	//	floodFillBFS(pipeMaze, startingPoint)

	fmt.Println("Farthest point: ", len(loopPoints)/2)
	fmt.Println("Enclosed points: ", enclosedPoints(loopPoints))

	fmt.Println("Time elapsed: ", time.Since(startTime))
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day10/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	PipeMaze(input)
}
