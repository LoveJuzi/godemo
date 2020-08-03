package leetcode862

// 滑动窗口算法，但是我没有写出来
func shortestSubarray(A []int, K int) int {
	rt := -1

	s := 0
	l := 0
	for i := 0; i < len(A); i++ {
		if s <= 0 {
			s = A[i]
			l = 1
		} else {
			s += A[i]
			if rt != -1 {
				l++
				if l == rt+1 {
					s -= A[i-rt]
					l--
				}

				for j := i - l + 1; j < i; j++ {
					if A[j] <= 0 {
						s -= A[j]
						l--
					} else {
						break
					}
				}
			}
		}

		if s >= K || rt != -1 {
			l, _ = f2(A, i, K)
			if rt == -1 {
				rt = l
			} else if rt > l {
				rt = l
			}
		}
	}

	return rt
}

func f2(A []int, k int, K int) (int, int) {
	var l int
	var s int
	for i := k; i >= 0; i-- {
		l++
		s += A[i]
		if s >= K {
			break
		}
	}

	return l, s
}
