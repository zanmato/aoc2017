package aoc2017

import (
	"math"
)

func calculateManhattanDistance(i int) int {
	var (
		steps int
		turns int
		x     float64
		y     float64
	)

	for steps < i-1 {
		length := (turns / 2) + 1
		for j := 0; j < length; j++ {
			if steps == i-1 {
				break
			}
			steps++
			direction := turns % 4
			switch direction {
			case 0:
				x++
			case 1:
				y++
			case 2:
				x--
			case 3:
				y--
			}
		}
		turns++
	}

	return int(math.Abs(x) + math.Abs(y))
}

func calculateLargerValue(i int) int {
	var (
		x         int
		y         int
		direction int
		turns     int
		grid      = map[int]map[int]int{
			0: map[int]int{
				0: 1,
			},
		}
	)

	for true {
		length := (turns / 2) + 1
		for j := 0; j < length; j++ {
			switch direction {
			case 0:
				x++
			case 1:
				y++
			case 2:
				x--
			case 3:
				y--
			}

			sum := 0

			if val, exists := grid[x+1][y]; exists {
				sum += val
			}
			if val, exists := grid[x+1][y+1]; exists {
				sum += val
			}
			if val, exists := grid[x][y+1]; exists {
				sum += val
			}
			if val, exists := grid[x-1][y+1]; exists {
				sum += val
			}
			if val, exists := grid[x-1][y]; exists {
				sum += val
			}
			if val, exists := grid[x-1][y-1]; exists {
				sum += val
			}
			if val, exists := grid[x][y-1]; exists {
				sum += val
			}
			if val, exists := grid[x+1][y-1]; exists {
				sum += val
			}

			if sum > i {
				return sum
			}

			if _, exists := grid[x]; !exists {
				grid[x] = map[int]int{}
			}
			grid[x][y] = sum
		}

		turns++
		direction = turns % 4
	}

	return 1
}
