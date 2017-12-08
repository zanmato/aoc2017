package aoc2017

import (
	"strconv"
)

func solveCaptcha(captcha string) int {
	r := []rune(captcha)

	p := -1
	s := 0
	for _, c := range r {
		n, _ := strconv.Atoi(string(c))
		if n == p {
			s += n
		}
		p = n
	}

	if r[0] == r[len(r)-1] {
		n, _ := strconv.Atoi(string(r[0]))
		s += n
	}

	return s
}

func solveHalfwayCaptcha(captcha string) int {
	l := len(captcha) / 2
	r := []rune(captcha + captcha[0:l])
	k := len(r)

	s := 0
	for i, c := range r {
		n, _ := strconv.Atoi(string(c))
		if i+l < k && c == r[i+l] {
			s += n
		}
	}

	return s
}
