package aoc2017

import (
	"io/ioutil"
	"testing"
)

func TestHexGridDistance(t *testing.T) {
	inp, err := ioutil.ReadFile("input/11dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := map[int]string{
		0:   "ne,ne,sw,sw",
		2:   "ne,ne,s,s",
		3:   "se,sw,se,sw,sw",
		784: string(inp),
	}

	for e, i := range data {
		if r, _ := hexGridDistance(i); r != e {
			t.Error("Expected", e, "for input", i, ", got", r)
		}
	}
}

func BenchmarkHexGridDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hexGridDistance("se,sw,se,sw,sw")
	}
}

func TestHexGridFurthestDistance(t *testing.T) {
	inp, err := ioutil.ReadFile("input/11dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := map[int]string{
		1558: string(inp),
	}

	for e, i := range data {
		if _, r := hexGridDistance(i); r != e {
			t.Error("Expected", e, "for input", i, ", got", r)
		}
	}
}
