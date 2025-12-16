package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const day = "10"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	for _, l := range lines {
		m := buildMachine(l, true)
		total += m.solve(true)
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, l := range lines {
		m := buildMachine(l, false)
		total += m.solve(false)
	}
	return total
}

func buildMachine(line string, lights bool) machine {
	parts := strings.Split(line, " ")
	m := machine{}
	for _, p := range parts {
		if strings.HasPrefix(p, "[") && lights {
			for _, c := range strings.Trim(p, "[]") {
				if c == '#' {
					m.tgt = append(m.tgt, 1)
				} else {
					m.tgt = append(m.tgt, 0)
				}
			}
		} else if strings.HasPrefix(p, "(") {
			button := []int{}
			btn := 0
			if p != "" {
				for s := range strings.SplitSeq(strings.Trim(p, "()"), ",") {
					v, _ := strconv.Atoi(s)
					button = append(button, v)
					btn |= 1 << v
				}
			}
			m.buttons = append(m.buttons, button)
		} else if strings.HasPrefix(p, "{") && !lights {
			for s := range strings.SplitSeq(strings.Trim(p, "{}"), ",") {
				v, _ := strconv.Atoi(s)
				m.tgt = append(m.tgt, v)
			}
		}
	}
	return m
}

func (m machine) allPatterns(mod bool) []pattern {
	nT, nB := len(m.tgt), len(m.buttons)
	patterns := []pattern{}
	for k := range 1 << nB {
		tgt := make(target, nT)
		cpt := 0
		for i := range nB {
			if k&(1<<i) != 0 {
				cpt++
				for _, idx := range m.buttons[i] {
					tgt[idx]++
				}
			}
		}
		if mod {
			for i := range tgt {
				tgt[i] %= 2
			}
		}
		patterns = append(patterns, pattern{tgt: tgt, cpt: cpt})
	}
	return patterns
}

func (m machine) solve(lights bool) int {
	cpt, _ := m.tgt.solve(m.allPatterns(lights))
	return cpt
}

func (t target) solve(patterns []pattern) (int, bool) {
	if t.isZero() {
		return 0, true
	}
	steps := math.MaxInt
	for _, p := range patterns {
		if !t.isSameParity(p.tgt) || p.tgt.isGreater(t) {
			continue
		}
		nextTgt := t.reduce(p.tgt)
		cpt, ok := nextTgt.solve(patterns)
		if !ok {
			continue
		}
		nSteps := 2*cpt + p.cpt
		if nSteps < steps {
			steps = nSteps
		}
	}
	if steps < math.MaxInt {
		return steps, true
	}
	return 0, false
}

type machine struct {
	tgt     target
	buttons [][]int
}

type target []int
type pattern struct {
	tgt target
	cpt int
}

func (t target) isZero() bool {
	for _, v := range t {
		if v != 0 {
			return false
		}
	}
	return true
}

func (t target) isSameParity(o target) bool {
	for i, v := range t {
		if v%2 != o[i]%2 {
			return false
		}
	}
	return true
}

func (t target) isGreater(o target) bool {
	for i, v := range t {
		if v > o[i] {
			return true
		}
	}
	return false
}

func (t target) reduce(o target) target {
	out := make(target, len(t))
	for i, v := range t {
		out[i] = (v - o[i]) / 2
	}
	return out
}
