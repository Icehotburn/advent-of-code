package main

import (
	"2023/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read input file
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(input)

	// first line is time
	sc.Scan()
	timeLine := sc.Text() // i.e "Time:      7  15   30"
	times := parseNumbers(timeLine[5:])

	// second line is distance
	sc.Scan()
	distanceLine := sc.Text() // i.e "Distance:  9  40  200"
	distances := parseNumbers(distanceLine[9:])

	// boatSpeed
	boatSpeedMMperMS := 1

	result := 0

	/* Potential optimization: this lookup is super slow
		Instead of checking every distance value, we can find the min and max times that break the record
	    Then we can return the number of values in the range between the min and max times
		This would greatly speed up the program.
	*/
	for race := 0; race < len(times); race++ {
		ways := 0

		for i := 0; i < times[race]; i++ {
			distanceTraveled := i * boatSpeedMMperMS * (times[race] - i)
			if distanceTraveled > distances[race] {
				ways += 1
			}
		}

		if ways == 0 {
			continue
		} else if result == 0 {
			result = ways
		} else {
			result = result * ways
		}
	}

	fmt.Println(result)

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
