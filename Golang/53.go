package main

func maxSubArray(nums []int) int {
	result := nums[0]
	if len(nums) == 1 {
		return result
	}
	i := 0
	for {
		if i == len(nums)-1 {
			break
		}
		if nums[i] > 0 {
			nums[i+1] += nums[i]
		}
		if nums[i+1] > result {
			result = nums[i+1]
		}
		i += 1
	}
	return result
}
