package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseNumbers(input string, deliminator string) []int {
	var numbers = []int{}
	for _, line := range strings.Split(input, deliminator) {
		if len(line) > 0 {
			if number, err := strconv.Atoi(line); err == nil {
				numbers = append(numbers, number)
			} else {
				panic(fmt.Errorf("could not convert string to int; %v (%s)", line, err))
			}
		}
	}
	return numbers
}
