package day11

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay11Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "55312"
	result := part1(input)
	if result != expect {
		t.Fatal("Day 11 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay11Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "65601038650482"
	result := part2(input)
	if result != expect {
		t.Fatal("Day 11 - Part 2 output should be", expect, "but got", result)
	}
}
