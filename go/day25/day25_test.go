package day25

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay25Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "3"
	result := part1(input)
	if result != expect {
		t.Fatal("Day 25 - Part 1 output should be", expect, "but got", result)
	}
}
