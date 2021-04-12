package main

func twoSum(numbers []int, target int) []int {
	per, end := 0, len(numbers)-1
	for per < end {
		if numbers[per]+numbers[end] > target {
			end = end - 1
		} else if numbers[per]+numbers[end] < target {
			per = per + 1
		} else {
			break
		}
	}
	return []int{per + 1, end + 1}
}
