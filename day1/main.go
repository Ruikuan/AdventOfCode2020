package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1() {
	nums := map[int]bool{}
	for _, l := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(l)
		if nums[2020-num] {
			fmt.Println(num * (2020 - num))
			break
		}
		nums[num] = true
	}
}

func part2() {
	lines := strings.Split(input, "\n")
	numsMap := map[int]bool{}
	nums := make([]int, len(lines))
	for i, l := range lines {
		num, _ := strconv.Atoi(l)
		numsMap[num] = true
		nums[i] = num
	}

	for i, n1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			n2 := nums[j]
			if numsMap[2020-n1-n2] {
				fmt.Println(n1 * n2 * (2020 - n1 - n2))
				return
			}
		}
	}
}

func main() {
	part1()
	part2()
}
