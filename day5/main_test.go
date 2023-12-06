package main

import (
	"testing"
)

var example = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func Test_day5_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  35,
		},
		{
			name:  "actual",
			input: input,
			want:  340994526,
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

func Test_day5_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  46,
		},
		{
			name:  "actual",
			input: input,
			want:  52210644,
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

func Test_findValueInMap(t *testing.T) {
	type args struct {
		value        int
		mapOverrides [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "seed-to-soil",
			args: args{
				value: 79,
				mapOverrides: [][]int{
					{50, 98, 2},
					{52, 50, 48},
				},
			},
			want: 81,
		},
		{
			name: "soil-to-fertilizer",
			args: args{
				value: 81,
				mapOverrides: [][]int{
					{0, 15, 37},
					{37, 52, 2},
					{39, 0, 15},
				},
			},
			want: 81,
		},
		{
			name: "fertilizer-to-water",
			args: args{
				value: 81,
				mapOverrides: [][]int{
					{49, 53, 8},
					{0, 11, 42},
					{42, 0, 7},
					{57, 7, 4},
				},
			},
			want: 81,
		},
		{
			name: "water-to-light",
			args: args{
				value: 81,
				mapOverrides: [][]int{
					{88, 18, 7},
					{18, 25, 70},
				},
			},
			want: 74,
		},
		{
			name: "light-to-temperature",
			args: args{
				value: 74,
				mapOverrides: [][]int{
					{45, 77, 23},
					{81, 45, 19},
					{68, 64, 13},
				},
			},
			want: 78,
		},
		{
			name: "temperature-to-humidity",
			args: args{
				value: 78,
				mapOverrides: [][]int{
					{0, 69, 1},
					{1, 0, 69},
				},
			},
			want: 78,
		},
		{
			name: "humidity-to-location",
			args: args{
				value: 78,
				mapOverrides: [][]int{
					{60, 56, 37},
					{56, 93, 4},
				},
			},
			want: 82,
		},
		{
			name: "soil-to-fertilizer",
			args: args{
				value: 14,
				mapOverrides: [][]int{
					{0, 15, 37},
					{37, 52, 2},
					{39, 0, 15},
				},
			},
			want: 53,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findValueInMap(tt.args.value, tt.args.mapOverrides); got != tt.want {
				t.Errorf("findValueInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
