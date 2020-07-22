package leetcode605

import (
	"testing"
)

func Test_canPlaceFlowers(t *testing.T) {
	cases := []struct {
		flowerbed []int
		n         int
		expected  bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{1, 0, 0, 0, 1, 0, 0}, 2, true},
		{[]int{0, 0, 0, 0}, 1, true},
	}

	for _, c := range cases {
		result := canPlaceFlowers(c.flowerbed, c.n)
		if result != c.expected {
			t.Fatalf("\n canPlaceFlowers: flowerbed is %v, n is %v, expected is %v, result is %v \n", c.flowerbed, c.n, c.expected, result)
		}
	}
}
