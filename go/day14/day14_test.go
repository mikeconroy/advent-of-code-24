package day14

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay14Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "12"
	result := part1(input, 11, 7)
	if result != expect {
		t.Fatal("Day 14 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay14Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "7138"
	result := part2(input, 101, 103)
	if result != expect {
		t.Fatal("Day 14 - Part 2 output should be", expect, "but got", result)
	}
}
