package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
	"golang.org/x/exp/maps"
)

func createHappinessMap(input string) (happinessMap map[string]map[string]int) {
	happinessMap = make(map[string]map[string]int)
	parser := regexp.MustCompile(`(?P<from>\w+) would (?P<sign>lose|gain) (?P<amount>\d+) happiness units by sitting next to (?P<to>\w+)`)

	for _, line := range strings.Split(input, "\n") {
		parsedLine := util.CreateMap(line, parser)

		happiness, _ := strconv.Atoi(parsedLine["amount"])
		if parsedLine["sign"] == "lose" {
			happiness = happiness * -1
		}

		_, ok := happinessMap[parsedLine["from"]]
		if !ok {
			happinessMap[parsedLine["from"]] = make(map[string]int)
		}

		happinessMap[parsedLine["from"]][parsedLine["to"]] = happiness
	}
	return
}

func part1(input string) int {
	happinessMap := createHappinessMap(input)
	people := maps.Keys(happinessMap)
	permuationList := permutateList(people)

	maxPermHappiness := 0
	var bestPerm []string
	for _, currPerm := range permuationList {
		currPermHappiness := 0
		for i := 0; i < len(currPerm)-1; i++ {
			currPermHappiness += happinessMap[currPerm[i]][currPerm[i+1]]
			currPermHappiness += happinessMap[currPerm[i+1]][currPerm[i]]
		}
		currPermHappiness += happinessMap[currPerm[len(currPerm)-1]][currPerm[0]]
		currPermHappiness += happinessMap[currPerm[0]][currPerm[len(currPerm)-1]]

		if currPermHappiness > maxPermHappiness {
			maxPermHappiness = currPermHappiness
			bestPerm = currPerm
		}

	}

	fmt.Println(bestPerm)
	return maxPermHappiness
}

func part2(input string) int {
	happinessMap := createHappinessMap(input)
	people := maps.Keys(happinessMap)

	// Include "me" in a new happinessMap
	for _, from := range happinessMap {
		from["me"] = 0
	}
	happinessMap["me"] = make(map[string]int)
	for _, person := range people {
		happinessMap["me"][person] = 0
	}

	people = append(people, "me")
	permuationList := permutateList(people)

	maxPermHappiness := 0
	var bestPerm []string
	for _, currPerm := range permuationList {
		currPermHappiness := 0
		for i := 0; i < len(currPerm)-1; i++ {
			currPermHappiness += happinessMap[currPerm[i]][currPerm[i+1]]
			currPermHappiness += happinessMap[currPerm[i+1]][currPerm[i]]
		}
		currPermHappiness += happinessMap[currPerm[len(currPerm)-1]][currPerm[0]]
		currPermHappiness += happinessMap[currPerm[0]][currPerm[len(currPerm)-1]]

		if currPermHappiness > maxPermHappiness {
			maxPermHappiness = currPermHappiness
			bestPerm = currPerm
		}

	}

	fmt.Println(bestPerm)
	return maxPermHappiness
}

func permutateList(input []string) (output [][]string) {
	if len(input) == 0 {
		return [][]string{}
	}
	if len(input) == 1 {
		return [][]string{input}
	}

	output = [][]string{}

	for i, val := range input {
		copyInput := make([]string, len(input))
		copy(copyInput, input)
		remainingList := util.Remove(copyInput, i)
		for _, p := range permutateList(remainingList) {
			p2 := append([]string{val}, p...)
			output = append(output, p2)
		}
	}
	return output
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
