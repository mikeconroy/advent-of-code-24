package day10

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day10/input")
	return part1(input), part2(input)
}

type Grid map[Point]int

type Point struct {
	x int
	y int
}

func (grid Grid) get(x int, y int) int {
	return grid[Point{x, y}]
}

func loadGrid(input []string) Grid {
	grid := make(Grid)
	for y, row := range input {
		for x, val := range row {
			grid[Point{x, y}] = int(val - '0')
		}
	}
	return grid
}

func getTrailheads(point Point, currVal int, grid Grid, endPoints map[Point]int) map[Point]int {
	val, ok := grid[point]
	if !ok {
		return endPoints
	}
	if val != currVal+1 {
		return endPoints
	}
	if val == 9 {
		endPoints[point] += 1
		return endPoints
	}

	//Up
	endPoints = getTrailheads(Point{point.x, point.y - 1}, val, grid, endPoints)
	//Down
	endPoints = getTrailheads(Point{point.x, point.y + 1}, val, grid, endPoints)
	//Left
	endPoints = getTrailheads(Point{point.x - 1, point.y}, val, grid, endPoints)
	//Right
	endPoints = getTrailheads(Point{point.x + 1, point.y}, val, grid, endPoints)

	return endPoints
}

func part1(input []string) string {
	grid := loadGrid(input)
	totalScore := 0
	for point, height := range grid {
		if height == 0 {
			endPoints := make(map[Point]int)
			endPoints = getTrailheads(point, -1, grid, endPoints)
			totalScore += len(endPoints)
		}
	}
	return fmt.Sprint(totalScore)
}

func part2(input []string) string {
	grid := loadGrid(input)
	totalScore := 0
	for point, height := range grid {
		if height == 0 {
			endPoints := make(map[Point]int)
			endPoints = getTrailheads(point, -1, grid, endPoints)
			for _, val := range endPoints {
				totalScore += val
			}
		}
	}
	return fmt.Sprint(totalScore)
}
