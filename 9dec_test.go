package aoc2017

import (
	"io/ioutil"
	"testing"
)

func TestCalculateGroupScore(t *testing.T) {
	inp, err := ioutil.ReadFile("input/9dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := map[int]string{
		1:    "{<a>,<a>,<a>,<a>}",
		6:    "{{{}}}",
		16:   "{{{},{},{{}}}}",
		3:    "{{<a!>},{<a!>},{<a!>},{<ab>}}",
		7640: string(inp),
	}

	for e, i := range data {
		if r, _ := calculateGroupScore(i); r != e {
			t.Error("Expected", e, "for input", i, ", got", r)
		}
	}
}

func BenchmarkCalculateGroupScore(b *testing.B) {
	for n := 0; n < b.N; n++ {
		calculateGroupScore("{{<a!>},{<a!>},{<a!>},{<ab>}}")
	}
}

func TestCountGarbage(t *testing.T) {
	inp, err := ioutil.ReadFile("input/9dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := map[int]string{
		0:    "<!!!>>",
		10:   "<{o\"i!a,<{i<a>",
		3:    "<<<<>",
		17:   "<random characters>",
		4368: string(inp),
	}

	for e, i := range data {
		if _, r := calculateGroupScore(i); r != e {
			t.Error("Expected", e, "for input", i, ", got", r)
		}
	}
}
