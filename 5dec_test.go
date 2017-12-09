package aoc2017

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func inputToIntArray(i string, t testing.TB) []int {
	t.Helper()
	rows := strings.Split(i, "\n")
	inst := make([]int, len(rows))
	for j, v := range rows {
		inst[j], _ = strconv.Atoi(strings.TrimSpace(v))
	}

	return inst
}

func TestCalculateInstructionSteps(t *testing.T) {
	inp, err := ioutil.ReadFile("input/5dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := map[int]string{
		5:      "0\n3\n0\n1\n-3",
		373160: string(inp),
	}

	for e, i := range data {
		in := inputToIntArray(i, t)
		if r := calculateInstructionSteps(in); r != e {
			t.Error("Expected", e, "for input", i, ", got", r)
		}
	}
}

func BenchmarkCalculateInstructionSteps(b *testing.B) {
	in := inputToIntArray("0\n3\n0\n1\n-3", b)
	for n := 0; n < b.N; n++ {
		calculateInstructionSteps(in)
	}
}

func TestCalculateInstructionStepsPartTwo(t *testing.T) {
	inp, err := ioutil.ReadFile("input/5dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := map[int]string{
		10:       "0\n3\n0\n1\n-3",
		26395586: string(inp),
	}

	for e, i := range data {
		in := inputToIntArray(i, t)
		if r := calculateInstructionSteps(in, true); r != e {
			t.Error("Expected", e, "for input", i, ", got", r)
		}
	}
}

func BenchmarkCalculateInstructionStepsPartTwo(b *testing.B) {
	in := inputToIntArray("0\n3\n0\n1\n-3", b)
	for n := 0; n < b.N; n++ {
		calculateInstructionSteps(in, true)
	}
}
