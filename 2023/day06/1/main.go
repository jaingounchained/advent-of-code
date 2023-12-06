package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type race struct {
	time     int
	distance int
}

func parseInput(input []string) []race {
	timeString := strings.Split(input[0], ":")[1]
	timeFields := strings.Fields(timeString)

	distanceString := strings.Split(input[1], ":")[1]
	distanceFields := strings.Fields(distanceString)

	races := make([]race, len(timeFields))
	for i := 0; i < len(races); i++ {
		time, _ := strconv.Atoi(timeFields[i])
		distance, _ := strconv.Atoi(distanceFields[i])
		races[i] = race{
			time,
			distance,
		}
	}

	return races
}

func waitForIt(input []string) {
	startTime := time.Now()

	races := parseInput(input)

	ways := 1

	for _, race := range races {
		validRaces := 0
		for i := 0; i <= race.time; i++ {
			if ((race.time - i) * i) > race.distance {
				validRaces++
			}
		}

		fmt.Println(race, validRaces)

		ways *= validRaces
	}

	fmt.Println(ways)
	fmt.Println("Time elapsed: ", time.Since(startTime))
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day06/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	waitForIt(input)
}
