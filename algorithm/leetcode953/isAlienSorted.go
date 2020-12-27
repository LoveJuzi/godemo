package leetcode953

func compare(a string, b string, T map[byte]int) bool {
	i := 0
	j := 0
	M := len(a)
	N := len(b)
	for {
		if i == M {
			return true
		}
		if j == N {
			return false
		}

		if T[a[i]] > T[b[j]] {
			return false
		} else if T[a[i]] < T[b[j]] {
			return true
		}
		i++
		j++
	}
}

func isAlienSorted(words []string, order string) bool {
	N := len(order)
	T := map[byte]int{}
	for i := 0; i < N; i++ {
		T[order[i]] = i
	}

	a := ""
	b := ""
	for i := 1; i < len(words); i++ {
		a = words[i-1]
		b = words[i]
		if !compare(a, b, T) {
			return false
		}
	}

	return true
}
