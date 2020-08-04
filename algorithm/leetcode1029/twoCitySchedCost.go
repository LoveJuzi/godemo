package leetcode1029

import (
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	AORB := false
	diff := 0
	cost := 0
	sum := 0
	delta := 0

	A := []int{}
	B := []int{}

	for _, v := range costs {
		cost, AORB, diff = f(v)
		sum += cost
		if AORB {
			A = append(A, diff)
		} else {
			B = append(B, diff)
		}
	}

	if len(A) == len(B) {
		return sum
	}

	if len(A) < len(B) {
		A, B = B, A
	}
	delta = len(A) - len(B)
	delta >>= 1
	sort.Ints(A)
	for i := 0; i < delta; i++ {
		sum += A[i] // extra cost
	}

	return sum
}

func f(a []int) (int, bool, int) {
	if a[0] <= a[1] {
		return a[0], true, a[1] - a[0]
	}

	return a[1], false, a[0] - a[1]
}
