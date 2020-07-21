package leetcode467

import (
	"testing"
)

func Test_findSubstringInWraproundString(t *testing.T) {
	cases := []struct {
		p        string
		expected int
	}{
		// {"a", 1},
		// {"cac", 2},
		{"zab", 6},
	}

	for _, c := range cases {
		result := findSubstringInWraproundString(c.p)
		if result != c.expected {
			t.Fatalf("\nfindSubstringInWraproundString: p is \"%s\", excepted is %d, result is %d\n", c.p, c.expected, result)
		}
	}
}
