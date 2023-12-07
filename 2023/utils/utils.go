package utils

import (
	"bufio"
	"os"
)

func ReadInputFile(file *os.File) string {
	var input string

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		input += sc.Text() + "\n"
	}

	return input
}
