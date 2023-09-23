package main

import (
	"fmt"

	"github.com/GarrettWells/AdventOfCode/util"
)

func part1(input string) int {
	currFloor := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			currFloor++
		}
		if input[i] == ')' {
			currFloor--
		}
	}
	return currFloor
}

// getPositionOfFloor gets the position (1 indexed) of the character in the provided string that results in the indicated floor when following the advent of code day 1 challenge's rule
func part2(input string, desiredFloor int) int {
	currFloor := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			currFloor++
		}
		if input[i] == ')' {
			currFloor--
		}
		if currFloor == desiredFloor {
			return i + 1
		}
	}
	return -1
}

func main() {
	// Read input from command line arguments
	input := util.ReadFile("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input, -1))
}
