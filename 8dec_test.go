package aoc2017

import (
	"io/ioutil"
	"testing"
)

func TestGetLargestValueInRegister(t *testing.T) {
	inp, err := ioutil.ReadFile("8dec.txt")

	if err != nil {
		t.Error(err)
	}

	data := map[int]string{
		1:    "b inc 5 if a > 1\na inc 1 if b < 5\nc dec -10 if a >= 1\nc inc -20 if c == 10",
		5849: string(inp),
	}

	for e, i := range data {
		r, _, err := getLargestValueInRegister(i)
		if err != nil {
			t.Error(err)
		}

		if r != e {
			t.Error("Expected", e, "for input", i[:20], "...", "got", r)
		}
	}
}

func BenchmarkGetLargestValueInRegister(b *testing.B) {
	inp, err := ioutil.ReadFile("8dec.txt")
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		getLargestValueInRegister(string(inp))
	}
}

func TestGetLargestValueInRegisterPartTwo(t *testing.T) {
	inp, err := ioutil.ReadFile("8dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := map[int]string{
		10:   "b inc 5 if a > 1\na inc 1 if b < 5\nc dec -10 if a >= 1\nc inc -20 if c == 10",
		6702: string(inp),
	}

	for e, i := range data {
		_, r, err := getLargestValueInRegister(i)
		if err != nil {
			t.Error(err)
		}

		if r != e {
			t.Error("Expected", e, "for input", i[:20], "...", "got", r)
		}
	}
}
