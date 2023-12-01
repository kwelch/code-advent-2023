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
	var sum int = 0
	for _, line := range strings.Split(input, "\n") {
		var calibrationNumbers = extract_calibration_numbers(line)

		if lineValue, err := strconv.Atoi(strings.Join(calibrationNumbers, "")); err == nil {
			sum += lineValue
		} else {
			panic(fmt.Errorf("could not convert string to int; %v (%s)", calibrationNumbers, err))
		}
	}
	return sum
}

func part2(input string) int {
	var sum int = 0
	for _, line := range strings.Split(input, "\n") {
		line = convert_words_in_string_with_numbers(line)
		var calibrationNumbers = extract_calibration_numbers(line)

		if lineValue, err := strconv.Atoi(strings.Join(calibrationNumbers, "")); err == nil {
			sum += lineValue
		} else {
			panic(fmt.Errorf("could not convert string to int; %v (%s)", calibrationNumbers, err))
		}
	}
	return sum
}

var wordNumbers = map[string]string{
	"one":   "one1one",
	"two":   "two2two",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}

func extract_calibration_numbers(input string) []string {
	regexNum := regexp.MustCompile(`\d`)

	foundNumbers := regexNum.FindAllString(input, -1)

	return []string{foundNumbers[0], foundNumbers[len(foundNumbers)-1]}
}

func extract_numbers(input string) []string {
	input = convert_words_in_string_with_numbers(input)
	regexNum := regexp.MustCompile(`\d`)

	return regexNum.FindAllString(input, -1)
}

func convert_words_in_string_with_numbers(input string) string {
	for word := range wordNumbers {
		input = strings.ReplaceAll(input, word, wordNumbers[word])
	}
	return input
}
