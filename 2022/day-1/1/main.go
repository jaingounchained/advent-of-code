package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calorieCounting(input []string) {
	localCalorieCount, highestCalorieCount := 0, 0

	for _, line := range input {
		calorie, _ := strconv.Atoi(line)
		localCalorieCount += calorie

		if line == "" {
			fmt.Println(localCalorieCount, highestCalorieCount)
			if localCalorieCount > highestCalorieCount {
				highestCalorieCount = localCalorieCount
			}
			localCalorieCount = 0
		}
	}
	if localCalorieCount > highestCalorieCount {
		highestCalorieCount = localCalorieCount
	}

	fmt.Println(highestCalorieCount)
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2022/day-1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	calorieCounting(input)
}
