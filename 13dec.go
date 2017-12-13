package aoc2017

import (
	"strconv"
	"strings"
)

func packetScanners(i string) (int, int) {
	var (
		layers   = map[int]int{}
		max      int
		severity int
		delay    int
	)

	// Parse input
	rows := strings.Split(i, "\n")

	for _, row := range rows {
		p := strings.Split(row, ": ")
		layer, _ := strconv.Atoi(p[0])
		depth, _ := strconv.Atoi(p[1])

		layers[layer] = depth
		if layer > max {
			max = layer
		}
	}

DelayLoop:
	for ; ; delay++ {
		for i := 0; i <= max; i++ {
			if depth, exists := layers[i]; exists {
				if (i+delay)%((depth-1)*2) == 0 {
					if delay == 0 {
						severity += i * depth
						continue
					}
					continue DelayLoop
				}
			}
		}

		if delay > 0 {
			break
		}
	}

	return severity, delay
}
