package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func RunDay1() {
	fmt.Println("Day 1")
	file, err := os.Open("inputs/2023_01.txt")
	if nil != err {
		panic("Failed to ope file")
	}

	var sum int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		left := -1
		for i := 0; i < len(line); i++ {
			number := getNumber(line, i)
			if 0 < number {
				left = number
				break
			}
		}

		if 0 > left {
			panic(-1)
		}

		right := -1
		for i := len(line) - 1; i >= 0; i-- {
			number := getNumber(line, i)
			if 0 < number {
				right = number
				break
			}
		}

		if 0 > right {
			panic(-1)
		}

		coord := left*10 + right
		sum += coord
	}

	fmt.Println("Sum=", sum)
}

func getNumber(line string, offset int) int {
	numberWords := [10]string{"_", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	if unicode.IsNumber(rune(line[offset])) {
		number, _ := strconv.Atoi(string(line[offset]))
		return number
	}

	for i := 0; i < len(numberWords); i++ {
		if strings.HasPrefix(line[offset:], numberWords[i]) {
			return i
		}
	}

	return -1
}
