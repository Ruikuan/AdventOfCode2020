package main

import (
	"fmt"
	"sort"
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
	list := sort.IntSlice{}
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
		list = append(list, seatID)
	}
	list.Sort()
	seatID := 0
	for i := 0; i < len(list)-1; i++ {
		if list[i+1] != list[i]+1 {
			seatID = list[i] + 1
			break
		}
	}
	fmt.Println(seatID)
}

func main() {
	part1()
	part2()
}
