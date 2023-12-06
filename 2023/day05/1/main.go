package main

import (
	"bufio"
	"fmt"
	"math"
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
	sourceRangeStart      int
	destinationRangeStart int
	rangeLength           int
}

func parseS2DMap(input []string) (s2dMap []S2DMap) {
	s2dMap = make([]S2DMap, 0)

	for _, line := range input {
		values := strings.Fields(line)
		sourceRangeStart, _ := strconv.Atoi(values[1])
		destinationRangeStart, _ := strconv.Atoi(values[0])
		rangeLength, _ := strconv.Atoi(values[2])
		s2dMap = append(s2dMap, S2DMap{
			sourceRangeStart,
			destinationRangeStart,
			rangeLength,
		})
	}
	return
}

func findEmptyLines(input []string) int {
	for i, line := range input {
		if line == "" {
			return i
		}
	}

	return -1
}

func parseInput(input []string) ([]int, []S2DMap, []S2DMap, []S2DMap, []S2DMap, []S2DMap, []S2DMap, []S2DMap) {
	seedsString := strings.Split(input[0], ":")[1]
	seedsIntermediate := strings.Fields(seedsString)
	seeds := make([]int, len(seedsIntermediate))
	for i, seed := range seedsIntermediate {
		seeds[i], _ = strconv.Atoi(seed)
	}

	input = input[3:]

	emptyLine := findEmptyLines(input)
	seedToSoilMap := parseS2DMap(input[:emptyLine])
	input = input[emptyLine+2:]

	emptyLine = findEmptyLines(input)
	soilToFertilizerMap := parseS2DMap(input[:emptyLine])
	input = input[emptyLine+2:]

	emptyLine = findEmptyLines(input)
	fertilizerToWaterMap := parseS2DMap(input[:emptyLine])
	input = input[emptyLine+2:]

	emptyLine = findEmptyLines(input)
	waterToLightMap := parseS2DMap(input[:emptyLine])
	input = input[emptyLine+2:]

	emptyLine = findEmptyLines(input)
	lightToTemperatureMap := parseS2DMap(input[:emptyLine])
	input = input[emptyLine+2:]

	emptyLine = findEmptyLines(input)
	temperatureToHumidityMap := parseS2DMap(input[:emptyLine])
	input = input[emptyLine+2:]

	humidityToLocationMap := parseS2DMap(input)

	return seeds, seedToSoilMap, soilToFertilizerMap, fertilizerToWaterMap, waterToLightMap, lightToTemperatureMap, temperatureToHumidityMap, humidityToLocationMap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findDestination(seed int, s2dmap []S2DMap) int {
	for _, mapValue := range s2dmap {
		if seed >= mapValue.sourceRangeStart && seed < (mapValue.sourceRangeStart+mapValue.rangeLength) {
			return (seed - mapValue.sourceRangeStart) + mapValue.destinationRangeStart
		}
	}

	return seed
}

func youGiveASeedAFertilizer(input []string) {
	startTime := time.Now()

	seeds, seedToSoilMap, soilToFertilizerMap, fertilizerToWaterMap, waterToLightMap, lightToTemperatureMap, temperatureToHumidityMap, humidityToLocationMap := parseInput(input)

	minimumLocation := math.MaxInt

	for _, seed := range seeds {
		soil := findDestination(seed, seedToSoilMap)
		fertilizer := findDestination(soil, soilToFertilizerMap)
		water := findDestination(fertilizer, fertilizerToWaterMap)
		light := findDestination(water, waterToLightMap)
		temperature := findDestination(light, lightToTemperatureMap)
		humidity := findDestination(temperature, temperatureToHumidityMap)
		location := findDestination(humidity, humidityToLocationMap)

		fmt.Println(seed, soil, fertilizer, water, light, temperature, humidity, location)

		minimumLocation = min(minimumLocation, location)
	}

	fmt.Println(minimumLocation)
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
