package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseInput(input []string) [][]int {
	sequences := make([][]int, 0)

	for _, line := range input {
		sequence := make([]int, 0)
		fields := strings.Fields(line)
		for _, n := range fields {
			number, _ := strconv.Atoi(n)
			sequence = append(sequence, number)
		}
		sequences = append(sequences, sequence)
	}

	return sequences
}

func extrapolateBackward(numbers []int) int {
	allZeros := true
	for _, n := range numbers {
		if n != 0 {
			allZeros = false
		}
	}
	if allZeros {
		return 0
	}

	nextSequence := make([]int, 0)
	for i := 0; i < len(numbers)-1; i++ {
		nextSequence = append(nextSequence, numbers[i+1]-numbers[i])
	}

	return numbers[0] - extrapolateBackward(nextSequence)
}

func MirageMaintainance(input []string) {
	startTime := time.Now()

	valueSum := 0

	sequences := parseInput(input)
	for _, sequence := range sequences {
		nextValue := extrapolateBackward(sequence)
		valueSum += nextValue
		fmt.Println(nextValue)
	}

	fmt.Println(valueSum)

	fmt.Println("Time elapsed: ", time.Since(startTime))
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day09/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	MirageMaintainance(input)
}
