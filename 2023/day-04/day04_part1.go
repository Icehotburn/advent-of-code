package main

import (
	"2023/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	//Read input file
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(input)

	// Sum of all the winning card values
	totalValue := 0

	for sc.Scan() {
		line := sc.Text()
		split := strings.Split(line, ":")

		game := split[1] // i.e " 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
		gameInfo := strings.Split(game, "|")
		winningNumbers := parseNumbers(gameInfo[0]) // i.e [41, 48, 83, 86, 17]
		cardNumbers := parseNumbers(gameInfo[1])    // i.e [83, 86,  6, 31, 17,  9, 48, 53]"

		cardValue := calculateCardValue(winningNumbers, cardNumbers)

		totalValue += cardValue
	}

	// Print the answer
	fmt.Println(totalValue)

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
			result = append(result, utils.ParseIntOrPanic(value))
		}
	}

	return result
}

func calculateCardValue(winningNumbers []int, cardNumbers []int) int {
	var winners []int

	for _, num := range cardNumbers {
		if contains(winningNumbers, num) {
			winners = append(winners, num)
		}
	}

	if len(winners) == 0 {
		return 0
	} else {
		return int(math.Pow(float64(2), float64(len(winners)-1)))
	}

}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
