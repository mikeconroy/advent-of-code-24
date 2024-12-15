package day13

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay13Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "480"
	result := part1(input)
	if result != expect {
		t.Fatal("Day 13 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay13Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "875318608908"
	result := part2(input)
	if result != expect {
		t.Fatal("Day 13 - Part 2 output should be", expect, "but got", result)
	}
}
