package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	for i := 1; i < len(nums)-1; i++ {
		s := nums[i] - nums[i-1]
		s2 := nums[i] - nums[i+1]
		if s > 0 && s2 > 0 {
			return i
		}
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}
	return -1
}
