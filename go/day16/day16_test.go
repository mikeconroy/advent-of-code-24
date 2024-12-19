package day16

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay16Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "7036"
	result := part1(input)
	if result != expect {
		t.Fatal("Day 16 - Part 1 output should be", expect, "but got", result)
	}

	input = utils.ReadFileIntoSlice("input_test_1")
	expect = "11048"
	result = part1(input)
	if result != expect {
		t.Fatal("Day 16 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay16Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "2"
	result := part2(input)
	if result != expect {
		t.Fatal("Day 16 - Part 2 output should be", expect, "but got", result)
	}
}
