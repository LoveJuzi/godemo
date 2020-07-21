package leetcode1486

import "testing"

func Test_xorOperation(t *testing.T) {
	cases := []struct {
		n        int
		start    int
		expected int
	}{
		{5, 0, 8},
		{4, 3, 8},
		{1, 7, 7},
		{10, 5, 2},
	}

	for _, c := range cases {
		result := xorOperation(c.n, c.start)
		if result != c.expected {
			t.Fatalf("\n xorOperation: n is %d, start is %d, excepted is %d, result is %d\n", c.n, c.start, c.expected, result)
		}
	}
}
