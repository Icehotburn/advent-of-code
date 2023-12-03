package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type coordinates struct {
	row int
	col int
}

func main() {
	//Read input file
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(input)

	// Store the answer
	gearRatioSum := 0

	// Store the input as an array
	schematic := make([]string, 0)

	for sc.Scan() {
		line := sc.Text()
		schematic = append(schematic, line)
	}

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

	// Print the answer
	fmt.Println(gearRatioSum)

	// Close Input
	err = input.Close()
	if err != nil {
		panic(err)
	}
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
