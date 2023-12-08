package main

import (
	"2023/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinates struct {
	row int
	col int
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
	// Store the answer
	partNumberSum := 0

	// Store the input as an array
	schematic := strings.Split(input, "\n")
	// Remove the last line of the schematic since it is empty
	schematic = schematic[:len(schematic)-1]

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

	return partNumberSum
}

func Part2(input string) int {
	// Store the answer
	gearRatioSum := 0

	// Store the input as an array
	schematic := strings.Split(input, "\n")
	// Remove the last line of the schematic since it is empty
	schematic = schematic[:len(schematic)-1]

	// Store a map of gear coordinates and their adjacent numbers
	var gearsMap map[coordinates][]int
	gearsMap = make(map[coordinates][]int)

	// Parse the schematic
	for row := 0; row < len(schematic); row++ {
		for col := 0; col < len(schematic[row]); col++ {

			offset := 0
			num := 0
			gearCoordinates := coordinates{-1, -1}

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

					//check if it touches a gear
					if adjacentGearCoordinates(row, col+offset, schematic).row != -1 && adjacentGearCoordinates(row, col+offset, schematic).col != -1 {
						gearCoordinates = adjacentGearCoordinates(row, col+offset, schematic)
					}

					offset += 1
					continue

				} else { // else increment counter and move to next char
					col = col + offset
					break
				}
			}

			// if num != 0 and isAdjacentToGear add it to the gearsMap
			if num > 0 && gearCoordinates.row != -1 && gearCoordinates.col != -1 {
				gearsMap[gearCoordinates] = append(gearsMap[gearCoordinates], num)
			}
		}
	}

	for _, value := range gearsMap {
		if len(value) == 2 {
			gearRatio := value[0] * value[1]
			gearRatioSum += gearRatio
		}
	}

	return gearRatioSum
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
	// check if it is a digit or "."
	_, err := strconv.Atoi(s)
	if s == "." || err == nil {
		return false
	}

	return true
}

func adjacentGearCoordinates(row int, col int, schematic []string) coordinates {

	// check upper left
	if col > 0 && row > 0 {
		if isGear(string(schematic[row-1][col-1])) {
			return coordinates{row - 1, col - 1}
		}
	}

	// check direct left
	if col > 0 {
		if isGear(string(schematic[row][col-1])) {
			return coordinates{row, col - 1}
		}
	}

	// check lower left
	if col > 0 && row < len(schematic)-1 {
		if isGear(string(schematic[row+1][col-1])) {
			return coordinates{row + 1, col - 1}
		}
	}

	// check up
	if row > 0 {
		if isGear(string(schematic[row-1][col])) {
			return coordinates{row - 1, col}
		}
	}

	// check down
	if row < len(schematic)-1 {
		if isGear(string(schematic[row+1][col])) {
			return coordinates{row + 1, col}
		}
	}

	// check upper right
	if col < len(schematic[row])-1 && row > 0 {
		if isGear(string(schematic[row-1][col+1])) {
			return coordinates{row - 1, col + 1}
		}
	}

	// check right
	if col < len(schematic[row])-1 {
		if isGear(string(schematic[row][col+1])) {
			return coordinates{row, col + 1}
		}
	}

	// check lower right
	if col < len(schematic[row])-1 && row < len(schematic)-1 {
		if isGear(string(schematic[row+1][col+1])) {
			return coordinates{row + 1, col + 1}
		}
	}

	//base case return false
	return coordinates{-1, -1}
}

func isGear(s string) bool {
	return s == "*"
}
