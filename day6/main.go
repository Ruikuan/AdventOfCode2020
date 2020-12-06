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

func part2() {

}

func main() {
	part1()
	part2()
}
