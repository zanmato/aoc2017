package aoc2017

import (
	"fmt"
)

func calculateStepsUntilInfiniteLoop(inst []int, rounds int) int {
	var (
		x               int
		mi              int
		mv              int
		current         string
		previousRecords = []string{
			fmt.Sprint(inst, ","),
		}
	)

	for {
		mi = 0
		mv = 0

		for j, v := range inst {
			// Get max index
			if v > mv {
				mi = j
				mv = v
			}
		}

		inst[mi] = 0
		x = mi

		for mv > 0 {
			x = (x + 1) % len(inst)
			inst[x]++
			mv--
		}

		current = fmt.Sprint(inst, ",")
		for _, p := range previousRecords {
			if p == current {
				rounds--
				if rounds == 0 {
					return len(previousRecords)
				}

				previousRecords = []string{
					current,
				}
			}
		}

		previousRecords = append(previousRecords, current)
	}
}
