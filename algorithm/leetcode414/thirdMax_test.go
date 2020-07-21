package leetcode414

import "testing"

func Test_thirdMax(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 2, 3, 1}, 1},
	}

	for _, c := range cases {
		result := thirdMax(c.nums)
		if result != c.expected {
			t.Fatalf("\n thirdMax: nums is %v, expected is %d, result is %d \n", c.nums, c.expected, result)
		}
	}
}
