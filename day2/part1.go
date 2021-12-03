package day2

import (
	"aoc2021/utils"
	"strings"
)

func Part1(test bool) int {
	return part1(test)
}

func part1(test bool) int {
	input := utils.GetInput("2", test)
	x, y := 0, 0

	for _, instruction := range input {
		part := strings.Split(instruction, " ")
		command := part[0]
		value := utils.ToInt(part[1])

		switch command {
		case "forward": x += value
		case "up": y -= value
		case "down": y += value
		}
	}
	return x * y
}
