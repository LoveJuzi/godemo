package leetcode1037

func isBoomerang(points [][]int) bool {
	A := points[0]
	B := points[1]
	C := points[2]

	if A[0] == B[0] && A[1] == B[1] {
		return false
	}
	if A[0] == C[0] && A[1] == C[1] {
		return false
	}
	if (B[0]-A[0])*(C[1]-A[1]) == (B[1]-A[1])*(C[0]-A[0]) {
		return false
	}

	return true
}
