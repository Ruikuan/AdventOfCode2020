package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1() {
	insts := strings.Split(input, "\n")
	execed := make([]bool, len(insts))
	accumulator := 0
	idx := 0
	for {
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
	fmt.Println(accumulator)
}

func main() {
	part1()
}
