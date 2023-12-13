package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type SpringInfo struct {
	springList            []byte
	contiguousDamagedList []int
}

func parseInput(input []string) []SpringInfo {
	springInfos := make([]SpringInfo, 0)

	for _, line := range input {
		springInfoFields := strings.Fields(line)
		springString := []byte(springInfoFields[0])
		damagedIntermediate := strings.Split(springInfoFields[1], ",")
		damagedList := make([]int, 0)
		for _, damaged := range damagedIntermediate {
			number, _ := strconv.Atoi(damaged)
			damagedList = append(damagedList, number)
		}
		springInfos = append(springInfos, SpringInfo{
			springString,
			damagedList,
		})
	}

	return springInfos
}

func generateSpringPermutations(springs []byte, unknownSpringPositions []int, length int, acc [][]byte) [][]byte {
	if length == 0 {
		springCopy := append([]byte(nil), springs...)
		acc = append(acc, springCopy)
		return acc
	}

	springs[unknownSpringPositions[length-1]] = '.'
	acc = generateSpringPermutations(springs, unknownSpringPositions, length-1, acc)
	springs[unknownSpringPositions[length-1]] = '#'
	acc = generateSpringPermutations(springs, unknownSpringPositions, length-1, acc)

	return acc
}

func calculateContiguousDamagedList(springs []byte) []int {
	contiguousDamagedList := make([]int, 0)

	damagedCount := 0
	for _, spring := range springs {
		if spring == '#' {
			damagedCount++
		} else if spring == '.' && damagedCount > 0 {
			contiguousDamagedList = append(contiguousDamagedList, damagedCount)
			damagedCount = 0
		}
	}

	if damagedCount > 0 {
		contiguousDamagedList = append(contiguousDamagedList, damagedCount)
	}

	return contiguousDamagedList
}

func IntArrayEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func calculateUnknownSpringPositions(springs []byte) []int {
	unknownPositions := make([]int, 0)
	for i, spring := range springs {
		if spring == '?' {
			unknownPositions = append(unknownPositions, i)
		}
	}

	return unknownPositions
}

func IntArraySum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func countDamagedSprings(springs []byte) int {
	count := 0
	for _, spring := range springs {
		if spring == '#' {
			count++
		}
	}
	return count
}

func HotSprings(input []string) {
	startTime := time.Now()

	springInfos := parseInput(input)

	totalPermutations := 0

	for _, springInfo := range springInfos {
		validPermutations := 0

		unknownSpringPositions := calculateUnknownSpringPositions(springInfo.springList)
		springListCopy := append([]byte(nil), springInfo.springList...)

		springPermutations := generateSpringPermutations(
			springListCopy,
			unknownSpringPositions,
			len(unknownSpringPositions),
			[][]byte{},
		)

		totalDamagedSprings := IntArraySum(springInfo.contiguousDamagedList)

		for _, springPermutation := range springPermutations {
			if totalDamagedSprings == countDamagedSprings(springPermutation) &&
				IntArrayEquals(
					calculateContiguousDamagedList(springPermutation),
					springInfo.contiguousDamagedList,
				) {
				validPermutations++
			}
		}
		fmt.Println(string(springInfo.springList), validPermutations)
		totalPermutations += validPermutations
	}

	fmt.Println("Total permutations: ", totalPermutations)

	fmt.Println("Time elapsed: ", time.Since(startTime))
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day12/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	HotSprings(input)
}
