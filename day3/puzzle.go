package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "3"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	for _, bank := range lines {
		total += highestJoltage(bank, 2)
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, bank := range lines {
		total += highestJoltage(bank, 12)
	}
	return total
}

func highestJoltage(bank string, size int) int {
	cumul := 0
	idx := -1
	for n := range size {
		max := 0
		cumul *= 10
		for i := idx + 1; i <= len(bank)-(size-n); i++ {
			id := int(bank[i] - '0')
			if id == 9 {
				idx = i
				max = 9
				break
			}
			if id > max {
				max = id
				idx = i
			}
		}
		cumul += max
	}
	return cumul
}
