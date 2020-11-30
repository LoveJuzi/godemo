package leetcode139

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	if len(wordDict) == 0 {
		return false
	}

	return f(s, 0, buildHT(wordDict), buildT(len(s)))
}

func buildT(n int) []int {
	var T []int

	T = make([]int, n)

	return T
}

func buildHT(wordDict []string) map[string]struct{} {
	var HT map[string]struct{}

	HT = make(map[string]struct{})

	for i := 0; i < len(wordDict); i++ {
		HT[wordDict[i]] = struct{}{}
	}

	return HT
}

func f(s string, i int, HT map[string]struct{}, T []int) bool {
	if i >= len(s) {
		return true
	}

	if T[i] == 1 {
		return true
	}

	if T[i] == 2 {
		return false
	}

	T[i] = 2

	for ip := i; ip < len(s); ip++ {
		if _, ok := HT[s[i:ip+1]]; !ok {
			continue
		}

		if f(s, ip+1, HT, T) {
			T[i] = 1
			break
		}
	}

	if T[i] == 1 {
		return true
	}

	return false
}
