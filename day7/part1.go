package day7

import (
	"aoc2021/utils"
	"math"
	"sort"
	"strings"
)

func Part1(test bool) float64 {
	return part1(test)
}

func parse_position(s string) []int {
	parts := strings.Split(s, ",")
	return utils.ToIntArr(parts)
}

func part1(test bool) float64 {

	input := utils.GetInput("7", test)

	positions := parse_position(input[0])

	sort.Ints(positions)
	position := positions[len(positions) / 2]

	total := 0.0
	for _, p := range positions {
		total += math.Abs(float64(p) - float64(position))
	}
	return total
}
