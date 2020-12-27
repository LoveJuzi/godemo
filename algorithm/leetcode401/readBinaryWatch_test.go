package leetcode401

import "testing"

func isEqual(result []string, expected []string) bool {
	N := len(result)
	M := len(expected)

	if N != M {
		return false
	}

	T := map[string]int{}
	for i := 0; i < M; i++ {
		T[expected[i]] = 1
	}

	for i := 0; i < N; i++ {
		if v, ok := T[result[i]]; !ok || v == 0 {
			return false
		}
		T[result[i]] = 0
	}
	return true
}

func Test_readBinaryWatch(t *testing.T) {
	cases := []struct {
		n        int
		expected []string
	}{
		{1, []string{"1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"}},
	}

	for _, c := range cases {
		result := readBinaryWatch(c.n)
		if !isEqual(result, c.expected) {
			t.Fatalf("\n readBinaryWatch: n is %v, expected is %v, result is %v \n", c.n, c.expected, result)
		}
	}
}
