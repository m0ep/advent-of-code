package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// wrong: 1219390,

type Card struct {
	id            int
	winingNumbers []int
	numbers       []int
}

func RunDay4() {
	fmt.Println("Day 4")
	file, err := os.Open("inputs/2023_04.txt")
	if nil != err {
		panic("Failed to ope file")
	}

	var cards []Card
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		card := parseCard(line)
		//fmt.Print(card)
		fmt.Printf("  %d - %d | %d\n", card.id, len(card.winingNumbers), len(card.numbers))
		cards = append(cards, card)
	}

	// task 1
	pointsSum := 0
	for _, card := range cards {
		count := countWiningNumbers(card)
		points := calcPoints(count)
		pointsSum += points
	}
	fmt.Printf("pointsSum = %d\n", pointsSum)

	// task 2

	var workingStack []Card
	copyCount := make([]int, len(cards))
	for i, card := range cards {
		copyCount[i] = 1

		count := countWiningNumbers(card)
		if 0 < count {
			workingStack = append(workingStack, card)
		}
	}

	for 0 != len(workingStack) {
		var copyStack []Card
		for i := 0; i < len(workingStack); i++ {
			count := countWiningNumbers(workingStack[i])
			if 0 < count {
				cardIndex := workingStack[i].id - 1
				copies := extractCopies(cards, cardIndex, count)
				for _, cpy := range copies {
					copyCount[cpy.id-1]++
					copyStack = append(copyStack, cpy)
				}
			}
		}

		if 0 < len(copyStack) {
			workingStack = make([]Card, len(copyStack))
			copy(workingStack, copyStack)
		} else {
			break
		}
	}

	sumCopies := 0
	for _, count := range copyCount {
		sumCopies += count
	}

	fmt.Printf("sumCopies = %d\n", sumCopies)
}

func parseCard(line string) Card {
	lineParts := strings.Split(line, ":")
	cardId, _ := strconv.Atoi(strings.TrimSpace(lineParts[0][5:]))

	cardParts := strings.Split(lineParts[1], "|")

	var winingNumbers []int
	winingNumbersPart := strings.TrimSpace(cardParts[0])
	winingNumbersSplit := strings.Split(winingNumbersPart, " ")
	for i := 0; i < len(winingNumbersSplit); i++ {
		trimmed := strings.TrimSpace(winingNumbersSplit[i])
		if 0 == len(trimmed) {
			continue
		}

		number, _ := strconv.Atoi(trimmed)
		winingNumbers = append(winingNumbers, number)
	}

	var numbers []int
	numberPart := strings.TrimSpace(cardParts[1])
	numbersSlit := strings.Split(numberPart, " ")
	for i := 0; i < len(numbersSlit); i++ {
		trimmed := strings.TrimSpace(numbersSlit[i])
		if 0 == len(trimmed) {
			continue
		}

		number, _ := strconv.Atoi(trimmed)
		numbers = append(numbers, number)
	}

	return Card{cardId, winingNumbers, numbers}
}

func countWiningNumbers(card Card) int {
	count := 0
	for w := 0; w < len(card.winingNumbers); w++ {
		winningNumber := card.winingNumbers[w]
		for n := 0; n < len(card.numbers); n++ {
			number := card.numbers[n]
			if winningNumber == number {
				count++
			}
		}
	}

	return count
}

func calcPoints(count int) int {

	if 0 == count {
		return 0
	} else if 1 == count {
		return 1
	}

	points := 1
	for i := 1; i < count; i++ {
		points *= 2
	}

	return points
}

func extractCopies(cards []Card, index int, count int) []Card {
	if 0 < count {
		var copies []Card
		for i := index + 1; i < len(cards); i++ {
			copies = append(copies, cards[i])
			if count == len(copies) {
				return copies
			}
		}

		return copies
	}

	return []Card{}
}
