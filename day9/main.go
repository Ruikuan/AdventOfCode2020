package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isValid(ring *[25]int, val int) bool {
	dic := make(map[int]bool, 25)
	for _, v := range *ring {
		if dic[val-v] {
			return true
		}
		dic[v] = true
	}
	return false
}

func part1() int {
	ring := [25]int{}
	for idx, val := range strings.Split(input, "\n") {
		ival, _ := strconv.Atoi(val)
		if idx >= 25 {
			if !isValid(&ring, ival) {
				return ival
			}
		}
		ring[idx%25] = ival
	}
	panic("no such number.")
}

func part2() {
	nums := strings.Split(input, "\n")
	slice := []int{}
	sum := 0
	r := 0
	target := part1()
	for r < len(nums) {
		if sum == target {
			min, max := slice[0], slice[0]
			for _, v := range slice {
				if min > v {
					min = v
				}
				if max < v {
					max = v
				}
			}
			fmt.Println(min + max)
			return
		}
		for sum < target {
			rVal, _ := strconv.Atoi(nums[r])
			sum += rVal
			slice = append(slice, rVal)
			r++
		}
		for sum > target {
			sum -= slice[0]
			slice = slice[1:]
		}
	}
	panic("no such range.")
}

func main() {
	fmt.Println(part1())
	part2()
}
