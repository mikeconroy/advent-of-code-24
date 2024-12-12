package day11

import (
	"fmt"
	"github.com/mikeconroy/advent-of-code-24/utils"
	"strconv"
	"strings"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day11/input")
	return part1(input), part2(input)
}

func Blink(stones []int) []int {
	newStones := []int{}
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if stoneStr := strconv.Itoa(stone); len(stoneStr)%2 == 0 {
			half := len(stoneStr) / 2
			lhs, _ := strconv.Atoi(stoneStr[0:half])
			rhs, _ := strconv.Atoi(stoneStr[half:])
			newStones = append(newStones, lhs, rhs)
		} else {
			newStones = append(newStones, stone*2024)
		}
	}
	return newStones

}

func parseInput(input string) []int {
	result := []int{}
	for _, val := range strings.Split(input, " ") {
		num, _ := strconv.Atoi(val)
		result = append(result, num)
	}
	return result
}

type Key struct {
	stone, movesLeft int
}

var cache = make(map[Key]int)

func blinkXTimes(stone int, x int) int {
	if x == 0 {
		return 1
	}
	key := Key{stone, x}
	result, ok := cache[key]
	if ok {
		return result
	}

	result = 0
	if stone == 0 {
		result = blinkXTimes(1, x-1)
	} else if stoneStr := strconv.Itoa(stone); len(stoneStr)%2 == 0 {
		half := len(stoneStr) / 2
		lhs, _ := strconv.Atoi(stoneStr[0:half])
		rhs, _ := strconv.Atoi(stoneStr[half:])
		result = blinkXTimes(lhs, x-1)
		result += blinkXTimes(rhs, x-1)
	} else {
		result = blinkXTimes(stone*2024, x-1)
	}
	cache[key] = result

	return result
}

func part1(input []string) string {
	stones := parseInput(input[0])
	for i := 0; i < 25; i++ {
		stones = Blink(stones)
	}
	return fmt.Sprint(len(stones))
}
func part2(input []string) string {
	stones := parseInput(input[0])
	result := 0
	for _, stone := range stones {
		ans := blinkXTimes(stone, 75)
		result += ans
	}
	return fmt.Sprint(result)
}
