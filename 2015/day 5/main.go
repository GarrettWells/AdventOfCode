package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
)

func part1(input string) int {
	numNiceStrings := 0
	for _, line := range strings.Split(input, "\n") {
		numVowels := 0
		hasDouble := false
		hasBadString := false
		prevChar := rune(0)

		vowels := []rune{'a', 'e', 'i', 'o', 'u'}
		badStrings := []string{"ab", "cd", "pq", "xy"}

		for _, currChar := range line {
			if slices.Contains(vowels, currChar) {
				numVowels++
			}
			if prevChar == currChar {
				hasDouble = true
			}
			if slices.Contains(badStrings, string(prevChar)+string(currChar)) {
				hasBadString = true
			}
			prevChar = currChar
		}

		if numVowels >= 3 && hasDouble && !hasBadString {
			numNiceStrings++
		}
	}
	return numNiceStrings
}

type Pair struct {
	prev, curr rune
}

func part2(input string) int {
	numNiceStrings := 0
	for _, line := range strings.Split(input, "\n") {
		// maps pairs to the index the second rune exist at
		pairs := map[Pair]int{}
		containsPairs := false
		containsSkipRepeat := false
		prevChar := rune(line[1])
		pairs[Pair{rune(line[0]), rune(line[1])}] = 1

		for index, currChar := range line[2:] {
			index += 2
			val, ok := pairs[Pair{prevChar, currChar}]

			if ok {
				if val+1 != index {
					containsPairs = true
				}
			} else {
				pairs[Pair{prevChar, currChar}] = index
			}

			if rune(line[index-2]) == currChar {
				containsSkipRepeat = true
			}

			prevChar = currChar
		}

		if containsPairs && containsSkipRepeat {
			numNiceStrings++
		}
	}
	return numNiceStrings
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
