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

	// Sum of all the minimum cubes power values
	powerSum := 0

	for sc.Scan() {
		line := sc.Text()
		split := strings.Split(line, ":")

		game := split[1] // i.e " 18 red, 8 green, 7 blue; 15 red, 4 blue, 1 green; 2 green, 17 red, 6 blue; 5 green, 1 blue, 11 red; 18 red, 1 green, 14 blue; 8 blue"

		minCubes := minCubes(game)

		minCubesPower := 1
		for _, value := range minCubes {
			minCubesPower = minCubesPower * value
		}

		powerSum += minCubesPower
	}

	// Print the answer
	fmt.Println(powerSum)

	// Close Input
	err = input.Close()
	if err != nil {
		panic(err)
	}

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
			parsed := strings.Split(value[1:], " ")  // i.e ["18", "red"]
			color := parsed[1]                       // i.e "red"
			quantity, err := strconv.Atoi(parsed[0]) // i.e "18"
			if err != nil {
				panic(err)
			}

			if quantity > result[color] {
				result[color] = quantity
			}
		}
	}
	return result
}
