package leetcode914

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b>0  {
		a, b = b, a%b
	}
	return a
}

func hasGroupsSizeX(deck []int) bool {
	T := map[int]int{}
	for _, v := range deck {
		T[v]++
	}

	tmp := 0
	for _, v := range T{
		tmp = gcd(tmp, v)
	}

	if tmp <= 1 {
		return false
	}

	return true
}