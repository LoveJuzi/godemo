package leetcode1037

import (
	"testing"
)

func Test_isBoomerang(t *testing.T) {
	cases := []struct {
		points   [][]int
		expected bool
	}{
		{[][]int{{1, 1}, {2, 3}, {3, 2}}, true},
		{[][]int{{1, 1}, {2, 2}, {3, 3}}, false},
	}

	for _, c := range cases {
		result := isBoomerang(c.points)
		if result != c.expected {
			t.Fatalf("\n isBoomerang: points is %v, expected is %v, result is %v \n", c.points, c.expected, result)
		}
	}
}
