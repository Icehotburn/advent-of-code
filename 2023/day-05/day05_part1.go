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

	//first line is seeds
	sc.Scan()
	seedsLine := sc.Text() // i.e "seeds: 79 14 55 13"
	seedsInput := parseNumbers(seedsLine[7:])

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
				seedToSoil = parseMap(sc)
			case "soil-to-fertilizer map:":
				soilToFertilizer = parseMap(sc)
			case "fertilizer-to-water map:":
				fertilizerToWater = parseMap(sc)
			case "water-to-light map:":
				waterToLight = parseMap(sc)
			case "light-to-temperature map:":
				lightToTemperature = parseMap(sc)
			case "temperature-to-humidity map:":
				temperatureToHumidity = parseMap(sc)
			case "humidity-to-location map:":
				humidityToLocation = parseMap(sc)
			}
		}
	}

	min := math.MaxInt
	for _, seed := range seedsInput {

		soil := lookupValue(seed, seedToSoil)
		fertilizer := lookupValue(soil, soilToFertilizer)
		water := lookupValue(fertilizer, fertilizerToWater)
		light := lookupValue(water, waterToLight)
		temperature := lookupValue(light, lightToTemperature)
		humidity := lookupValue(temperature, temperatureToHumidity)
		location := lookupValue(humidity, humidityToLocation)

		if location < min {
			min = location
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

func parseNumbers(s string) []int {
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

func parseMap(sc *bufio.Scanner) [][]int {
	var result [][]int
	mapLine := sc.Text()
	for mapLine != "" {
		result = append(result, parseNumbers(mapLine))
		sc.Scan()
		mapLine = sc.Text()
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
