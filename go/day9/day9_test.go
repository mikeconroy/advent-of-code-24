package day9

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay9Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "1928"
	if part1(input) != expect {
		t.Fatal("Day 9 - Part 1 output should be " + expect)
	}
}

func TestDay9Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "2" {
		t.Fatal("Day 9 - Part 2 output should be 2")
	}
}
