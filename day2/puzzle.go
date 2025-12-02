package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "2"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		for rStr := range strings.SplitSeq(line, ",") {
			r := NewRange(rStr)
			for id := range r.invalidIds() {
				total += id
			}
		}
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		for rStr := range strings.SplitSeq(line, ",") {
			r := NewRange(rStr)
			for id := range r.invalidIds2() {
				total += id
			}
		}
	}
	return total
}

type Range struct {
	from, to int
}

// generator
func (r Range) invalidIds() func(func(int) bool) {
	return func(f func(int) bool) {
		i := r.from
		for i <= r.to {
			if !isValid(i) && !f(i) {
				return
			}
			i++
		}
	}
}
func (r Range) invalidIds2() func(func(int) bool) {
	return func(f func(int) bool) {
		i := r.from
		for i <= r.to {
			if !isValid2(i) && !f(i) {
				return
			}
			i++
		}
	}
}

func isValid(id int) bool {
	str := strconv.Itoa(id)
	if len(str)%2 != 0 {
		return true
	}
	half := len(str) / 2
	return str[:half] != str[half:]
}

func isValid2(id int) bool {
	idStr := strconv.Itoa(id)
	h := len(idStr) / 2
	for i := range h {
		prefix := idStr[:i+1]
		if len(strings.ReplaceAll(idStr, prefix, "")) == 0 {
			return false
		}
	}
	return true
}

func NewRange(rangeStr string) Range {
	parts := strings.Split(rangeStr, "-")
	from := utils.ParseInt(parts[0])
	to := utils.ParseInt(parts[1])
	return Range{from: from, to: to}
}
