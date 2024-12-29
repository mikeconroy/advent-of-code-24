package day20

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay20Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "4"
	result := part1(input, 30)
	if result != expect {
		t.Fatal("Day 20 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay20Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "7"
	result := part2(input, 73)
	if result != expect {
		t.Fatal("Day 20 - Part 2 output should be", expect, "but got", result)
	}
}
