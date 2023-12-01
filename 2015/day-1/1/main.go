package main

import (
	"bufio"
	"fmt"
	"os"
)

func notQuiteLisp(input string) {
	floor := 0

	for _, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	fmt.Println(floor)
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
