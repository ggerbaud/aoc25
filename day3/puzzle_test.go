package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 357
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 3121910778619
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func TestHighestJoltage(t *testing.T) {
	for _, data := range hjdata() {
		result := highestJoltage(data.bank, data.size)
		if result != data.expect {
			t.Fatalf("HighestJoltage returns %d for %s,%d, we expect %d", result, data.bank, data.size, data.expect)
		}
	}
}

func hjdata() []hjtest {
	return []hjtest{
		{"987654321111111", 2, 98},
		{"811111111111119", 2, 89},
		{"234234234234278", 2, 78},
		{"818181911112111", 2, 92},
		{"987654321111111", 12, 987654321111},
		{"811111111111119", 12, 811111111119},
		{"234234234234278", 12, 434234234278},
		{"818181911112111", 12, 888911112111},
	}
}

type hjtest struct {
	bank   string
	size   int
	expect int
}

func testData() []string {
	return []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
}
