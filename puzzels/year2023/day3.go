package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type SymbolPos struct {
	yOffset, xOffset int
	ch               int32
}

func RunDay3() {
	fmt.Println("Day 3")
	file, err := os.Open("inputs/2023_03.txt")
	if nil != err {
		panic("Failed to ope file")
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	symbols := findSymbols(lines)

	sum := 0
	gearRatioSum := 0
	for _, symbol := range symbols {
		numbers := findAdjacentNumbers(lines, symbol)
		for _, number := range numbers {
			sum += number
		}

		if '*' == symbol.ch && 1 < len(numbers) {
			gearRatio := 1
			for _, number := range numbers {
				gearRatio *= number
			}

			gearRatioSum += gearRatio
		}
	}

	fmt.Printf("sum=%d\n", sum)
	fmt.Printf("gearRatioSum=%d\n", gearRatioSum)
}

func findSymbols(lines []string) []SymbolPos {
	var result []SymbolPos

	for lineNo, line := range lines {
		for offset, element := range line {
			if '.' != element && !unicode.IsNumber(element) {
				result = append(result, SymbolPos{lineNo, offset, element})
			}
		}
	}

	return result
}

func findAdjacentNumbers(lines []string, symbol SymbolPos) []int {
	var result []int
	for y := maxInt(0, symbol.yOffset-1); y < minInt(symbol.yOffset+2, len(lines)); y++ {
		line := lines[y]

		lastFoundNumber := -1
		for x := maxInt(0, symbol.xOffset-1); x < minInt(symbol.xOffset+2, len(line)); x++ {
			if unicode.IsNumber(rune(line[x])) {
				number, _ := extractNumber(line, x)
				if lastFoundNumber != number {
					lastFoundNumber = number
					result = append(result, number)
				}
			}
		}
	}

	return result
}

func extractNumber(line string, pos int) (int, error) {
	left := -1
	for i := pos; i >= 0; i-- {
		left = i
		if !unicode.IsNumber(rune(line[i])) {
			left = i + 1
			break
		}
	}

	right := -1
	for i := pos; i < len(line); i++ {
		right = i
		if !unicode.IsNumber(rune(line[i])) {
			right = i - 1
			break
		}
	}

	return strconv.Atoi(line[left : right+1])
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
