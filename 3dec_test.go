package aoc2017

import "testing"

func TestManhattanDistance(t *testing.T) {
	data := map[int]int{
		1:      0,
		12:     3,
		23:     2,
		1024:   31,
		361527: 326,
	}

	for i, e := range data {
		if r := calculateManhattanDistance(i); r != e {
			t.Error("Expected", e, "for input", i, ", got", r)
		}
	}
}

func BenchmarkManhattanDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		calculateManhattanDistance(1024)
	}
}

func TestCalculateLargerValue(t *testing.T) {
	data := map[int]int{
		361527: 363010,
	}

	for i, e := range data {
		if r := calculateLargerValue(i); r != e {
			t.Error("Expected", e, "for input", i, ", got", r)
		}
	}
}

func BenchmarkCalculateLargerValue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		calculateLargerValue(361527)
	}
}
