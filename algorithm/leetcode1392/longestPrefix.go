package leetcode1392

func panding(s string, m int, l int) bool {
	for i, j := 0, m; j < l; i, j = i+1, j+1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func longestPrefix(s string) string {
	l := len(s)

	for i := l - 1; i > 0; i-- {
		if panding(s, l-i, l) == true {
			return s[:i]
		}
	}

	return ""
}
