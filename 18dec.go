package aoc2017

import (
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func duet(input string) int {
	var (
		register     = map[string]int{}
		instructions = strings.Split(input, "\n")
		p            []string
		v            int
		lastPlayed   int
		err          error
	)

	for i := 0; i < len(instructions); i++ {
		p = strings.Split(instructions[i], " ")

		v = 0
		if len(p) == 3 {
			v, err = strconv.Atoi(p[2])
			if err != nil {
				v = register[p[2]]
			}
		}

		switch p[0] {
		case "set":
			register[p[1]] = v
		case "mul":
			register[p[1]] *= v
		case "jgz":
			if register[p[1]] > 0 {
				i--
				i += v
			}
		case "add":
			register[p[1]] += v
		case "mod":
			register[p[1]] %= v
		case "snd":
			lastPlayed = register[p[1]]
		case "rcv":
			if register[p[1]] > 0 {
				return lastPlayed
			}
		}
	}

	return 0
}

func duetTwo(input string) int {
	var (
		instructions = strings.Split(input, "\n")
		count        uint64
		wg           sync.WaitGroup
	)

	dp0 := make(chan int, 100)
	dp1 := make(chan int, 100)

	wg.Add(2)

	go duetProgram(0, instructions, dp1, dp0, &wg, nil)
	go duetProgram(1, instructions, dp0, dp1, &wg, &count)

	wg.Wait()

	return int(count)
}

func duetProgram(num int, instructions []string, snd chan<- int, rcv <-chan int, wg *sync.WaitGroup, count *uint64) {
	var (
		register = map[string]int{"p": num}
		p        []string
		v        int
		err      error
		started  bool
	)

	for i := 0; i < len(instructions); i++ {
		p = strings.Split(instructions[i], " ")

		v = 0
		if len(p) == 3 {
			v, err = strconv.Atoi(p[2])
			if err != nil {
				v = register[p[2]]
			}
		}

		switch p[0] {
		case "set":
			register[p[1]] = v
		case "mul":
			register[p[1]] *= v
		case "jgz":
			reg, err := strconv.Atoi(p[1])
			if err == nil && reg > 0 {
				i--
				i += v
			} else if register[p[1]] > 0 {
				i--
				i += v
			}
		case "add":
			register[p[1]] += v
		case "mod":
			register[p[1]] %= v
		case "snd":
			if count != nil {
				atomic.AddUint64(count, 1)
			}

			v, err = strconv.Atoi(p[1])
			if err != nil {
				v = register[p[1]]
			}

			if !started {
				wg.Done()
				started = true
			}

			wg.Add(1)

			snd <- v
		case "rcv":
			r := <-rcv
			register[p[1]] = r
			wg.Done()
		}
	}
}
