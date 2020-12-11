package main

import (
	"fmt"
	"strings"
)

var changes = [8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func getInput() [][]rune {
	lines := strings.Split(input, "\n")
	state := make([][]rune, len(lines))
	for i, l := range lines {
		state[i] = []rune(l)
	}
	return state
}

func applyRule(state [][]rune) [][]rune {
	nextState := make([][]rune, len(state))
	for i := 0; i < len(state); i++ {
		l := len(state[i])
		nextState[i] = make([]rune, l)
		for j := 0; j < l; j++ {
			if state[i][j] == '.' {
				nextState[i][j] = '.'
			} else if state[i][j] == 'L' {
				nextState[i][j] = '#'
				for _, ch := range changes {
					x, y := i+ch[0], j+ch[1]
					if x >= 0 && x < len(state) && y >= 0 && y < l {
						if state[x][y] == '#' {
							nextState[i][j] = 'L'
							break
						}
					}
				}
			} else if state[i][j] == '#' {
				occAdj := 0
				nextState[i][j] = '#'
				for _, ch := range changes {
					x, y := i+ch[0], j+ch[1]
					if x >= 0 && x < len(state) && y >= 0 && y < l {
						if state[x][y] == '#' {
							occAdj++
							if occAdj >= 4 {
								nextState[i][j] = 'L'
								break
							}
						}
					}
				}
			}
		}
	}
	return nextState
}

func isEqual(state1 [][]rune, state2 [][]rune) bool {
	for row, line := range state1 {
		for col, char := range line {
			if char != state2[row][col] {
				return false
			}
		}
	}
	return true
}

// printState for debug
func printState(state [][]rune) {
	for i := 0; i < len(state); i++ {
		l := len(state[i])
		for j := 0; j < l; j++ {
			fmt.Printf("%s", string(state[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func runPart(fn func([][]rune) [][]rune) int {
	state := getInput()
	eq := false
	for !eq {
		//printState(state)
		nextState := fn(state)
		eq = isEqual(state, nextState)
		state = nextState
	}
	occ := 0
	for _, line := range state {
		for _, c := range line {
			if c == '#' {
				occ++
			}
		}
	}
	return occ
}

func part1() int {
	return runPart(applyRule)
}

func applyRule2(state [][]rune) [][]rune {
	nextState := make([][]rune, len(state))
	for i := 0; i < len(state); i++ {
		l := len(state[i])
		nextState[i] = make([]rune, l)
		for j := 0; j < l; j++ {
			if state[i][j] == '.' {
				nextState[i][j] = '.'
			} else if state[i][j] == 'L' {
				nextState[i][j] = '#'
			l1:
				for _, ch := range changes {
					for scale := 1; scale < l || scale < len(state); scale++ {
						x, y := i+scale*ch[0], j+scale*ch[1]
						if x < 0 || y < 0 || x >= len(state) || y >= l {
							break
						}
						if state[x][y] == 'L' {
							break
						} else if state[x][y] == '#' {
							nextState[i][j] = 'L'
							break l1
						}
					}
				}
			} else if state[i][j] == '#' {
				occAdj := 0
				nextState[i][j] = '#'
			l2:
				for _, ch := range changes {
					for scale := 1; scale < l || scale < len(state); scale++ {
						x, y := i+scale*ch[0], j+scale*ch[1]
						if x < 0 || y < 0 || x >= len(state) || y >= l {
							break
						}
						if state[x][y] == 'L' {
							break
						}
						if state[x][y] == '#' {
							occAdj++
							if occAdj >= 5 {
								nextState[i][j] = 'L'
								break l2
							}
							break
						}
					}
				}
			}
		}
	}
	return nextState
}

func part2() int {
	return runPart(applyRule2)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
