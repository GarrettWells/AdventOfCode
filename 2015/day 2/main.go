package main

import (
	"fmt"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
)

func part2(input string) int {
	ribbonLength := 0
	for _, element := range strings.Split(input, "\n") {
		var l, w, h int
		fmt.Sscanf(element, "%dx%dx%d", &l, &w, &h)
		ribbonLength += l * w * h

		maxNum := l
		if w > maxNum {
			maxNum = w
		}
		if h > maxNum {
			maxNum = h
		}
		ribbonLength += 2*w + 2*l + 2*h - 2*maxNum
	}

	return ribbonLength
}

func part1(input string) int {
	squareFeet := 0
	for _, element := range strings.Split(input, "\n") {
		var l, w, h int
		fmt.Sscanf(element, "%dx%dx%d", &l, &w, &h)
		squareFeet += 2*l*w + 2*l*h + 2*w*h

		maxNum := l
		if w > maxNum {
			maxNum = w
		}
		if h > maxNum {
			maxNum = h
		}
		squareFeet += l * w * h / maxNum
	}
	return squareFeet
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
