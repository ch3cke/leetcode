package main

import "strconv"

func summaryRanges(nums []int) []string {
	tmp_str := ""
	result := []string{}
	if len(nums) == 0 {
		return result
	}
	start := nums[0]
	if len(nums) == 1 {
		result = append(result, strconv.Itoa(start))
		return result
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			if nums[i-1] == start {
				tmp_str += strconv.Itoa(start)
				result = append(result, tmp_str)
				tmp_str = ""
				start = nums[i]
			} else {
				tmp_str += strconv.Itoa(start)
				tmp_str += "->"
				tmp_str += strconv.Itoa(nums[i-1])
				result = append(result, tmp_str)
				tmp_str = ""
				start = nums[i]
			}
		}
		if i == len(nums)-1 {
			if start != nums[i] {
				tmp_str += strconv.Itoa(start)
				tmp_str += "->"
				tmp_str += strconv.Itoa(nums[len(nums)-1])
				result = append(result, tmp_str)
			} else {
				tmp_str += strconv.Itoa(start)
				result = append(result, tmp_str)
			}
		}
	}
	return result
}
