package main

import (
	"fmt"
	"strings"
)

// https://adventofcode.com/2020/day/3

type downArg struct {
	right int
	step1 bool
}

func part2() {
	args := []downArg{{right: 1, step1: true}, {right: 3, step1: true}, {right: 5, step1: true}, {right: 7, step1: true}, {right: 1, step1: false}}
	sum := 1

	for _, arg := range args {
		index := 0
		takeNext := true
		localSum := 0
		for _, line := range strings.Split(input, "\n") {
			if takeNext {
				localIndex := index % len(line)
				if line[localIndex] == '#' {
					localSum++
				}
				index += arg.right
			}
			takeNext = takeNext == arg.step1
		}
		sum *= localSum
	}
	fmt.Println(sum)
}

func main() {
	part2()
}
