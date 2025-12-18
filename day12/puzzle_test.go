package main

import "testing"

func TestPart1(t *testing.T) {
	lines := testData()
	result := part1(lines)
	expect := 3
	if result != expect {
		t.Fatalf("Part1 returns %d, we expect %d", result, expect)
	}
}

func TestPart2(t *testing.T) {
	lines := testData()
	result := part2(lines)
	expect := 0
	if result != expect {
		t.Fatalf("Part2 returns %d, we expect %d", result, expect)
	}
}

func testData() []string {
	return []string{
		"0:",
		"###",
		"##.",
		"##.",
		"",
		"1:",
		"###",
		"##.",
		".##",
		"",
		"2:",
		".##",
		"###",
		"##.",
		"",
		"3:",
		"##.",
		"###",
		"##.",
		"",
		"4:",
		"###",
		"#..",
		"###",
		"",
		"5:",
		"###",
		".#.",
		"###",
		"",
		"4x4: 0 0 0 0 2 0",
		"12x5: 1 0 1 0 2 2",
		"12x5: 1 0 1 0 3 2",
	}
}
