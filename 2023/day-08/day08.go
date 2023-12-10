package main

import (
	"2023/utils"
	"fmt"
	"os"
	"strings"
)

type Location struct {
	left  string
	right string
}

func main() {
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
	directions := lines[0] // i.e LLR

	locationsMap := initLocationMap(lines[1:])

	return distanceFromTerminalLocation("ZZZ", directions, "AAA", locationsMap)
}

func distanceFromTerminalLocation(terminalSuffix string, directions string, currentLocation string, locationsMap map[string]Location) int {
	steps := 0

	for !strings.HasSuffix(currentLocation, terminalSuffix) {
		for _, direction := range directions {
			steps += 1
			if string(direction) == "R" {
				currentLocation = locationsMap[currentLocation].right
			} else { // direction == "L"
				currentLocation = locationsMap[currentLocation].left
			}

			if strings.HasSuffix(currentLocation, terminalSuffix) {
				return steps
			}
		}
	}

	return steps
}

// Part2 Note: this solution assumes that all starting locations are cycles with a total length less than or equal to the length of directions
func Part2(input string) int {
	lines := strings.Split(input, "\n")
	directions := lines[0] // i.e LLR

	locationsMap := initLocationMap(lines[1:])
	var startLocationDistancesFromTerminalLocations []int

	for location := range locationsMap {
		if strings.HasSuffix(location, "A") {
			startLocationDistancesFromTerminalLocations = append(startLocationDistancesFromTerminalLocations, distanceFromTerminalLocation("Z", directions, location, locationsMap))
		}
	}

	return leastCommonMultiple(startLocationDistancesFromTerminalLocations)
}

func greatestCommonDenominator(a int, b int) int {
	if b == 0 {
		return a
	}

	return greatestCommonDenominator(b, a%b)

}

func leastCommonMultiple(nums []int) int {
	result := 1
	for i := 0; i < len(nums); i += 1 {
		result = (nums[i] * result) / greatestCommonDenominator(nums[i], result)
	}

	return result
}

func initLocationMap(lines []string) map[string]Location {
	locationsMap := make(map[string]Location)

	for _, line := range lines {
		if line == "" {
			continue
		}

		source, destinations := parseLine(line)

		locationsMap[source] = destinations
	}

	return locationsMap
}

func parseLine(line string) (string, Location) {
	//line is in format AAA = (BBB, CCC)
	splits := strings.Split(line, "=")
	leftSide := strings.TrimSpace(splits[0]) // i.e AAA
	rightSide := splits[1]                   // i.e (BBB, CCC)
	rightSide = strings.ReplaceAll(rightSide, "(", "")
	rightSide = strings.ReplaceAll(rightSide, ")", "")

	// rightSide is now BBB, CCC
	splitRightSide := strings.Split(rightSide, ",")
	left := strings.TrimSpace(splitRightSide[0])
	right := strings.TrimSpace(splitRightSide[1])

	return leftSide, Location{left: left, right: right}
}
