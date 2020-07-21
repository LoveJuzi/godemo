package leetcode414

import "math"

func thirdMax(nums []int) int {
	firstM := math.MinInt64
	secondM := math.MinInt64
	thirdM := math.MinInt64
	for _, n := range nums {
		if n == firstM || n == secondM {
			continue
		}
		switch {
		case firstM < n:
			thirdM, secondM, firstM = secondM, firstM, n
		case secondM < n:
			thirdM, secondM = secondM, n
		case thirdM < n:
			thirdM = n
		}
	}
	if thirdM == math.MinInt64 {
		return firstM
	}
	return thirdM
}
