package main

import (
	"advent/utils"
	"fmt"
	"strconv"

	"github.com/peterstace/simplefeatures/geom"
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
	points := make([]geom.XY, 0)
	data := make([]float64, 0)
	areas := utils.PriorityQueue[geom.Geometry]{}
	for _, line := range lines {
		coords := utils.ListOfNumbers(line, ",")
		pt := geom.XY{X: float64(coords[0]), Y: float64(coords[1])}
		data = append(data, pt.X, pt.Y)
		for _, p := range points {
			r := geom.NewEnvelope(p, pt)
			area := r.Area()
			_ = area // useless
			a := (utils.Abs(int(p.X-pt.X)) + 1) * (utils.Abs(int(p.Y-pt.Y)) + 1)
			areas.GPush(r.AsGeometry(), -a)
		}
		points = append(points, pt)
	}
	points = append(points, points[0])
	data = append(data, data[0], data[1])
	polygone := geom.NewPolygonXY(data).AsGeometry()
	for len(areas) > 0 {
		rect, a := areas.GPop()
		if ok, err := geom.Contains(polygone, rect); err == nil && ok {
			return -a
		}
	}
	return 0
}

func rect(p1, p2 geom.Point) geom.Envelope {
	xy1, _ := p1.XY()
	xy2, _ := p2.XY()
	return geom.NewEnvelope(xy1, xy2)
}
