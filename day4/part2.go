package day4

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func Part2(test bool) int {
	return part2(test)
}

func part2(test bool) int {
	input := utils.GetInput("4", test)

	numbers, matrices := parse_input(input)

	lookup := build_lookup_dict(matrices)

	var countRow, countCol = make(map[string]int), make(map[string]int)
	for mId, _ := range matrices {
		for id := 0; id < 5; id++ {
			key := fmt.Sprintf("%v_%v", mId, id)
			countRow[key] = 0
			countCol[key] = 0
		}
	}

	seenNumbers := make(map[string]int)
	won := make(map[int]int)

	for _, i := range numbers {
		if seenNumbers[i] > 0 {
			continue
		}
		seenNumbers[i] = 1

		positions, existed := lookup[i]
		if existed {
			for _, position := range positions {
				ids := strings.Split(position, "_")
				idR := fmt.Sprintf("%s_%s", ids[0], ids[1])
				idC := fmt.Sprintf("%s_%s", ids[0], ids[2])

				mId, _ := strconv.Atoi(ids[0])

				countRow[idR] += 1
				countCol[idC] += 1
				if countRow[idR] == 5 || countCol[idC] == 5 {
					if _, alreadyWon := won[mId]; !alreadyWon {
						won[mId] = 1
						if len(won) == len(matrices) {
							return calculatePoint(seenNumbers, matrices[mId], i)
						}
					}
				}
			}
		}
	}

	return 0
}
