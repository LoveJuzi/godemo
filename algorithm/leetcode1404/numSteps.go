package leetcode1404

func numSteps(s string) int {
	cnt := 0
	for {
		if s == "1" {
			break
		}

		cnt++

		if isOdd(s) {
			s = addOne(s)
		} else {
			s = diviedTwo(s)
		}
	}

	return cnt
}

func isOdd(s string) bool {
	return s[len(s)-1] == '1'
}

func diviedTwo(s string) string {
	return s[0 : len(s)-1]
}

func addOne(s string) string {
	rt := ""
	c := '1'
	b := []byte(s)
	for i := len(b) - 1; i >= 0 && c == '1'; i-- {
		if b[i] == '0' {
			b[i] = '1'
			c = '0'
		} else {
			b[i] = '0'
		}
	}

	if c == '1' {
		rt = "1"
	}

	rt += string(b)

	return rt
}
