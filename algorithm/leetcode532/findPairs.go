package leetcode532

func findPairs(nums []int, k int) int {
	N := len(nums)
	if N <= 0 {
		return 0
	}

	rt := 0
	T := map[int]int{}

	if k == 0 {
		for i := 0; i < N; i++ {
			T[nums[i]]++
			if T[nums[i]] == 2 {
				rt++
			}
		}
		return rt
	} else if k < 0 {
		return 0
	}

	T[nums[0]] = 1
	t := 0
	for i := 1; i < N; i++ {
		if T[nums[i]] == 1 {
			continue
		}
		t = nums[i] - k
		if T[t] == 1 {
			rt++
		}
		t = nums[i] + k
		if T[t] == 1 {
			rt++
		}
		T[nums[i]] = 1
	}

	return rt
}
