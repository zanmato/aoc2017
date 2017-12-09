package aoc2017

import (
	"io/ioutil"
	"testing"
)

func TestCheckPassphrases(t *testing.T) {
	inp, err := ioutil.ReadFile("input/4dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := map[int]string{
		0:   "aa bb cc dd aa",
		1:   "aa bb cc dd aaa",
		455: string(inp),
	}

	for e, i := range data {
		if r := checkPassphrases(i); r != e {
			t.Error("Expected", e, "for input", i, ", got", r)
		}
	}
}

func BenchmarkCheckPassphrases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		checkPassphrases("aa bb cc dd ee")
	}
}

func TestCheckAnagramPassphrases(t *testing.T) {
	inp, err := ioutil.ReadFile("input/4dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := map[int]string{
		1:   "iiii oiii ooii oooi oooo",
		0:   "oiii ioii iioi iiio",
		186: string(inp),
	}

	for e, i := range data {
		if r := checkAnagramPassphrases(i); r != e {
			t.Error("Expected", e, "for input", i, ", got", r)
		}
	}
}

func BenchmarkCheckAnagramPassphrases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		checkAnagramPassphrases("abcde fghij")
	}
}
