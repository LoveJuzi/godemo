package leetcode54

func spiralOrder(matrix [][]int) []int {
	M := len(matrix)
	if M == 0 {
		return []int{}
	}
	N := len(matrix[0])
	rt := []int{}
	i, j, a, b, c, d, direct := 0, 0, -1, -1, N, M, 0
	ok := true
	for ok {
		rt = append(rt, matrix[i][j])
		i, j, a, b, c, d, direct, ok = f(i, j, a, b, c, d, direct, 2)
	}
	return rt
}

func f(i, j, a, b, c, d int, direct int, cnt int) (int, int, int, int, int, int, int, bool) {
	if cnt == 0 {
		return i, j, a, b, c, d, direct, false
	}

	if direct == 0 {
		j = j + 1
	} else if direct == 1 {
		i = i + 1
	} else if direct == 2 {
		j = j - 1
	} else if direct == 3 {
		i = i - 1
	}

	if g(i, j, a, b, c, d) {
		return i, j, a, b, c, d, direct, true
	}

	if direct == 0 {
		j = j - 1
		direct = 1
		b = b + 1
	} else if direct == 1 {
		i = i - 1
		direct = 2
		c = c - 1
	} else if direct == 2 {
		j = j + 1
		direct = 3
		d = d - 1
	} else if direct == 3 {
		i = i + 1
		direct = 0
		a = a + 1
	}

	return f(i, j, a, b, c, d, direct, cnt-1)
}

func g(i, j, a, b, c, d int) bool {
	return a < j && j < c && b < i && i < d
}
