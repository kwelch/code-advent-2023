package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/kwelch/code-advent-2023/utils"
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

var mapOrder = []string{
	"seed",
	"soil",
	"fertilizer",
	"water",
	"light",
	"temperature",
	"humidity",
	"location",
}

func part1(input string) int {
	var lines = strings.Split(input, "\n")
	var seedsList = utils.ParseNumbers(strings.Split(lines[0], ":")[1], " ")
	var maps = map[string][][]int{}
	var mapRegEx = regexp.MustCompile(`^[a-z]+-to-([a-z]+) map:$`)
	var currMap string
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		var matches = mapRegEx.FindStringSubmatch(line)
		if matches == nil {
			maps[currMap] = append(maps[currMap], utils.ParseNumbers(line, " "))
		} else {
			currMap = matches[1]

			maps[currMap] = [][]int{}
		}
	}

	var seedLocations = []int{}

	for _, seed := range seedsList {
		var currValue = seed
		for _, mapName := range mapOrder[1:] {
			var mapOverrides = maps[mapName]
			currValue = findValueInMap(currValue, mapOverrides)
		}
		seedLocations = append(seedLocations, currValue)
	}

	return utils.LowestValue(seedLocations)
}

func findValueInMap(value int, mapOverrides [][]int) int {
	for _, override := range mapOverrides {
		var valueStart = override[0]
		var searchStart = override[1]
		var searchRange = override[2]
		if value >= searchStart && value < searchStart+searchRange {
			var retVal = valueStart + (value - searchStart)
			return retVal
		}
	}
	return value
}

func part2(input string) int {
	var lines = strings.Split(input, "\n")
	var seedsData = utils.ParseNumbers(strings.Split(lines[0], ":")[1], " ")

	var maps = map[string][][]int{}
	var mapRegEx = regexp.MustCompile(`^[a-z]+-to-([a-z]+) map:$`)
	var currMap string
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		var matches = mapRegEx.FindStringSubmatch(line)
		if matches == nil {
			maps[currMap] = append(maps[currMap], utils.ParseNumbers(line, " "))
		} else {
			currMap = matches[1]

			maps[currMap] = [][]int{}
		}
	}

	var seedLocations = []int{}

	for i := 0; i < len(seedsData); i += 2 {
		var seedStart = seedsData[i]
		var seedRange = seedsData[i+1]
		var seedEnd = (seedStart + seedRange) - 1
		fmt.Printf("seedStart: %v, seedRange: %v, seedEnd: %v\n", seedStart, seedRange, seedEnd)
		for seedVal := seedStart; seedVal < seedEnd; seedVal++ {
			var currValue = seedVal
			for _, mapName := range mapOrder[1:] {
				var mapOverrides = maps[mapName]
				currValue = findValueInMap(currValue, mapOverrides)
			}
			seedLocations = append(seedLocations, currValue)
		}
	}

	return utils.LowestValue(seedLocations)
	return 0
}
