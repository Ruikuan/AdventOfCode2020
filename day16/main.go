package main

import (
	"fmt"
	"strconv"
	"strings"
)

type validator struct {
	name  string
	from1 int
	to1   int
	from2 int
	to2   int
}

func part1() int {
	segments := strings.Split(input, "\n\n")
	scanners := []validator{}
	for _, line := range strings.Split(segments[0], "\n") {
		parts := strings.Split(line, ": ")
		v := validator{name: parts[0]}
		ranges := strings.Split(parts[1], " or ")
		v.from1, _ = strconv.Atoi(strings.Split(ranges[0], "-")[0])
		v.to1, _ = strconv.Atoi(strings.Split(ranges[0], "-")[1])
		v.from2, _ = strconv.Atoi(strings.Split(ranges[1], "-")[0])
		v.to2, _ = strconv.Atoi(strings.Split(ranges[1], "-")[1])
		scanners = append(scanners, v)
	}
	errRate := 0
	firstLine := true
	for _, line := range strings.Split(segments[2], "\n") {
		if firstLine {
			firstLine = false
			continue
		}
		nums := strings.Split(line, ",")
		for _, numStr := range nums {
			num, _ := strconv.Atoi(numStr)
			numValid := false
			for _, v := range scanners {
				if (v.from1 <= num && num <= v.to1) || (v.from2 <= num && num <= v.to2) {
					numValid = true
					break
				}
			}
			if !numValid {
				errRate += num
			}
		}
	}
	return errRate
}

func main() {
	fmt.Println(part1())
}
