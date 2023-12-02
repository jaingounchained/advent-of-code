package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cubeConondrum(input []string) {
	powerOfCubeSets := 0

	for _, line := range input {
		maxRedCubes, maxBlueCubes, maxGreenCubes := 0, 0, 0

		split := strings.Split(line, ":")

		sets := strings.Split(split[1], ";")
		for _, set := range sets {
			draws := strings.Split(set, ",")

			for _, draw := range draws {
				drawSplit := strings.Split(draw, " ")
				number, _ := strconv.Atoi(drawSplit[1])
				color := drawSplit[2]

				switch color {
				case "red":
					if number > maxRedCubes {
						maxRedCubes = number
					}
				case "blue":
					if number > maxBlueCubes {
						maxBlueCubes = number
					}
				case "green":
					if number > maxGreenCubes {
						maxGreenCubes = number
					}
				}
			}
		}

		powerOfCubeSets += maxRedCubes * maxBlueCubes * maxGreenCubes
	}

	fmt.Println(powerOfCubeSets)
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day02/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	cubeConondrum(input)
}

