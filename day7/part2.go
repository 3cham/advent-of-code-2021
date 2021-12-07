package day7

import (
	"aoc2021/utils"
	"sort"
)

func Part2(test bool) int {
	return part2(test)
}

func cost(positions []int, pos int) int {

	sum := 0
	for _, p := range positions {
		x := p - pos
		if x < 0 {
			x *= -1
		}

		sum += x * (x + 1) / 2
	}

	return sum
}

func part2(test bool) int {

	input := utils.GetInput("7", test)

	positions := parse_position(input[0])

	sort.Ints(positions)

	sum := 0
	for _, p := range positions {
		sum += p
	}

	pos := sum / len(positions)

	if cost(positions, pos) < cost(positions, pos + 1) {
		return cost(positions, pos)
	}

	return cost(positions, pos + 1)
}