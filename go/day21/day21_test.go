package day21

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay21Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "126384"
	result := part1(input)
	if result != expect {
		t.Fatal("Day 21 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay21Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "154115708116294"
	result := part2(input)
	if result != expect {
		t.Fatal("Day 21 - Part 2 output should be", expect, "but got", result)
	}
}
