package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
)

func part1(input string) int {
	characterLiterals, charsIgnored := 0, 0

	escapeCharRegex := regexp.MustCompile(`(?:\\{2})|(?:\\")|(?:\\x[[:xdigit:]]{2})`)
	for _, line := range strings.Split(input, "\n") {

		lineLength := len(line)
		characterLiterals += lineLength

		output := escapeCharRegex.FindAllString(line, -1)

		// Every lines starts and ends in quotes
		charsIgnored += 2
		if output != nil {
			for _, match := range output {
				if len(match) > 2 {
					charsIgnored += 3
				} else {
					charsIgnored += 1
				}
			}
		}
	}

	return charsIgnored
}

func part2(input string) int {
	charsAdded := 0

	escapeCharRegex := regexp.MustCompile(`[\\"]`)
	for _, line := range strings.Split(input, "\n") {
		charsAdded += 2
		output := escapeCharRegex.FindAllString(line, -1)

		// Every lines starts and ends in quotes
		if output != nil {
			charsAdded += len(output)
		}
	}

	return charsAdded
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
