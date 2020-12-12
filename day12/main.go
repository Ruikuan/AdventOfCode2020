package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	directions = [4][2]int{
		{1, 0},
		{0, -1},
		{-1, 0},
		{0, 1},
	}
)

func part1() int {
	curDir := 0
	pos := [2]int{0, 0}
	for _, ins := range strings.Split(input, "\n") {
		act := ins[0]
		val, _ := strconv.Atoi(ins[1:])
		switch act {
		case 'N':
			pos[1] = pos[1] + val
			break
		case 'S':
			pos[1] = pos[1] - val
			break
		case 'E':
			pos[0] = pos[0] + val
			break
		case 'W':
			pos[0] = pos[0] - val
			break
		case 'F':
			pos[0] = pos[0] + directions[curDir][0]*val
			pos[1] = pos[1] + directions[curDir][1]*val
			break
		case 'L':
			val = val % 360
			val = val / 90
			curDir = (curDir + 4 - val) % 4
			break
		case 'R':
			val = val % 360
			val = val / 90
			curDir = (curDir + val) % 4
			break
		}
	}
	return int(math.Abs(float64(pos[0]))) + int(math.Abs(float64(pos[1])))
}

func rRotate90(x, y int) (int, int) {
	return y, -x
}
func lRotate90(x, y int) (int, int) {
	return -y, x
}

func part2() int {
	waypoint := [2]int{10, 1}
	pos := [2]int{0, 0}
	for _, ins := range strings.Split(input, "\n") {
		act := ins[0]
		val, _ := strconv.Atoi(ins[1:])
		switch act {
		case 'N':
			waypoint[1] = waypoint[1] + val
			break
		case 'S':
			waypoint[1] = waypoint[1] - val
			break
		case 'E':
			waypoint[0] = waypoint[0] + val
			break
		case 'W':
			waypoint[0] = waypoint[0] - val
			break
		case 'F':
			pos[0] = pos[0] + waypoint[0]*val
			pos[1] = pos[1] + waypoint[1]*val
			break
		case 'L':
			val = val % 360
			val = val / 90
			for i := 0; i < val; i++ {
				waypoint[0], waypoint[1] = lRotate90(waypoint[0], waypoint[1])
			}
			break
		case 'R':
			val = val % 360
			val = val / 90
			for i := 0; i < val; i++ {
				waypoint[0], waypoint[1] = rRotate90(waypoint[0], waypoint[1])
			}
			break
		}
	}
	return int(math.Abs(float64(pos[0]))) + int(math.Abs(float64(pos[1])))
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
