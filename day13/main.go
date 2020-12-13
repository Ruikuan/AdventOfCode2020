package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1() int {
	lines := strings.Split(input, "\n")
	depart, _ := strconv.Atoi(lines[0])
	validBus := []int{}
	for _, id := range strings.Split(lines[1], ",") {
		if num, err := strconv.Atoi(id); err == nil {
			validBus = append(validBus, num)
		}
	}

	wait := depart
	bus := 0
	for _, busID := range validBus {
		w := busID - depart%busID
		if wait > w {
			wait = w
			bus = busID
		}
	}
	return wait * bus
}

func main() {
	fmt.Println(part1())
}
