package main

import (
	"bufio"
	"fmt"
	"strings"
)

// https://adventofcode.com/2020/day/5

func part1() {
	highest := 0
	for _, seat := range strings.Split(input, "\n") {
		rl, rr := 0, 127
		for _, v := range seat[0:7] {
			if rune(v) == 'F' {
				rr = (rl + rr) / 2
			} else {
				rl = (rl+rr)/2 + 1
			}
		}
		cl, cr := 0, 7
		for _, v := range seat[7:] {
			if rune(v) == 'L' {
				cr = (cl + cr) / 2
			} else {
				cl = (cl+cr)/2 + 1
			}
		}
		seatID := rl*8 + cl
		if highest < seatID {
			highest = seatID
		}
	}
	fmt.Println(highest)
}

func part2() {
	minSeatID, maxSeatID := 9999, 0
	sum := 0
	reader := bufio.NewReader(strings.NewReader(input))

	for {
		seat, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		rl, rr := 0, 127
		for _, v := range seat[0:7] {
			if rune(v) == 'F' {
				rr = (rl + rr) / 2
			} else {
				rl = (rl+rr)/2 + 1
			}
		}
		cl, cr := 0, 7
		for _, v := range seat[7:] {
			if rune(v) == 'L' {
				cr = (cl + cr) / 2
			} else {
				cl = (cl+cr)/2 + 1
			}
		}
		seatID := rl*8 + cl

		if seatID < minSeatID {
			minSeatID = seatID
		}
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
		sum += seatID
	}
	seatID := (maxSeatID-minSeatID+1)*(maxSeatID+minSeatID)/2 - sum
	fmt.Println(seatID)
}

func main() {
	part1()
	part2()
}
