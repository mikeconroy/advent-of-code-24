package day17

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay17Part1(t *testing.T) {
	comp := Computer{
		a:            0,
		b:            0,
		c:            9,
		instructions: []int{2, 6},
		pointer:      0,
		output:       []int{},
	}
	for !comp.tick() {
	}
	if comp.b != 1 {
		t.Fatal("Day 17 - Computer B Register should be 1.")
	}

	comp = Computer{
		a:            10,
		b:            0,
		c:            0,
		instructions: []int{5, 0, 5, 1, 5, 4},
		pointer:      0,
		output:       []int{},
	}
	for !comp.tick() {
	}
	expectArray := []int{0, 1, 2}
	for i, val := range comp.output {
		if val != expectArray[i] {
			t.Fatal("Day 17 - Computer output should be 0,1,2.")
		}
	}

	comp = Computer{
		a:            2024,
		b:            0,
		c:            0,
		instructions: []int{0, 1, 5, 4, 3, 0},
		pointer:      0,
	}
	for !comp.tick() {
	}
	if comp.a != 0 {
		t.Fatal("Day 17 - Computer Register A is not 0")
	}
	expectArray = []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}
	for i, val := range comp.output {
		if val != expectArray[i] {
			t.Fatal("Day 17 - Computer output should be 0,1,2.")
		}
	}

	comp = Computer{
		a:            0,
		b:            29,
		c:            0,
		instructions: []int{1, 7},
	}
	for !comp.tick() {
	}
	if comp.b != 26 {
		t.Fatal("Day 17 - Computer Register B is not 26.")
	}

	comp = Computer{
		a:            0,
		b:            2024,
		c:            43690,
		instructions: []int{4, 0},
	}
	for !comp.tick() {
	}
	if comp.b != 44354 {
		t.Fatal("Day 17 - Computer Register B is no 44354.")
	}

	input := utils.ReadFileIntoSlice("input_test")
	expect := "4635635210"
	result := part1(input)
	if result != expect {
		t.Fatal("Day 17 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay17Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test_1")
	expect := "117440"
	result := part2(input)
	if result != expect {
		t.Fatal("Day 17 - Part 2 output should be", expect, "but got", result)
	}
}
