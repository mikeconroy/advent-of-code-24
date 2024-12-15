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
	bMax := machine.prize.x / machine.b.x
	multiple := 1
	firstXFound := -1
	for b := bMax; b >= 0; b -= multiple {

		bX := b * machine.b.x
		if bX > machine.prize.x {
			continue
		}
		if bX == machine.prize.x {
			//Machine can be solved with only B presses:
			if b*machine.b.y == machine.prize.y {
				fmt.Println("A:", 0, "B:", b)
				return (b * bCost)
			}
			continue
		}
		a := (machine.prize.x - bX) / machine.a.x
		aX := (machine.a.x * a)
		if (bX + aX) == machine.prize.x {
			bY := b * machine.b.y
			yAns := bY + a*machine.a.y

			if multiple == 1 && firstXFound == -1 {
				firstXFound = b
			} else if multiple != 1 {
				multiple = firstXFound - b
			}

			if yAns == machine.prize.y {
				return (b * bCost) + (a * aCost)
			}
		}
	}
	return 0
}

// aX + bX = prize.x
// aX = prize.x - bX
// a = (prize.x- bX) / X
// Substitute into Y Equation
// aY + bY = prize.y
// Y((prize.x+bX)/X) + bY = prize.y

func (machine *Machine) solveEquation() int {
	pX := machine.prize.x
	pY := machine.prize.y
	aX := machine.a.x
	bX := machine.b.x
	aY := machine.a.y
	bY := machine.b.y
	bCount := ((aX * pY) - (aY * pX)) / ((-aY * bX) + (aX * bY))
	aCount := (pX - (bX * bCount)) / aX
	xAns := (aCount * machine.a.x) + (bCount * machine.b.x)
	yAns := (aCount * machine.a.y) + (bCount * machine.b.y)
	if (xAns == machine.prize.x) && (yAns == machine.prize.y) {
		return (aCount * aCost) + (bCount * bCost)
	}
	return 0
}

func part1(input []string) string {
	machines := parseInput(input)
	tokens := 0
	for _, machine := range machines {
		tokens += machine.solveEquation()
	}
	return fmt.Sprint(tokens)
}

func part2(input []string) string {
	machines := parseInput(input)
	tokens := 0
	for _, machine := range machines {
		machine.prize.x += 10_000_000_000_000
		machine.prize.y += 10_000_000_000_000
		tokens += machine.solveEquation()
	}
	return fmt.Sprint(tokens)
}
