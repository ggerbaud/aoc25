package main

import (
	"advent/utils"
	"fmt"
	"slices"
	"strconv"
)

const day = "5"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	ranges := Ranges{make([]*Range, 0)}
	idx := 0
	for i, line := range lines {
		if line == "" {
			idx = i + 1
			break
		}
		ranges.addRange(NewRange(line))
		if !ranges.debug() {
			panic(fmt.Sprintf("ranges are overlapping after adding line %s (%d)", line, i))
		}
	}
	total := 0
	for _, line := range lines[idx:] {
		n, _ := strconv.Atoi(line)
		if ranges.isIn(n) {
			total++
		}
	}
	return total
}

func part2(lines []string) int {
	ranges := Ranges{make([]*Range, 0)}
	for _, line := range lines {
		if line == "" {
			break
		}
		ranges.addRange(NewRange(line))
	}
	total := 0
	for _, r := range ranges.data {
		total += r.max - r.min + 1
	}
	return total
}

type Ranges struct {
	data []*Range
}

func (rs *Ranges) isIn(n int) bool {
	for _, r := range rs.data {
		if n < r.min {
			return false
		}
		if n <= r.max {
			return true
		}
	}
	return false
}

func (rs *Ranges) addRange(r *Range) {
	overlaping := false
	for i, existing := range rs.data {
		if overlaping {
			if r.max < existing.min {
				// fin du cas C
				return
			}
			if r.max >= existing.max {
				// cas F
				rs.data = slices.Delete(rs.data, i, i+1)
				continue
			}
			if r.max <= existing.max {
				// cas D
				rs.data[i-1].max = existing.max
				rs.data = slices.Delete(rs.data, i, i+1)
				return
			}
		} else {
			if r.max < existing.min {
				// cas A
				rs.data = slices.Insert(rs.data, i, r)
				return
			}
			if r.min > existing.max {
				continue
			}
			if r.min < existing.min {
				// cas B
				existing.min = r.min
			}
			if r.max > existing.max {
				// cas C ou E
				existing.max = r.max
				if i+1 < len(rs.data) {
					overlaping = true
					continue
				}
			}
		}
		return
	}
	if !overlaping {
		// cas G
		rs.data = append(rs.data, r)
	}
}

func (rs *Ranges) debug() bool {
	for i, r := range rs.data {
		if i > 0 && r.min <= rs.data[i-1].max {
			return false
		}
	}
	return true
}

type Range struct {
	min, max int
}

func (r *Range) isIn(n int) bool {
	return n >= r.min && n <= r.max
}

func NewRange(s string) *Range {
	var r Range
	fmt.Sscanf(s, "%d-%d", &r.min, &r.max)
	return &r
}
