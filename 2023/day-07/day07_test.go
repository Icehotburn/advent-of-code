package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"

	want := 6440
	answer := Part1(input)
	if want != answer {
		t.Fatalf(`Part1() failed. Wanted %d received %d`, want, answer)
	}
}

func TestPart2(t *testing.T) {
	input := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"

	want := 5905
	answer := Part2(input)
	if want != answer {
		t.Fatalf(`Part2() failed. Wanted %d received %d`, want, answer)
	}
}
