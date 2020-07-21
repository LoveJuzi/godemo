package utils

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		x        int
		y        int
		excepted int
	}{
		{1, 2, 3},
		{2, 2, 4},
	}

	for _, c := range cases {
		result := Sum(c.x, c.y)
		if result != c.excepted {
			t.Fatalf("Sum function failed, x: %d, y: %d, expected: %d, result: %d", c.x, c.y, c.excepted, result)
		}
	}
}
