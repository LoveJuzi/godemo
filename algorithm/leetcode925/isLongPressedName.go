package leetcode925

func isLongPressedName(name string, typed string) bool {
	N := len(name)
	M := len(typed)

	j := 0
	for i := 0; i < M; i++ {
		if j < N {
			if name[j] == typed[i] {
				j++
				continue
			}
			if j-1 < 0 {
				return false
			}
		}
		if name[j-1] != typed[i] {
			return false
		}
	}

	return j == N
}
