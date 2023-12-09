package main

import (
	"bufio"
	"fmt"
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
	r, _ := regexp.Compile("^([A-Z]{3}) = \\(([A-Z]{3}), ([A-Z]{3})\\)$")
	for _, line := range input {
		matches := r.FindStringSubmatch(line)
		nodeMap[matches[1]] = Location{
			matches[2],
			matches[3],
		}
	}

	return directions, nodeMap
}

func hauntedWasteland(input []string) {
	startTime := time.Now()

	directions, nodeMap := parseInput(input)

	currentNode, endingNode := "AAA", "ZZZ"
	steps := 0
	for i := 0; currentNode != endingNode; i++ {
		i %= (len(directions))

		switch directions[i] {
		case 'L':
			currentNode = nodeMap[currentNode].left
		case 'R':
			currentNode = nodeMap[currentNode].right
		}
		steps++
		if currentNode == endingNode {
			break
		}
	}

	fmt.Println(steps)
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
