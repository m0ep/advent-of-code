package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunDay9() {
	fmt.Println("Day 9")
	file, err := os.Open("inputs/2023_09.txt")
	if nil != err {
		panic("Failed to open file")
	}

	scanner := bufio.NewScanner(file)

	forward := 0
	backward := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := lineToNumbers(line)

		b, f := extrapolate(numbers)
		forward += numbers[len(numbers)-1] + f
		backward += numbers[0] - b
	}

	fmt.Printf("Answer 1: %d\n", forward)
	fmt.Printf("Answer 2: %d\n", backward)
}

func lineToNumbers(line string) []int {
	parts := strings.Split(line, " ")
	var result []int
	for _, val := range parts {
		number, _ := strconv.Atoi(val)
		result = append(result, number)
	}

	return result
}

func extrapolate(numbers []int) (int, int) {
	diffs := make([]int, len(numbers)-1)
	for i := 0; i < len(diffs); i++ {
		diffs[i] = numbers[i+1] - numbers[i]
	}

	if allZeros(diffs) {
		return 0, 0
	}

	b, f := extrapolate(diffs)

	f = diffs[len(diffs)-1] + f
	b = diffs[0] - b
	return b, f
}

func allZeros(numbers []int) bool {
	for _, n := range numbers {
		if 0 != n {
			return false
		}
	}

	return true
}
