package main

import (
	"bufio"
	"fmt"
	"os"
)

func notQuiteLisp(input string) {
	floor := 0
	basementPosition := 0
	enteredBasementFirstTime := false

	for i, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 && !enteredBasementFirstTime {
			basementPosition = i + 1
			enteredBasementFirstTime = true
		}
	}

	fmt.Println(floor, basementPosition)
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2015/day-1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = scanner.Text()
	}

	notQuiteLisp(input)
}
