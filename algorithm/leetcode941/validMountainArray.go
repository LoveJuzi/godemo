package leetcode941

func validMountainArray(A []int) bool {
	N := len(A)

	if N < 3 || A[0] >= A[1] || A[N-2] <= A[N-1] {
		return false
	}

	i := 2
	// 检查严格递增
	for i < N {
		if A[i-1] < A[i] {
			i++
		} else {
			break
		}
	}

	if A[i-1] == A[i] {
		return false
	}

	// 检查严格递减
	for i < N-2 {
		if A[i] > A[i+1] {
			i++
		} else {
			return false
		}
	}

	return true
}
