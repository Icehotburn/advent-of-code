package main

import (
	"2023/utils"
	"fmt"
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

	// first line is time
	raceDurationsLine := lines[0] // i.e "Time:      7  15   30"
	raceDurations := utils.ParseNumbersFromString(raceDurationsLine[5:], " ")

	// second line is distance
	distancesToBeatLine := lines[1] // i.e "Distance:  9  40  200"
	distancesToBeat := utils.ParseNumbersFromString(distancesToBeatLine[9:], " ")

	boatSpeed := 1

	totalWaysToWin := 0

	/* Potential optimization: this lookup is super slow
		Instead of checking every distance value, we can find the min and max raceDurations that break the record
	    Then we can return the number of values in the range between the min and max raceDurations
		This would greatly speed up the program.
	*/
	for race := 0; race < len(raceDurations); race++ {
		waysToWin := calculateWaysToWin(raceDurations[race], distancesToBeat[race], boatSpeed)

		if waysToWin == 0 {
			continue
		} else if totalWaysToWin == 0 {
			totalWaysToWin = waysToWin
		} else {
			totalWaysToWin = totalWaysToWin * waysToWin
		}
	}

	return totalWaysToWin
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	// first line is raceDuration
	raceDurationLine := lines[0] // i.e "Time:      7  15   30"
	raceDuration := utils.ParseIntOrPanic(strings.ReplaceAll(raceDurationLine[5:], " ", ""))

	// second line is distanceTobeat
	distanceToBeatLine := lines[1] // i.e "Distance:  9  40  200"
	distanceToBeat := utils.ParseIntOrPanic(strings.ReplaceAll(distanceToBeatLine[9:], " ", ""))

	boatSpeed := 1

	return calculateWaysToWin(raceDuration, distanceToBeat, boatSpeed)
}

func calculateWaysToWin(raceDuration int, distanceToBeat int, boatSpeed int) int {
	waysToWin := 0
	for i := 0; i < raceDuration; i++ {
		distanceTraveled := i * boatSpeed * (raceDuration - i)
		if distanceTraveled > distanceToBeat {
			waysToWin += 1
		}
	}

	return waysToWin
}
