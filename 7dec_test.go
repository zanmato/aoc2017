package aoc2017

import (
	"io/ioutil"
	"testing"
)

func TestGetRootNode(t *testing.T) {
	inp, err := ioutil.ReadFile("7dec.txt")
	if err != nil {
		t.Error(err)
	}

	e := "svugo"
	if r := getRootNode(string(inp)); r.Name != e {
		t.Error("Expected", e, "for input 7dec.txt, got", r)
	}
}

func BenchmarkGetRootNode(b *testing.B) {
	inp, err := ioutil.ReadFile("7dec.txt")
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		getRootNode(string(inp))
	}
}

func TestGetNewWeight(t *testing.T) {
	inp, err := ioutil.ReadFile("7dec.txt")
	if err != nil {
		t.Error(err)
	}

	e := 1152
	if r := getNewWeight(string(inp)); r != e {
		t.Error("Expected", e, "for input 7dec.txt, got", r)
	}
}

func BenchmarkGetNewWeight(b *testing.B) {
	inp, err := ioutil.ReadFile("7dec.txt")
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		getNewWeight(string(inp))
	}
}
