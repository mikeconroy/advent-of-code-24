package day12

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day12/input")
	return part1(input), part2(input)
}

type Grid = map[Point]rune
type Point struct {
	x int
	y int
}

func parseInput(input []string) Grid {
	grid := make(map[Point]rune)
	for y, line := range input {
		for x, char := range line {
			grid[Point{x, y}] = char
		}
	}
	return grid
}

type Region struct {
	id    rune
	cells map[Point]bool
}
type Regions = map[rune]Region

func (region *Region) getPerimeter(grid Grid) int {
	perimeter := 0
	for cell := range region.cells {
		perimeter += countPerimeter(region.id, cell, grid)
	}
	return perimeter
}

var deltas []Point = []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func countPerimeter(id rune, point Point, grid Grid) int {
	perimeter := 0
	for _, delta := range deltas {
		if grid[Point{point.x + delta.x, point.y + delta.y}] != id {
			perimeter += 1
		}
	}
	return perimeter
}

func getPointsInRegion(id rune, cell Point, grid Grid, points map[Point]bool) map[Point]bool {
	for _, delta := range deltas {
		neighbourPoint := Point{cell.x + delta.x, cell.y + delta.y}
		if _, ok := points[neighbourPoint]; ok {
			continue
		}
		if grid[neighbourPoint] == id {
			points[neighbourPoint] = true
			points = getPointsInRegion(id, neighbourPoint, grid, points)
		}
	}
	return points
}

func getRegions(grid Grid) []Region {
	regions := []Region{}
	cellsDone := make(map[Point]bool)
	for point, id := range grid {
		// If this cell has already been included in a region then move on.
		if _, ok := cellsDone[point]; ok {
			continue
		}
		region := Region{
			id:    id,
			cells: make(map[Point]bool),
		}
		region.cells[point] = true
		points := getPointsInRegion(id, point, grid, region.cells)
		for point := range points {
			region.cells[point] = true
			cellsDone[point] = true
		}
		regions = append(regions, region)
	}
	// printRegions(regions)
	return regions
}

func printRegions(regions []Region) {
	for _, region := range regions {
		fmt.Println("ID:", string(region.id), "Cells:", region.cells)
	}
}

func part1(input []string) string {
	grid := parseInput(input)
	regions := getRegions(grid)
	totalPrice := 0
	for _, region := range regions {
		totalPrice += region.getPerimeter(grid) * len(region.cells)
	}
	return fmt.Sprint(totalPrice)
}

func (region *Region) getSides(grid Grid) int {
	// Total Sides = Total Corners
	// Each corner is represented as the top left of the cell
	corners := make(map[Point]int)
	for point := range region.cells {
		//Top left is a corner if above and left are both empty or both in region.
		north := Point{point.x, point.y - 1}
		south := Point{point.x, point.y + 1}
		west := Point{point.x - 1, point.y}
		east := Point{point.x + 1, point.y}
		northeast := Point{point.x + 1, point.y - 1}
		northwest := Point{point.x - 1, point.y - 1}
		southeast := Point{point.x + 1, point.y + 1}
		southwest := Point{point.x - 1, point.y + 1}
		_, northInRegion := region.cells[north]
		_, southInRegion := region.cells[south]
		_, westInRegion := region.cells[west]
		_, eastInRegion := region.cells[east]
		_, northeastInRegion := region.cells[northeast]
		_, northwestInRegion := region.cells[northwest]
		_, southeastInRegion := region.cells[southeast]
		_, southwestInRegion := region.cells[southwest]

		if (!northInRegion && !westInRegion) || (northInRegion && westInRegion && !northwestInRegion) {
			//Account for corners meeting on diagonals...
			if !northInRegion && !westInRegion && northwestInRegion {
				corners[point] = 2
			} else {
				corners[point] = 1
			}
		}
		if (!northInRegion && !eastInRegion) || (northInRegion && eastInRegion && !northeastInRegion) {
			if !northInRegion && !eastInRegion && northeastInRegion {

				corners[Point{point.x + 1, point.y}] = 2
			} else {
				corners[Point{point.x + 1, point.y}] = 1
			}
		}
		if (!southInRegion && !westInRegion) || (southInRegion && westInRegion && !southwestInRegion) {
			if !southInRegion && !westInRegion && southwestInRegion {
				corners[Point{point.x, point.y + 1}] = 2
			} else {

				corners[Point{point.x, point.y + 1}] = 1
			}
		}
		if (!southInRegion && !eastInRegion) || (southInRegion && eastInRegion && !southeastInRegion) {
			if !southInRegion && !eastInRegion && southeastInRegion {
				corners[Point{point.x + 1, point.y + 1}] = 2
			} else {
				corners[Point{point.x + 1, point.y + 1}] = 1
			}
		}
	}

	totalCorners := 0
	for _, count := range corners {
		totalCorners += count
	}
	return totalCorners
}

// 837163 is too low
func part2(input []string) string {
	grid := parseInput(input)
	regions := getRegions(grid)

	totalPrice := 0
	for _, region := range regions {
		sides := region.getSides(grid)
		totalPrice += sides * len(region.cells)
	}
	return fmt.Sprint(totalPrice)
}
