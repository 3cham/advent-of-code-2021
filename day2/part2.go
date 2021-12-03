package day2

import (
	"aoc2021/utils"
	"strings"
)

func Part2(test bool) int {
	return part2(test)
}

func part2(test bool) int {
	input := utils.GetInput("2", test)
	x, y, aim := 0, 0, 0

	for _, instruction := range input {
		part := strings.Split(instruction, " ")
		command := part[0]
		value := utils.ToInt(part[1])

		switch command {
		case "forward":
			x += value
			y += aim * value
		case "up": aim -= value
		case "down": aim += value
		}
	}
	return x * y
}
