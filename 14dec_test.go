package aoc2017

import (
	"testing"
)

func TestDiskDefragmentation(t *testing.T) {
	data := []struct {
		Used    int
		Regions int
		Input   string
	}{
		{
			Used:    8108,
			Regions: 1242,
			Input:   "flqrgnkx",
		},
		{
			Used:    8194,
			Regions: 1141,
			Input:   "uugsqrei",
		},
	}

	for _, i := range data {
		used, regions := diskDefragmentation(i.Input)

		if used != i.Used {
			t.Error("Expected", i.Used, "for input '", i.Input, "...', got", used)
		}

		if regions != i.Regions {
			t.Error("Expected", i.Regions, "for input '", i.Input, "...', got", regions)
		}
	}
}

func BenchmarkDiskDefragmentation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		diskDefragmentation("uugsqrei")
	}
}
