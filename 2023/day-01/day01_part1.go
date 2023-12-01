package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Read input file
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(input)

	// Sum of all the calibration values
	calibrationSum := 0

	for sc.Scan() {
		line := sc.Text()
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

	// Print the answer
	fmt.Println(calibrationSum)

	// Close Input
	err = input.Close()
	if err != nil {
		panic(err)
	}

}
