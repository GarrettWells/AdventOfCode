package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
)

type LightAction int

const (
	Toggle LightAction = iota
	Off
	On
)

func createLightAction(input string) LightAction {
	if strings.Contains(input, "turn on") {
		return On
	}
	if strings.Contains(input, "turn off") {
		return Off
	}
	return Toggle
}

func part1(input string) int {
	lights := [1000][1000]bool{}

	actionRegex := regexp.MustCompile(`^([^\d]*)`)
	coordsRegex := regexp.MustCompile(`(?P<x1>\d*),(?P<y1>\d*) through (?P<x2>\d*),(?P<y2>\d*)`)
	for _, line := range strings.Split(input, "\n") {
		action := createLightAction(actionRegex.FindString(line))
		coords := coordsRegex.FindStringSubmatch(line)
		x1, _ := strconv.Atoi(coords[coordsRegex.SubexpIndex("x1")])
		x2, _ := strconv.Atoi(coords[coordsRegex.SubexpIndex("x2")])
		y1, _ := strconv.Atoi(coords[coordsRegex.SubexpIndex("y1")])
		y2, _ := strconv.Atoi(coords[coordsRegex.SubexpIndex("y2")])

		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				if action == Toggle {
					lights[i][j] = !lights[i][j]
				} else if action == On {
					lights[i][j] = true
				} else if action == Off {
					lights[i][j] = false
				}
			}
		}
	}

	numOn := 0

	for _, row := range lights {
		for _, light := range row {
			if light {
				numOn++
			}
		}
	}

	return numOn
}

func part2(input string) int {
	lights := [1000][1000]int{}

	actionRegex := regexp.MustCompile(`^([^\d]*)`)
	coordsRegex := regexp.MustCompile(`(?P<x1>\d*),(?P<y1>\d*) through (?P<x2>\d*),(?P<y2>\d*)`)
	for _, line := range strings.Split(input, "\n") {
		action := createLightAction(actionRegex.FindString(line))
		coords := coordsRegex.FindStringSubmatch(line)
		x1, _ := strconv.Atoi(coords[coordsRegex.SubexpIndex("x1")])
		x2, _ := strconv.Atoi(coords[coordsRegex.SubexpIndex("x2")])
		y1, _ := strconv.Atoi(coords[coordsRegex.SubexpIndex("y1")])
		y2, _ := strconv.Atoi(coords[coordsRegex.SubexpIndex("y2")])

		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				if action == Toggle {
					lights[i][j] += 2
				} else if action == On {
					lights[i][j] += 1
				} else if action == Off {
					if lights[i][j] > 0 {
						lights[i][j] -= 1
					}
				}
			}
		}
	}

	totalBrightness := 0

	for _, row := range lights {
		for _, light := range row {
			totalBrightness += light
		}
	}

	return totalBrightness
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
