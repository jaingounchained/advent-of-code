package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scratchcards(input []string) {
	totalPoints := 0

	for _, line := range input {
		points := 0

		numbers := strings.Split(line, ":")[1]

		numberSplit := strings.Split(numbers, "|")
		winningNumberStrings := strings.Fields(numberSplit[0])
		myNumberStrings := strings.Fields(numberSplit[1])

		for _, myNumber := range myNumberStrings {
			for _, winningNumber := range winningNumberStrings {
				if myNumber == winningNumber {
					if points == 0 {
						points = 1
					} else {
						points *= 2
					}
					break
				}
			}
		}

		totalPoints += points
	}

	fmt.Println(totalPoints)
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day04/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	scratchcards(input)
}
