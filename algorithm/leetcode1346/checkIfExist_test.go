package leetcode1346

import "testing"

func Test_checkIfExist(t *testing.T) {
	cases := []struct {
		arr      []int
		expected bool
	}{
		{[]int{10, 2, 5, 3}, true},
		{[]int{7, 1, 14, 11}, true},
		{[]int{3, 1, 7, 11}, false},
	}

	for _, c := range cases {
		result := checkIfExist(c.arr)
		if result != c.expected {
			t.Fatalf("\n checkIfExist: arr is %v, expected is %v, result is %v \n", c.arr, c.expected, result)
		}
	}
}
