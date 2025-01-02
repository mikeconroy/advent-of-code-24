package day24

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day24/input")
	return part1(input), part2(input)
}

// Wires[x00] = true | 1
// Wires[x01] = false | 0
type Wires map[string]bool
type Gate struct {
	lhs, rhs string
	result   string
	op       func(bool, bool) bool
}

func xor(lhs, rhs bool) bool {
	return lhs != rhs
}

func or(lhs, rhs bool) bool {
	return lhs || rhs
}

func and(lhs, rhs bool) bool {
	return lhs && rhs
}

func parseInput(input []string) (Wires, []Gate) {
	wires := make(Wires)
	gates := []Gate{}

	for _, line := range input {
		if line == "" {
			continue
		}
		if strings.Contains(line, ":") {
			split := strings.Split(line, ": ")
			wire := split[0]
			val := split[1]
			if val == "1" {
				wires[wire] = true
			} else {
				wires[wire] = false
			}
		} else {
			split := strings.Split(line, " ")
			gate := Gate{
				lhs:    split[0],
				rhs:    split[2],
				result: split[4],
			}
			if split[1] == "XOR" {
				gate.op = xor
			} else if split[1] == "AND" {
				gate.op = and
			} else if split[1] == "OR" {
				gate.op = or
			}
			gates = append(gates, gate)
		}
	}
	return wires, gates
}

// [0] = Least Significant Bit
// [0,0,1] = 4
func boolsToInt(bools []bool) int {
	result := 0
	for index, bit := range bools {
		if bit {
			result += int(math.Pow(2, float64(index)))
		}
	}
	return result
}

func part1(input []string) string {
	wires, gates := parseInput(input)
	processedGates := make(map[string]bool)

	for len(processedGates) != len(gates) {
		for _, gate := range gates {
			if lhsVal, ok := wires[gate.lhs]; ok {
				if rhsVal, ok := wires[gate.rhs]; ok {
					wires[gate.result] = gate.op(lhsVal, rhsVal)
					processedGates[gate.result] = true
				}
			}
		}
	}

	zWires := []string{}
	for wire, _ := range wires {
		if strings.Contains(wire, "z") {
			zWires = append(zWires, wire)
		}
	}

	sort.Strings(zWires)
	zBools := []bool{}
	for _, wire := range zWires {
		zBools = append(zBools, wires[wire])
	}

	result := boolsToInt(zBools)

	return fmt.Sprint(result)
}

func part2(input []string) string {
	return fmt.Sprint(2)
}
