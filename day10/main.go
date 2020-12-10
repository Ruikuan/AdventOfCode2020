package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func part1() int {
	lines := strings.Split(input, "\n")
	nums := sort.IntSlice{}
	for _, l := range lines {
		num, _ := strconv.Atoi(l)
		nums = append(nums, num)
	}
	nums.Sort()
	d1, d3 := 0, 0
	pre := 0
	for _, n := range nums {
		diff := n - pre
		if diff == 1 {
			d1++
		} else if diff == 3 {
			d3++
		}
		pre = n
	}
	d3++
	return d1 * d3
}

func main() {
	fmt.Println(part1())
}
