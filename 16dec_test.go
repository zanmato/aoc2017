package aoc2017

import (
	"io/ioutil"
	"testing"
)

func TestPermutationPromenade(t *testing.T) {
	inp, err := ioutil.ReadFile("input/16dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := []struct {
		Result string
		Input  string
		Moves  string
		Rounds int
	}{
		{
			Result: "baedc",
			Input:  "abcde",
			Moves:  "s1,x3/4,pe/b",
			Rounds: 1,
		},
		{
			Result: "namdgkbhifpceloj",
			Input:  "abcdefghijklmnop",
			Moves:  string(inp),
			Rounds: 1,
		},
		{
			Result: "ibmchklnofjpdeag",
			Input:  "abcdefghijklmnop",
			Moves:  string(inp),
			Rounds: 1000000000,
		},
	}

	for _, i := range data {
		result := permutationPromenade(i.Rounds, i.Input, i.Moves)

		if result != i.Result {
			t.Error("Expected", i.Result, "for input '", i.Input, "...', got", result)
		}
	}
}

func BenchmarkPermutationPromenade(b *testing.B) {
	for n := 0; n < b.N; n++ {
		permutationPromenade(1, "abcde", "s1,x3/4,pe/b")
	}
}
