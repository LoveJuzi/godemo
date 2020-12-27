package leetcode368

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	T := map[int][]int{} // use a table to optimize algorithm
	rt := []int{}
	for i := 0; i < len(nums); i++ { // find all child problem
		tmp := f(T, nums, i)
		if len(rt) < len(tmp) {
			rt = tmp
		}
	}
	return rt
}

func f(T map[int][]int, nums []int, idx int) []int {
	if v, ok := T[nums[idx]]; ok { // direct retrun if child problem has been resolved
		return v
	}

	rt := []int{}

	for j := idx + 1; j < len(nums); j++ { // find all child problem
		if nums[j]%nums[idx] != 0 { // empty child problem
			continue
		}
		// resolve child problem
		tmp := f(T, nums, j)
		if len(rt) < len(tmp) {
			rt = []int{}
			rt = append(rt, tmp...)
		}
	}

	rt = append(rt, nums[idx])

	T[nums[idx]] = rt
	return T[nums[idx]]
}
