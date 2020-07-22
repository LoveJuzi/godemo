package leetcode507

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	i := 2
	j := num

	s := 1
	for i < j {
		j = num / i
		if num%i == 0 {
			s += i
			s += j
		}
		i++
	}
	return s == num
}
