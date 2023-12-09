package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func Contains[T comparable](arr []T, val T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func ParseNumbersFromString(s string, delimiter string) []int {
	s = strings.TrimSpace(s)
	values := strings.Split(s, delimiter)

	var result []int

	for _, value := range values {
		if value != "" {
			result = append(result, ParseIntOrPanic(value))
		}
	}

	return result
}
