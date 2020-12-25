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

func part2() int {
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

	nearTicks := [][]int{}
	firstLine := true
	for _, line := range strings.Split(segments[2], "\n") {
		if firstLine {
			firstLine = false
			continue
		}
		lineNums := []int{}
		nums := strings.Split(line, ",")
		lineValid := true
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
				lineValid = false
				break
			}
			lineNums = append(lineNums, num)
		}
		if lineValid { //only do something when line valid
			nearTicks = append(nearTicks, lineNums)
		}
	}

	validScanners := [][]validator{}
	for i := 0; i < len(scanners); i++ {
		validScanners = append(validScanners, []validator{})
		for _, scanner := range scanners {
			fit := true
			for _, tick := range nearTicks {
				if !validate(tick[i], &scanner) {
					fit = false
					break
				}
			}
			if fit {
				validScanners[i] = append(validScanners[i], scanner)
			}
		}
	}
	theScanner := make([]validator, len(scanners))
	arranged := 0
	for arranged < len(scanners) {
		for i, vs := range validScanners {
			if len(vs) == 1 {
				theScanner[i] = vs[0]
				arranged++
				for j, vs2 := range validScanners {
					if len(vs2) >= 1 {
						newVs := []validator{}
						for _, v := range vs2 {
							if v != vs[0] {
								newVs = append(newVs, v)
							}
						}
						validScanners[j] = newVs
					}
				}
			}
		}
	}

	yourTick := strings.Split(strings.Split(segments[1], "\n")[1], ",")

	result := 1
	for i, v := range theScanner {
		if strings.Index(v.name, "departure") >= 0 {
			num, _ := strconv.Atoi(yourTick[i])
			result *= num
		}
	}
	return result
}

func validate(num int, v *validator) bool {
	return (v.from1 <= num && num <= v.to1) || (v.from2 <= num && num <= v.to2)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
