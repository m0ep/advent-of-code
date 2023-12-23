package year2023

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	name  string
	left  string
	right string
}

func RunDay8() {
	fmt.Println("Day 8")
	file, err := os.Open("inputs/2023_08.txt")
	if nil != err {
		panic("Failed to open file")
	}

	foundSep := false
	var instructions []string
	nodeMap := make(map[string]Node)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if 0 == len(line) {
			foundSep = true
			continue
		}

		if foundSep {
			node := parseNode(line)
			nodeMap[node.name] = node
		} else {
			for _, ch := range line {
				instructions = append(instructions, string(ch))
			}
		}
	}

	//fmt.Printf("instructions: (%d) %v\n", len(instructions), instructions)
	//fmt.Printf("nodes: %v\n", nodeMap)

	task1(nodeMap, instructions)
	task2(nodeMap, instructions)
}

func parseNode(line string) Node {
	pattern := regexp.MustCompile(`^(\w+) = \((\w+), (\w+)\)$`)
	result := pattern.FindStringSubmatch(line)
	return Node{result[1], result[2], result[3]}
}

func task1(nodeMap map[string]Node, instructions []string) {
	steps := 0
	currentNode := nodeMap["AAA"]
	for {
		instruction := instructions[steps%len(instructions)]
		steps++

		if "L" == instruction {
			currentNode = nodeMap[currentNode.left]
		} else {
			currentNode = nodeMap[currentNode.right]
		}

		if "ZZZ" == currentNode.name {
			fmt.Printf("Answer 1: %d\n", steps)
			break
		}
	}
}

func task2(nodeMap map[string]Node, instructions []string) {
	steps := 0

	currentNodes := filterValues(nodeMap, func(v Node) bool {
		return strings.HasSuffix(v.name, "A")
	})

	nSteps := make([]int, len(currentNodes))
	for !allMatch(currentNodes, reachDestination) {
		instruction := instructions[steps%len(instructions)]

		for i, currentNode := range currentNodes {
			if reachDestination(currentNode) {
				continue
			}

			if "L" == instruction {
				currentNodes[i] = nodeMap[currentNode.left]
			} else {
				currentNodes[i] = nodeMap[currentNode.right]
			}

			if reachDestination(currentNodes[i]) {
				nSteps[i] = steps + 1
			}
		}
		steps++
	}

	fmt.Printf("Answer 2: %d\n", steps)
	fmt.Println(nSteps)

	// peeked from https://github.com/MarkusFreitag/advent-of-code/blob/master/puzzles/year2023/day8/day8.go
	// :(
	fmt.Println(LCM(nSteps[0], nSteps[1], nSteps[2:]...))
}

func filterValues[K comparable, V any](m map[K]V, fn func(v V) bool) []V {
	res := make([]V, 0)
	for _, v := range m {
		if fn(v) {
			res = append(res, v)
		}
	}

	return res
}

func allMatch[V any](items []V, fn func(v V) bool) bool {
	for _, v := range items {
		if !fn(v) {
			return false
		}
	}

	return true
}

func reachDestination(node Node) bool {
	return strings.HasSuffix(node.name, "Z")
}

func GCD[T int](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM[T int](a, b T, nums ...T) T {
	result := a * b / GCD(a, b)
	for _, num := range nums {
		result = LCM(result, num)
	}
	return result
}
