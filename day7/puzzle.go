package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "7"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	fl := lines[0]
	beams := make(map[int]bool)
	for i, ch := range fl {
		if ch == 'S' {
			beams[i] = true
		}
	}
	splits := 0
	for _, line := range lines[1:] {
		for i, ch := range line {
			if ch == '^' && beams[i] {
				splits++
				beams[i] = false
				if i > 0 {
					beams[i-1] = true
				}
				if i < len(line)-1 {
					beams[i+1] = true
				}
			}
		}
	}
	return splits
}

func part2(lines []string) int {
	fl := lines[0]
	beams := make(map[int]int)
	for i, ch := range fl {
		if ch == 'S' {
			beams[i]++
		}
	}
	for _, line := range lines[1:] {
		nbeams := make(map[int]int)
		for i, ch := range line {
			if ch == '^' {
				n := beams[i]
				nbeams[i] = 0
				if i > 0 {
					nbeams[i-1] += n
				}
				if i < len(line)-1 {
					nbeams[i+1] += n
				}
			} else {
				nbeams[i] += beams[i]
			}
		}
		beams = nbeams
	}
	splits := 0
	for _, n := range beams {
		splits += n
	}
	return splits
}
