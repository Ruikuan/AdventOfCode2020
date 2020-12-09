package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isValid(ring *[25]int, val int) bool {
	dic := make(map[int]bool, 25)
	for _, v := range *ring {
		if dic[val-v] {
			return true
		}
		dic[v] = true
	}
	return false
}

func part1() {
	ring := [25]int{}
	for idx, val := range strings.Split(input, "\n") {
		ival, _ := strconv.Atoi(val)
		if idx >= 25 {
			if !isValid(&ring, ival) {
				fmt.Println(ival)
				return
			}
			ring[idx%25] = ival

		} else {
			ring[idx] = ival
		}
	}
}

func main() {
	part1()
}
