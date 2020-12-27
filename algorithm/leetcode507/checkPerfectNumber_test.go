package leetcode507

import "testing"

func Test_checkPerfectNumber(t *testing.T) {
	cases := []struct {
		num      int
		expected bool
	}{
		{28, true},
		{1, false},
	}

	for _, c := range cases {
		result := checkPerfectNumber(c.num)
		if result != c.expected {
			t.Fatalf("\n checkPerfectNumber: num is %v, expected is %v, result is %v \n", c.num, c.expected, result)
		}
	}
}
