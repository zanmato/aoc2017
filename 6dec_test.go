package aoc2017

import (
	"strconv"
	"strings"
	"testing"
)

func tabInputToIntArray(i string, t testing.TB) []int {
	t.Helper()
	cols := strings.Split(i, "\t")
	inst := make([]int, len(cols))
	for j, v := range cols {
		inst[j], _ = strconv.Atoi(strings.TrimSpace(v))
	}

	return inst
}

func TestCalculateStepsUntilInfiniteLoop(t *testing.T) {
	data := map[string]int{
		`0	2	7	0`: 5,
		`5	1	10	0	1	7	13	14	3	12	8	10	7	12	0	6`: 5042,
	}

	for i, e := range data {
		in := tabInputToIntArray(i, t)
		if r := calculateStepsUntilInfiniteLoop(in, 1); r != e {
			t.Error("Expected", e, "for input", i, ", got", r)
		}
	}
}

func BenchmarkCalculateStepsUntilInfiniteLoop(b *testing.B) {
	in := tabInputToIntArray("5	1	10	0	1	7	13	14	3	12	8	10	7	12	0	6", b)
	for n := 0; n < b.N; n++ {
		calculateStepsUntilInfiniteLoop(in, 1)
	}
}

func TestCalculateStepsUntilInfiniteLoopPartTwo(t *testing.T) {
	data := map[string]int{
		`0	2	7	0`: 4,
		`5	1	10	0	1	7	13	14	3	12	8	10	7	12	0	6`: 1086,
	}

	for i, e := range data {
		in := tabInputToIntArray(i, t)
		if r := (calculateStepsUntilInfiniteLoop(in, 2) - 1); r != e {
			t.Error("Expected", e, "for input", i, ", got", r)
		}
	}
}

func BenchmarkCalculateStepsUntilInfiniteLoopPartTwo(b *testing.B) {
	in := tabInputToIntArray("0	2	7	0", b)
	for n := 0; n < b.N; n++ {
		calculateStepsUntilInfiniteLoop(in, 2)
	}
}
