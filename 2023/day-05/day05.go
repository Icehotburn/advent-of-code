package main

import (
	"2023/utils"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	//Open input file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	// Close Input File
	defer func() {
		err = inputFile.Close()
		if err != nil {
			panic(err)
		}
	}()

	input := utils.ReadInputFile(inputFile)

	// Print the answer
	fmt.Println("Part 1 answer: ", Part1(input))
	fmt.Println("Part 2 answer: ", Part2(input))

}

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	seedsLine := lines[0] // i.e "seeds: 79 14 55 13"
	seedsInput := utils.ParseNumbersFromString(seedsLine[7:], " ")

	// input maps
	var seedToSoil [][]int // i.e [[50, 98, 2]], [52, 50, 48]]
	var soilToFertilizer [][]int
	var fertilizerToWater [][]int
	var waterToLight [][]int
	var lightToTemperature [][]int
	var temperatureToHumidity [][]int
	var humidityToLocation [][]int

	for pos := 1; pos < len(lines)-1; {
		line := lines[pos]
		if strings.HasSuffix(line, "map:") {
			//advance 1 line to get to the map body
			pos += 1
			switch line {
			case "seed-to-soil map:":
				seedToSoil = parseMap(lines[pos:])
			case "soil-to-fertilizer map:":
				soilToFertilizer = parseMap(lines[pos:])
			case "fertilizer-to-water map:":
				fertilizerToWater = parseMap(lines[pos:])
			case "water-to-light map:":
				waterToLight = parseMap(lines[pos:])
			case "light-to-temperature map:":
				lightToTemperature = parseMap(lines[pos:])
			case "temperature-to-humidity map:":
				temperatureToHumidity = parseMap(lines[pos:])
			case "humidity-to-location map:":
				humidityToLocation = parseMap(lines[pos:])
			}
		} else {
			pos += 1
		}

	}

	minLocation := math.MaxInt
	for _, seed := range seedsInput {

		soil := lookupValue(seed, seedToSoil)
		fertilizer := lookupValue(soil, soilToFertilizer)
		water := lookupValue(fertilizer, fertilizerToWater)
		light := lookupValue(water, waterToLight)
		temperature := lookupValue(light, lightToTemperature)
		humidity := lookupValue(temperature, temperatureToHumidity)
		location := lookupValue(humidity, humidityToLocation)

		if location < minLocation {
			minLocation = location
		}
	}

	return minLocation
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	seedsLine := lines[0] // i.e "seeds: 79 14 55 13"
	seedsInput := utils.ParseNumbersFromString(seedsLine[7:], " ")

	// input maps
	var seedToSoil [][]int // i.e [[50, 98, 2]], [52, 50, 48]]
	var soilToFertilizer [][]int
	var fertilizerToWater [][]int
	var waterToLight [][]int
	var lightToTemperature [][]int
	var temperatureToHumidity [][]int
	var humidityToLocation [][]int

	for pos := 1; pos < len(lines)-1; {
		line := lines[pos]
		if strings.HasSuffix(line, "map:") {
			//advance 1 line to get to the map body
			pos += 1
			switch line {
			case "seed-to-soil map:":
				seedToSoil = parseMap(lines[pos:])
			case "soil-to-fertilizer map:":
				soilToFertilizer = parseMap(lines[pos:])
			case "fertilizer-to-water map:":
				fertilizerToWater = parseMap(lines[pos:])
			case "water-to-light map:":
				waterToLight = parseMap(lines[pos:])
			case "light-to-temperature map:":
				lightToTemperature = parseMap(lines[pos:])
			case "temperature-to-humidity map:":
				temperatureToHumidity = parseMap(lines[pos:])
			case "humidity-to-location map:":
				humidityToLocation = parseMap(lines[pos:])
			}
		} else {
			pos += 1
		}

	}

	/* Potential optimization: this lookup is super slow
		Instead of checking every seed value, we should invert the operation so that only check if each location's seed id is present in the Seed Ids.
	    Then we can return the minimum seedId or the minimumLocation... whichever one is less.
		This would greatly speed up the program.
	*/
	minLocation := math.MaxInt
	for i := 0; i < len(seedsInput); i += 2 {
		fmt.Println("Checking seed input:", i/2, "of", len(seedsInput)/2)
		start := seedsInput[i]
		length := seedsInput[i+1]

		for x := 0; x < length; x++ {
			seed := start + x

			soil := lookupValue(seed, seedToSoil)
			fertilizer := lookupValue(soil, soilToFertilizer)
			water := lookupValue(fertilizer, fertilizerToWater)
			light := lookupValue(water, waterToLight)
			temperature := lookupValue(light, lightToTemperature)
			humidity := lookupValue(temperature, temperatureToHumidity)
			location := lookupValue(humidity, humidityToLocation)

			if location < minLocation {
				minLocation = location
			}
		}

	}

	return minLocation
}

func parseMap(lines []string) [][]int {
	var result [][]int
	for pos := 0; lines[pos] != ""; pos++ {
		result = append(result, utils.ParseNumbersFromString(lines[pos], " "))
	}

	return result
}

func lookupValue(id int, valuesMap [][]int) int {
	for _, value := range valuesMap {
		destinationRangeStart := value[0]
		sourceRangeStart := value[1]
		rangeLength := value[2]

		if id >= sourceRangeStart && id < sourceRangeStart+rangeLength {
			offset := id - sourceRangeStart
			return destinationRangeStart + offset
		}
	}

	return id
}
