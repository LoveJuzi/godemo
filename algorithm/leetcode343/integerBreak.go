package leetcode343

func integerBreak(n int) int {
	return f(n, buildT(n))
}

func buildT(n int) []int {
	var T []int
	T = make([]int, n+1)
	T[1] = 1
	return T
}

// 获得 n 的最大值
func f(n int, T []int) int {
	if T[n] != 0 {
		return T[n]
	}

	var rt int

	rt = 1

	for i := 1; i <= n/2; i++ {
		t := g(f(i, T), i) * g(f(n-i, T), n-i)
		if rt < t {
			rt = t
		}
	}

	T[n] = rt

	return T[n]
}

func g(a, b int) int {
	if a > b {
		return a
	}

	return b
}
