package main

import (
	"fmt"
	"strings"
)

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
	seatTaken := map[int]bool{}
	minSeatID, maxSeatID := 9999, 0
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
		seatTaken[seatID] = true
		if seatID < minSeatID {
			minSeatID = seatID
		}
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	seatID := 0
	for i := minSeatID; i < maxSeatID; i++ {
		if !seatTaken[i] {
			seatID = i
			break
		}
	}
	fmt.Println(seatID)
}

func main() {
	part1()
	part2()
}
