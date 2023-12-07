package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func printStringList(input []string) {
	for _, line := range input {
		fmt.Println(line)
	}
	fmt.Println()
}

type S2DMap struct {
	destinationRangeStart int
	sourceRange           Range
}

type Range struct {
	start  int
	length int
}

func parseInput(input []string) ([]Range, [][]S2DMap) {
	seedsString := strings.Split(input[0], ":")[1]
	seedsIntermediate := strings.Fields(seedsString)
	seedRanges := make([]Range, len(seedsIntermediate)/2)
	for i := 0; i < len(seedRanges); i++ {
		seedsStart, _ := strconv.Atoi(seedsIntermediate[i*2])
		rangeLength, _ := strconv.Atoi(seedsIntermediate[i*2+1])
		seedRanges[i] = Range{
			seedsStart,
			rangeLength,
		}
	}

	input = input[3:]
	input = append(input, "")
	s2dMapList := make([][]S2DMap, 0)
	s2dMap := make([]S2DMap, 0)
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			s2dMapList = append(s2dMapList, s2dMap)
			s2dMap = make([]S2DMap, 0)
			i++
			continue
		}

		values := strings.Fields(input[i])
		sourceRangeStart, _ := strconv.Atoi(values[1])
		destinationRangeStart, _ := strconv.Atoi(values[0])
		rangeLength, _ := strconv.Atoi(values[2])
		s2dMap = append(s2dMap, S2DMap{
			destinationRangeStart,
			Range{
				sourceRangeStart,
				rangeLength,
			},
		})
	}

	return seedRanges, s2dMapList
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func minRange(ranges ...Range) Range {
	if len(ranges) == 0 {
		return Range{}
	}

	minRange := ranges[0]
	for _, r := range ranges {
		if r.start < minRange.start {
			minRange = r
		}
	}

	return minRange
}

func youGiveASeedAFertilizer(input []string) {
	startTime := time.Now()

	seedRanges, s2dMapList := parseInput(input)

	sourceRanges := seedRanges
	for _, s2dMap := range s2dMapList {
		destinationRanges := make([]Range, 0)

		for len(sourceRanges) > 0 {
			seedRange := sourceRanges[len(sourceRanges)-1]
			sourceRanges = sourceRanges[:len(sourceRanges)-1]

			flag := false
			for _, s2d := range s2dMap {
				divideStart := Max(seedRange.start, s2d.sourceRange.start)
				divideEnd := Min(seedRange.start+seedRange.length, s2d.sourceRange.start+s2d.sourceRange.length)

				if divideStart < divideEnd {
					destinationRanges = append(destinationRanges, Range{
						divideStart - s2d.sourceRange.start + s2d.destinationRangeStart,
						divideEnd - divideStart,
					})

					if divideStart > seedRange.start {
						sourceRanges = append(sourceRanges, Range{
							seedRange.start,
							divideStart - seedRange.start,
						})
					}

					if divideEnd < seedRange.start+seedRange.length {
						sourceRanges = append(sourceRanges, Range{
							divideEnd,
							seedRange.start + seedRange.length - divideEnd,
						})
					}
					flag = true
					break
				}
			}
			if flag == false {
				destinationRanges = append(destinationRanges, Range{
					seedRange.start,
					seedRange.length,
				})
			}
		}
		sourceRanges = destinationRanges
	}

	fmt.Println(minRange(sourceRanges...).start)

	fmt.Println("Time elapsed: ", time.Since(startTime))
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day05/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	youGiveASeedAFertilizer(input)
}
