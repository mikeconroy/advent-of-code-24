package day20

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day20/input")
	return part1(input, 100), part2(input, 100)
}

type Point struct {
	x, y int
}

func getDistances(grid [][]rune, start Point, end Point) map[Point]int {
	distances := make(map[Point]int)
	currDistance := 0
	currPos := start
	prevPos := currPos
	for currPos != end {
		distances[currPos] = currDistance
		currDistance += 1

		nextPos := nextPosition(grid, currPos, prevPos)
		prevPos = currPos
		currPos = nextPos
	}
	distances[end] = currDistance

	return distances
}

var neighbours []Point = []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func nextPosition(grid [][]rune, pos Point, prevPos Point) Point {
	for _, neighbour := range neighbours {
		newPos := Point{pos.x + neighbour.x, pos.y + neighbour.y}
		if newPos != prevPos {
			if grid[newPos.y][newPos.x] != '#' {
				return newPos
			}
		}
	}
	return Point{}
}

func findCheatWithin(pos Point, distances map[Point]int, cheatDistance int, minimumSaved int) int {
	posDistance := distances[pos]
	count := 0
	for k, val := range distances {
		if k == pos {
			continue
		}

		diffX := absDiff(pos.x, k.x)
		diffY := absDiff(pos.y, k.y)
		diff := diffX + diffY
		if diff <= cheatDistance {
			if posDistance < val+diff {
				if (val - posDistance - diff) >= minimumSaved {
					count++
				}
			}
		}
	}

	return count
}

func absDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func part1(input []string, minPicoseconds int) string {
	grid, start, end := parseInput(input)
	distances := getDistances(grid, start, end)

	count := 0
	for k := range distances {
		if k == end {
			continue
		}

		count += findCheatWithin(k, distances, 2, minPicoseconds)
	}
	return fmt.Sprint(count)
}

func part2(input []string, minPicoseconds int) string {
	grid, start, end := parseInput(input)
	distances := getDistances(grid, start, end)

	count := 0
	for k := range distances {
		if k == end {
			continue
		}

		count += findCheatWithin(k, distances, 20, minPicoseconds)
	}
	return fmt.Sprint(count)
}

func parseInput(input []string) ([][]rune, Point, Point) {
	grid := make([][]rune, len(input)-1)
	start := Point{}
	end := Point{}
	for y, line := range input {
		if line == "" {
			continue
		}
		row := make([]rune, len(line))
		for x, val := range line {
			row[x] = val
			if val == 'S' {
				start.x = x
				start.y = y
			} else if val == 'E' {
				end.x = x
				end.y = y
			}
		}
		grid[y] = row
	}
	return grid, start, end
}
