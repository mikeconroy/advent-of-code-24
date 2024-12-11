package day9

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay9Part1(t *testing.T) {
	expect := "60"
	result := part1([]string{"12345"})
	if result != expect {
		t.Fatal("Day 9 - Part 1 12345 output should be", expect, "but got", result)
	}
	input := utils.ReadFileIntoSlice("input_test")
	expect = "1928"
	result = part1(input)
	if result != expect {
		t.Fatal("Day 9 - Part 1 output should be " + expect + "but got" + result)
	}
}

func TestChecksum(t *testing.T) {
	expect := 10
	list := &DoublyLinkedList{}
	list.append(&Node{
		id:   0,
		size: 2,
	})
	list.append(&Node{
		id:   2,
		size: 2,
	})
	result := calculateChecksum(list)
	if result != expect {
		t.Fatal("Day 9 - Checksum should be", expect, "but was", result)
	}
}

func TestDay9Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	if part2(input) != "2858" {
		t.Fatal("Day 9 - Part 2 output should be 2858 but got", part2(input))
	}
}
