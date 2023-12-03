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

	// Store the answer
	partNumberSum := 0

	// Store the input as an array
	schematic := make([]string, 0)

	for sc.Scan() {
		line := sc.Text()
		schematic = append(schematic, line)
	}

	// Parse the schematic
	for row := 0; row < len(schematic); row++ {
		for col := 0; col < len(schematic[row]); col++ {

			offset := 0
			num := 0
			isPartNumber := false

			// for each num find adjacent nums until we hit EOL or a non digit char
			for col+offset < len(schematic[row]) {
				// get the current char
				char := string(schematic[row][col+offset])
				// try to convert it to a digit
				digit, err := strconv.Atoi(char)

				// if it is a digit...
				if err == nil {

					//add it to the current num
					num = (num * 10) + digit

					//check if it touches a symbol
					if adjacentToSymbol(row, col+offset, schematic) {
						isPartNumber = true
					}

					offset += 1
					continue

				} else { // else increment counter and move to next char
					col = col + offset
					break
				}
			}

			// if num != 0 and isPartNumber add it to the partNumberSum
			if num > 0 && isPartNumber {
				partNumberSum += num
			}
		}
	}

	// Print the answer
	fmt.Println(partNumberSum)

	// Close Input
	err = input.Close()
	if err != nil {
		panic(err)
	}
}

func adjacentToSymbol(row int, col int, schematic []string) bool {

	// check upper left
	if col > 0 && row > 0 {
		if isSymbol(string(schematic[row-1][col-1])) {
			return true
		}
	}

	// check direct left
	if col > 0 {
		if isSymbol(string(schematic[row][col-1])) {
			return true
		}
	}

	// check lower left
	if col > 0 && row < len(schematic)-1 {
		if isSymbol(string(schematic[row+1][col-1])) {
			return true
		}
	}

	// check up
	if row > 0 {
		if isSymbol(string(schematic[row-1][col])) {
			return true
		}
	}

	// check down
	if row < len(schematic)-1 {
		if isSymbol(string(schematic[row+1][col])) {
			return true
		}
	}

	// check upper right
	if col < len(schematic[row])-1 && row > 0 {
		if isSymbol(string(schematic[row-1][col+1])) {
			return true
		}
	}

	// check right
	if col < len(schematic[row])-1 {
		if isSymbol(string(schematic[row][col+1])) {
			return true
		}
	}

	// check lower right
	if col < len(schematic[row])-1 && row < len(schematic)-1 {
		if isSymbol(string(schematic[row+1][col+1])) {
			return true
		}
	}

	//base case return false
	return false
}

func isSymbol(s string) bool {
	// check if it is a "."
	if s == "." {
		return false
	}

	// check if it is a digit
	_, err := strconv.Atoi(s)
	if err == nil {
		return false
	}

	return true
}
