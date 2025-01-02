package day22

import (
	"fmt"
	"strconv"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day22/input")
	return part1(input), part2(input)
}

func generateSecret(secret int64) int64 {
	result := secret * 64
	secret ^= result
	secret %= 16_777_216

	result = secret / 32
	secret ^= result
	secret %= 16_777_216

	result = secret * 2048
	secret ^= result
	secret %= 16_777_216
	return secret
}

func nthSecret(secret int64, n int) int64 {
	for i := 0; i < n; i++ {
		secret = generateSecret(secret)
	}
	return secret
}

func part1(input []string) string {
	var sum int64 = 0
	for _, line := range input {
		if line == "" {
		}
		secret32, _ := strconv.Atoi(line)
		secret := int64(secret32)
		sum += nthSecret(secret, 2000)

	}

	return fmt.Sprint(sum)
}

func part2(input []string) string {
	return fmt.Sprint(2)
}
