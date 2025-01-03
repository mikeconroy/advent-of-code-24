package day24

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay24BoolsToInt(t *testing.T) {
	expect := 28
	bools := []bool{false, false, true, true, true}
	result := boolsToInt(bools)
	if result != expect {
		t.Fatal("Day 24 - Bools to Int conversion should be", expect, "but got", result)
	}
}

func TestDay24Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "4"
	result := part1(input)
	if result != expect {
		t.Fatal("Day 24 - Part 1 output should be", expect, "but got", result)
	}

	input = utils.ReadFileIntoSlice("input_test_1")
	expect = "2024"
	result = part1(input)
	if result != expect {
		t.Fatal("Day 24 - Part 1 output should be", expect, "but got", result)
	}
}
