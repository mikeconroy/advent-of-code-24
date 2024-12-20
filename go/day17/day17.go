package day17

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day17/input")
	return part1(input), part2(input)
}

type Computer struct {
	a, b, c      int
	instructions []int
	pointer      int
	output       []int
}

func (comp *Computer) getComboOperand() int {
	operand := comp.instructions[comp.pointer+1]
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return comp.a
	case 5:
		return comp.b
	case 6:
		return comp.c
	case 7:
		fmt.Println("INVALID OPERAND 7")
	}
	return -1
}

func (comp *Computer) adv() {
	numerator := comp.a
	operand := float64(comp.getComboOperand())
	denominator := int(math.Pow(2, operand))
	comp.a = numerator / denominator
	comp.pointer += 2
}

func (comp *Computer) bxl() {
	b := comp.b
	operand := comp.instructions[comp.pointer+1]
	comp.b = b ^ operand
	comp.pointer += 2
}

func (comp *Computer) bst() {
	operand := comp.getComboOperand()
	comp.b = operand % 8
	comp.pointer += 2
}

func (comp *Computer) jnz() {
	if comp.a == 0 {
		comp.pointer += 2
		return
	}

	comp.pointer = comp.instructions[comp.pointer+1]
}

func (comp *Computer) bxc() {
	b := comp.b
	c := comp.c
	comp.b = b ^ c
	comp.pointer += 2
}

func (comp *Computer) out() {
	comp.output = append(comp.output, comp.getComboOperand()%8)
	comp.pointer += 2
}

func (comp *Computer) bdv() {
	numerator := comp.a
	operand := float64(comp.getComboOperand())
	denominator := int(math.Pow(2, operand))
	comp.b = numerator / denominator
	comp.pointer += 2
}

func (comp *Computer) cdv() {
	numerator := comp.a
	operand := float64(comp.getComboOperand())
	denominator := int(math.Pow(2, operand))
	comp.c = numerator / denominator
	comp.pointer += 2
}

func (comp *Computer) tick() bool {
	if comp.pointer >= len(comp.instructions) {
		return true
	}
	op := comp.instructions[comp.pointer]
	switch op {
	case 0:
		fmt.Println("ADV")
		comp.adv()
	case 1:
		fmt.Println("BXL")
		comp.bxl()
	case 2:
		fmt.Println("BST")
		comp.bst()
	case 3:
		fmt.Println("JNZ")
		comp.jnz()
	case 4:
		fmt.Println("BXC")
		comp.bxc()
	case 5:
		fmt.Println("OUT")
		comp.out()
	case 6:
		fmt.Println("BDV")
		comp.bdv()
	case 7:
		fmt.Println("CDV")
		comp.cdv()
	}

	return false
}

func (comp *Computer) print() {
	fmt.Println()
	fmt.Println("A:", comp.a, "B:", comp.b, "C:", comp.c, "Program:", comp.instructions, "Pointer:", comp.pointer)
	fmt.Println("--------------------------------------")
}

func parseInput(input []string) Computer {
	a, _ := strconv.Atoi(strings.Split(input[0], "A: ")[1])
	b, _ := strconv.Atoi(strings.Split(input[1], "B: ")[1])
	c, _ := strconv.Atoi(strings.Split(input[2], "C: ")[1])

	instructions := []int{}
	program := strings.Split(input[4], ": ")[1]
	for _, char := range strings.Split(program, ",") {
		instruction, _ := strconv.Atoi(char)
		instructions = append(instructions, instruction)
	}
	return Computer{
		a:            a,
		b:            b,
		c:            c,
		instructions: instructions,
		pointer:      0,
		output:       []int{},
	}
}

// 210172503 - wrong?
func part1(input []string) string {
	comp := parseInput(input)
	for !comp.tick() {
	}
	result := ""
	for _, val := range comp.output {
		result += fmt.Sprint(val)
	}
	return fmt.Sprint(result)
}

func part2(input []string) string {
	return fmt.Sprint(2)
}
