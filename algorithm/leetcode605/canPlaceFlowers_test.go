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
	}

}
