package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "4"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	d := newDiagram(lines)
	for p := range d.visit() {
		if d.isRoll(p) && d.hasLessThanFourAround(p) {
			total++
		}
	}
	return total
}

func part2(lines []string) int {
	total := 0
	d := newDiagram(lines)
	next := d.copy()
	for {
		partial := 0
		for p := range d.visit() {
			if d.isRoll(p) && d.hasLessThanFourAround(p) {
				partial++
				next.remove(p)
			}
		}
		if partial == 0 {
			break
		}
		total += partial
		d = next
		next = d.copy()
	}
	return total
}

type diagram struct {
	data [][]bool
}

func (d diagram) isIn(p utils.Point) bool {
	if p.X < 0 || p.Y < 0 || p.Y >= len(d.data) || p.X >= len(d.data[p.Y]) {
		return false
	}
	return true
}

func (d diagram) isRoll(p utils.Point) bool {
	return d.isIn(p) && d.data[p.Y][p.X]
}

func (d diagram) remove(p utils.Point) {
	if d.isIn(p) {
		d.data[p.Y][p.X] = false
	}
}

func (d diagram) hasLessThanFourAround(p utils.Point) bool {
	total := 0
	for _, np := range []utils.Point{p.Up().Left(), p.Up(), p.Up().Right(), p.Left(), p.Right(), p.Down().Left(), p.Down(), p.Down().Right()} {
		if d.isRoll(np) {
			total++
		}
		if total >= 4 {
			return false
		}
	}
	return true
}

func (d diagram) visit() func(func(utils.Point) bool) {
	return func(f func(utils.Point) bool) {
		for y, row := range d.data {
			for x, _ := range row {
				if !f(utils.Point{X: x, Y: y}) {
					return
				}
			}
		}
	}
}

func (d diagram) copy() diagram {
	newD := diagram{data: make([][]bool, len(d.data))}
	for y, row := range d.data {
		newD.data[y] = make([]bool, len(row))
		copy(newD.data[y], row)
	}
	return newD
}

func newDiagram(lines []string) diagram {
	d := diagram{data: make([][]bool, len(lines))}
	for y, line := range lines {
		d.data[y] = make([]bool, len(line))
		for x, c := range line {
			if c == '@' {
				d.data[y][x] = true
			} else {
				d.data[y][x] = false
			}
		}
	}
	return d
}
