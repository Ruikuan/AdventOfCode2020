package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1() int {
	lastTurns := map[int]int{}
	//preLastTurns := map[int]int{}
	i := 0
	num := 0
	for _, v := range strings.Split(input, ",") {
		i++
		num, _ = strconv.Atoi(v)
		//preLastTurns[num] = lastTurns[num]
		lastTurns[num] = i
	}
	num = 0
	i++
	for i++; i <= 2020; i++ {
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
func main() {
	fmt.Println(part1())
}
