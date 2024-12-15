package day12

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay12Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "1930"
	result := part1(input)
	if result != expect {
		t.Fatal("Day 12 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay12Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "1206"
	result := part2(input)
	if result != expect {
		t.Fatal("Day 12 - Part 2 output should be", expect, "but got", result)
	}

	input = utils.ReadFileIntoSlice("input_test_1")
	expect = "236"
	result = part2(input)
	if result != expect {
		t.Fatal("Day 12 - Part 2 output should be", expect, "but got", result)
	}

	input = utils.ReadFileIntoSlice("input_test_2")
	expect = "368"
	result = part2(input)
	if result != expect {
		t.Fatal("Day 12 - Part 2 output should be", expect, "but got", result)
	}
}
