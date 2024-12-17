package day15

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day15/input")
	return part1(input), part2(input)
}

type Grid = [][]rune

func parseInput(input []string) (Grid, []rune) {
	isInstructions := false
	grid := Grid{}
	instructions := []rune{}
	for _, line := range input {
		if line == "" {
			isInstructions = true
			continue
		}
		if isInstructions {
			for _, char := range line {
				instructions = append(instructions, char)
			}
			continue
		}

		row := []rune{}
		for _, char := range line {
			row = append(row, char)
		}
		grid = append(grid, row)

	}
	return grid, instructions
}

func print(grid Grid) {
	for _, y := range grid {
		for _, x := range y {
			fmt.Print(string(x))
		}
		fmt.Println()
	}
}

type Point struct {
	x, y int
}

func getStartPoint(grid Grid) Point {
	for y, line := range grid {
		for x, val := range line {
			if val == '@' {
				return Point{x, y}
			}
		}
	}
	return Point{}
}

func moveRobot(pos Point, grid Grid, instruction rune) Point {
	var direction Point
	switch instruction {
	case '<':
		direction = Point{-1, 0}
	case '^':
		direction = Point{0, -1}
	case '>':
		direction = Point{1, 0}
	case 'v':
		direction = Point{0, 1}
	}

	newPoint := Point{pos.x + direction.x, pos.y + direction.y}
	if grid[newPoint.y][newPoint.x] == '.' {
		grid[pos.y][pos.x] = '.'
		grid[newPoint.y][newPoint.x] = '@'
		return newPoint
	}
	if grid[newPoint.y][newPoint.x] == '#' {
		return pos
	}
	val := grid[newPoint.y][newPoint.x]
	for val != '#' {
		if val == 'O' {
			newPoint.x += direction.x
			newPoint.y += direction.y
			val = grid[newPoint.y][newPoint.x]
		} else if val == '.' {
			grid[newPoint.y][newPoint.x] = 'O'
			grid[pos.y+direction.y][pos.x+direction.x] = '@'
			grid[pos.y][pos.x] = '.'
			return Point{pos.x + direction.x, pos.y + direction.y}
		}
	}
	return pos
}

func calculateGps(grid Grid) int {
	gps := 0
	for y, row := range grid {
		for x, val := range row {
			if val == 'O' {
				gps += (100 * y) + x
			}
		}
	}
	return gps
}

func part1(input []string) string {
	grid, instructions := parseInput(input)
	robotPos := getStartPoint(grid)
	for _, instruction := range instructions {
		robotPos = moveRobot(robotPos, grid, instruction)
	}
	result := calculateGps(grid)
	return fmt.Sprint(result)
}

func part2(input []string) string {
	return fmt.Sprint(2)
}
