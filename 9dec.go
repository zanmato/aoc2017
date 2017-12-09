package aoc2017

func calculateGroupScore(i string) (int, int) {
	var (
		s           int
		groupscore  int
		garbage     int
		r           string
		skipNext    bool
		skipGarbage bool
	)

	runes := []rune(i)
	skipNext = false

	for _, rune := range runes {
		r = string(rune)

		if skipNext {
			skipNext = false
			continue
		}

		switch r {
		case "!":
			skipNext = true
		case ">":
			skipGarbage = false
			continue
		}

		if skipGarbage {
			if !skipNext {
				garbage++
			}
			continue
		}

		switch r {
		case "{":
			groupscore++
			s += groupscore
		case "}":
			groupscore--
		case "<":
			skipGarbage = true
		}
	}

	return s, garbage
}
