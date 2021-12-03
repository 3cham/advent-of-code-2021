package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetInput(day string, isTest bool) []string {

	cwd, _ := os.UserHomeDir()

	var inputFile string

	if !isTest {
		inputFile = fmt.Sprintf("%v/workspace/aoc2021/utils/input_files/%v.txt", cwd, day)
	} else {
		inputFile = fmt.Sprintf("%v/workspace/aoc2021/utils/input_files/%v_test.txt", cwd, day)
	}
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(content), "\n")
	return input
}
