package aoc2017

import (
	"testing"
)

func TestKnotHashRound(t *testing.T) {
	data := map[int][]byte{
		38415: []byte{189, 1, 111, 246, 254, 2, 0, 120, 215, 93, 255, 50, 84, 15, 94, 62},
	}

	for e, i := range data {
		if r := knotHashRound(i, 1); r[0]*r[1] != e {
			t.Error("Expected", e, "for input", i, ", got", r[0]*r[1])
		}
	}
}

func BenchmarkKnotHashRound(b *testing.B) {
	inp := []byte{3, 4, 1, 5}
	for n := 0; n < b.N; n++ {
		knotHashRound(inp, 1)
	}
}

func TestKnotHash(t *testing.T) {
	data := map[string]string{
		"33efeb34ea91902bb2f59c9920caa6cd": "AoC 2017",
		"9de8846431eef262be78f590e39a4848": "189,1,111,246,254,2,0,120,215,93,255,50,84,15,94,62",
	}

	for e, i := range data {
		if r := knotHash(i); r != e {
			t.Error("Expected", e, "for input", i, ", got", r)
		}
	}
}

func BenchmarkKnot(b *testing.B) {
	for n := 0; n < b.N; n++ {
		knotHash("AoC 2017")
	}
}
