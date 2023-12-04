package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scratchcards(input []string) {
	scratchCardCopies := make([]int, len(input)*2)
	for i := 0; i < len(input); i++ {
		scratchCardCopies[i]++
	}

	for i, line := range input {
		numbers := strings.Split(line, ":")[1]

		numberSplit := strings.Split(numbers, "|")
		winningNumberStrings := strings.Fields(numberSplit[0])
		myNumberStrings := strings.Fields(numberSplit[1])

		points := 0
		for _, myNumber := range myNumberStrings {
			for _, winningNumber := range winningNumberStrings {
				if myNumber == winningNumber {
					points += 1
					break
				}
			}
		}

		for j := i + 1; j <= i+points; j++ {
			scratchCardCopies[j] += scratchCardCopies[i]
		}

	}

	totalPoints := 0
	for i := 0; i < len(input); i++ {
		totalPoints += scratchCardCopies[i]
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
