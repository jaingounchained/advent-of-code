package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ifCharIsSymbol(char byte) bool {
	if (char < '0' || char > '9') && char != '.' {
		return true
	}
	return false
}

func checkIfNumberIncluded(numStartingIndex, numEndingIndex, numj int, input []string) bool {
	for j := numj - 1; j <= numj+1; j++ {
		if ifCharIsSymbol(input[j][numStartingIndex-1]) {
			return true
		}
		if ifCharIsSymbol(input[j][numEndingIndex+1]) {
			return true
		}
	}

	for i := numStartingIndex; i <= numEndingIndex; i++ {
		if ifCharIsSymbol(input[numj-1][i]) {
			return true
		}
		if ifCharIsSymbol(input[numj+1][i]) {
			return true
		}
	}
	return false
}

func gearRatios(input []string) {
	sum := 0

	for j, line := range input {
		for i := 0; i < len(line); i++ {
			numStartingIndex := i
			if line[i] >= '0' && line[i] <= '9' {
				num := 0
				for k := i; k < len(line); k, i = k+1, i+1 {
					if line[i] < '0' || line[i] > '9' {
						break
					}
					num = num*10 + int(line[k]-'0')
				}
				numEndingIndex := i - 1
				if checkIfNumberIncluded(numStartingIndex, numEndingIndex, j, input) {
					sum += num
				}
			}
		}
	}

	fmt.Println(sum)
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
