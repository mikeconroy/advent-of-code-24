package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day13/input")
	return part1(input), part2(input)
}

type Point struct {
	x, y int
}
type Machine struct {
	a     Point
	b     Point
	prize Point
}

func parseInput(input []string) []Machine {
	machines := []Machine{}
	a, b := Point{}, Point{}
	for _, line := range input {
		if strings.Contains(line, "Button") {
			split := strings.Split(line, ", Y+")
			y, _ := strconv.Atoi(split[1])
			x, _ := strconv.Atoi(strings.Split(split[0], "X+")[1])
			if strings.Contains(line, "Button A:") {
				a = Point{x, y}
			} else {
				b = Point{x, y}
			}
		} else if strings.Contains(line, "Prize") {
			split := strings.Split(line, ", Y=")
			y, _ := strconv.Atoi(split[1])
			x, _ := strconv.Atoi(strings.Split(split[0], "X=")[1])
			machine := Machine{
				a:     a,
				b:     b,
				prize: Point{x, y},
			}
			machines = append(machines, machine)
		}
	}
	return machines
}

var aCost = 3
var bCost = 1

func (machine *Machine) solve() int {
	for b := 100; b >= 0; b-- {
		if b*machine.b.x > machine.prize.x {
			continue
		}
		if b*machine.b.x == machine.prize.x {
			//Machine can be solved with only B presses:
			if b*machine.b.y == machine.prize.y {
				return (b * bCost)
			}
			continue
		}

		for a := 100; a >= 0; a-- {
			xAns := b*machine.b.x + a*machine.a.x
			if xAns == machine.prize.x {
				yAns := b*machine.b.y + a*machine.a.y
				if yAns == machine.prize.y {
					return (b * bCost) + (a * aCost)
				}
			}
			if xAns < machine.prize.x {
				break
			}
		}

	}

	return 0
}

func part1(input []string) string {
	machines := parseInput(input)
	tokens := 0
	for _, machine := range machines {
		tokens += machine.solve()
	}
	return fmt.Sprint(tokens)
}

func part2(input []string) string {
	return fmt.Sprint(2)
}
