package aoc2017

import (
	"io/ioutil"
	"testing"
)

func TestSeriesOfTubes(t *testing.T) {
	data := []struct {
		Result string
		Steps  int
		Input  string
	}{
		{
			Result: "ABCDEF",
			Steps:  38,
			Input:  "input/19dec1.txt",
		},
		{
			Result: "GEPYAWTMLK",
			Steps:  17628,
			Input:  "input/19dec2.txt",
		},
	}

	for _, i := range data {
		inp, err := ioutil.ReadFile(i.Input)
		if err != nil {
			t.Error(err)
		}

		result, steps := seriesOfTubes(string(inp))
		if result != i.Result {
			t.Error("Expected", i.Result, "for input '", i.Input, "', got", result)
		}

		if steps != i.Steps {
			t.Error("Expected", i.Steps, "for input '", i.Steps, "', got", steps)
		}
	}
}

func BenchmarkSeriesOfTubes(b *testing.B) {
	inp, err := ioutil.ReadFile("input/19dec1.txt")
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		seriesOfTubes(string(inp))
	}
}
