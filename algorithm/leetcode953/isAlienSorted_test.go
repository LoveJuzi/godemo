package leetcode953

import "testing"

func Test_isAlienSorted(t *testing.T) {
	cases := []struct {
		words    []string
		order    string
		expected bool
	}{
		{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
		{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
		{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
	}

	for _, c := range cases {
		result := isAlienSorted(c.words, c.order)
		if result != c.expected {
			t.Fatalf("\n isAlienSorted: words is %v, order is %v, expected is %v, result is %v \n", c.words, c.order, c.expected, result)
		}
	}
}
