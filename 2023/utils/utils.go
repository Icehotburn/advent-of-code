package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInputFile(file *os.File) string {
	var input string

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		input += sc.Text() + "\n"
	}

	return input
}

func ParseIntOrPanic(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}
