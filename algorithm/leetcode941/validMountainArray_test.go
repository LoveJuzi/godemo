package leetcode941

import "testing"

func Test_validMountainArray(t *testing.T) {
	cases := []struct {
		A        []int
		expected bool
	}{
		{[]int{2, 1}, false},
		{[]int{3, 5, 5}, false},
		{[]int{0, 3, 2, 1}, true},
	}

	for _, c := range cases {
		result := validMountainArray(c.A)
		if result != c.expected {
			t.Fatalf("\n validMountainArray: A is %v, expected is %v, result is %v \n", c.A, c.expected, result)
		}
	}
}
