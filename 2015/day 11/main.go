package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
)

func part1(input string) string {
	input = strings.Trim(input, " \n")
	invalidLetters := regexp.MustCompile("[iol]")
	for {
		has3ConsecutiveLetters := false
		input = incrementString(input)
		inputRunes := []rune(input)
		for i := 0; i < len(inputRunes)-2; i++ {
			if inputRunes[i] == inputRunes[i+1]-1 && inputRunes[i] == inputRunes[i+2]-2 {
				has3ConsecutiveLetters = true
				break
			}
		}
		if !has3ConsecutiveLetters {
			continue
		}

		if invalidLetters.MatchString(input) {
			continue
		}

		numPairs := 0
		for i := 0; i < len(inputRunes)-1; i++ {
			if inputRunes[i] == inputRunes[i+1] {
				numPairs++
				i++
			}
		}

		if numPairs < 2 {
			continue
		}

		return input

	}
}

func part2(input string) string {
	return part1(part1(input))
}

func incrementString(input string) string {
	arr := []rune(input)
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == rune('z') {
			arr[i] = rune('a')
		} else {
			arr[i]++
			return string(arr)
		}
	}
	panic("failure")
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
