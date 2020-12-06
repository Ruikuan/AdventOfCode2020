package main

import (
	"fmt"
	"strings"
)

func part1() {
	gmap := [26]bool{}
	sum := 0
	count := 0
	for _, answer := range strings.Split(input, "\n") {
		if answer == "" {
			sum += count
			count = 0
			gmap = [26]bool{}
			continue
		}
		for _, a := range answer {
			if !gmap[a-'a'] {
				count++
				gmap[a-'a'] = true
			}
		}
	}
	sum += count
	fmt.Println(sum)
}

func countYesAndReset(gmap *[26]bool) int {
	count := 0
	for i, b := range *gmap {
		if b {
			count++
		}
		gmap[i] = true
	}
	return count
}

func part2() {
	gmap := [26]bool{}
	for i := 0; i < 26; i++ {
		gmap[i] = true
	}

	sum := 0

	for _, answer := range strings.Split(input, "\n") {
		if answer == "" {
			sum += countYesAndReset(&gmap)
			continue
		}
		lmap := [26]bool{}
		for _, a := range answer {
			lmap[a-'a'] = true
		}
		for i := 0; i < 26; i++ {
			gmap[i] = gmap[i] && lmap[i]
		}
	}
	sum += countYesAndReset(&gmap)
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
