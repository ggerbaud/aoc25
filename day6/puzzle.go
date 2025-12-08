package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "6"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	pbms := buildProblems(lines)
	total := 0
	for _, pbm := range pbms {
		res := 1
		if pbm.add {
			res = 0
		}
		for _, n := range pbm.data {
			if pbm.add {
				res += n
			} else {
				res *= n
			}
		}
		total += res
	}
	return total
}

func part2(lines []string) int {
	digits := make([]int, len(lines[0]))
	for _, line := range lines[:len(lines)-1] {
		for j, ch := range line {
			if ch == ' ' {
				continue
			}
			n := int(ch - '0')
			digits[j] = 10*digits[j] + n
		}
	}
	total := 0
	add := false
	for i, ch := range lines[len(lines)-1] {
		if ch == ' ' {
			continue
		}
		add = ch == '+'
		res := 1
		if add {
			res = 0
		}
		for k := i; k < len(digits); k++ {
			if digits[k] == 0 {
				total += res
				res = 0
				break
			}
			if add {
				res += digits[k]
			} else {
				res *= digits[k]
			}
		}
		if res > 0 {
			total += res
		}
	}
	return total
}

type problem struct {
	data []int
	add  bool
}

func buildProblems(lines []string) []*problem {
	pbms := make([]*problem, 0)
	for i, line := range lines {
		parts := strings.Split(line, " ")
		if i == len(lines)-1 {
			idx := 0
			for _, part := range parts {
				if part == "" {
					continue
				}
				pbms[idx].add = strings.TrimSpace(part) == "+"
				idx++
			}
		} else {
			idx := 0
			for _, part := range parts {
				if part == "" {
					continue
				}
				if i == 0 {
					pbms = append(pbms, &problem{make([]int, 0), true})
				}
				n := utils.ParseInt(strings.TrimSpace(part))
				pbms[idx].data = append(pbms[idx].data, n)
				idx++
			}
		}
	}
	return pbms
}
