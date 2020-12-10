package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func sortedNums() sort.IntSlice {
	lines := strings.Split(input, "\n")
	nums := sort.IntSlice{}
	for _, l := range lines {
		num, _ := strconv.Atoi(l)
		nums = append(nums, num)
	}
	nums.Sort()
	return nums
}

func part1() int {
	nums := sortedNums()
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

func part2() int64 {
	nums := sortedNums()
	var dp1, dp2, dp3 int64 = 1, 0, 0
	pre1, pre2, pre3 := 0, 0, 0
	for _, num := range nums {
		var dp int64
		if num-pre3 <= 3 {
			dp = dp1 + dp2 + dp3
		} else if num-pre2 <= 2 {
			dp = dp1 + dp2
		} else {
			dp = dp1
		}
		dp3 = dp2
		dp2 = dp1
		dp1 = dp

		pre3 = pre2
		pre2 = pre1
		pre1 = num
	}
	return dp1
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
