package day6

import (
	"aoc2021/utils"
	"strconv"
	"strings"
)

func Part1(test bool) int {
	days := 80
	return part1(test, days)
}

func parse_state(s string) []int {
	var result []int
	parts := strings.Split(s, ",")
	for _, s := range parts {
		x, _ := strconv.Atoi(s)
		result = append(result, x)
	}
	return result
}

func part1(test bool, days int) int {
	input := utils.GetInput("6", test)
	row := parse_state(input[0])
	state := make(map[int]int)

	for i := 0; i <= 8; i ++ {
		state[i] = 0
	}

	for _, x := range row {
        state[x] += 1
    }

	for day := 0; day < days; day ++ {
		new_state := make(map[int]int)

		for i := 1; i <= 8; i++ {
			new_state[i - 1] = state[i]
		}
		new_state[8] = state[0]
		new_state[6] += state[0]

		state = new_state
	}

	total := 0
	for key := range state {
		total += state[key]
	}
	return total
}
