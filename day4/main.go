package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/4

func part1() {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	isValidPassport := func(fields map[string]bool) bool {
		valid := true
		for _, rf := range requiredFields {
			if !fields[rf] {
				valid = false
				break
			}
		}
		return valid
	}
	fields := map[string]bool{}
	sum := 0

	for _, line := range strings.Split(inputText, "\n") {
		if line == "" {
			if isValidPassport(fields) {
				sum++
			}
			fields = make(map[string]bool)
			continue
		}
		for _, pair := range strings.Split(line, " ") {
			fields[strings.Split(pair, ":")[0]] = true
		}
	}
	if isValidPassport(fields) {
		sum++
	}
	fmt.Println(sum)
}

func part2() {
	byr := func(val string) bool {
		ival, _ := strconv.Atoi(val)
		return 1920 <= ival && ival <= 2002
	}
	iyr := func(val string) bool {
		ival, _ := strconv.Atoi(val)
		return 2010 <= ival && ival <= 2020
	}
	eyr := func(val string) bool {
		ival, _ := strconv.Atoi(val)
		return 2020 <= ival && ival <= 2030
	}
	hgt := func(val string) bool {
		l := len(val)
		suffix := val[l-2:]
		val = val[:l-2]
		ival, _ := strconv.Atoi(val)
		if suffix == "cm" {
			return 150 <= ival && ival <= 193
		} else if suffix == "in" {
			return 59 <= ival && ival <= 76
		}
		return false
	}
	hcl := func(val string) bool {
		if len(val) == 7 && rune(val[0]) == '#' {
			for _, v := range val[1:] {
				if !(('0' <= rune(v) && rune(v) <= '9') || ('a' <= rune(v) && rune(v) <= 'z')) {
					return false
				}
			}
			return true
		}
		return false
	}
	ecls := map[string]bool{
		"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true,
	}
	ecl := func(val string) bool {
		return ecls[val]
	}
	pid := func(val string) bool {
		if len(val) != 9 {
			return false
		}
		for _, v := range val {
			if !('0' <= rune(v) && rune(v) <= '9') {
				return false
			}
		}
		return true
	}

	fieldsValidChecks := map[string]func(string) bool{
		"byr": byr, "iyr": iyr, "eyr": eyr, "hgt": hgt, "hcl": hcl, "ecl": ecl, "pid": pid,
	}

	isValidPassport := func(fields map[string]string) bool {
		for k, v := range fieldsValidChecks {
			if fields[k] == "" {
				return false
			}
			if !v(fields[k]) {
				return false
			}
		}
		return true
	}

	fields := map[string]string{}
	sum := 0

	for _, line := range strings.Split(inputText, "\n") {
		if line == "" {
			if isValidPassport(fields) {
				sum++
			}
			fields = make(map[string]string)
			continue
		}
		for _, pair := range strings.Split(line, " ") {
			kv := strings.Split(pair, ":")
			fields[kv[0]] = kv[1]
		}
	}
	if isValidPassport(fields) {
		sum++
	}
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
