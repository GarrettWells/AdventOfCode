package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/GarrettWells/AdventOfCode/util"
)

func part1(input string) int {
	numberRegex := regexp.MustCompile(`-?\d+`)
	matches := numberRegex.FindAllString(input, -1)

	sum := 0
	for _, str := range matches {
		num, _ := strconv.Atoi(str)
		sum += num
	}
	return sum
}

func stripExcludes(input map[string]any) {
	for _, value := range input {
		switch value.(type) {
		case map[string]any:
			stripExcludes(value.(map[string]any))
		case []any:
			stripExcludesArr(value.([]any))
		case any:
			if value == "red" {
				clear(input)
			}
		}
	}
}

func stripExcludesArr(input []any) {
	for _, value := range input {
		switch value.(type) {
		case map[string]any:
			stripExcludes(value.(map[string]any))
		case []any:
			stripExcludesArr(value.([]any))
		case any:
			// We don't do anything in this case
		}
	}
}

func part2(input string) int {
	var result map[string]any
	json.Unmarshal([]byte(input), &result)

	stripExcludes(result)
	for key, value := range result {
		fmt.Println(key, value)
	}
	toReturn, _ := json.Marshal(result)
	return part1(string(toReturn))
}

func main() {
	input := util.ReadFile("input.txt")
	// fmt.Println(part1(input))
	fmt.Println(part2(input))
}
