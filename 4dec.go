package aoc2017

import (
	"sort"
	"strings"
)

type runeSlice []rune

func (p runeSlice) Len() int {
	return len(p)
}

func (p runeSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p runeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func checkPassphrases(i string) int {
	rows := strings.Split(i, "\n")

	s := 0
LoopDeLoop:
	for _, row := range rows {
		cols := strings.Split(strings.TrimSpace(row), " ")

		for i, cell := range cols {
			for j, c2 := range cols {
				if i == j {
					continue
				}

				if c2 == cell {
					continue LoopDeLoop
				}
			}
		}

		s++
	}

	return s
}

func checkAnagramPassphrases(i string) int {
	rows := strings.Split(i, "\n")

	s := 0
LoopDeLoop:
	for _, row := range rows {
		cols := strings.Split(strings.TrimSpace(row), " ")

		for i, cell := range cols {
			runes := []rune(cell)
			sort.Sort(runeSlice(runes))
			cols[i] = string(runes)
		}

		for i, cell := range cols {
			for j, c2 := range cols {
				if i == j {
					continue
				}

				if c2 == cell {
					continue LoopDeLoop
				}
			}
		}

		s++
	}

	return s
}
