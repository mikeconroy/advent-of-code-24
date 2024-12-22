package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day18/input")
	return part1(input, 70, 1024), part2(input, 70, 1024)
}

type Grid = [][]rune

func loadGrid(in []string, size int, bytes int) Grid {
	grid := Grid{}
	for y := 0; y <= size; y++ {
		row := []rune{}
		for x := 0; x <= size; x++ {
			row = append(row, '.')
		}
		grid = append(grid, row)
	}

	for i, coords := range in {
		if i >= bytes || coords == "" {
			continue
		}
		split := strings.Split(coords, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		grid[y][x] = '#'
	}

	return grid
}

func printGrid(grid Grid) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
}

type Point struct {
	x, y int
}

type PointCost struct {
	point Point
	cost  int
}

func walk(grid Grid, start Point, end Point) int {
	visited := make(map[Point]int)
	// Changing to a PriorityQueue/MinHeap is more efficient
	// as guarantees we visit the lowest values first
	toVisit := []PointCost{{start, 0}}

	for len(toVisit) > 0 {
		// fmt.Println(toVisit)
		// fmt.Println(visited)
		curr := toVisit[0]
		toVisit = toVisit[1:]
		cost, alreadyVisited := visited[curr.point]

		if !alreadyVisited || curr.cost < cost {
			visited[curr.point] = curr.cost
			if curr.point == end {
				continue
			}

			neighbours := getNeighbours(curr, grid, visited)
			toVisit = append(toVisit, neighbours...)
		}
	}
	return visited[end]
}

func getNeighbours(point PointCost, grid Grid, visited map[Point]int) []PointCost {
	dirs := []Point{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}

	neighbours := []PointCost{}

	for _, dir := range dirs {
		neighbour := PointCost{
			point: Point{point.point.x + dir.x, point.point.y + dir.y},
			cost:  point.cost + 1,
		}
		if neighbour.point.y >= 0 && neighbour.point.y < len(grid) &&
			neighbour.point.x >= 0 && neighbour.point.x < len(grid[0]) {
			if grid[neighbour.point.y][neighbour.point.x] != '#' {
				if val, ok := visited[neighbour.point]; !ok || val < point.cost+1 {
					neighbours = append(neighbours, neighbour)
				}
			}
		}
	}

	return neighbours
}

func part1(input []string, size int, bytes int) string {
	grid := loadGrid(input, size, bytes)
	start := Point{0, 0}
	end := Point{size, size}
	// printGrid(grid)
	res := walk(grid, start, end)
	return fmt.Sprint(res)
}

func part2(input []string, size int, bytes int) string {
	grid := loadGrid(input, size, bytes)
	start := Point{0, 0}
	end := Point{size, size}
	res := walk(grid, start, end)
	return fmt.Sprint(res)
}
