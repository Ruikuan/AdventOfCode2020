package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1() int {
	lines := strings.Split(input, "\n")
	depart, _ := strconv.Atoi(lines[0])
	validBus := []int{}
	for _, id := range strings.Split(lines[1], ",") {
		if num, err := strconv.Atoi(id); err == nil {
			validBus = append(validBus, num)
		}
	}

	wait := depart
	bus := 0
	for _, busID := range validBus {
		r := depart % busID
		if r == 0 {
			return 0
		}
		w := busID - r
		if wait > w {
			wait = w
			bus = busID
		}
	}
	return wait * bus
}

func part2() int64 {
	busMap := map[int64]int64{}
	for i, id := range strings.Split(strings.Split(input, "\n")[1], ",") {
		if num, err := strconv.Atoi(id); err == nil {
			busMap[int64(num)] = int64((num - i) % num)
		}
	}
	return solveChineseRemainder(busMap)
}

// https://zh.wikipedia.org/wiki/%E4%B8%AD%E5%9B%BD%E5%89%A9%E4%BD%99%E5%AE%9A%E7%90%86
func solveChineseRemainder(nums map[int64]int64) int64 {
	var sum int64 = 0
	var M int64 = 1
	for k := range nums {
		M *= k
	}
	//fmt.Printf("M:%d\n", M)
	for m1, a1 := range nums {
		//fmt.Printf("a1:%d m1:%d ", a1, m1)
		M1 := M / m1
		//fmt.Printf("M1:%d ", M1)
		t1 := int64(1)
		for { // loop to find coprime, hope not too large
			if (t1*M1)%m1 != 1 {
				t1++
			} else {
				break
			}
		}
		//fmt.Printf("t1:%d \n", t1)
		sum += a1 * t1 * M1
	}
	return sum % M
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
