package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "8"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines, 1000, 3)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string, n, k int) int {
	q := utils.PriorityQueue[[2]*box]{}
	boxes := make([]*box, 0)
	for _, line := range lines {
		pt := NewPt3(utils.ListOfNumbers(line, ","))
		b := &box{pt: pt}
		for _, b2 := range boxes {
			q.GPush([2]*box{b, b2}, b2.pt.sqDist(pt))
		}
		boxes = append(boxes, b)
	}
	circuits := make([]*circuit, 0)
	idx := 0
	for len(q) > 0 && idx < n {
		pair, _ := q.GPop()
		b1, b2 := pair[0], pair[1]
		c1, c2 := b1.c, b2.c
		if c1 == nil && c2 == nil {
			// new circuit
			c := &circuit{boxes: []*box{b1, b2}}
			b1.c = c
			b2.c = c
			circuits = append(circuits, c)
		} else if c1 != nil && c2 == nil {
			// add b2 to c1
			c1.boxes = append(c1.boxes, b2)
			b2.c = c1
		} else if c1 == nil && c2 != nil {
			// add b1 to c2
			c2.boxes = append(c2.boxes, b1)
			b1.c = c2
		} else if c1 != c2 {
			// merge circuits
			c1.boxes = append(c1.boxes, c2.boxes...)
			for _, b := range c2.boxes {
				b.c = c1
			}
			// remove c2 from circuits
			for i, c := range circuits {
				if c == c2 {
					circuits = append(circuits[:i], circuits[i+1:]...)
					break
				}
			}
		}
		idx++
	}
	sizes := utils.PriorityQueue[int]{}
	for _, c := range circuits {
		sizes.GPush(len(c.boxes), -len(c.boxes))
	}
	total := 1
	for _ = range k {
		size, _ := sizes.GPop()
		total *= size
	}
	return total
}

func part2(lines []string) int {
	q := utils.PriorityQueue[[2]*box]{}
	boxes := make([]*box, 0)
	circuits := make([]*circuit, 0)
	for _, line := range lines {
		pt := NewPt3(utils.ListOfNumbers(line, ","))
		b := &box{pt: pt}
		for _, b2 := range boxes {
			q.GPush([2]*box{b, b2}, b2.pt.sqDist(pt))
		}
		boxes = append(boxes, b)
		c := &circuit{boxes: []*box{b}}
		circuits = append(circuits, c)
		b.c = c
	}
	idx := 0
	for len(q) > 0 {
		pair, _ := q.GPop()
		b1, b2 := pair[0], pair[1]
		c1, c2 := b1.c, b2.c
		if c1 != c2 {
			// merge circuits
			c1.boxes = append(c1.boxes, c2.boxes...)
			for _, b := range c2.boxes {
				b.c = c1
			}
			// remove c2 from circuits
			for i, c := range circuits {
				if c == c2 {
					circuits = append(circuits[:i], circuits[i+1:]...)
					break
				}
			}
		}
		if len(circuits) == 1 {
			return b1.pt.X * b2.pt.X
		}
		idx++
	}
	return 0
}

type box struct {
	pt Pt3
	c  *circuit
}

type circuit struct {
	boxes []*box
}

type Pt3 struct {
	X, Y, Z int
}

func (p Pt3) sqDist(o Pt3) int {
	// distance euclidienne au carré
	return (p.X-o.X)*(p.X-o.X) + (p.Y-o.Y)*(p.Y-o.Y) + (p.Z-o.Z)*(p.Z-o.Z)
}

func NewPt3(data []int) Pt3 {
	return Pt3{X: data[0], Y: data[1], Z: data[2]}
}
