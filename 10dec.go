package aoc2017

import (
	"encoding/hex"
)

func knotHashRound(lengths []byte, rounds int) []int {
	var (
		listLen = 256
		n       int
		pos     int
		skip    int
		list    = make([]int, listLen)
	)

	for i := 0; i < listLen; i++ {
		list[i] = i
	}

	for j := 0; j < rounds; j++ {
		for _, k := range lengths {
			n = int(k) - 1
			for x := 0; x < n; x++ {
				list[(pos+x)%listLen], list[(pos+n)%listLen] = list[(pos+n)%listLen], list[(pos+x)%listLen]
				n--
			}
			pos += int(k) + skip
			skip++
		}
	}

	return list
}

func knotHash(inp string) string {
	var (
		b      []byte
		suffix = []byte{17, 31, 73, 47, 23}
		rounds = 64
		blocks = make([]byte, 16)
		s      int
	)

	b = []byte(inp)
	b = append(b, suffix...)

	l := knotHashRound(b, rounds)

	for i := 0; i < len(l); i += 16 {
		s = 0
		for k := 0; k < 16; k++ {
			s ^= l[i+k]
		}
		blocks[i/16] = byte(s)
	}

	return hex.EncodeToString(blocks)
}
