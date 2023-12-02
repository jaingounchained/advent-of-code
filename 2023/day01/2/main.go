package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func trebuchet(input []string) {
	calibrationCount := 0

	for _, line := range input {
		firstDigit, secondDigit := 0, 0
		firstDigitIndex, secondDigitIndex := -1, -1

		for i, char := range line {
			if char >= '0' && char <= '9' {
				firstDigit = int(char) - '0'
				firstDigitIndex = i
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				secondDigit = int(line[i]) - '0'
				secondDigitIndex = i
				break
			}
		}

		subStringFirstDigit, subStringSecondDigit := "", ""

		if firstDigitIndex < 0 {
			subStringFirstDigit, subStringSecondDigit = line, line
			firstDigitIndex, secondDigitIndex = len(line)-1, 0
		} else {
			subStringFirstDigit, subStringSecondDigit = line[:firstDigitIndex+1], line[secondDigitIndex:]
		}

		globalFirstDigitIndex := firstDigitIndex

		for k, v := range digitMap {
			i := strings.Index(subStringFirstDigit, k)
			if i != -1 && i <= globalFirstDigitIndex {
				firstDigit = v
				globalFirstDigitIndex = i
			}
		}

		globalSecondDigitIndex := 0

		for k, v := range digitMap {
			i := strings.LastIndex(subStringSecondDigit, k)
			if i != -1 && i >= globalSecondDigitIndex {
				secondDigit = v
				globalSecondDigitIndex = i
			}
		}

		calibration := firstDigit*10 + secondDigit
		calibrationCount += calibration
	}

	fmt.Println(calibrationCount)
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day-/input.txt")
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

	// trebuchet([]string{"twosixmpnpdzmjxlmjsjdnkmnhmdtdgxrbknkplsixss"})
}
