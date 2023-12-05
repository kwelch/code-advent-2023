package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
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

func part1(input string) int {
	var foundNumbers = []int{}
	var lines = strings.Split(input, "\n")
	for lineIndex, line := range lines {
		var numRegex = regexp.MustCompile(`\d+`)
		for _, num := range numRegex.FindAllStringIndex(line, -1) {
			if _, found := searchAroundForMatch(lines, lineIndex, num, isCharSymbol); found {
				var number, _ = strconv.Atoi(line[num[0]:num[1]])
				foundNumbers = append(foundNumbers, number)
			}
		}
	}
	return utils.SumValues(foundNumbers)
}

func getCharAt(lines []string, lineIndex int, charIndex int) byte {
	if lineIndex < 0 || lineIndex >= len(lines) {
		return '.'
	}
	if charIndex < 0 || charIndex >= len(lines[lineIndex]) {
		return '.'
	}

	return lines[lineIndex][charIndex]
}

func isCharSymbol(char byte) bool {
	if _, err := strconv.Atoi(string(char)); err == nil {
		return false
	}
	return char != '.'
}

func searchAroundForMatch(lines []string, lineIndex int, num []int, matchFunc func(char byte) bool) (foundIndex []int, found bool) {
	for checkLineIndex := lineIndex - 1; checkLineIndex <= lineIndex+1; checkLineIndex++ {
		for checkCharIndex := num[0] - 1; checkCharIndex < num[1]+1; checkCharIndex++ {
			if currentCheckChar := getCharAt(lines, checkLineIndex, checkCharIndex); matchFunc(currentCheckChar) {
				return []int{checkLineIndex, checkCharIndex}, true
			}
		}
	}
	return []int{}, false
}

func part2(input string) int {
	var total = 0
	var foundGears = map[string][]int{}
	var lines = strings.Split(input, "\n")
	for lineIndex, line := range lines {
		var numRegex = regexp.MustCompile(`\d+`)
		for _, num := range numRegex.FindAllStringIndex(line, -1) {
			if symbolIndex, found := searchAroundForMatch(lines, lineIndex, num, func(char byte) bool { return char == '*' }); found {
				var number, _ = strconv.Atoi(line[num[0]:num[1]])
				var symbolGearIndex = fmt.Sprintf("%d,%d", symbolIndex[0], symbolIndex[1])
				foundGears[symbolGearIndex] = append(foundGears[symbolGearIndex], number)
			}
		}
	}

	for _, gearNumbers := range foundGears {
		if len(gearNumbers) > 1 {
			total += utils.PowerValues(gearNumbers)
		}
	}
	return total
}
