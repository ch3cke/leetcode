package main

func singleNumber(nums []int) int {
	tmp := 0
	for i := 0; i < len(nums); i++ {
		tmp = tmp ^ nums[i]
	}
	return tmp
}
