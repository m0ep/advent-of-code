package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeSet struct {
	red, green, blue int
}

type Game struct {
	id    int
	draws []CubeSet
}

func RunDay2() {
	fmt.Println("Day 2")
	file, err := os.Open("inputs/2023_02.txt")
	if nil != err {
		panic("Failed to ope file")
	}

	testRed := 12
	testGreen := 13
	testBlue := 14

	var gameIdSum int = 0
	var powerSum = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game := parseLine(line)

		// Task 1
		impossible := false
		for i := 0; i < len(game.draws); i++ {
			draw := game.draws[i]
			if draw.red > testRed || draw.green > testGreen || draw.blue > testBlue {
				impossible = true
				break
			}
		}

		if !impossible {
			gameIdSum += game.id
		}

		// Task 2
		minCubeSet := minimalNumberOfCubesInBag(game)
		power := minCubeSet.red * minCubeSet.green * minCubeSet.blue
		powerSum += power
	}

	fmt.Printf("ID Sum     = %d\n", gameIdSum)
	fmt.Printf("Powser Sum = %d\n", powerSum)
}

func parseLine(line string) Game {
	lineParts := strings.Split(line, ":")
	if 2 != len(lineParts) {
		panic("invalid line")
	}

	gameId, _ := strconv.Atoi(lineParts[0][5:])

	var draws []CubeSet
	gameParts := strings.Split(lineParts[1], ";")
	for i := 0; i < len(gameParts); i++ {
		draw := strings.TrimSpace(gameParts[i])

		red := 0
		green := 0
		blue := 0
		drawParts := strings.Split(draw, ",")
		for e := 0; e < len(drawParts); e++ {
			cube := strings.TrimSpace(drawParts[e])

			cubeParts := strings.Split(cube, " ")
			cubeCount, _ := strconv.Atoi(cubeParts[0])
			cubeColor := cubeParts[1]
			if "red" == cubeColor {
				red = cubeCount
			} else if "green" == cubeColor {
				green = cubeCount
			} else if "blue" == cubeColor {
				blue = cubeCount
			}
		}

		draws = append(draws, CubeSet{red, green, blue})
	}

	return Game{gameId, draws}
}

func minimalNumberOfCubesInBag(game Game) CubeSet {
	red := 0
	green := 0
	blue := 0

	for i := 0; i < len(game.draws); i++ {
		draw := game.draws[i]

		if red < draw.red {
			red = draw.red
		}

		if green < draw.green {
			green = draw.green
		}

		if blue < draw.blue {
			blue = draw.blue
		}
	}

	return CubeSet{red, green, blue}
}
