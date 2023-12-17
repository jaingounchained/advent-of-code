package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func parseInput(input []string) []string {
	steps := strings.Split(input[0], ",")

	return steps
}

func calculateHash(step string) int {
	hash := 0
	for _, c := range step {
		hash += int(c)
		hash = (hash * 17) % 256
	}
	return hash
}

func ParabolicReflectorDish(input []string) {
	startTime := time.Now()

	steps := parseInput(input)

	hashSum := 0
	for _, step := range steps {
		hashSum += calculateHash(step)
	}

	fmt.Println(hashSum)

	fmt.Println("Time elapsed: ", time.Since(startTime))
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day15/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	ParabolicReflectorDish(input)
}
