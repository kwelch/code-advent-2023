package main

import (
	"fmt"
	"testing"
)

var example = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func Test_day1_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  142,
		},

		{
			name:  "actual",
			input: input,
			want:  54630,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

var example_2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func Test_day1_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example_2,
			want:  281,
		},

		{
			name:  "actual",
			input: input,
			want:  54770,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extract_numbers(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			"xtwone3four",
			[]string{"2", "1", "3", "4"},
		},
		{
			"ninesevensrzxkzpmgz8kcjxsbdftwoner",
			[]string{"9", "7", "8", "2", "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := extract_numbers(tt.input); fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("extract_numbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
