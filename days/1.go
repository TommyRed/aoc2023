package days

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func findFirstAndLast(value string) (int, int) {
	var first, last int = -1, -1

	for _, char := range value {
		if unicode.IsDigit(char) {
			if first == -1 {
				first = int(char - '0')
			}

			last = int(char - '0')
		}
	}

	return first, last
}

func FirstDay() {
	dir, _ := os.Getwd()
	content, _ := os.ReadFile(dir + "/data/1.txt")

	inpt := strings.Split(string(content), "\n")

	sum := 0

	for _, str := range inpt {
		first, last := findFirstAndLast(str)
		sum = sum + (first*10 + last)
	}

	fmt.Println(sum)
}
