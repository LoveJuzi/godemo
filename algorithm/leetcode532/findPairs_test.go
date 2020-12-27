package leetcode532

import (
	"testing"
)

func Test_findPairs(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{3, 1, 4, 1, 5}, 2, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 4},
		{[]int{1, 3, 1, 5, 4}, 0, 1},
		{[]int{1, 2, 3, 4, 5}, -1, 0},
		{[]int{1, 1, 1, 1, 1}, 0, 1},
	}

	for _, c := range cases {
		result := findPairs(c.nums, c.k)
		if result != c.expected {
			t.Fatalf("\n findPairs: nums is %v, k is %v, expected is %v, result is %v", c.nums, c.k, c.expected, result)
		}
	}
}
