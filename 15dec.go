package aoc2017

func duelingGenerators(a int, b int) int {
	const (
		factorA = 16807
		factorB = 48271
		validA  = 4
		validB  = 8
		divisor = 2147483647
	)

	var (
		count int
	)

	for j := 0; j < 40000000; j++ {
		a = (a * factorA) % divisor
		b = (b * factorB) % divisor

		if int16(a&0xFFFF) == int16(b&0xFFFF) {
			count++
		}
	}

	return count
}

func duelingGenerators2(a int, b int) int {
	const (
		factorA = 16807
		factorB = 48271
		validA  = 4
		validB  = 8
		divisor = 2147483647
	)

	var (
		count   int
		av      bool
		bv      bool
		matches int
	)

	for matches < 5000000 {
		if !av {
			a = (a * factorA) % divisor
		}

		if !bv {
			b = (b * factorB) % divisor
		}

		if a%validA == 0 {
			av = true
		}

		if b%validB == 0 {
			bv = true
		}

		if av && bv {
			matches++
			av = false
			bv = false
			if int16(a&0xFFFF) == int16(b&0xFFFF) {
				count++
			}
		}
	}

	return count
}
