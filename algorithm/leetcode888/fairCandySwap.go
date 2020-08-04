package leetcode888

func fairCandySwap(A []int, B []int) []int {
	ac := 0
	T := map[int]struct{}{}
	rt := [2]int{}

	func() {
		s1 := 0
		s2 := 0
		for _, v := range A {
			s1 += v
		}
		for _, v := range B {
			s2 += v
		}
		ac = (s1 - s2) >> 1
	}()

	func() {
		for _, v := range B {
			T[v] = struct{}{}
		}
	}()

	func() {
		for _, v := range A {
			if _, ok := T[v-ac]; ok {
				rt[0] = v
				rt[1] = v - ac
				break
			}
		}
	}()

	return rt[:]
}
