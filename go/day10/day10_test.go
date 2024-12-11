package day10

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay10Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "36"
	result := part1(input)
	if result != expect {
		t.Fatal("Day 10 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay10Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "81"
	result := part2(input)
	if result != expect {
		t.Fatal("Day 10 - Part 2 output should be", expect, "but got", result)
	}
}
