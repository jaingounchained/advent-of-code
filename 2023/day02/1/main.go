package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cubeConondrum(input []string) {
	idSum := 0

	for _, line := range input {
		gameValid := true

		split := strings.Split(line, ":")

		gameIdInfo := strings.Split(split[0], " ")
		gameId, _ := strconv.Atoi(gameIdInfo[1])

		sets := strings.Split(split[1], ";")
		for _, set := range sets {
			draws := strings.Split(set, ",")

			for _, draw := range draws {
				drawSplit := strings.Split(draw, " ")
				number, _ := strconv.Atoi(drawSplit[1])
				color := drawSplit[2]

				switch color {
				case "red":
					if number > 12 {
						gameValid = false
						break
					}
				case "blue":
					if number > 14 {
						gameValid = false
						break
					}
				case "green":
					if number > 13 {
						gameValid = false
						break
					}
				}
			}

			if !gameValid {
				break
			}
		}

		if gameValid {
			idSum += gameId
		}
	}

	fmt.Println(idSum)
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
