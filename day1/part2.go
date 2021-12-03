package day1

import "aoc2021/utils"

func Part2(test bool) int {
	return day2(test)
}

func consum(input []int, start, length int) int {
    var sum int
	if start + length > len(input) {
		return 0
	}
	for i := start; i < start+length; i++ {
        sum += input[i]
    }

    return sum
}

func day2(test bool) int {
	input := utils.ToIntArr(utils.GetInput("1", test))
	count := 0
	length := 3

	for i := 1; i < len(input); i ++ {
		if consum(input, i, length) > consum(input, i - 1, length) {
			count += 1
		}
	}
	return count
}
