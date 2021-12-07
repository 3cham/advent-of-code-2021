package day4

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func Part1(test bool) int {
	return part1(test)
}

func parse_input(input []string) ([]string, [][][]string) {
	numbers := strings.Split(input[0], ",")

	var matrices [][][]string

	for blockStart := 2; blockStart < len(input); blockStart += 6 {
		if blockStart >= len(input) {
			break
		}
		var m [][]string
		for row := 0; row < 5; row++ {

			m = append(m, strings.Split(strings.Trim(strings.Replace(input[blockStart+row], "  ", " ", -1), " "), " "))
		}
		matrices = append(matrices, m)
	}
	return numbers, matrices
}

func build_lookup_dict(m [][][]string) map[string][]string {
	l := make(map[string][]string)

	for mId, matric := range m {
		for rowId, row := range matric {
			for colId, cell := range row {
				if _, inserted := l[cell]; !inserted {
					l[cell] = []string{}
				}
				l[cell] = append(l[cell], fmt.Sprintf("%v_%v_%v", mId, rowId, colId))
			}
		}
	}
	return l
}

func calculatePoint(seenNumbers map[string]int, m [][]string, lastNumber string) int {
	l, _ := strconv.Atoi(lastNumber)
	s := 0
	for _, row := range m {
		for _, cell := range row {
			if _, existed := seenNumbers[cell]; !existed {
				c, _ := strconv.Atoi(cell)
				s += c
			}
		}
	}

	return l * s
}

func part1(test bool) int {
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
					return calculatePoint(seenNumbers, matrices[mId], i)
				}
			}
		}
	}

	return 0
}
