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

type Sequence struct {
	one, two, three, four int64
}

// Prices are the ones digit of the secret numbers
// 123 -> 3
// 15887950 -> 0 (-3)
// 16495136 -> 6 (6)
// Sequence of 4 consecutive changes -> Then sells when it sees that sequence

// totalBananas[sequence]amount
// totalBananas[-2, 1,0,-1] = 9
// Stop once a sequence is found.
func part2(input []string) string {
	totals := make(map[Sequence]int64)
	for _, line := range input {
		if line == "" {
			continue
		}

		secret32, _ := strconv.Atoi(line)
		secret := int64(secret32)
		seq := Sequence{-10, -10, -10, -10}
		var prevPrice int64 = endDigit(secret)
		sequenceHit := make(map[Sequence]bool)
		for i := 0; i <= 2000; i++ {
			newSecret := generateSecret(secret)
			price := endDigit(newSecret)
			change := price - prevPrice

			newSeq := Sequence{seq.two, seq.three, seq.four, change}

			if i > 2 {
				if _, ok := sequenceHit[newSeq]; !ok {
					sequenceHit[newSeq] = true
					totals[newSeq] = totals[newSeq] + price
				}
			}

			secret = newSecret
			prevPrice = price
			seq = newSeq
		}

	}

	var highestPrice int64 = 0
	for _, v := range totals {
		if v > highestPrice {
			highestPrice = v
		}
	}
	return fmt.Sprint(highestPrice)
}

func endDigit(n int64) int64 {
	return n % 10
}
