package leetcode1486

func xor(a int, b int) int {
	N := 31
	c := 0
	for i := 0; i < N; i++ {
		a1 := a & 0x1
		b1 := b & 0x1
		c <<= 1
		if a1 != b1 {
			c |= 0x1
		}
		a >>= 1
		b >>= 1
	}
	d := 0
	for i := 0; i < N; i++ {
		c1 := c & 0x01
		d <<= 1
		d |= c1
		c >>= 1
	}
	return d
}

func xorOperation(n int, start int) int {
	a := start
	b := start
	for i := 1; i < n; i++ {
		b = b + 2
		a = xor(a, b)
	}
	return a
}
