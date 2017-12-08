package aoc2017

func calculateInstructionSteps(inst []int, partTwo ...bool) int {
	var (
		s int
		v int
		n int
	)

	for n >= 0 && n < len(inst) {
		v = inst[n]

		// Increment or decrement!?!!!
		if v >= 3 && len(partTwo) > 0 {
			inst[n]--
		} else {
			inst[n]++
		}
		n += v
		s++
	}
	return s
}
