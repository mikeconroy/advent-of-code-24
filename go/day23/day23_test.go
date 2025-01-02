package day23

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay23Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "7"
	result := part1(input)
	if result != expect {
		t.Fatal("Day 23 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay23Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "co,de,ka,ta"
	result := part2(input)
	if result != expect {
		t.Fatal("Day 23 - Part 2 output should be", expect, "but got", result)
	}
}
