package day3

import (
	"aoc2021/utils"
)

func Part2(test bool) int {
	return part2(test)
}

func filters(input []string, pos int, isOxyGen bool) string {
	if len(input) == 1 {
		return input[0]
	}

	count := 0
	for _, line := range input {
		if line[pos] == '1' {
			count += 1
		}
	}

	var foundBit uint8 = '0'
	if count >= len(input)/2+len(input)%2 {
		foundBit = '1'
	} else {
		foundBit = '0'
	}

	if !isOxyGen {
		foundBit = '1' - foundBit + '0'
	}

	var newInput []string
	for _, line := range input {
		if line[pos] == foundBit {
			newInput = append(newInput, line)
		}
	}
	return filters(newInput, pos+1, isOxyGen)
}

func getVal(s string) int {
	value := 0
	for _, c := range s {
		value = value*2 + int(c-'0')
	}
	return value
}

func part2(test bool) int {
	input := utils.GetInput("3", test)

	x := getVal(filters(input, 0, true))
	y := getVal(filters(input, 0, false))
	return x * y
}
