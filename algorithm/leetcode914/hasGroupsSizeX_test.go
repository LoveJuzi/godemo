package leetcode914

import "testing"

func Test_hasGroupsSizeX(t *testing.T) {
	cases := []struct {
		deck     []int
		expected bool
	} {
		{[]int{1,2,3,4,4,3,2,1}, true},
		{[]int{1,1,1,2,2,2,3,3}, false},
		{[]int{1}, false},
		{[]int{1,1}, true},
		{[]int{1,1,2,2,2,2}, true},
	}

	for _, c := range cases {
		result := hasGroupsSizeX(c.deck)
		if result != c.expected {
			t.Fatalf("\n hasGroupsSizeX: deck is %v, expected is %v, result is %v \n", c.deck, c.expected, result)
		}
	}
}

func Test_gcd(t *testing.T) {
	cases := []struct{
		a int
		b int
		e int
	} {
		{2, 3, 1},
		{4, 10, 2},
		{10, 4, 2},
	}

	for _, c := range cases {
		r := gcd(c.a, c.b)
		if r != c.e {
			t.Fatalf("\n gcd: a is %v, b is %v , e is %v, r is %v \n", c.a, c.b, c.e, r)
		}
	}
}