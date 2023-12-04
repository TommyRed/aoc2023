package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var settings = map[string]int{"red": 12, "green": 13, "blue": 14}

func isSubsetInvalid(index int, subset string) bool {
	if index == 0 {
		subset = strings.Split(subset, ":")[1]
	}

	draws := strings.Split(subset, ",")

	for _, draw := range draws {
		val := strings.Split(strings.TrimSpace(draw), " ")

		number, _ := strconv.Atoi(strings.TrimSpace(val[0]))
		color := strings.TrimSpace(val[1])

		if settings[color] < number {
			return true
		}
	}

	return false
}

func Day02PartOne() {
	dir, _ := os.Getwd()
	content, _ := os.ReadFile(dir + "/days/day02/data.txt")
	input := strings.Split(string(content), "\n")

	sumOfIndices := 0

	for i, game := range input {
		subsets := strings.Split(game, ";")

		invalidGame := false

		for j, subset := range subsets {
			invalidGame = invalidGame || isSubsetInvalid(j, subset)
		}

		if !invalidGame {
			sumOfIndices += i + 1
		}
	}

	fmt.Println("SUM", sumOfIndices)
}
