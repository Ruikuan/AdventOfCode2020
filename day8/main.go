package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1() {
	insts := strings.Split(input, "\n")
	acc, _ := exec(insts)
	fmt.Println(acc)
}

func exec(insts []string) (accumulator int, terminate bool) {
	end := len(insts)
	accumulator = 0
	execed := make([]bool, len(insts))
	idx := 0
	for {
		if idx > end {
			break
		}
		if idx == end {
			return accumulator, true
		}
		if execed[idx] {
			break
		}
		execed[idx] = true
		pair := strings.Split(insts[idx], " ")
		inst := pair[0]
		switch inst {
		case "acc":
			arg, _ := strconv.Atoi(pair[1])
			accumulator += arg
			break
		case "jmp":
			arg, _ := strconv.Atoi(pair[1])
			idx += arg
			continue
		case "nop":
			break
		}
		idx++
	}
	return accumulator, false
}

func part2() {
	insts := strings.Split(input, "\n")
	for i, v := range insts {
		edited := false
		if strings.HasPrefix(v, "nop") {
			insts[i] = strings.Replace(v, "nop", "jmp", -1)
			edited = true
		} else if strings.HasPrefix(v, "jmp") {
			insts[i] = strings.Replace(v, "jmp", "nop", -1)
			edited = true
		}
		if edited {
			if acc, t := exec(insts); t {
				fmt.Println(acc)
				break
			}
		}
		insts[i] = v
	}
}

func main() {
	part1()
	part2()
}
