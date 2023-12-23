package year2023

import (
	"AdventOfCode2023/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Mapping struct {
	destStart int
	srcStart  int
	len       int
}

func RunDay5() {
	fmt.Println("Day 5")
	file, err := os.Open("inputs/2023_05.txt")
	if nil != err {
		panic("Failed to ope file")
	}

	state := 0
	var seeds []int
	var seedToSoil []Mapping
	var soilToFertilizer []Mapping
	var fertilizerToWater []Mapping
	var waterToLight []Mapping
	var lightToTemperature []Mapping
	var tempToHumid []Mapping
	var humidToLocation []Mapping

	// # Task 1
	// Parsing
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if 0 == state {
			if strings.HasPrefix(line, "seeds:") {
				seeds = toIntArray(line)
				state++
			}
		} else if 1 == state && "seed-to-soil map:" == line {
			state++
		} else if 2 == state {
			if 0 == len(line) {
				state++
				continue
			}

			numbers := toIntArray(line)
			seedToSoil = append(seedToSoil,
				Mapping{numbers[0], numbers[1], numbers[2]})
		} else if 3 == state && "soil-to-fertilizer map:" == line {
			state++
		} else if 4 == state {
			if 0 == len(line) {
				state++
				continue
			}
			numbers := toIntArray(line)
			soilToFertilizer = append(soilToFertilizer,
				Mapping{numbers[0], numbers[1], numbers[2]})
		} else if 5 == state && "fertilizer-to-water map:" == line {
			state++
		} else if 6 == state {
			if 0 == len(line) {
				state++
				continue
			}
			numbers := toIntArray(line)
			fertilizerToWater = append(fertilizerToWater,
				Mapping{numbers[0], numbers[1], numbers[2]})
		} else if 7 == state && "water-to-light map:" == line {
			state++
		} else if 8 == state {
			if 0 == len(line) {
				state++
				continue
			}
			numbers := toIntArray(line)
			waterToLight = append(waterToLight,
				Mapping{numbers[0], numbers[1], numbers[2]})
		} else if 9 == state && "light-to-temperature map:" == line {
			state++
		} else if 10 == state {
			if 0 == len(line) {
				state++
				continue
			}
			numbers := toIntArray(line)
			lightToTemperature = append(lightToTemperature,
				Mapping{numbers[0], numbers[1], numbers[2]})
		} else if 11 == state && "temperature-to-humidity map:" == line {
			state++
		} else if 12 == state {
			if 0 == len(line) {
				state++
				continue
			}
			numbers := toIntArray(line)
			tempToHumid = append(tempToHumid,
				Mapping{numbers[0], numbers[1], numbers[2]})
		} else if 13 == state && "humidity-to-location map:" == line {
			state++
		} else if 14 == state {
			if 0 == len(line) {
				state++
				continue
			}
			numbers := toIntArray(line)
			humidToLocation = append(humidToLocation,
				Mapping{numbers[0], numbers[1], numbers[2]})
		}
	}

	fmt.Printf("seeds:                   %v\n", seeds)
	fmt.Printf("seed-to-soil:            %v\n", seedToSoil)
	fmt.Printf("soil-to-fertilizer:      %v\n", soilToFertilizer)
	fmt.Printf("fertilizer-to-water:     %v\n", fertilizerToWater)
	fmt.Printf("water-to-light:          %v\n", waterToLight)
	fmt.Printf("light-to-temperature:    %v\n", lightToTemperature)
	fmt.Printf("temperature-to-humidity: %v\n", tempToHumid)
	fmt.Printf("humidity-to-location:    %v\n", humidToLocation)

	minLocation := math.MaxInt
	for s := 0; s < len(seeds); s++ {
		soil := findDestId(seeds[s], seedToSoil)
		fertilizer := findDestId(soil, soilToFertilizer)
		water := findDestId(fertilizer, fertilizerToWater)
		light := findDestId(water, waterToLight)
		temp := findDestId(light, lightToTemperature)
		humid := findDestId(temp, tempToHumid)
		location := findDestId(humid, humidToLocation)

		minLocation = utils.MinInt(minLocation, location)
	}

	fmt.Printf("task1 answer: %d\n", minLocation)

	// Task 2

	minLocation = math.MaxInt
	for index := 0; index < len(seeds); index += 2 {
		for offset := 0; offset < seeds[index+1]; offset++ {
			soil := findDestId(seeds[index]+offset, seedToSoil)
			fertilizer := findDestId(soil, soilToFertilizer)
			water := findDestId(fertilizer, fertilizerToWater)
			light := findDestId(water, waterToLight)
			temp := findDestId(light, lightToTemperature)
			humid := findDestId(temp, tempToHumid)
			location := findDestId(humid, humidToLocation)

			minLocation = utils.MinInt(minLocation, location)
		}
	}

	fmt.Printf("task2 answer: %d\n", minLocation)
}

func toIntArray(line string) []int {
	var result []int
	start := 0
	for i := 0; i < len(line); i++ {
		if !unicode.IsNumber(rune(line[i])) {
			if 0 < i-start {
				number, _ := strconv.Atoi(line[start:i])
				result = append(result, number)
			}

			start = i + 1
		} else if i == len(line)-1 {
			number, _ := strconv.Atoi(line[start:])
			result = append(result, number)
		}
	}

	return result
}

func findDestId(srcId int, mappings []Mapping) int {
	for m := 0; m < len(mappings); m++ {
		mapping := mappings[m]
		if mapping.srcStart <= srcId && srcId < mapping.srcStart+mapping.len {
			return mapping.destStart + (srcId - mapping.srcStart)
		}
	}

	return srcId
}
