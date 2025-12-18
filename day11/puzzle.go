package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "11"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	devices := buildDevice(lines)
	// suppose no loop
	return findPath(devices, devices["you"], devices["out"])
}

func part2(lines []string) int {
	devices := buildDevice(lines)
	memory := make(map[query]int)
	return findPathWithQuery(memory, devices, query{"svr", true, true})
}

func findPath(devices map[string]*device, start, exit *device) int {
	total := 0
	for _, d := range start.out {
		if d == exit {
			total++
		} else {
			total += findPath(devices, d, exit)
		}
	}
	return total
}

func findPathWithQuery(memory map[query]int, devices map[string]*device, q query) int {
	if v, ok := memory[q]; ok {
		return v
	}
	total := 0
	start := devices[q.start]
	for _, d := range start.out {
		if d.name == "out" {
			if !q.fft && !q.dac {
				total++
			}
		} else {
			nq := query{d.name, q.dac, q.fft}
			if d.name == "fft" {
				nq.fft = false
			} else if d.name == "dac" {
				nq.dac = false
			}
			total += findPathWithQuery(memory, devices, nq)
		}
	}
	memory[q] = total
	return total
}

func buildDevice(lines []string) map[string]*device {
	devices := make(map[string]*device)
	for _, l := range lines {
		parts := strings.Split(l, ": ")
		name := parts[0]
		out := strings.Split(parts[1], " ")
		d, ok := devices[name]
		if !ok {
			d = &device{name: name, out: make([]*device, 0)}
			devices[name] = d
		}
		for _, o := range out {
			clearO := strings.TrimSpace(o)
			if len(clearO) == 0 {
				continue
			}
			do, ok := devices[clearO]
			if !ok {
				do = &device{name: clearO, out: make([]*device, 0)}
				devices[clearO] = do
			}
			d.out = append(d.out, do)
		}
	}
	return devices
}

type device struct {
	name string
	out  []*device
}

type query struct {
	start    string
	dac, fft bool
}
