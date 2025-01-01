package day21

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day21/input")
	return part1(input), part2(input)
}

var numPad string = "789456123_0A"
var dirPad string = "_^A<v>"

func createKeypad(pad string, holeRow int, holeCol int) func(string, string) string {
	return func(current string, target string) string {
		currPos := strings.Index(pad, current)
		targetPos := strings.Index(pad, target)
		rowDiff := (targetPos / 3) - (currPos / 3)
		colDiff := (targetPos % 3) - (currPos % 3)

		colDir := ""
		rowDir := ""
		if colDiff > 0 {
			colDir = ">"
		} else {
			colDir = "<"
		}
		if rowDiff > 0 {
			rowDir = "v"
		} else {
			rowDir = "^"
		}
		rowMoves := strings.Repeat(rowDir, abs(rowDiff))
		colMoves := strings.Repeat(colDir, abs(colDiff))

		if targetPos/3 == holeRow && currPos%3 == holeCol {
			return colMoves + rowMoves
		} else if currPos/3 == holeRow && targetPos%3 == holeCol {
			return rowMoves + colMoves
		} else {
			if strings.Contains(colMoves, "<") {
				return colMoves + rowMoves
			} else {
				return rowMoves + colMoves
			}
		}
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return a * -1
}

func pressKeypadsIterative(code string, pressFuncs []func(string, string) string) int {
	sequence := code
	for _, pressFunc := range pressFuncs {
		current := "A"
		newSequence := ""
		for _, target := range sequence {
			newSequence += pressFunc(current, string(target)) + "A"
			current = string(target)
		}
		sequence = newSequence
	}
	return len(sequence)
}

// Based on: https://www.reddit.com/r/adventofcode/comments/1hqm79a/2024_day_21_a_maybe_simpler_solution_to_a_hard/
func part1(input []string) string {
	pressNumpad := createKeypad(numPad, 3, 0)
	pressDirpad := createKeypad(dirPad, 0, 0)
	result := 0
	for _, code := range input {
		if code == "" {
			continue
		}
		length := pressKeypadsIterative(code, []func(string, string) string{pressNumpad, pressDirpad, pressDirpad})
		codeNum, _ := strconv.Atoi(code[:3])
		result += length * codeNum

	}
	return fmt.Sprint(result)
}

type CacheKey struct {
	current string
	target  string
	level   int
}

var totalLevels int
var cache map[CacheKey]int
var pressKey []func(string, string) string

func numPresses(current string, target string, level int) int {
	cacheKey := CacheKey{current, target, level}
	if presses, ok := cache[cacheKey]; ok {
		return presses
	}
	sequence := pressKey[level](current, target) + "A"
	if level == totalLevels-1 {
		return len(sequence)
	}

	length := 0
	current = "A"
	for _, target := range sequence {
		length += numPresses(current, string(target), level+1)
		current = string(target)
	}
	cache[cacheKey] = length

	return length
}

func pressKeypadsRecursive(code string) int {
	length := 0
	current := "A"
	for _, target := range code {
		length += numPresses(current, string(target), 0)
		current = string(target)
	}
	return length
}

func part2(input []string) string {
	pressNumpad := createKeypad(numPad, 3, 0)
	pressDirpad := createKeypad(dirPad, 0, 0)

	totalLevels = 26
	cache = make(map[CacheKey]int)
	pressKey = make([]func(string, string) string, totalLevels)
	pressKey[0] = pressNumpad
	for i := 1; i < totalLevels; i++ {
		pressKey[i] = pressDirpad
	}

	result := 0
	for _, code := range input {
		if code == "" {
			continue
		}
		presses := pressKeypadsRecursive(code)
		codeNum, _ := strconv.Atoi(code[:3])
		result += presses * codeNum

	}

	return fmt.Sprint(result)
}
