package leetcode485

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
	}

	for _, c := range cases {
		result := findMaxConsecutiveOnes(c.nums)
		if result != c.expected {
			t.Fatalf("\n findMaxConsecutiveOnes: nums is %v, expected is %v, result is %v \n", c.nums, c.expected, result)
		}
	}
}
