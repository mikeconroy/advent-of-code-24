package day18

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay18Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "22"
	result := part1(input, 6, 12)
	if result != expect {
		t.Fatal("Day 18 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay18Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "6,1"
	result := part2(input, 6, 12)
	if result != expect {
		t.Fatal("Day 18 - Part 2 output should be", expect, "but got", result)
	}
}
