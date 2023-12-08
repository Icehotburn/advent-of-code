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

	// Create the bag contents map
	bagContents := InitializeBagContents(12, 13, 14)

	// Print the answer
	fmt.Println("Part 1 answer: ", Part1(bagContents, input))
	fmt.Println("Part 2 answer: ", Part2(input))

}

func Part1(bagContents map[string]int, input string) int {
	// Sum of all the valid game id values
	gameIdSum := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {

		if line == "" {
			continue
		}

		split := strings.Split(line, ":")

		// get the game id
		gameInfo := split[0]                          // i.e "Game 1"
		gameId := utils.ParseIntOrPanic(gameInfo[5:]) // i.e "1"

		game := split[1] // i.e " 18 red, 8 green, 7 blue; 15 red, 4 blue, 1 green; 2 green, 17 red, 6 blue; 5 green, 1 blue, 11 red; 18 red, 1 green, 14 blue; 8 blue"

		if isValidGame(game, bagContents) {
			gameIdSum += gameId
		}
	}

	return gameIdSum
}

func InitializeBagContents(red int, green int, blue int) map[string]int {
	var bagContents map[string]int
	bagContents = make(map[string]int)
	bagContents["red"] = red
	bagContents["green"] = green
	bagContents["blue"] = blue
	return bagContents
}

func isValidGame(game string, contents map[string]int) bool {
	// get the turns
	turns := strings.Split(game, ";") // i.e [" 18 red, 8 green, 7 blue"]

	// for each turn
	for _, turn := range turns {
		// check that the count/color value is less than the input map
		values := strings.Split(turn, ",") // i.e "[" 18 red", " 8 green", " 7 blue"]
		for _, value := range values {
			parsed := strings.Split(value[1:], " ")      // i.e ["18", "red"]
			color := parsed[1]                           // i.e "red"
			quantity := utils.ParseIntOrPanic(parsed[0]) // i.e "18"

			if quantity > contents[color] {
				return false
			}
		}
	}
	return true
}

func Part2(input string) int {
	// Sum of all the minimum cubes power values
	powerSum := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {

		if line == "" {
			continue
		}

		split := strings.Split(line, ":")

		game := split[1] // i.e " 18 red, 8 green, 7 blue; 15 red, 4 blue, 1 green; 2 green, 17 red, 6 blue; 5 green, 1 blue, 11 red; 18 red, 1 green, 14 blue; 8 blue"

		minCubes := minCubes(game)

		minCubesPower := 1
		for _, value := range minCubes {
			minCubesPower = minCubesPower * value
		}

		powerSum += minCubesPower
	}
	return powerSum
}

func minCubes(game string) map[string]int {
	var result map[string]int
	result = make(map[string]int)

	// get the turns
	turns := strings.Split(game, ";") // i.e [" 18 red, 8 green, 7 blue"]

	// for each turn
	for _, turn := range turns {
		values := strings.Split(turn, ",") // i.e "[" 18 red", " 8 green", " 7 blue"]
		for _, value := range values {
			parsed := strings.Split(value[1:], " ")      // i.e ["18", "red"]
			color := parsed[1]                           // i.e "red"
			quantity := utils.ParseIntOrPanic(parsed[0]) // i.e "18"

			if quantity > result[color] {
				result[color] = quantity
			}
		}
	}
	return result
}
