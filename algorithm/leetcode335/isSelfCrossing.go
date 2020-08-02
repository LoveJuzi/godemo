package leetcode335

func isSelfCrossing(x []int) bool {
	for i := 3; i < len(x); i++ {
		a := x[i-3]
		b := x[i-2]
		c := x[i-1]
		d := x[i]

		if a >= c && b <= d {
			return true
		}
	}

	for i := 4; i < len(x); i++ {
		a := x[i-4]
		b := x[i-3]
		c := x[i-2]
		d := x[i-1]
		e := x[i]

		if b == d && c <= a+e {
			return true
		}
	}

	for i := 5; i < len(x); i++ {
		a := x[i-5]
		b := x[i-4]
		c := x[i-3]
		d := x[i-2]
		e := x[i-1]
		f := x[i]

		if b < d && c >= e && e >= c-a && a < c && b+f >= d {
			return true
		}
	}

	return false
}
