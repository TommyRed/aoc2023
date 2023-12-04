package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSubsetInvalidS(index int, subset string, sum map[string]int) {
	if index == 0 {
		subset = strings.Split(subset, ":")[1]
	}

	draws := strings.Split(subset, ",")

	for _, draw := range draws {
		val := strings.Split(strings.TrimSpace(draw), " ")

		number, _ := strconv.Atoi(strings.TrimSpace(val[0]))
		color := strings.TrimSpace(val[1])

		if sum[color] < number {
			sum[color] = number
		}

	}
}

func Day02PartTwo() {
	dir, _ := os.Getwd()
	content, _ := os.ReadFile(dir + "/days/day02/data.txt")
	input := strings.Split(string(content), "\n")

	result := 0

	for _, game := range input {
		sum := map[string]int{"red": 0, "green": 0, "blue": 0}

		subsets := strings.Split(game, ";")

		for j, subset := range subsets {
			isSubsetInvalidS(j, subset, sum)
		}

		result += sum["red"] * sum["green"] * sum["blue"]
	}

	fmt.Println("result", result)
}
