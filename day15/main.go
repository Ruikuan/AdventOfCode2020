package main

import (
	"fmt"
	"strconv"
	"strings"
)

func takeTurn(turn int) int {
	lastTurns := map[int]int{}
	i := 0
	num := 0
	for _, v := range strings.Split(input, ",") {
		i++
		num, _ = strconv.Atoi(v)
		lastTurns[num] = i
	}
	num = 0
	i++
	for i++; i <= turn; i++ {
		next := 0
		if lastTurns[num] != 0 {
			next = i - 1 - lastTurns[num]
		} else {
			next = 0
		}
		lastTurns[num] = i - 1
		num = next
	}
	return num
}

func part1() int {
	return takeTurn(2020)
}

func part2() int {
	return takeTurn(30000000)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
