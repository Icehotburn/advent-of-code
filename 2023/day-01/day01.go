package main

import (
	"2023/utils"
	"errors"
	"fmt"
	"os"
	"strconv"
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
	// Sum of all the calibration values
	calibrationSum := 0

	lines := strings.Split(input, "\n")

	for _, line := range lines {

		// Skip the last line
		if line == "" {
			continue
		}

		/*
			Potential optimization:
			Only store two digits per line since only the first and last digits are needed
		*/
		var digits []int
		for _, char := range line {
			digit, err := strconv.Atoi(string(char))

			if err == nil {
				digits = append(digits, digit)
			}
		}

		first := digits[0]
		last := digits[len(digits)-1]

		calibration := first*10 + last

		calibrationSum += calibration
	}

	return calibrationSum
}

func Part2(input string) int {

	// Sum of all the calibration values
	calibrationSum := 0

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		/*
			Potential optimization:
			Only store two digits per line since only the first and last digits are needed
		*/
		var digits []int

		// Skip the last line
		if line == "" {
			continue
		}

		for pos, char := range line {
			/*
				Potential optimization:
				The current approach checks if the current char is an integer and adds it to the digits array if it is one.
				If it is not, it attempts to check all substrings starting with the current char to see if the form a number as a word.
				This can be greatly optimized since we know that the word must be between 3 and 5 characters, so there is no need to check other possibilities.
				Additionally, we can exit the substring check if one of the characters is identified as a number (i.e 1),
				since we know that a number as a word cannot contain an integer.
			*/

			// try to convert the char literal to an int i.e "1" -> 1
			digit, err := strconv.Atoi(string(char))

			if err == nil {
				digits = append(digits, digit)
				continue
			}

			// try to convert the word to an int i.e "one" -> 1
			word := ""
			for _, c := range line[pos:] {
				word = word + string(c)

				digit, err = convertWordToInt(word)

				if err == nil {
					digits = append(digits, digit)
					continue
				}
			}
		}

		first := digits[0]
		last := digits[len(digits)-1]

		calibration := first*10 + last
		calibrationSum += calibration
	}

	return calibrationSum
}

func convertWordToInt(s string) (int, error) {
	switch s {
	case "zero":
		return 0, nil
	case "one":
		return 1, nil
	case "two":
		return 2, nil
	case "three":
		return 3, nil
	case "four":
		return 4, nil
	case "five":
		return 5, nil
	case "six":
		return 6, nil
	case "seven":
		return 7, nil
	case "eight":
		return 8, nil
	case "nine":
		return 9, nil
	default:
		return -1, errors.New("not a digit")
	}
}
