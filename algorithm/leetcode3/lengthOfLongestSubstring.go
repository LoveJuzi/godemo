package leetcode3

func lengthOfLongestSubstring(s string) int {
	i := 0
	var cnt int
	var rt int
	for i < len(s) {
		i, cnt = f(i, s)
		if rt < cnt {
			rt = cnt
		}
	}
	return rt
}

func f(i int, s string) (int, int) {
	T := map[byte]int{}

	ni := i + 1
	cnt := 0
	for j := i; j < len(s); j++ {
		if v, ok := T[s[j]]; ok {
			ni = v + 1
			break
		}
		T[s[j]] = i
		cnt++
	}

	return ni, cnt
}
