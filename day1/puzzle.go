package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "1"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	size := 100
	total := 0
	safe := 50
	for _, line := range lines {
		digit, err := strconv.Atoi(line[1:])
		utils.CheckErrorP(err)
		if line[0] == 'R' {
			safe = (safe + digit) % size
		} else {
			safe = (safe - digit + size) % size
		}
		if safe == 0 {
			total++
		}
	}
	return total
}

func part2(lines []string) int {
	size := 100
	total := 0
	safe := 50
	for _, line := range lines {
		digit, err := strconv.Atoi(line[1:])
		utils.CheckErrorP(err)
		if line[0] == 'R' {
			safe = safe + digit
			for safe >= size {
				safe -= size
				total++
			}
		} else {
			if safe == 0 {
				total--
			}
			safe = safe - digit
			for safe < 0 {
				safe += size
				total++
			}
			if safe == 0 {
				total++
			}
		}
	}
	return total
}
