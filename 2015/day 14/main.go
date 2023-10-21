package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
)

type Reindeer struct {
	speed, flightTime, restTime, distanceTraveled, timer, points int
	name                                                         string
	isResting                                                    bool
}

func createReindeer(input string) (output []Reindeer) {
	reindeerRegex := regexp.MustCompile(`(?P<reindeer>\w+) can fly (?P<speed>\d+) km/s for (?P<flightTime>\d+) seconds, but then must rest for (?P<restTime>\d+) seconds.`)
	for _, line := range strings.Split(input, "\n") {
		params := util.CreateMap(line, reindeerRegex)
		var reindeer Reindeer
		reindeer.speed, _ = strconv.Atoi(params["speed"])
		reindeer.flightTime, _ = strconv.Atoi(params["flightTime"])
		reindeer.restTime, _ = strconv.Atoi(params["restTime"])
		reindeer.name, _ = params["name"]

		output = append(output, reindeer)
	}
	return
}

func reindeerStep(reindeer *Reindeer) {
	if reindeer.isResting {
		if reindeer.timer >= reindeer.restTime {
			reindeer.timer = 0
			reindeer.isResting = !reindeer.isResting
		}
	} else {
		if reindeer.timer >= reindeer.flightTime {
			reindeer.timer = 0
			reindeer.isResting = !reindeer.isResting
		}
	}

	if !reindeer.isResting {
		reindeer.distanceTraveled += reindeer.speed
	}
	reindeer.timer++

}

func part1(input string) int {
	reindeers := createReindeer(input)
	for i := 0; i < 2503; i++ {
		for idx, _ := range reindeers {
			reindeerStep(&reindeers[idx])
		}
	}

	maxDistance := 0
	for _, reindeer := range reindeers {
		if maxDistance < reindeer.distanceTraveled {
			maxDistance = reindeer.distanceTraveled
		}
	}
	return maxDistance
}

func part2(input string) int {
	reindeers := createReindeer(input)
	for i := 0; i < 2503; i++ {
		for idx, _ := range reindeers {
			reindeerStep(&reindeers[idx])
		}

		maxDistance := 0
		var reindeerWithMaxDistance *Reindeer
		for _, reindeer := range reindeers {
			if maxDistance < reindeer.distanceTraveled {
				maxDistance = reindeer.distanceTraveled
				reindeerWithMaxDistance = &reindeer
			}
		}
		reindeerWithMaxDistance.points++
	}

	mostPoints := 0
	for _, reindeer := range reindeers {
		if mostPoints < reindeer.points {
			mostPoints = reindeer.points
		}
	}
	return mostPoints
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
