package main

import (
	"fmt"

	"github.com/GarrettWells/AdventOfCode/util"
	mapset "github.com/deckarep/golang-set/v2"
)

type House struct {
	x int
	y int
}

type Santa struct {
	x int
	y int
}

func (santa *Santa) Up() {
	santa.y++
}
func (santa *Santa) Down() {
	santa.y--
}
func (santa *Santa) Right() {
	santa.x++
}
func (santa *Santa) Left() {
	santa.x--
}

func part2(input string) int {
	set := mapset.NewSet[House]()
	santa, roboSanta := Santa{0, 0}, Santa{0, 0}
	var currSanta = &santa

	set.Add(House{0, 0})

	for _, currChar := range input {
		if currChar == '^' {
			currSanta.Up()
		} else if currChar == '>' {
			currSanta.Right()
		} else if currChar == 'v' {
			currSanta.Down()
		} else if currChar == '<' {
			currSanta.Left()
		}

		fmt.Println(*currSanta)

		set.Add(House{currSanta.x, currSanta.y})
		if currSanta == &santa {
			currSanta = &roboSanta
		} else {
			currSanta = &santa
		}
	}

	return set.Cardinality()
}

func part1(input string) int {
	set := mapset.NewSet[House]()

	currX, currY := 0, 0
	set.Add(House{currX, currY})

	for _, currChar := range input {
		if currChar == '^' {
			currY++
		} else if currChar == '>' {
			currX++
		} else if currChar == 'v' {
			currY--
		} else if currChar == '<' {
			currX--
		}

		set.Add(House{currX, currY})
	}

	return set.Cardinality()

}

func main() {
	input := util.ReadFile("input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))

}
