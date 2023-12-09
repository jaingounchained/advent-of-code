package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"time"
)

type Location struct {
	left  string
	right string
}

type NodeMap map[string]Location

func parseInput(input []string) ([]byte, NodeMap) {
	directions := []byte(input[0])

	input = input[2:]

	nodeMap := make(NodeMap)
	r, _ := regexp.Compile("^(.{3}) = \\((.{3}), (.{3})\\)$")
	for _, line := range input {
		matches := r.FindStringSubmatch(line)
		nodeMap[matches[1]] = Location{
			matches[2],
			matches[3],
		}
	}

	return directions, nodeMap
}

var primes []int

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func calculateNextPrime() {
	if len(primes) == 0 {
		primes = append(primes, 2)
		return
	}

	for i := primes[len(primes)-1] + 1; ; i++ {
		for j := 0; j < len(primes); j++ {
			if i%primes[j] == 0 {
				break
			}
		}

		if isPrime(i) {
			primes = append(primes, i)
			return
		}
	}
}

func calculatePrimeFactorWeights(number int) map[int]int {
	primeFactorWeights := make(map[int]int)

	for i := 0; number != 1; i++ {
		if len(primes) <= i {
			calculateNextPrime()
		}

		if number%primes[i] == 0 {
			for j := 0; ; j++ {
				if number%primes[i] != 0 {
					break
				}
				primeFactorWeights[primes[i]]++
				number /= primes[i]
			}
		}
	}

	return primeFactorWeights
}

func lowestCommonDenominator(numbers ...int) int {
	weights := make([]map[int]int, 0)
	for _, number := range numbers {
		weights = append(weights, calculatePrimeFactorWeights(number))
	}

	lcmMap := make(map[int]int)
	for _, weight := range weights {
		for k, v := range weight {
			if _, ok := lcmMap[k]; !ok {
				lcmMap[k] = v
				continue
			}

			if lcmMap[k] < v {
				lcmMap[k] = v
			}
		}
	}

	lcm := 1
	for k, v := range lcmMap {
		lcm *= int(math.Pow(float64(k), float64(v)))
	}

	return lcm
}

func hauntedWasteland(input []string) {
	startTime := time.Now()

	primes = make([]int, 0)

	directions, nodeMap := parseInput(input)

	startingNodes := make([]string, 0)
	for k := range nodeMap {
		if k[2] == 'A' {
			startingNodes = append(startingNodes, k)
		}
	}

	stepsArray := make([]int, 0)

	for _, startingNode := range startingNodes {
		steps := 0

		nextNode := startingNode
		for i := 0; ; i++ {
			i %= (len(directions))

			switch directions[i] {
			case 'L':
				nextNode = nodeMap[nextNode].left
			case 'R':
				nextNode = nodeMap[nextNode].right
			}

			steps++

			if nextNode[2] == 'Z' {
				break
			}
		}

		stepsArray = append(stepsArray, steps)
	}

	fmt.Println(stepsArray)
	fmt.Println(lowestCommonDenominator(stepsArray...))

	fmt.Println("Time elapsed: ", time.Since(startTime))
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day08/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	hauntedWasteland(input)
}
