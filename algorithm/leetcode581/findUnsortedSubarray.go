package leetcode581

import "sort"

func findUnsortedSubarray(nums []int) int {
	N := len(nums)
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(nums)
	i := 0
	for i < N {
		if tmp[i] != nums[i] {
			break
		}
		i++
	}
	if i == N {
		return 0
	}
	j := N - 1
	for j >= 0 {
		if tmp[j] != nums[j] {
			break
		}
		j--
	}
	return j - i + 1
}
