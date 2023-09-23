package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strconv"

	"github.com/GarrettWells/AdventOfCode/util"
)

func part1(input string) int {
	i := 1
	for {
		md5Input := []byte(input + strconv.Itoa(i))
		md5Hash := md5.Sum(md5Input)

		// fmt.Printf("%x\n", md5Hash)
		if bytes.Compare(md5Hash[:3], []byte{0x00, 0x00, 0x0F}) <= 0 {
			fmt.Printf("%s\n", md5Input)
			fmt.Printf("%x\n", md5Hash)
			return i
		}
		// if i%1000 == 0 {
		// 	fmt.Println(i)
		// }
		i++
	}
}

func part2(input string) int {
	i := 1
	for {
		md5Input := []byte(input + strconv.Itoa(i))
		md5Hash := md5.Sum(md5Input)

		// fmt.Printf("%x\n", md5Hash)
		if bytes.Compare(md5Hash[:3], []byte{0x00, 0x00, 0x00}) <= 0 {
			fmt.Printf("%s\n", md5Input)
			fmt.Printf("%x\n", md5Hash)
			return i
		}
		// if i%1000 == 0 {
		// 	fmt.Println(i)
		// }
		i++
	}
}

func main() {
	input := util.ReadFile("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
