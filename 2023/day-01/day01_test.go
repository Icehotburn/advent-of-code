package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"

	want := 142
	answer := Part1(input)
	if want != answer {
		t.Fatalf(`Part1() failed. Wanted %d received %d`, want, answer)
	}
}

func TestPart2(t *testing.T) {
	input := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"

	want := 281
	answer := Part2(input)
	if want != answer {
		t.Fatalf(`Part2() failed. Wanted %d received %d`, want, answer)
	}
}
