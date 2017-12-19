package aoc2017

import (
	"strings"
)

func seriesOfTubes(input string) (string, int) {
	var (
		start int
		steps = 1
	)

	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))

	for y, line := range lines {
		grid[y] = make([]string, len(line))
		for x, c := range line {
			grid[y][x] = string(c)

			if y == 0 && string(c) == "|" {
				start = x
			}
		}
	}

	n := 0
	x := start
	y := 0
	direction := 1
	letters := []string{}

	yMax := len(grid)
	xMax := len(grid[0])

NotAnNeverEndingLoop:
	for n < 1000 {
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

		if x < 0 || x >= xMax || y < 0 || y >= yMax {
			// Out of bounds
			break
		}

		switch grid[y][x] {
		case " ":
			// 'Tis the end
			break NotAnNeverEndingLoop
		case "|", "-":
			// Continue on the path
			steps++
			continue
		case "+":
			// Time to switch direction
			steps++
			for j := 0; j < 4; j++ {
				if j == direction || direction%2 == j%2 {
					continue
				}

				if j == 0 && x+1 < xMax && (grid[y][x+1] == "-" || isUppercaseLetter(grid[y][x+1])) {
					direction = j
					break
				} else if j == 1 && y+1 < yMax && (grid[y+1][x] == "|" || isUppercaseLetter(grid[y+1][x])) {
					direction = j
					break
				} else if j == 2 && x-1 >= 0 && (grid[y][x-1] == "-" || isUppercaseLetter(grid[y][x-1])) {
					direction = j
					break
				} else if j == 3 && y-1 >= 0 && (grid[y-1][x] == "|" || isUppercaseLetter(grid[y-1][x])) {
					direction = j
					break
				}
			}
		default:
			if isUppercaseLetter(grid[y][x]) {
				steps++
				// It's a letter!
				letters = append(letters, grid[y][x])
			}
		}

		n++
	}

	return strings.Join(letters, ""), steps
}

func isUppercaseLetter(input string) bool {
	return input[0] >= 65 && input[0] <= 90
}
