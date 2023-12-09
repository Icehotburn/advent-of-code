package main

import (
	"2023/utils"
	"fmt"
	"math"
	"os"
	"strings"
)

type Card struct {
	copies         int
	winningNumbers []int
	numbers        []int
}

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

	cards := getCards(input)

	totalValue := 0

	for _, card := range cards {
		winners := getWinningNumbers(card.winningNumbers, card.numbers)
		cardValue := int(math.Pow(float64(2), float64(len(winners)-1)))

		totalValue += cardValue
	}

	return totalValue
}

func Part2(input string) int {
	cards := getCards(input)

	for pos, card := range cards {
		cardValue := len(getWinningNumbers(card.winningNumbers, card.numbers))

		for i := 1; i <= cardValue; i++ {
			cards[pos+i].copies += card.copies
		}
	}

	totalValue := 0

	for _, card := range cards {
		totalValue += card.copies
	}

	return totalValue
}

func getCards(input string) []Card {
	var cards []Card
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		cards = append(cards, parseLineToCard(line))
	}

	return cards
}

func parseLineToCard(line string) Card {
	split := strings.Split(line, ":")

	game := split[1] // i.e " 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	gameInfo := strings.Split(game, "|")
	winningNumbers := utils.ParseNumbersFromString(gameInfo[0], " ") // i.e [41, 48, 83, 86, 17]
	cardNumbers := utils.ParseNumbersFromString(gameInfo[1], " ")    // i.e [83, 86,  6, 31, 17,  9, 48, 53]"

	return Card{1, winningNumbers, cardNumbers}
}

func getWinningNumbers(winningNumbers []int, cardNumbers []int) []int {
	var winners []int

	for _, num := range cardNumbers {
		if utils.Contains(winningNumbers, num) {
			winners = append(winners, num)
		}
	}

	return winners
}
