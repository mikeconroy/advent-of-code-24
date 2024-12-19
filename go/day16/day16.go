package day16

import (
	"fmt"
	"github.com/mikeconroy/advent-of-code-24/utils"
	"math"
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

func walk(grid Grid, start Point, end Point, dir Direction, visited map[Point]Position, maxCost int) Position {
	visited[start] = Position{start, right, 0}
	toVisit := []Position{
		{
			pos:  add(start, up),
			dir:  up,
			cost: getCost(dir, up),
		},
		{
			pos:  add(start, down),
			dir:  down,
			cost: getCost(dir, down),
		}, {
			pos:  add(start, left),
			dir:  left,
			cost: getCost(dir, left),
		}, {
			pos:  add(start, right),
			dir:  right,
			cost: getCost(dir, right),
		},
	}
	for len(toVisit) != 0 {
		// fmt.Println(toVisit)
		currPos := toVisit[0]
		// fmt.Println("visiting:", currPos)
		if grid[currPos.pos.y][currPos.pos.x] != '#' {

			if visitedPos, ok := visited[currPos.pos]; !ok || (currPos.cost < visitedPos.cost && currPos.cost <= maxCost) {
				visited[currPos.pos] = Position{currPos.pos, currPos.dir, currPos.cost}

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
	// fmt.Println(visited)
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
	visited := make(map[Point]Position)
	grid, start, end := parseInput(input)
	return fmt.Sprint(walk(grid, start, end, right, visited, math.MaxInt).cost)
}

// Either bruteforce Dijkstras algo (part1) by checking every cell
// to see whether it can reach the end in the best time - cache visited map as
// we go to speed it up.
// Or DFS through each path seeing if it's possible to reach the end within the best cost.
// Or update the walk function to include the paths taken and then add these to a set at the end of the algo.
func part2(input []string) string {
	grid, start, end := parseInput(input)
	visited := make(map[Point]Position)
	best := walk(grid, start, end, right, visited, math.MaxInt).cost
	fmt.Println(best)
	count := 0

	logCounter := 0
	for _, val := range visited {
		// fmt.Println(logCounter, "/", len(visited))
		logCounter += 1
		result := val.cost
		result += walk(grid, val.pos, end, val.dir, make(map[Point]Position), best).cost
		if result <= best {
			count++
		}
	}
	return fmt.Sprint(count)
}
