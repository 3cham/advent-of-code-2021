package day3

import (
	"aoc2021/utils"
	"math"
)

func Part1(test bool) float64 {
	return part1(test)
}

func part1(test bool) float64 {
	input := utils.GetInput("3", test)
	N := len(input[0])

	countOne := make([]int, N)

	for _, line := range input {
		for pos, char := range line {
			if char == '1' {
				countOne[pos]++
			}
		}
	}

	value := 0.0
	for _, v := range countOne {
		value *= 2
		if v >= len(input)/2 {
			value += 1
		}
	}
	return (math.Pow(2, float64(N)) - value - 1) * value
}
