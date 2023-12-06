package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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
	var myPoints = []int{}
	for _, line := range strings.Split(input, "\n") {
		var cardData = strings.Split(line, ":")[1]
		var cards = strings.Split(cardData, "|")
		var winningNumbers = getMatches(utils.ParseNumbers(cards[0], " "), utils.ParseNumbers(cards[1], " "))
		myPoints = append(myPoints, calculatePoints(winningNumbers))
	}
	return utils.SumValues(myPoints)
}

func calculatePoints(winningNumbers []int) int {
	var total = 0
	if len(winningNumbers) > 2 {
		total = int(math.Pow(2, float64(len(winningNumbers)-1)))
	} else {
		total = len(winningNumbers)
	}
	return total
}

func getMatches(winningNumbers []int, myNumbers []int) []int {
	var matches = []int{}
	for _, myNumber := range myNumbers {
		for _, winningNumber := range winningNumbers {
			if myNumber == winningNumber {
				matches = append(matches, myNumber)
			}
		}
	}
	return matches
}

func part2(input string) int {
	var playedCards = map[int]int{}
	for index, line := range strings.Split(input, "\n") {
		var cardData = strings.Split(line, ":")[1]
		var cardId = index + 1
		var copies = playedCards[cardId]
		var cards = strings.Split(cardData, "|")
		var winningNumbers = getMatches(utils.ParseNumbers(cards[0], " "), utils.ParseNumbers(cards[1], " "))
		playedCards[cardId] += 1
		for idx := range winningNumbers {
			playedCards[cardId+(idx+1)] += copies + 1
		}
	}

	fmt.Println(playedCards)
	var total = 0
	for _, wins := range playedCards {
		total += wins
	}
	return total
}
