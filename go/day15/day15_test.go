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
	input := utils.ReadFileIntoSlice("input_test_2")
	expect := "406"
	result := part2(input)
	if result != expect {
		t.Fatal("Day 15 - Part 2 output should be", expect, "but got", result)
	}
	input = utils.ReadFileIntoSlice("input_test_3")
	expect = "509"
	result = part2(input)
	if result != expect {
		t.Fatal("Day 15 - Part 2 output should be", expect, "but got", result)
	}
	input = utils.ReadFileIntoSlice("input_test_1")
	expect = "618"
	result = part2(input)
	if result != expect {
		t.Fatal("Day 15 - Part 2 output should be", expect, "but got", result)
	}
	input = utils.ReadFileIntoSlice("input_test")
	expect = "9021"
	result = part2(input)
	if result != expect {
		t.Fatal("Day 15 - Part 2 output should be", expect, "but got", result)
	}
}
