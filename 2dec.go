package aoc2017

import (
	"strconv"
	"strings"
)

func calculateChecksum(i string) int {
	rows := strings.Split(i, "\n")

	s := 0
	for _, row := range rows {
		cols := strings.Split(row, "\t")

		min := 9999
		max := 0
		for _, cell := range cols {
			n, err := strconv.Atoi(cell)
			if err != nil {
				continue
			}

			if n < min {
				min = n
			}

			if n > max {
				max = n
			}
		}
		s += max - min
	}

	return s
}

func calculateDivisibleChecksum(i string) int {
	rows := strings.Split(i, "\n")

	s := 0
	for _, row := range rows {
		cols := strings.Split(row, "\t")

		for i, cell := range cols {
			n, err := strconv.Atoi(cell)
			if err != nil {
				continue
			}

			for j, c2 := range cols {
				if i == j {
					continue
				}

				n2, err := strconv.Atoi(c2)
				if err != nil {
					continue
				}

				if n2%n == 0 {
					s += n2 / n
					break
				}
			}
		}
	}

	return s
}
