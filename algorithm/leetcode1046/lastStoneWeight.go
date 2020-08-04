package leetcode1046

func lastStoneWeight(stones []int) int {
	build(stones)
	l := len(stones)
	var A int
	var B int
	for l > 1 {
		A = stones[0]
		stones[0] = stones[l-1]
		adjust(stones, 0, l-1)
		l--
		B = stones[0]
		if A == B {
			stones[0] = stones[l-1]
			adjust(stones, 0, l-1)
			l--
			continue
		}
		stones[0] = A - B
		adjust(stones, 0, l)
	}
	if l == 0 {
		return 0
	}
	return stones[0]
}

func top(A []int, l int) (int, int) {
	v := A[0]
	l--
	A[0], A[l] = A[l], A[0]
	adjust(A, 0, l)
	return v, l
}

func f(A []int, idx int, l int) int {
	nidx := idx
	L := (idx << 1) + 1
	R := (idx << 1) + 2
	if L < l && A[L] > A[nidx] {
		nidx = L
	}
	if R < l && A[R] > A[nidx] {
		nidx = R
	}
	return nidx
}

func adjust(A []int, idx int, l int) {
	nidx := f(A, idx, l)
	if nidx == idx {
		return
	}
	A[idx], A[nidx] = A[nidx], A[idx]
	adjust(A, nidx, l)
}

func build(A []int) {
	for i := len(A) >> 1; i >= 0; i-- {
		adjust(A, i, len(A))
	}
}
