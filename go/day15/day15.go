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

func parseInput(input []string, expanded bool) (Grid, []rune) {
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
			if expanded {
				switch char {
				case '#':
					row = append(row, []rune{'#', '#'}...)
				case '.':
					row = append(row, []rune{'.', '.'}...)
				case '@':
					row = append(row, []rune{'@', '.'}...)
				case 'O':
					row = append(row, []rune{'[', ']'}...)

				}
			} else {
				row = append(row, char)
			}
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

func getDirection(instruction rune) (direction Point) {
	switch instruction {
	case '<':
		direction = left
	case '^':
		direction = up
	case '>':
		direction = right
	case 'v':
		direction = down
	}
	return
}

func moveRobot(pos Point, grid Grid, instruction rune) Point {
	direction := getDirection(instruction)

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
			if val == 'O' || val == '[' {
				gps += (100 * y) + x
			}
		}
	}
	return gps
}

func part1(input []string) string {
	grid, instructions := parseInput(input, false)
	robotPos := getStartPoint(grid)
	for _, instruction := range instructions {
		robotPos = moveRobot(robotPos, grid, instruction)
	}
	result := calculateGps(grid)
	return fmt.Sprint(result)
}

var (
	left  Point = Point{-1, 0}
	right Point = Point{1, 0}
	up    Point = Point{0, -1}
	down  Point = Point{0, 1}
)

func canMove(from Point, direction Point, grid Grid) bool {
	newPos := Point{x: from.x + direction.x, y: from.y + direction.y}
	newVal := grid[newPos.y][newPos.x]
	if newVal == '#' {
		return false
	}
	if newVal == '.' {
		return true
	}
	if newVal == '[' {
		if direction == right {
			return canMove(Point{newPos.x + 1, newPos.y}, direction, grid)
		} else {
			return canMove(newPos, direction, grid) && canMove(Point{newPos.x + 1, newPos.y}, direction, grid)
		}
	}
	if newVal == ']' {
		if direction == left {
			return canMove(Point{newPos.x - 1, newPos.y}, direction, grid)
		} else {
			return canMove(newPos, direction, grid) && canMove(Point{newPos.x - 1, newPos.y}, direction, grid)
		}
	}

	return false
}

func move(from Point, direction Point, grid Grid) {
	newPos := Point{from.x + direction.x, from.y + direction.y}
	newVal := grid[newPos.y][newPos.x]
	currVal := grid[from.y][from.x]
	if newVal == '[' {
		move(newPos, direction, grid)
		if direction != right {
			move(Point{newPos.x + 1, newPos.y}, direction, grid)
		}
	} else if newVal == ']' {
		move(newPos, direction, grid)
		if direction != left {
			move(Point{newPos.x - 1, newPos.y}, direction, grid)
		}
	}
	grid[newPos.y][newPos.x] = currVal
	grid[from.y][from.x] = '.'
}

func part2(input []string) string {
	grid, instructions := parseInput(input, true)
	robotPos := getStartPoint(grid)
	for _, instruction := range instructions {
		direction := getDirection(instruction)
		if canMove(robotPos, direction, grid) {
			move(robotPos, direction, grid)
			robotPos = Point{robotPos.x + direction.x, robotPos.y + direction.y}
		}
	}
	result := calculateGps(grid)
	return fmt.Sprint(result)
}
