package leetcode581

import "testing"

func Test_findUnsortedSubarray(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
		{[]int{1, 2, 3, 4}, 0},
	}

	for _, c := range cases {
		result := findUnsortedSubarray(c.nums)
		if result != c.expected {
			t.Fatalf("\n findUnsortedSubarray: nums is %v, expected is %v, result is %v \n", c.nums, c.expected, result)
		}
	}
}
