package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
)

func part1(input string) int {
	input = strings.Trim(input, " \n")
	for i := 1; i <= 40; i++ {
		matchingRune := rune(input[0])
		numMatches := 1
		nextString := ""
		for _, currRune := range input[1:] {
			if matchingRune == currRune {
				numMatches++
			} else {
				nextString += strconv.Itoa(numMatches) + string(matchingRune)
				matchingRune = currRune
				numMatches = 1
			}
		}
		nextString += strconv.Itoa(numMatches) + string(matchingRune)
		input = nextString
	}
	return len(input)
}

func part2(input string) int {
	input = strings.Trim(input, " \n")
	for i := 1; i <= 50; i++ {
		matchingRune := rune(input[0])
		numMatches := 1
		var nextString strings.Builder
		for _, currRune := range input[1:] {
			if matchingRune == currRune {
				numMatches++
			} else {
				nextString.WriteString(fmt.Sprintf("%d%s", numMatches, string(matchingRune)))
				matchingRune = currRune
				numMatches = 1
			}
		}
		nextString.WriteString(fmt.Sprintf("%d%s", numMatches, string(matchingRune)))
		input = nextString.String()
	}
	return len(input)
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
