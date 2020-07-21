package leetcode1518

import "testing"

func Test_numWaterBottles(t *testing.T) {
	cases := []struct {
		numBottles  int
		numExchange int
		expected    int
	}{
		{5, 5, 6},
		{2, 3, 2},
		{9, 3, 13},
		{15, 4, 19},
		{15, 8, 17},
	}

	for _, c := range cases {
		result := numWaterBottles(c.numBottles, c.numExchange)
		if result != c.expected {
			t.Fatalf("\n numWaterBottles: numBottles is %v, numExchange is %v, expected is %v, result is %v", c.numBottles, c.numExchange, c.expected, result)
		}
	}
}
