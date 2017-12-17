package aoc2017

func spinlock(insertions int, steps int) int {
	var (
		ins = []int{0}
		i   int
		j   int
	)

	for x := 1; x <= insertions; x++ {
		j = (i + steps) % len(ins)
		ins = append(ins, 0)
		copy(ins[j+1:], ins[j:])
		ins[j+1] = x
		i = j + 1
	}

	return ins[i+1]
}

func spinlockAngry(insertions int, steps int) int {
	var (
		v int
		j int
	)

	for x := 1; x <= insertions; x++ {
		j = (j+steps)%x + 1
		if j == 1 {
			v = x
		}
	}

	return v
}
