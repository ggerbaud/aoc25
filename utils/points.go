package utils

type Point struct {
	X, Y int
}

func (p Point) Up() Point {
	return Point{X: p.X, Y: p.Y - 1}
}

func (p Point) Down() Point {
	return Point{X: p.X, Y: p.Y + 1}
}

func (p Point) Right() Point {
	return Point{X: p.X + 1, Y: p.Y}
}

func (p Point) Left() Point {
	return Point{X: p.X - 1, Y: p.Y}
}

func (p Point) Dist(pt Point) int {
	return Abs(p.X-pt.X) + Abs(p.Y-pt.Y)
}

func (p Point) Vect(pt Point) Point {
	return Point{X: pt.X - p.X, Y: pt.Y - p.Y}
}

type Mover func(p Point) Point

func (m Mover) Revert() Mover {
	pt := Point{X: 0, Y: 0}
	if pt.Up() == m(pt) {
		return Downer
	}
	if pt.Down() == m(pt) {
		return Upper
	}
	if pt.Left() == m(pt) {
		return Righter
	}
	if pt.Right() == m(pt) {
		return Lefter
	}
	panic("invalid mover")
}

var (
	EmptyPoint = Point{}
	Upper      = func(p Point) Point { return p.Up() }
	Downer     = func(p Point) Point { return p.Down() }
	Lefter     = func(p Point) Point { return p.Left() }
	Righter    = func(p Point) Point { return p.Right() }
)
