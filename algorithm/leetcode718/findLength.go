package leetcode718

func findLength(A []int, B []int) int {
	n1, n2 := len(A), len(B)

	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	res := 0
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if A[i] == B[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				res = max(res, dp[i+1][j+1])
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func findLength(A []int, B []int) int {
// 	N := len(A)
// 	M := len(B)
// 	T := make(map[int]map[int]int)

// 	for i := 0; i < N; i++ {
// 		if _, ok := T[A[i]]; ok {
// 			continue
// 		}
// 		T[A[i]] = map[int]int{}
// 		for j := 0; j < M; j++ {
// 			if A[i] == B[j] {
// 				T[A[i]][j] = 1
// 			}
// 		}
// 	}

// 	rt := 0
// 	cm := 1
// 	c := 1
// 	i := 0
// 	pk := 0
// 	nk := 0
// 	for i < N {
// 		cm = 0
// 		for k := range T[A[i]] {
// 			c = 1
// 			nk = k + 1
// 			for j := i + 1; j < N; j++ {
// 				if _, ok := T[A[j]][nk]; !ok {
// 					break
// 				}
// 				nk++
// 				c++
// 			}
// 			pk = k - 1
// 			for j := i - 1; j >= 0; j-- {
// 				if _, ok := T[A[j]][pk]; !ok {
// 					break
// 				}
// 				pk--
// 				c++
// 			}
// 			if cm < c {
// 				cm = c
// 			}
// 		}
// 		if cm <= 1 {
// 			i++
// 		} else {
// 			i += cm - 1
// 		}
// 		// i++
// 		if rt < cm {
// 			rt = cm
// 		}
// 	}

// 	return rt
// }
