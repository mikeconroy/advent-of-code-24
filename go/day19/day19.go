package day19

import (
	"fmt"
	"strings"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day19/input")
	return part1(input), part2(input)
}

func isPossible(pattern string, towels []string, cache map[string]bool) bool {
	if val, ok := cache[pattern]; ok {
		return val
	}
	// fmt.Println("Pattern:", pattern)

	for _, towel := range towels {
		// If towel matches first n chars of pattern
		// remove those chars and call isPossible again
		if len(pattern) >= len(towel) {
			if strings.Index(pattern, towel) == 0 {
				newPattern := pattern[len(towel):]
				if newPattern == "" {
					cache[pattern] = true
					return true
				}
				// fmt.Println("Index:", strings.Index(pattern, towel), "Towel:", towel, "Pattern:", pattern, "New:", newPattern)
				if isPossible(newPattern, towels, cache) {
					cache[newPattern] = true
					return true
				}
			}
		}
	}
	cache[pattern] = false
	return false
}

func part1(input []string) string {
	towels, patterns := parseInput(input)
	count := 0
	cache := make(map[string]bool)
	for _, pattern := range patterns {
		if isPossible(pattern, towels, cache) {
			count += 1
		}
	}
	return fmt.Sprint(count)
}

func countCombinations(pattern string, towels []string, cache map[string]int) int {
	if val, ok := cache[pattern]; ok {
		return val
	}

	count := 0

	for _, towel := range towels {
		if len(pattern) >= len(towel) {
			if pattern == towel {
				count += 1
			} else {
				if strings.Index(pattern, towel) == 0 {
					newPattern := pattern[len(towel):]
					count += countCombinations(newPattern, towels, cache)
				}
			}
		}
	}

	cache[pattern] = count
	return count
}

func part2(input []string) string {
	towels, patterns := parseInput(input)
	count := 0
	cache := make(map[string]int)
	for _, pattern := range patterns {
		count += countCombinations(pattern, towels, cache)
	}

	return fmt.Sprint(count)
}

func parseInput(input []string) ([]string, []string) {
	towels := strings.Split(input[0], ", ")
	patterns := make([]string, len(input)-3)
	for i, line := range input[2:] {
		if line != "" {
			patterns[i] = line
		}

	}
	return towels, patterns
}
