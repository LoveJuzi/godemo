package leetcode485

func findMaxConsecutiveOnes(nums []int) int {
	N := len(nums)

	res := 0
	c := 0
	for i := 0; i < N; i++ {
		if nums[i] == 0 {
			c = 0
			continue
		} else {
			c++
			if res < c {
				res = c
			}
		}
	}
	return res
}
