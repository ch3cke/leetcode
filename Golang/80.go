package main

func removeDuplicates(nums []int) int {
	flag := 0
	i := 0
	for {
		if i >= len(nums)-1 {
			break
		}
		if nums[i] == nums[i+1] {
			flag++
			i = i + 1
		} else {
			flag = 0
			i = i + 1
		}
		if flag > 1 {
			nums = append(nums[:i], nums[i+1:]...)
			i = i - 1
		}
	}
	return len(nums)
}
