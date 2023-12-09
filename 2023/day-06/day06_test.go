package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := "Time:      7  15   30\nDistance:  9  40  200"
	want := 288
	answer := Part1(input)
	if want != answer {
		t.Fatalf(`Part1() failed. Wanted %d received %d`, want, answer)
	}
}

func TestPart2(t *testing.T) {
	input := "Time:      7  15   30\nDistance:  9  40  200"
	want := 71503
	answer := Part2(input)
	if want != answer {
		t.Fatalf(`Part2() failed. Wanted %d received %d`, want, answer)
	}
}
