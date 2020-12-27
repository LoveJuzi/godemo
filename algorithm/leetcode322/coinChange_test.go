package leetcode322

import "testing"

func Test_coinChange(t *testing.T) {
	cases := []struct {
		coins    []int
		amount   int
		expected int
	}{
		// {[]int{1, 2, 5}, 11, 3},
		// {[]int{186, 419, 83, 408}, 6249, 20},
		{[]int{1, 5, 7}, 25, 5},
	}

	var result int
	for _, c := range cases {
		result = coinChange(c.coins, c.amount)
		if result != c.expected {
			t.Fatalf("\n coinChange: coins is %v, amount is %v, expected is %v, result is %v \n", c.coins, c.amount, c.expected, result)
		}
	}
}
