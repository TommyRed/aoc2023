package day01

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func findFirstAndLastPartOne(value string) (int, int) {
	var first, last int = -1, -1

	for _, char := range value {
		if unicode.IsDigit(char) {
			num := int(char - '0')

			if first == -1 {
				first = num
			}

			last = num
		}
	}

	return first, last
}

func FirstDayPartOne() {
	dir, _ := os.Getwd()
	content, _ := os.ReadFile(dir + "/days/day01/data.txt")
	input := strings.Split(string(content), "\n")

	sum := 0

	for _, str := range input {
		first, last := findFirstAndLastPartOne(str)
		sum = sum + (first*10 + last)
	}

	fmt.Println(sum)
}
