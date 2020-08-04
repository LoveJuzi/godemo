package leetcode1496

func isPathCrossing(path string) bool {
	T := map[[2]int]struct{}{}

	pos := [2]int{0, 0}
	T[pos] = struct{}{}
	for _, v := range path {
		pos = nextpos(v, pos)
		if _, ok := T[pos]; ok {
			return true
		}
		T[pos] = struct{}{}
	}

	return false
}

func nextpos(order rune, pos [2]int) [2]int {
	nextpos := [2]int{}
	copy(nextpos[:], pos[:])

	if order == 'N' {
		nextpos[1]++
	} else if order == 'W' {
		nextpos[0]--
	} else if order == 'S' {
		nextpos[1]--
	} else {
		nextpos[0]++
	}

	return nextpos
}
