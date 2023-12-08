package main

import (
	"2023/utils"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Hand struct {
	cards string
	wager int
	score int
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
	cardValuesMap := initializeCardValuesMap(false)
	hands := loadHands(input, false)
	sortHands(hands, cardValuesMap)
	return calculateWinnings(hands)
}

func Part2(input string) int {
	cardValuesMap := initializeCardValuesMap(true)
	hands := loadHands(input, true)
	sortHands(hands, cardValuesMap)
	return calculateWinnings(hands)
}

func sortHands(hands []Hand, cardValuesMap map[string]int) {
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(cardValuesMap, hands[i], hands[j])
	})
}

func loadHands(input string, jackIsWild bool) []Hand {
	var hands []Hand
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		splits := strings.Split(line, " ")
		cards := splits[0]
		wager := splits[1]

		hands = append(hands, Hand{cards, utils.ParseIntOrPanic(wager), scoreHand(splits[0], jackIsWild)})
	}

	return hands
}

func compareHands(cardValuesMap map[string]int, hand1 Hand, hand2 Hand) bool {
	if hand1.score == hand2.score {
		for pos := range hand1.cards {
			if cardValuesMap[string(hand1.cards[pos])] < cardValuesMap[string(hand2.cards[pos])] {
				return true
			} else if cardValuesMap[string(hand1.cards[pos])] > cardValuesMap[string(hand2.cards[pos])] {
				return false
			} else {
				continue
			}
		}
	}

	return hand1.score < hand2.score
}

func calculateWinnings(hands []Hand) int {
	winnings := 0
	for pos, hand := range hands {
		winnings += (pos + 1) * hand.wager
	}

	return winnings
}

func scoreHand(hand string, jackIsWild bool) int {
	var cardCounts map[string]int
	cardCounts = make(map[string]int)

	if jackIsWild {
		numWilds := 0

		for _, char := range hand {
			if string(char) == "J" {
				numWilds += 1
			} else {
				cardCounts[string(char)] += 1
			}
		}

		if numWilds == 5 {
			// 5 of a kind -> 5
			return 5
		}

		mostFrequentChar := ""
		for char, count := range cardCounts {
			if count > cardCounts[mostFrequentChar] {
				mostFrequentChar = char
			}
		}

		cardCounts[mostFrequentChar] += numWilds
	} else {
		for _, char := range hand {
			cardCounts[string(char)] += 1
		}
	}

	return scoreCardCounts(cardCounts)
}

func scoreCardCounts(cardCounts map[string]int) int {
	var counts map[int]int
	counts = make(map[int]int)
	for _, count := range cardCounts {
		counts[count] += 1
	}

	var score int

	if counts[5] == 1 {
		// 5 of a kind -> 5
		score = 5
	} else if counts[4] == 1 {
		// 4 of a kind -> 4
		score = 4
	} else if counts[3] == 1 && counts[2] == 1 {
		// fullHouse -> 3
		score = 3
	} else if counts[3] == 1 {
		// three of a kind -> 2
		score = 2
	} else if counts[2] == 2 {
		// 2 pair -> 1
		score = 1
	} else if counts[2] == 1 {
		// 1 pair -> 0
		score = 0
	} else {
		// nothing -> -1
		score = -1
	}

	return score
}

func initializeCardValuesMap(jackIsWild bool) map[string]int {
	cardValuesMap := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	if jackIsWild {
		cardValuesMap["J"] = 1
	} else {
		cardValuesMap["J"] = 11
	}

	return cardValuesMap
}
