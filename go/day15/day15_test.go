package day15

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay15Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "10092"
	result := part1(input)
	if result != expect {
		t.Fatal("Day 15 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay15Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "2"
	result := part2(input)
	if result != expect {
		t.Fatal("Day 15 - Part 2 output should be", expect, "but got", result)
	}
}
