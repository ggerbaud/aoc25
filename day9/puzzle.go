package main

import (
	"advent/utils"
	"fmt"
	"strconv"

	"github.com/tidwall/geometry"
)

const day = "9"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	max := 0
	points := make([]utils.Point, 0)
	for _, line := range lines {
		coords := utils.ListOfNumbers(line, ",")
		pt := utils.Point{X: coords[0], Y: coords[1]}
		for _, p := range points {
			area := (utils.Abs(pt.X-p.X) + 1) * (utils.Abs(pt.Y-p.Y) + 1)
			if area > max {
				max = area
			}
		}
		points = append(points, pt)
	}
	return max
}

func part2(lines []string) int {
	points := make([]geometry.Point, 0)
	areas := utils.PriorityQueue[geometry.Rect]{}
	for _, line := range lines {
		coords := utils.ListOfNumbers(line, ",")
		pt := geometry.Point{X: float64(coords[0]), Y: float64(coords[1])}
		for _, p := range points {
			r := geometry.Rect{Min: p, Max: pt}
			areas.GPush(r, -int(r.Area()))
		}
		points = append(points, pt)
	}
	polygone := geometry.NewPoly(points, [][]geometry.Point{}, geometry.NoIndexing)
	for len(areas) > 0 {
		rect, a := areas.GPop()
		if polygone.ContainsRect(rect) {
			return -a
		}
	}
	return 0
}
