package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Race struct {
	time     int
	distance int
}

func combineNumberStrings(numbers ...string) int {
	str := ""
	for _, number := range numbers {
		str += number
	}
	number, _ := strconv.Atoi(str)
	return number
}

func parseInput(input []string) Race {
	timeString := strings.Split(input[0], ":")[1]
	timeFields := strings.Fields(timeString)

	distanceString := strings.Split(input[1], ":")[1]
	distanceFields := strings.Fields(distanceString)

	return Race{
		combineNumberStrings(timeFields...),
		combineNumberStrings(distanceFields...),
	}
}

func waitForIt(input []string) {
	startTime := time.Now()

	race := parseInput(input)

	validRaces := 0
	for i := 0; i <= race.time; i++ {
		if ((race.time - i) * i) > race.distance {
			validRaces++
		}
	}

	fmt.Println(validRaces)
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
