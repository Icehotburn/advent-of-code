package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Read input file
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(input)

	// array of the inputted seeds

	//first line is seeds
	sc.Scan()
	seedsLine := sc.Text() // i.e "seeds: 79 14 55 13"
	seedsInput := parseNumbersToArray(seedsLine[7:])

	// input maps
	var seedToSoil [][]int // i.e [[50, 98, 2]], [52, 50, 48]]
	var soilToFertilizer [][]int
	var fertilizerToWater [][]int
	var waterToLight [][]int
	var lightToTemperature [][]int
	var temperatureToHumidity [][]int
	var humidityToLocation [][]int

	for sc.Scan() {
		line := sc.Text()

		if strings.HasSuffix(line, "map:") {
			//advance 1 line to get to the map body
			sc.Scan()
			switch line {
			case "seed-to-soil map:":
				seedToSoil = parseStringToMap(sc)
			case "soil-to-fertilizer map:":
				soilToFertilizer = parseStringToMap(sc)
			case "fertilizer-to-water map:":
				fertilizerToWater = parseStringToMap(sc)
			case "water-to-light map:":
				waterToLight = parseStringToMap(sc)
			case "light-to-temperature map:":
				lightToTemperature = parseStringToMap(sc)
			case "temperature-to-humidity map:":
				temperatureToHumidity = parseStringToMap(sc)
			case "humidity-to-location map:":
				humidityToLocation = parseStringToMap(sc)
			}
		}
	}

	/* Potential optimization: this lookup is super slow
		Instead of checking every seed value, we should invert the operation so that only check if each location's seed id is present in the Seed Ids.
	    Then we can return the minimum seedId or the minimumLocation... whichever one is less.
		This would greatly speed up the program.
	*/
	min := math.MaxInt
	for i := 0; i < len(seedsInput); i += 2 {
		fmt.Println("Checking seed input: ", i/2)
		start := seedsInput[i]
		length := seedsInput[i+1]

		for x := 0; x < length; x++ {
			seed := start + x

			soil := lookupValueFromMap(seed, seedToSoil)
			fertilizer := lookupValueFromMap(soil, soilToFertilizer)
			water := lookupValueFromMap(fertilizer, fertilizerToWater)
			light := lookupValueFromMap(water, waterToLight)
			temperature := lookupValueFromMap(light, lightToTemperature)
			humidity := lookupValueFromMap(temperature, temperatureToHumidity)
			location := lookupValueFromMap(humidity, humidityToLocation)

			if location < min {
				min = location
			}
		}

	}

	// Print the answer
	fmt.Println(min)

	// Close Input
	err = input.Close()
	if err != nil {
		panic(err)
	}

}

func parseNumbersToArray(s string) []int {
	s = strings.TrimSpace(s)
	values := strings.Split(s, " ")

	var result []int

	for _, value := range values {
		if value != "" {
			num, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			result = append(result, num)
		}
	}

	return result
}

func parseStringToMap(sc *bufio.Scanner) [][]int {
	var result [][]int
	mapLine := sc.Text()
	for mapLine != "" {
		result = append(result, parseNumbersToArray(mapLine))
		sc.Scan()
		mapLine = sc.Text()
	}
	return result
}

func lookupValueFromMap(id int, valuesMap [][]int) int {
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
