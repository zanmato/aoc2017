package aoc2017

import (
	"strconv"
	"strings"
)

func permutationPromenade(rounds int, input string, moves string) string {
	m := strings.Split(moves, ",")
	in := strings.Split(input, "")
	positions := map[string]int{}

	for j := 0; j < rounds; j++ {
		for _, move := range m {
			switch string(move[0]) {
			case "s":
				n, _ := strconv.Atoi(string(move[1:]))

				for i := 1; i <= n; i++ {
					in = append(in[len(in)-1:], in[0:len(in)-1]...)
				}
			case "p":
				p := strings.Split(move[1:], "/")
				a := indexOf(p[0], in)
				b := indexOf(p[1], in)
				in[a], in[b] = in[b], in[a]
			case "x":
				p := strings.Split(move[1:], "/")
				a, _ := strconv.Atoi(p[0])
				b, _ := strconv.Atoi(p[1])
				in[a], in[b] = in[b], in[a]
			}
		}

		if rounds > 1 {
			pos := strings.Join(in, "")
			if _, ok := positions[pos]; ok {
				for v, x := range positions {
					if x == (rounds%j)-1 {
						return v
					}
				}
			}
			positions[pos] = j
		}
	}

	return strings.Join(in, "")
}

func indexOf(find string, in []string) int {
	for i, v := range in {
		if v == find {
			return i
		}
	}

	return -1
}
