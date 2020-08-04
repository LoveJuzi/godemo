package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		i = f(nums, i)
	}

	rt := []int{}
	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			rt = append(rt, i+1)
		}
	}

	return rt
}

func f(nums []int, idx int) int {
	if nums[idx] == idx+1 {
		return idx + 1
	}

	if nums[nums[idx]-1] == nums[idx] {
		return idx + 1
	}

	nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]

	return idx
}
