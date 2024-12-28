package day19

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay19Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "6"
	result := part1(input)
	if result != expect {
		t.Fatal("Day 19 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay19Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "16"
	result := part2(input)
	if result != expect {
		t.Fatal("Day 19 - Part 2 output should be", expect, "but got", result)
	}
}
