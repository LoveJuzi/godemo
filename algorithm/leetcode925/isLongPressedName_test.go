package leetcode925

import (
	"testing"
)

func Test_isLongPressedName(t *testing.T) {
	cases := []struct {
		name     string
		typed    string
		expected bool
	}{
		{"alex", "aaleex", true},
		{"saeed", "ssaaedd", false},
		{"leelee", "lleeelee", true},
		{"laiden", "laiden", true},
		{"pyplrz", "ppyypllr", false},
	}

	for _, c := range cases {
		result := isLongPressedName(c.name, c.typed)
		if result != c.expected {
			t.Fatalf("\n isLongPressedName: name is %v, typed is %v, expected is %v, result is %v \n", c.name, c.typed, c.expected, result)
		}
	}
}
