package main

import (
	"fmt"
	"strconv"
	"strings"
)

func mask36(mask string, val uint64) uint64 {
	var result uint64 = 0
	for i := 0; i < 36; i++ {
		bit := (val >> (35 - i)) & 1
		if mask[i] == '1' || (bit == 1 && mask[i] == 'X') {
			result |= 1 << (35 - i)
		}
	}
	return result
}

func part1() uint64 {
	mem := map[uint64]uint64{}
	mask := ""
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " = ")
		if parts[0] == "mask" {
			mask = parts[1]
		} else {
			addr, _ := strconv.ParseUint(parts[0][4:len(parts[0])-1], 10, 64)
			val, _ := strconv.ParseUint(parts[1], 10, 64)
			mem[addr] = mask36(mask, val)
		}
	}
	var result uint64 = 0
	for _, v := range mem {
		result += v
	}
	return result
}

func maskAddr(mask string, addr uint64) []uint64 {
	result := []uint64{}
	helper(mask, addr, &result, 0, 0)
	return result
}

func helper(mask string, addr uint64, result *[]uint64, i int, cur uint64) {
	if i == 36 {
		*result = append(*result, cur)
		return
	}

	switch mask[i] {
	case '1':
		cur |= 1 << (35 - i)
		helper(mask, addr, result, i+1, cur)
		break
	case '0':
		bit := (addr >> (35 - i)) & 1
		if bit == 1 {
			cur |= 1 << (35 - i)
		}
		helper(mask, addr, result, i+1, cur)
		break
	case 'X':
		helper(mask, addr, result, i+1, cur)
		cur |= 1 << (35 - i)
		helper(mask, addr, result, i+1, cur)
		break
	}
}

func part2() uint64 {
	mem := map[uint64]uint64{}
	mask := ""
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " = ")
		if parts[0] == "mask" {
			mask = parts[1]
		} else {
			addr, _ := strconv.ParseUint(parts[0][4:len(parts[0])-1], 10, 64)
			val, _ := strconv.ParseUint(parts[1], 10, 64)
			for _, a := range maskAddr(mask, addr) {
				mem[a] = val
			}
		}
	}
	var result uint64 = 0
	for _, v := range mem {
		result += v
	}
	return result
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
