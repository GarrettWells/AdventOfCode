package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
)

type City = string

type Pair struct {
	cityA, cityB City
}

type CityDistances map[Pair]uint

func (self *CityDistances) getDistance(cityA string, cityB string) uint {
	val, ok := (*self)[Pair{cityA, cityB}]
	if ok {
		return val
	}
	val, _ = (*self)[Pair{cityB, cityA}]
	return val
}

func part1(input string) uint {
	edgeList, cityList := populateCities(input)

	minDistance := uint(math.MaxUint)
	minDistancePath := []string{}

	for index, currCity := range cityList {
		distance, path := part1_helper(currCity, edgeList, removeElementFromArray[string](cityList, index))
		if minDistance > distance {
			minDistance = distance
			minDistancePath = path
		}
	}

	fmt.Println(minDistancePath)
	return minDistance
}

func part1_helper(currCity string, edgeList *CityDistances, cityList []string) (uint, []string) {
	if len(cityList) == 0 {
		return 0, []string{currCity}
	}

	minDistance := uint(math.MaxUint)
	minDistancePath := []string{}

	for index, nextCity := range cityList {
		distance, path := part1_helper(nextCity, edgeList, removeElementFromArray[string](cityList, index))
		distance += edgeList.getDistance(currCity, nextCity)
		if minDistance > distance {
			minDistance = distance
			minDistancePath = path
		}
	}
	return minDistance, append(minDistancePath, currCity)
}

func part2(input string) uint {
	edgeList, cityList := populateCities(input)

	maxDistance := uint(0)
	maxDistancePath := []string{}

	for index, currCity := range cityList {
		distance, path := part2_helper(currCity, edgeList, removeElementFromArray[string](cityList, index))
		if maxDistance < distance {
			maxDistance = distance
			maxDistancePath = path
		}
	}

	fmt.Println(maxDistancePath)
	return maxDistance
}

func part2_helper(currCity string, edgeList *CityDistances, cityList []string) (uint, []string) {
	if len(cityList) == 0 {
		return 0, []string{currCity}
	}

	maxDistance := uint(0)
	maxDistancePath := []string{}

	for index, nextCity := range cityList {
		distance, path := part2_helper(nextCity, edgeList, removeElementFromArray[string](cityList, index))
		distance += edgeList.getDistance(currCity, nextCity)
		if maxDistance < distance {
			maxDistance = distance
			maxDistancePath = path
		}
	}
	return maxDistance, append(maxDistancePath, currCity)
}

func removeElementFromArray[V any](arr []V, index int) []V {
	return append(slices.Clone(arr[:index]), arr[index+1:]...)
}

func populateCities(input string) (*CityDistances, []City) {
	cityList := []City{}
	edgeList := CityDistances{}

	edgeRegex := regexp.MustCompile(`(\w+) to (\w+) = (\d+)`)

	for _, line := range strings.Split(input, "\n") {
		res := edgeRegex.FindStringSubmatch(line)
		// Just looking at the first city will find all cities based on the given input
		if !slices.Contains(cityList, res[1]) {
			cityList = append(cityList, res[1])
		}
		if !slices.Contains(cityList, res[2]) {
			cityList = append(cityList, res[2])
		}
		distance, _ := strconv.ParseUint(res[3], 10, 8)
		edgeList[Pair{res[1], res[2]}] = uint(distance)
	}
	return &edgeList, cityList
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
