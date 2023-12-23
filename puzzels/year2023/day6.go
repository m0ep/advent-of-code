package year2023

import (
	"fmt"
)

func RunDay6() {
	fmt.Println("Day 6")

	/*
		races := [][]int{
			{7, 9},
			{15, 40},
			{30, 200},
		}*/

	/*
		races := [][]int{
			{40, 215},
			{92, 1064},
			{97, 1505},
			{90, 1100},
		}*/

	races := [][]int{
		{40929790, 215106415051100},
	}

	answerTask1 := 1
	for _, race := range races {
		time := race[0]
		record := race[1]

		options := 0
		for pressMs := 0; pressMs <= time; pressMs++ {
			timeLeft := time - pressMs
			dist := pressMs * timeLeft

			if record < dist {
				options++
			}
			//fmt.Printf("pressMs:=%d timeLeft:%s dist::%d\n", pressMs, timeLeft, dist)
		}

		answerTask1 *= options
	}

	fmt.Printf("Anser: %d\n", answerTask1)
}
