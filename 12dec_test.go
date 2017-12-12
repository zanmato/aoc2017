package aoc2017

import (
	"io/ioutil"
	"testing"
)

func TestDigitalPlumber(t *testing.T) {
	inp, err := ioutil.ReadFile("input/12dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := []struct {
		Programs int
		Groups   int
		Input    string
	}{
		{
			Programs: 6,
			Groups:   2,
			Input:    "0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6\n5 <-> 6\n6 <-> 4, 5",
		},
		{
			Programs: 128,
			Groups:   209,
			Input:    string(inp),
		},
	}

	for _, i := range data {
		programs, groups := digitalPlumber(i.Input)

		if programs != i.Programs {
			t.Error("Expected", i.Programs, "for input '", i.Input[:20], "...', got", programs)
		}

		if groups != i.Groups {
			t.Error("Expected", i.Groups, "for input '", i.Input[:20], "...', got", groups)
		}
	}
}

func BenchmarkDigitalPlumber(b *testing.B) {
	for n := 0; n < b.N; n++ {
		digitalPlumber("0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6\n5 <-> 6\n6 <-> 4, 5")
	}
}
