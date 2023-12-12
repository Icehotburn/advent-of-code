package main

import (
	"testing"
)

func Test1Part1(t *testing.T) {
	input := "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45"

	want := 114
	answer := Part1(input)
	if want != answer {
		t.Fatalf(`Part1() failed. Wanted %d received %d`, want, answer)
	}
}

func TestPart2(t *testing.T) {
	input := "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45"

	want := 2
	answer := Part2(input)
	if want != answer {
		t.Fatalf(`Part2() failed. Wanted %d received %d`, want, answer)
	}
}
