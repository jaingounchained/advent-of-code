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

type Lens struct {
	label       string
	focalLength int
	next        *Lens
}

func deleteLens(lens *Lens, label string) *Lens {
	if lens == nil {
		return nil
	}

	if lens.label == label {
		return lens.next
	}

	lens.next = deleteLens(lens.next, label)
	return lens
}

func upsertLens(lens *Lens, label string, focalLength int) *Lens {
	if lens == nil {
		return &Lens{
			label,
			focalLength,
			nil,
		}
	}

	if lens.label == label {
		lens.focalLength = focalLength
		return lens
	}

	lens.next = upsertLens(lens.next, label, focalLength)
	return lens
}

func printBox(lens *Lens) {
	if lens == nil {
		return
	}

	fmt.Println(lens.label, lens.focalLength)
	printBox(lens.next)
}

func calculateFocusingPower(lens *Lens, boxNo, slot, acc int) int {
	if lens == nil {
		return acc
	}

	acc += boxNo * slot * lens.focalLength

	return calculateFocusingPower(lens.next, boxNo, slot+1, acc)
}

func ParabolicReflectorDish(input []string) {
	startTime := time.Now()

	steps := parseInput(input)

	boxes := make([]*Lens, 256)

	for _, step := range steps {
		if step[len(step)-1] == '-' {
			label := step[:len(step)-1]
			boxNo := calculateHash(label)

			boxes[boxNo] = deleteLens(boxes[boxNo], label)
		} else {
			label, focalLength := step[:len(step)-2], int(step[len(step)-1]-'0')
			boxNo := calculateHash(label)

			boxes[boxNo] = upsertLens(boxes[boxNo], label, focalLength)
		}
	}

	focusingPowers := 0

	for i, box := range boxes {
		if box != nil {
			focusingPowers += calculateFocusingPower(box, i+1, 1, 0)
		}
	}

	fmt.Println(focusingPowers)

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
