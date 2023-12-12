package main

import (
	"2023/utils"
	"fmt"
	"os"
	"slices"
	"strings"
)

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

	var histories [][]int

	answer := 0
	for _, line := range lines {
		history := utils.ParseNumbersFromString(line, " ")
		histories = append(histories, history)
		answer += extrapolate(history)
	}

	return answer
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	var histories [][]int

	answer := 0
	for _, line := range lines {
		history := utils.ParseNumbersFromString(line, " ")
		slices.Reverse(history)
		histories = append(histories, history)
		answer += extrapolate(history)
	}

	return answer
}

func extrapolate(history []int) int {
	var sequences [][]int
	sequences = append(sequences, history)

	for !isBottom(sequences[len(sequences)-1]) {
		previous := sequences[len(sequences)-1]

		var nextSequence []int

		for i := 1; i < len(previous); i++ {
			nextSequence = append(nextSequence, previous[i]-previous[i-1])
		}

		sequences = append(sequences, nextSequence)
	}

	nextValue := 0

	for _, sequence := range sequences {
		if len(sequence) == 0 {
			continue
		}
		nextValue += sequence[len(sequence)-1]
	}

	return nextValue
}

func isBottom(values []int) bool {
	for _, value := range values {
		if value != 0 {
			return false
		}
	}
	return true
}
