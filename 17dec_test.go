package aoc2017

import (
	"testing"
)

func TestSpinlock(t *testing.T) {
	data := []struct {
		Result     int
		Insertions int
		Steps      int
	}{
		{
			Result:     638,
			Insertions: 2017,
			Steps:      3,
		},
		{
			Result:     417,
			Insertions: 2017,
			Steps:      348,
		},
	}

	for _, i := range data {
		if result := spinlock(i.Insertions, i.Steps); result != i.Result {
			t.Error("Expected", i.Result, "for input '", i.Insertions, "...', got", result)
		}
	}
}

func BenchmarkSpinlock(b *testing.B) {
	for n := 0; n < b.N; n++ {
		spinlock(2017, 348)
	}
}

func TestSpinlockAngry(t *testing.T) {
	data := []struct {
		Result     int
		Insertions int
		Steps      int
	}{
		{
			Result:     34334221,
			Insertions: 50000000,
			Steps:      348,
		},
	}

	for _, i := range data {
		if result := spinlockAngry(i.Insertions, i.Steps); result != i.Result {
			t.Error("Expected", i.Result, "for input '", i.Insertions, "...', got", result)
		}
	}
}

func BenchmarkSpinlockAngry(b *testing.B) {
	for n := 0; n < b.N; n++ {
		spinlockAngry(2017, 348)
	}
}
