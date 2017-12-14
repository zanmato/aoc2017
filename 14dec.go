package aoc2017

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func knotHashToBin(h string) string {
	bi := big.NewInt(0)
	bi.SetString(h, 16)

	return fmt.Sprintf("%0128b", bi)
}

func diskDefragmentation(i string) (int, int) {
	var (
		used    int
		regions int
		grid    [128][128]int
		b       string
	)

	for j := 0; j < 128; j++ {
		b = knotHashToBin(knotHash(i + "-" + strconv.Itoa(j)))
		used += len(strings.Replace(b, "0", "", -1))

		for x, v := range b {
			grid[j][x], _ = strconv.Atoi(string(v))
		}
	}

	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			if grid[y][x] == 1 {
				regions++
				emptyAdjacent(y, x, &grid)
			}
		}
	}

	return used, regions
}

func emptyAdjacent(y int, x int, grid *[128][128]int) {
	if grid[y][x] == 0 {
		return
	}

	grid[y][x] = 0

	// Look ahead
	if x+1 < 128 {
		emptyAdjacent(y, x+1, grid)
	}

	// Look behind
	if x-1 >= 0 {
		emptyAdjacent(y, x-1, grid)
	}

	// Look down
	if y+1 < 128 {
		emptyAdjacent(y+1, x, grid)
	}

	// Look up
	if y-1 >= 0 {
		emptyAdjacent(y-1, x, grid)
	}
}
