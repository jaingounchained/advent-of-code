package main

import (
	"bufio"
	"fmt"
	"os"
)

func trebuchet(input []string) {
	calibrationCount := 0

	for _, line := range input {
		calibration := 0

		for _, char := range line {
			if char >= '0' && char <= '9' {
				calibration += 10 * (int(char) - '0')
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				calibration += int(line[i]) - '0'
				break
			}
		}

		calibrationCount += calibration
	}

	fmt.Println(calibrationCount)
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day-1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	trebuchet(input)
}
