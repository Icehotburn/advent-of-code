package main

import (
	"bufio"
	"fmt"
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

	// Create the bag contents map
	var contents map[string]int
	contents = make(map[string]int)
	contents["red"] = 12
	contents["green"] = 13
	contents["blue"] = 14

	// Sum of all the valid game id values
	gameIdSum := 0

	for sc.Scan() {
		line := sc.Text()
		split := strings.Split(line, ":")

		// get the game id
		gameInfo := split[0]                      // i.e "Game 1"
		gameId, err := strconv.Atoi(gameInfo[5:]) // i.e "1"
		if err != nil {
			panic(err)
		}

		game := split[1] // i.e " 18 red, 8 green, 7 blue; 15 red, 4 blue, 1 green; 2 green, 17 red, 6 blue; 5 green, 1 blue, 11 red; 18 red, 1 green, 14 blue; 8 blue"

		if isValidGame(game, contents) {
			gameIdSum += gameId
		}
	}

	// Print the answer
	fmt.Println(gameIdSum)

	// Close Input
	err = input.Close()
	if err != nil {
		panic(err)
	}

}

func isValidGame(game string, contents map[string]int) bool {
	// get the turns
	turns := strings.Split(game, ";") // i.e [" 18 red, 8 green, 7 blue"]

	// for each turn
	for _, turn := range turns {
		// check that the count/color value is less than the input map
		values := strings.Split(turn, ",") // i.e "[" 18 red", " 8 green", " 7 blue"]
		for _, value := range values {
			parsed := strings.Split(value[1:], " ")  // i.e ["18", "red"]
			color := parsed[1]                       // i.e "red"
			quantity, err := strconv.Atoi(parsed[0]) // i.e "18"
			if err != nil {
				panic(err)
			}

			if quantity > contents[color] {
				return false
			}
		}
	}
	return true
}
