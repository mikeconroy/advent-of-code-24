package day22

import (
	"testing"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func TestDay22SecretGenerator(t *testing.T) {
	expect := int64(15_887_950)
	secret := int64(123)
	result := generateSecret(secret)
	if result != expect {
		t.Fatal("Day 22 - Secret Generator output should be", expect, "but got", result)
	}

	expect = 16_495_136
	secret = 15_887_950
	result = generateSecret(secret)
	if result != expect {
		t.Fatal("Day 22 - Secret Generator output should be", expect, "but got", result)
	}

	expect = 5_908_254
	secret = 123
	result = nthSecret(secret, 10)

	if result != expect {
		t.Fatal("Day 22 - Secret Generator output should be", expect, "but got", result)
	}

}

func TestDay22Part1(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "37327623"
	result := part1(input)
	if result != expect {
		t.Fatal("Day 22 - Part 1 output should be", expect, "but got", result)
	}
}

func TestDay22Part2(t *testing.T) {
	input := utils.ReadFileIntoSlice("input_test")
	expect := "2"
	result := part2(input)
	if result != expect {
		t.Fatal("Day 22 - Part 2 output should be", expect, "but got", result)
	}
}
