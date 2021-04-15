package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt64
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= target {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start += 1
		}
		end += 1
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}
