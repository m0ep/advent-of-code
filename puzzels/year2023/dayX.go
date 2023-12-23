package year2023

import (
	"bufio"
	"fmt"
	"os"
)

func RunDayX() {
	fmt.Println("Day X")
	file, err := os.Open("inputs/2023_04.txt")
	if nil != err {
		panic("Failed to ope file")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%04d: %s\n", len(line), line)
	}
}
