package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isCharNumber(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

func calculateNumberLeft(y, x int, input []string) string {
	numStr := ""
	for i := x - 1; i >= 0; i-- {
		char := input[y][i]
		if !isCharNumber(char) {
			break
		}
		numStr = string(char) + numStr
	}
	return numStr
}

func calculateNumberRight(y, x int, input []string) string {
	numStr := ""
	for i := x + 1; i < len(input[y]); i++ {
		char := input[y][i]
		if !isCharNumber(char) {
			break
		}
		numStr = numStr + string(char)
	}
	return numStr
}

func combineNumbers(numbers ...string) int {
	str := ""
	for _, number := range numbers {
		str += number
	}
	number, _ := strconv.Atoi(str)
	return number
}

func gear(gearj, geari int, input []string) int {
	numbersAroundGear := []int{}

	if isCharNumber(input[gearj][geari-1]) {
		numbers := []string{calculateNumberLeft(gearj, geari, input)}
		numbersAroundGear = append(numbersAroundGear, combineNumbers(numbers...))
	}

	if isCharNumber(input[gearj][geari+1]) {
		numbers := []string{calculateNumberRight(gearj, geari, input)}
		numbersAroundGear = append(numbersAroundGear, combineNumbers(numbers...))
	}

	if isCharNumber(input[gearj-1][geari]) {
		numbers := []string{}
		if isCharNumber(input[gearj-1][geari-1]) {
			numbers = append(numbers, calculateNumberLeft(gearj-1, geari, input))
		}
		numbers = append(numbers, string(input[gearj-1][geari]))
		if isCharNumber(input[gearj-1][geari+1]) {
			numbers = append(numbers, calculateNumberRight(gearj-1, geari, input))
		}
		numbersAroundGear = append(numbersAroundGear, combineNumbers(numbers...))
	} else {
		if isCharNumber(input[gearj-1][geari-1]) {
			numbers := []string{calculateNumberLeft(gearj-1, geari, input)}
			numbersAroundGear = append(numbersAroundGear, combineNumbers(numbers...))
		}
		if isCharNumber(input[gearj-1][geari+1]) {
			numbers := []string{calculateNumberRight(gearj-1, geari, input)}
			numbersAroundGear = append(numbersAroundGear, combineNumbers(numbers...))
		}
	}
	if isCharNumber(input[gearj+1][geari]) {
		numbers := []string{}
		if isCharNumber(input[gearj+1][geari-1]) {
			numbers = append(numbers, calculateNumberLeft(gearj+1, geari, input))
		}
		numbers = append(numbers, string(input[gearj+1][geari]))
		if isCharNumber(input[gearj+1][geari+1]) {
			numbers = append(numbers, calculateNumberRight(gearj+1, geari, input))
		}
		numbersAroundGear = append(numbersAroundGear, combineNumbers(numbers...))
	} else {
		if isCharNumber(input[gearj+1][geari-1]) {
			numbers := []string{calculateNumberLeft(gearj+1, geari, input)}
			numbersAroundGear = append(numbersAroundGear, combineNumbers(numbers...))
		}
		if isCharNumber(input[gearj+1][geari+1]) {
			numbers := []string{calculateNumberRight(gearj+1, geari, input)}
			numbersAroundGear = append(numbersAroundGear, combineNumbers(numbers...))
		}
	}

	if len(numbersAroundGear) == 2 {
		return numbersAroundGear[0] * numbersAroundGear[1]
	}

	return 0
}

func gearRatios(input []string) {
	gearRatioSum := 0

	for j, line := range input {
		for i, char := range line {
			if char == '*' {
				gearRatioSum += gear(j, i, input)
			}
		}
	}

	fmt.Println(gearRatioSum)
}

func convertInput(input []string) []string {
	dy, dx := len(input), len(input[0])
	convertedInput := make([]string, dy+2)

	convertedInput[0] = strings.Repeat(".", dx+2)
	for j := 1; j <= dy; j++ {
		convertedInput[j] = "." + input[j-1] + "."
	}
	convertedInput[dy+1] = strings.Repeat(".", dx+2)

	return convertedInput
}

func print2dArray(input []string) {
	for _, line := range input {
		fmt.Println(line)
	}
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day03/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	convertedInput := convertInput(input)
	gearRatios(convertedInput)
}
