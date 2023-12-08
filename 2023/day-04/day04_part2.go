package main

import (
	"2023/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Card struct {
	copies int
	value  string
}

func main() {
	//Read input file
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(input)

	// Sum of all the winning card values
	totalValue := 0

	var cards []Card

	for sc.Scan() {
		line := sc.Text()

		cards = append(cards, Card{1, line})
	}

	for pos, card := range cards {
		split := strings.Split(card.value, ":")

		game := split[1] // i.e " 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
		gameInfo := strings.Split(game, "|")
		winningNumbers := parseNumbersToArray(gameInfo[0]) // i.e [41, 48, 83, 86, 17]
		cardNumbers := parseNumbersToArray(gameInfo[1])    // i.e [83, 86,  6, 31, 17,  9, 48, 53]"

		cardValue := calculateValue(winningNumbers, cardNumbers)

		for i := 1; i <= cardValue; i++ {
			cards[pos+i].copies += card.copies
		}
	}

	for _, card := range cards {
		totalValue += card.copies
	}

	// sum the total copies of all the cards

	// Print the answer
	fmt.Println(totalValue)

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
			result = append(result, utils.ParseIntOrPanic(value))
		}
	}

	return result
}

func calculateValue(winningNumbers []int, cardNumbers []int) int {
	var winners []int

	for _, num := range cardNumbers {
		if containsInt(winningNumbers, num) {
			winners = append(winners, num)
		}
	}

	return len(winners)
}

func containsInt(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
