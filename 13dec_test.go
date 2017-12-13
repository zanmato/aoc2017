package aoc2017

import (
	"io/ioutil"
	"testing"
)

func TestPacketScanners(t *testing.T) {
	inp, err := ioutil.ReadFile("input/13dec.txt")
	if err != nil {
		t.Error(err)
	}

	data := []struct {
		Severity int
		Delay    int
		Input    string
	}{
		{
			Severity: 24,
			Delay:    10,
			Input:    "0: 3\n1: 2\n4: 4\n6: 4",
		},
		{
			Severity: 2604,
			Delay:    3941460,
			Input:    string(inp),
		},
	}

	for _, i := range data {
		severity, delay := packetScanners(i.Input)

		if severity != i.Severity {
			t.Error("Expected", i.Severity, "for input '", i.Input[:10], "...', got", severity)
		}

		if delay != i.Delay {
			t.Error("Expected", i.Delay, "for input '", i.Input[:10], "...', got", delay)
		}
	}
}

func BenchmarkPacketScanners(b *testing.B) {
	for n := 0; n < b.N; n++ {
		packetScanners("0: 3\n1: 2\n4: 4\n6: 4")
	}
}
