package day5

import (
	"aoc2021/utils"
	"strconv"
	"strings"
)

func Part1(test bool) int {
	return part1(test)
}

func parse_point(s string) (int, int) {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x, y
}

func parse_row(input string) []int {
	items := strings.Split(input, " -> ")
	start := strings.Trim(items[0], " ")
	end := strings.Trim(items[1], " ")

	start_x, start_y := parse_point(start)
	end_x, end_y := parse_point(end)

	return []int{start_x, start_y, end_x, end_y}
}

func part1(test bool) int {
	input := utils.GetInput("5", test)
	count := make(map[int]map[int]int)

	for _, row := range input {
		line := parse_row(row)

		// consider only horizontal or vertical lines
		if line[0] == line[2] {
			start, end := line[1], line[3]

			if line[1] > line[3] {
				start = line[3]
				end = line[1]
			}

			for i := start; i <= end; i++ {
				if _, existed := count[line[0]]; !existed {
					count[line[0]] = make(map[int]int)
				}

				if _, existed := count[line[0]][i]; !existed {
					count[line[0]][i] = 0
				}

				count[line[0]][i] += 1
			}
		} else if line[1] == line[3] {
			start, end := line[0], line[2]

			if line[0] > line[2] {
				start = line[2]
				end = line[0]
			}

			for i := start; i <= end; i++ {
				if _, existed := count[i]; !existed {
					count[i] = make(map[int]int)
				}

				if _, existed := count[i][line[1]]; !existed {
					count[i][line[1]] = 0
				}

				count[i][line[1]] += 1
			}
		}
	}

	total := 0

	for x := range count {
		for y := range count[x] {
			if count[x][y] > 1 {
				total += 1
			}
		}
	}

	return total
}
