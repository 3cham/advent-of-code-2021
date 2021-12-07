package day1

import "aoc2021/utils"

func Solution(test bool) int {
	return day1(test)
}

func day1(test bool) int {
	input := utils.ToIntArr(utils.GetInput("1", test))
	count := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			count += 1
		}
	}
	return count
}
