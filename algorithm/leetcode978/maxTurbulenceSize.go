package leetcode978

func main() {

}

func maxTurbulenceSize(A []int) int {
	if len(A) == 0 {
		return 0
	}

	l := 0
	cnt := 0
	bf := [...]int{0, 0}
	for i := 1; i < len(A); i++ {
		bf[0] = bf[1]
		if A[i-1] < A[i] {
			bf[1] = -1
		} else if A[i-1] > A[i] {
			bf[1] = 1
		} else {
			bf[1] = 0
		}

		if bf[0] == 1 && bf[1] == -1 {
			cnt++
		} else if bf[0] == -1 && bf[1] == 1 {
			cnt++
		} else if bf[1] != 0 {
			cnt = 1
		} else {
			cnt = 0
		}

		if l < cnt {
			l = cnt
		}
	}

	return l + 1
}
