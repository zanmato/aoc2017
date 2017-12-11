package aoc2017

import (
	"math"
	"strings"
)

func hexGridDistance(i string) (int, int) {
	var (
		steps []string
		x     float64
		y     float64
		z     float64
		m     float64
	)

	steps = strings.Split(i, ",")

	for _, direction := range steps {
		switch direction {
		case "ne":
			x++
			z--
		case "n":
			z--
			y++
		case "nw":
			x--
			y++
		case "se":
			x++
			y--
		case "s":
			y--
			z++
		case "sw":
			z++
			x--
		}

		m = math.Max(m, (math.Abs(x)+math.Abs(y)+math.Abs(z))/2)
	}

	return int((math.Abs(x) + math.Abs(y) + math.Abs(z)) / 2), int(m)
}
