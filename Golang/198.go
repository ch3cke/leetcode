package main

func rob(nums []int) int {
	a := nums[0]
	if len(nums) == 1 {
		return a
	}
	b := max_num(nums[0], nums[1])
	if len(nums) == 2 {
		return b
	}
	for i := 2; i < len(nums); i++ {
		a, b = b, max_num(a+nums[i], b)
	}
	return b
}

func max_num(a, b int) int {
	if a > b {
		return a
	}
	return b
}
