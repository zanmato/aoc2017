package aoc2017

import (
	"testing"
)

func TestDuelingGenerators(t *testing.T) {
	data := []struct {
		Count  int
		InputA int
		InputB int
	}{
		{
			Count:  588,
			InputA: 65,
			InputB: 8921,
		},
		{
			Count:  569,
			InputA: 116,
			InputB: 299,
		},
	}

	for _, i := range data {
		count := duelingGenerators(i.InputA, i.InputB)

		if count != i.Count {
			t.Error("Expected", i.Count, "for input '", i.InputA, i.InputB, "...', got", count)
		}
	}
}

func BenchmarkDuelingGenerators(b *testing.B) {
	for n := 0; n < b.N; n++ {
		duelingGenerators(65, 8921)
	}
}

func TestDuelingGenerators2(t *testing.T) {
	data := []struct {
		Count  int
		InputA int
		InputB int
	}{
		{
			Count:  309,
			InputA: 65,
			InputB: 8921,
		},
		{
			Count:  298,
			InputA: 116,
			InputB: 299,
		},
	}

	for _, i := range data {
		count := duelingGenerators2(i.InputA, i.InputB)

		if count != i.Count {
			t.Error("Expected", i.Count, "for input '", i.InputA, i.InputB, "...', got", count)
		}
	}
}

func BenchmarkDuelingGenerators2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		duelingGenerators2(65, 8921)
	}
}
