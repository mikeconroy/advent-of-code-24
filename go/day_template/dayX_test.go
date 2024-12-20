package dayX

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDayXPart1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "1"
	result := part1(input)
	if result != expect {
		t.Fatal("Day X - Part 1 output should be", expect, "but got", result)
	}
}

func TestDayXPart2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "2"
	result := part2(input)
	if result != expect {
		t.Fatal("Day X - Part 2 output should be", expect, "but got", result)
	}
}
