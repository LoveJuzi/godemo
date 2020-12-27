package leetcode1360

import "testing"

//The given dates are valid dates between the years 1971 and 2100.

func Test_daysBetweenDates(t *testing.T) {
	cases := []struct {
		date1    string
		date2    string
		expected int
	}{
		{"2019-06-29", "2019-06-30", 1},
		{"2020-01-15", "2019-12-31", 15},
	}

	for _, c := range cases {
		result := daysBetweenDates(c.date1, c.date2)
		if result != c.expected {
			t.Fatalf("\n daysBetweenDates: date1 is %v, date2 is %v, expected is %v, result is %v \n", c.date1, c.date2, c.expected, result)
		}
	}
}
