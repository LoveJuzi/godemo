package leetcode55

func canJump(nums []int) bool {
	return f(nums, len(nums)-2, len(nums)-1)
}

func f(nums []int, i, j int) bool {
	if j <= 0 {
		return true
	}
	for ; i >= 0; i-- {
		d := j - i
		if nums[i] >= d {
			return f(nums, i-1, i)
		}
	}
	return false
}
