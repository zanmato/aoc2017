package aoc2017

import (
	"io/ioutil"
	"testing"
)

func TestDuet(t *testing.T) {
	inp, err := ioutil.ReadFile("input/18dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := []struct {
		Result int
		Input  string
	}{
		{
			Result: 4,
			Input:  "set a 1\nadd a 2\nmul a a\nmod a 5\nsnd a\nset a 0\nrcv a\njgz a -1\nset a 1\njgz a -2",
		},
		{
			Result: 8600,
			Input:  string(inp),
		},
	}

	for _, i := range data {
		if result := duet(i.Input); result != i.Result {
			t.Error("Expected", i.Result, "for input '", i.Input, "...', got", result)
		}
	}
}

func BenchmarkDuet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		duet("set a 1\nadd a 2\nmul a a\nmod a 5\nsnd a\nset a 0\nrcv a\njgz a -1\nset a 1\njgz a -2")
	}
}

func TestDuetTwo(t *testing.T) {
	inp, err := ioutil.ReadFile("input/18dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := []struct {
		Result int
		Input  string
	}{
		{
			Result: 3,
			Input:  "snd 1\nsnd 2\nsnd p\nrcv a\nrcv b\nrcv c\nrcv d",
		},
		{
			Result: 7239,
			Input:  string(inp),
		},
	}

	for _, i := range data {
		if result := duetTwo(i.Input); result != i.Result {
			t.Error("Expected", i.Result, "for input '", i.Input, "...', got", result)
		}
	}
}
