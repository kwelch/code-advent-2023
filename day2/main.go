package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	var cubes_max_values = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var validGameIds = []int{}

	for _, line := range strings.Split(input, "\n") {
		var gameIdRegEx = regexp.MustCompile(`Game (\d+): (.*)`)
		var gameId, _ = strconv.Atoi(gameIdRegEx.FindStringSubmatch(line)[1])
		var rawGameData = gameIdRegEx.FindStringSubmatch(line)[2]
		if isGameValid(rawGameData, cubes_max_values) {
			validGameIds = append(validGameIds, gameId)
		}
	}

	return sumValues(validGameIds)
}

func isGameValid(rawGameData string, cubes_max_values map[string]int) bool {
	for _, cubes := range strings.Split(rawGameData, ";") {
		for _, cube := range strings.Split(cubes, ",") {
			var cubeData = strings.Split(strings.TrimSpace(cube), " ")
			var cubeColor = cubeData[1]
			if cubeCount, err := strconv.Atoi(cubeData[0]); err == nil {
				if cubeCount > cubes_max_values[cubeColor] {
					return false
				}
			} else {
				panic(fmt.Errorf("could not convert string to int; %v (%s)", cubeData, err))
			}
		}
	}
	return true
}

func sumValues(values []int) int {
	var total int = 0
	for _, value := range values {
		total += value
	}
	return total
}

func powerValues(values []int) int {
	var total int = 0
	for _, value := range values {
		total *= value
	}
	return total
}

func part2(input string) int {
	var powerPerSet = []int{}

	for _, line := range strings.Split(input, "\n") {
		var gameIdRegEx = regexp.MustCompile(`Game (\d+): (.*)`)
		var rawGameData = gameIdRegEx.FindStringSubmatch(line)[2]
		var maxCubeColors = map[string]int{}

		for _, cubes := range strings.Split(rawGameData, ";") {
			for _, cube := range strings.Split(cubes, ",") {
				var cubeData = strings.Split(strings.TrimSpace(cube), " ")
				var cubeColor = cubeData[1]
				if cubeCount, err := strconv.Atoi(cubeData[0]); err == nil {
					if cubeCount > maxCubeColors[cubeColor] {
						maxCubeColors[cubeColor] = cubeCount
					}
				} else {
					panic(fmt.Errorf("could not convert string to int; %v (%s)", cubeData, err))
				}
			}
		}

		var gamePower int = 0
		for _, cubeCount := range maxCubeColors {
			if gamePower == 0 {
				gamePower = cubeCount
				continue
			}
			gamePower *= cubeCount
		}
		powerPerSet = append(powerPerSet, gamePower)
	}

	return sumValues(powerPerSet)
}
