package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calorieCounting(input []string) {
	calorieCount := []int{}

	localCalorieCount := 0

	for _, line := range input {
		calorie, _ := strconv.Atoi(line)
		localCalorieCount += calorie

		if line == "" {
			calorieCount = append(calorieCount, localCalorieCount)
			localCalorieCount = 0
		}
	}

	calorieCount = append(calorieCount, localCalorieCount)

	for i := 0; i < 3; i++ {
		largestIndex := i
		largest := calorieCount[i]
		for j := i + 1; j < len(calorieCount); j++ {
			if calorieCount[j] > largest {
				largest = calorieCount[j]
				largestIndex = j
			}
		}

		calorieCount[i], calorieCount[largestIndex] = calorieCount[largestIndex], calorieCount[i]
	}

	highestCalorieCount := calorieCount[0] + calorieCount[1] + calorieCount[2]

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
