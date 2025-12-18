package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "12"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	shapes, trees := buildData(lines)
	for _, t := range trees {
		area := t.w * t.h
		for n, c := range t.needs {
			area -= c * shapes[n].area
		}
		if area >= 0 {
			total++
		}
	}
	return total
}

func part2(lines []string) int {
	return 0
}

func buildData(lines []string) ([]shape, []tree) {
	shapes := make([]shape, 0)
	trees := make([]tree, 0)
	sh := shape{}
	for _, line := range lines {
		if len(line) == 0 {
			shapes = append(shapes, sh)
			sh = shape{}
		} else if strings.Contains(line, "x") {
			t := tree{needs: make([]int, 0)}
			parts := strings.Split(line, ":")
			wh := strings.Split(parts[0], "x")
			t.w = utils.ParseInt(wh[0])
			t.h = utils.ParseInt(wh[1])
			t.needs = utils.ListOfNumbers(parts[1], " ")
			trees = append(trees, t)
		} else {
			sh.area += strings.Count(line, "#")
		}
	}
	return shapes, trees
}

type shape struct {
	area int
}

type tree struct {
	w, h  int
	needs []int
}
