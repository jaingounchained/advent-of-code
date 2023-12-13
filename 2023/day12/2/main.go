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
	springRow []byte
	targets   []int
}

var springPermutationsCache map[string]int

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

func unFoldSprings(springInfo SpringInfo) SpringInfo {
	unfoldedSpring := SpringInfo{
		springInfo.springRow,
		springInfo.targets,
	}

	for i := 0; i < 4; i++ {
		unfoldedSpring.springRow = append(unfoldedSpring.springRow, '?')
		unfoldedSpring.springRow = append(unfoldedSpring.springRow, springInfo.springRow...)
		unfoldedSpring.targets = append(unfoldedSpring.targets, springInfo.targets...)
	}

	return unfoldedSpring
}

func IntSliceToString(numbers []int) string {
	str := ""
	for _, n := range numbers {
		str += string(n)
	}
	return str
}

func countSpringPermutations(springRow []byte, targets []int) int {
	// Base cases
	if len(springRow) == 0 {
		if len(targets) == 0 {
			return 1
		}
		return 0
	}

	if len(targets) == 0 {
		for _, spring := range springRow {
			if spring == '#' {
				return 0
			}
		}
		return 1
	}

	cacheKey := string(springRow) + IntSliceToString(targets)

	if v, ok := springPermutationsCache[cacheKey]; ok {
		return v
	}

	count := 0

	// Recursion
	if springRow[0] == '.' || springRow[0] == '?' {
		count += countSpringPermutations(springRow[1:], targets)
	}

	if springRow[0] == '#' || springRow[0] == '?' {
		if targets[0] > len(springRow) {
			return count
		}

		for i := 0; i < targets[0]; i++ {
			if springRow[i] == '.' {
				return count
			}
		}

		if targets[0] == len(springRow) || springRow[targets[0]] != '#' {
			if len(springRow) > targets[0] {
				count += countSpringPermutations(springRow[targets[0]+1:], targets[1:])
			} else {
				count += countSpringPermutations([]byte{}, targets[1:])
			}
		}
	}

	springPermutationsCache[cacheKey] = count

	return count
}

func HotSprings(input []string) {
	startTime := time.Now()

	springInfos := parseInput(input)

	springPermutationsCache = make(map[string]int)

	totalPermutations := 0

	for _, springInfo := range springInfos {
		unfoldedSpringInfo := unFoldSprings(springInfo)

		validPermutations := countSpringPermutations(unfoldedSpringInfo.springRow, unfoldedSpringInfo.targets)

		fmt.Println(string(unfoldedSpringInfo.springRow), validPermutations)
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
