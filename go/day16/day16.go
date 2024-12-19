package day16

import (
	"fmt"
	"math"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day16/input")
	return part1(input), part2(input)
}

type Grid = [][]rune

type Point struct {
	x, y int
}

func parseInput(input []string) (Grid, Point, Point) {
	grid := Grid{}
	start := Point{}
	end := Point{}
	for y, line := range input {
		row := []rune{}
		for x, val := range line {
			row = append(row, val)
			if val == 'S' {
				start = Point{x, y}
			} else if val == 'E' {
				end = Point{x, y}
			}
		}
		grid = append(grid, row)
	}
	return grid, start, end
}

func printGrid(grid Grid) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
}

type Direction = Point

var (
	left  = Point{-1, 0}
	right = Point{1, 0}
	up    = Point{0, -1}
	down  = Point{0, 1}
)

func add(a Point, b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

type Position struct {
	pos  Point
	dir  Direction
	cost int
}

func walk(grid Grid, start Point, end Point) int {
	visited := make(map[Point]int)
	visited[start] = 0
	toVisit := []Position{
		{
			pos:  add(start, up),
			dir:  up,
			cost: getCost(right, up),
		},
		{
			pos:  add(start, down),
			dir:  down,
			cost: getCost(right, down),
		}, {
			pos:  add(start, left),
			dir:  left,
			cost: getCost(right, left),
		}, {
			pos:  add(start, right),
			dir:  right,
			cost: getCost(right, right),
		},
	}
	for len(toVisit) != 0 {
		// fmt.Println(toVisit)
		currPos := toVisit[0]
		// fmt.Println("visiting:", currPos)
		if grid[currPos.pos.y][currPos.pos.x] != '#' {
			if _, ok := visited[currPos.pos]; !ok {
				visited[currPos.pos] = math.MaxInt
			}

			if currPos.cost < visited[currPos.pos] {
				visited[currPos.pos] = currPos.cost

				if currPos.pos != end {
					// Check surrounding tiles
					toVisit = append(toVisit, []Position{
						{
							pos:  add(currPos.pos, up),
							dir:  up,
							cost: currPos.cost + getCost(currPos.dir, up),
						}, {
							pos:  add(currPos.pos, down),
							dir:  down,
							cost: currPos.cost + getCost(currPos.dir, down),
						}, {
							pos:  add(currPos.pos, right),
							dir:  right,
							cost: currPos.cost + getCost(currPos.dir, right),
						}, {
							pos:  add(currPos.pos, left),
							dir:  left,
							cost: currPos.cost + getCost(currPos.dir, left),
						},
					}...)
				}
			}
		}
		// remove from to Visit
		toVisit = toVisit[1:]
	}
	return visited[end]
}

func getCost(from Direction, to Direction) int {
	if from == up && to == down {
		return 2001
	}
	if from == down && to == up {
		return 2001
	}
	if from == left && to == right {
		return 2001
	}
	if from == right && to == left {
		return 2001
	}
	if from == to {
		return 1
	}
	return 1001
}

func part1(input []string) string {
	grid, start, end := parseInput(input)
	return fmt.Sprint(walk(grid, start, end))
}

func part2(input []string) string {
	return fmt.Sprint(2)
}
