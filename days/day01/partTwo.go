package day01

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var words = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func findFirstAndLastPartTwo(value string) (int, int) {
	var first, last int = -1, -1

	for i, char := range value {
		if unicode.IsDigit(char) {
			num := int(char - '0')

			if first == -1 {
				first = num
			}

			last = num
		} else {
			for j := 3; j < 6; j++ {
				if i+j <= len(value) {
					a, exists := words[value[i:i+j]]

					if exists {
						if first == -1 {
							first = a
						}

						last = a
					}
				}
			}
		}
	}

	return first, last
}

func FirstDayPartTwo() {
	dir, _ := os.Getwd()
	content, _ := os.ReadFile(dir + "/days/day01/data.txt")

	input := strings.Split(string(content), "\n")

	sum := 0

	for _, str := range input {
		first, last := findFirstAndLastPartTwo(str)
		sum = sum + (first*10 + last)
	}

	fmt.Println(sum)
}
