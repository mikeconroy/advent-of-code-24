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
		comp.adv()
	case 1:
		comp.bxl()
	case 2:
		comp.bst()
	case 3:
		comp.jnz()
	case 4:
		comp.bxc()
	case 5:
		comp.out()
	case 6:
		comp.bdv()
	case 7:
		comp.cdv()
	}

	return false
}

func (comp *Computer) equals(compB Computer) bool {
	if comp.a == compB.a && comp.b == compB.b && comp.c == compB.c {
		for i, instruction := range comp.instructions {
			if compB.instructions[i] != instruction {
				return false
			}
		}
		return true
	}
	return false
}

func (comp *Computer) print() {
	fmt.Println()
	fmt.Println("A:", comp.a, "B:", comp.b, "C:", comp.c, "Pointer:", comp.pointer)
	fmt.Println("Program:", comp.instructions)
	fmt.Println("Output:", comp.output)
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

type CacheKey struct {
	a, b, c, pointer int
}

func compareIntSlices(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}

/*
- 2 4 1 7 7 5 0 3 4 4 1 7 5 5 3 0

- 2 4	B = A % 8

- 1 7	B = B XOR 7

- 7 5	C = A / (2^B)

- 0 3	A = A / (2^3) = A / 8

- 4 4	B = B XOR C

- 1 7	B = B XOR 7

- 5 5	OUT B%8

  - 3 0	JNZ 0
    *

- 0 3 5 4 3 0

- 0 3	A = A / 8

- 5 4	OUT A % 8

  - 3 0	JNZ 0
    *
    a: 117440 Octal a: 0o345300
    [0]
    a: 14680 Octal a: 0o34530
    [0 3]
    a: 1835 Octal a: 0o3453
    [0 3 5]
    a: 229 Octal a: 0o345
    [0 3 5 4]
    a: 28 Octal a: 0o34
    [0 3 5 4 3]
    a: 3 Octal a: 0o3
    [0 3 5 4 3 0]
    a: 0 Octal a: 0o0

    Working backwards - each A always increases by a single Octal Digit
    whilst the previous digits remain the same.
    0o3 = [0]
    0o34 = [0, 3]
    0o345 = [0, 3, 4]
    0o3453 = [0, 3, 4, 5]
    0o34530 = [0, 3, 4, 5, 3]
    0o345300 = [0, 3, 4, 5, 3, 0]

    x := 0o3
    fmt.Printf("%O %d\n", x, x)
    x = (x << 3)
    x = x + 0o4
    fmt.Printf("%O %d\n", x, x)
    x = (x << 3)
    x = x + 0o5
    fmt.Printf("%O %d\n", x, x)
*/
func findA(origComp Computer, a int, index int) int {
	comp := origComp
	comp.a = a
	for !comp.tick() {
	}

	if len(comp.output) > index && comp.output[0] == comp.instructions[len(comp.instructions)-(index+1)] {
		if len(comp.output) == len(comp.instructions) {
			return a
		}
		newA := a << 3
		for i := 0o0; i < 0o10; i++ {
			res := findA(origComp, newA+i, index+1)
			if res != -1 {
				return res
			}
		}

	}

	return -1
}

func part2(input []string) string {
	comp := parseInput(input)

	for a := 0o0; a < 0o10; a++ {
		result := findA(comp, a, 0)
		if result != -1 {
			return fmt.Sprint(result)
		}
	}
	return fmt.Sprint(-1)

}
