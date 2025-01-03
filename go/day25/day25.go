package day25

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day25/input")
	return part1(input), "Happy Xmas!"
}

// 5-Pins
func parseInput(input []string) ([][]int, [][]int) {
	keys := [][]int{}
	locks := [][]int{}

	currType := "none"
	pinNo := 0

	// Represents height of each pin
	pins := make([]int, 5)
	for _, line := range input {
		if currType == "none" {
			if line == "#####" {
				currType = "key"
			} else {
				currType = "lock"
			}
		}

		for i, char := range line {
			if char == '#' {
				if currType == "lock" {
					if pins[i] < 6-pinNo {
						pins[i] = 6 - pinNo
					}
				} else {
					pins[i] = pinNo
				}
			}
		}
		pinNo++

		if line == "" {
			if currType == "lock" {
				locks = append(locks, pins)
			} else {
				keys = append(keys, pins)
			}
			currType = "none"
			pinNo = 0
			pins = make([]int, 5)
		}
	}

	return keys, locks
}

// 182713 is too high
func part1(input []string) string {
	keys, locks := parseInput(input)

	count := 0
	for _, key := range keys {
		for _, lock := range locks {
			validCombo := true
			for pin := 0; pin < 5; pin++ {
				if key[pin]+lock[pin] > 5 {
					validCombo = false
					break
				}
			}
			if validCombo {
				count++
			}
		}
	}

	return fmt.Sprint(count)
}
