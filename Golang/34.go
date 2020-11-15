package main

func searchRange(nums []int, target int) []int {
	i := 0
	j := len(nums) - 1
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	if len(nums) == 1 {
		if nums[0] == target {
			result[0] = 0
			result[1] = 0
			return result
		}
		return result
	}
	for {
		if i > j {
			break
		}
		if result[0] != -1 && result[1] != -1 {
			return result
		}
		if nums[i] < target {
			i = i + 1
		} else if nums[i] == target {
			result[0] = i
		}
		if nums[j] > target {
			j = j - 1
		} else if nums[j] == target {
			result[1] = j
		}
	}
	return result
}
