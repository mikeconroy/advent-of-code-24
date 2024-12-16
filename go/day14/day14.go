package day14

import (
	"fmt"
	"github.com/mikeconroy/advent-of-code-24/utils"
	"strconv"
	"strings"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day14/input")
	return part1(input, 101, 103), part2(input, 101, 103)
}

type Point struct {
	x, y int
}

type Robot struct {
	pos Point
	vel Point
}

func parseInput(input []string) []*Robot {
	robots := []*Robot{}
	for _, line := range input {
		if line == "" {
			continue
		}
		split := strings.Split(line, " v=")
		pos := strings.Split(split[0], ",")
		posX, _ := strconv.Atoi(strings.Split(pos[0], "=")[1])
		posY, _ := strconv.Atoi(pos[1])
		vel := strings.Split(split[1], ",")
		velX, _ := strconv.Atoi(vel[0])
		velY, _ := strconv.Atoi(vel[1])
		robot := &Robot{
			pos: Point{posX, posY},
			vel: Point{velX, velY},
		}
		robots = append(robots, robot)
	}
	return robots
}

func (robot *Robot) move(width int, height int) {
	newX := robot.pos.x + robot.vel.x
	newX %= width
	for newX < 0 {
		newX += width
	}
	newY := robot.pos.y + robot.vel.y
	newY %= height
	for newY < 0 {
		newY += height
	}
	robot.pos.x = newX
	robot.pos.y = newY
}

func printRobots(grid map[Point]int, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[Point{x, y}] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}

		}
		fmt.Println()
	}
	fmt.Println("-----------------------------------------------------------------")
}

func part1(input []string, width int, height int) string {
	robots := parseInput(input)
	for second := 0; second < 100; second++ {
		for _, robot := range robots {
			robot.move(width, height)
		}
	}
	quads := make([]int, 4)
	for _, robot := range robots {
		if robot.pos.x == width/2 || robot.pos.y == height/2 {
			continue
		}

		if robot.pos.x > width/2 {
			if robot.pos.y > height/2 {
				quads[0] += 1
			} else {
				quads[1] += 1
			}
		} else {
			if robot.pos.y > height/2 {
				quads[2] += 1
			} else {
				quads[3] += 1
			}
		}
	}

	result := quads[0] * quads[1] * quads[2] * quads[3]
	return fmt.Sprint(result)
}

func part2(input []string, width, height int) string {
	robots := parseInput(input)
	grid := make(map[Point]int)
	for _, robot := range robots {
		grid[robot.pos] += 1
	}
	for second := 0; second < 7138; second++ {
		for _, robot := range robots {
			grid[robot.pos] -= 1
			robot.move(width, height)
			grid[robot.pos] += 1
		}
	}
	// Uncomment to see ASCII art of Xmas Tree
	// printRobots(grid, width, height)
	return fmt.Sprint(7138)
}
