package aoc2017

import (
	"strconv"
	"strings"
)

func digitalPlumber(i string) (int, int) {
	var (
		connectedPrograms = map[int]map[int]bool{0: {}}
	)

	// Parse input
	visits := strings.Split(i, "\n")
	programs := make([][]int, len(visits))
	for _, visit := range visits {
		p := strings.Split(visit, " <-> ")
		prog, _ := strconv.Atoi(p[0])

		l := strings.Split(p[1], ", ")
		vp := make([]int, len(l))
		for i, v := range l {
			vp[i], _ = strconv.Atoi(v)
		}

		programs[prog] = vp
	}

ProgramLoop:
	for p := range programs {
		for cp, sps := range connectedPrograms {
			if cp == p {
				continue
			}

			for sp := range sps {
				if sp == p {
					continue ProgramLoop
				}
			}
		}
		connectedPrograms[p] = map[int]bool{}
		findConnections(p, programs, connectedPrograms[p])
	}

	return len(connectedPrograms[0]), len(connectedPrograms)
}

func findConnections(index int, programs [][]int, connectedPrograms map[int]bool) {
	sps := programs[index]
	for _, sp := range sps {
		// Prevent infinite loop
		if _, ex := connectedPrograms[sp]; !ex {
			connectedPrograms[sp] = true
			findConnections(sp, programs, connectedPrograms)
		}
	}
}
