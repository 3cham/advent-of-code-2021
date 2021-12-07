package day5

import (
	"aoc2021/utils"
)

func Part2(test bool) int {
	return part2(test)
}

func part2(test bool) int {
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
		} else {
			if line[1] == line[3] {
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
			} else {
				startx, starty, endx, endy := line[0], line[1], line[2], line[3]

				for startx != endx {
					if _, existed := count[startx]; !existed {
						count[startx] = make(map[int]int)
					}

					if _, existed := count[startx][starty]; !existed {
						count[startx][starty] = 0
					}

					count[startx][starty] += 1

					if startx < endx {
						startx += 1
					} else {
						startx -= 1
					}

					if starty < endy {
						starty += 1
					} else {
						starty -= 1
					}
				}

				if _, existed := count[endx]; !existed {
					count[endx] = make(map[int]int)
				}

				if _, existed := count[endx][endy]; !existed {
					count[endx][endy] = 0
				}

				count[endx][endy] += 1
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
