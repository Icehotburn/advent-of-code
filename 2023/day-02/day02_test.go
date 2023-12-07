package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green\n"
	bagContents := InitializeBagContents(12, 13, 14)

	want := 8
	answer := Part1(bagContents, input)
	if want != answer {
		t.Fatalf(`Part1() failed. Wanted %d received %d`, want, answer)
	}
}

func TestPart2(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green\n"

	want := 2286
	answer := Part2(input)
	if want != answer {
		t.Fatalf(`Part2() failed. Wanted %d received %d`, want, answer)
	}
}
