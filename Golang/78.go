package main

func subsets(nums []int) [][]int {
	result := [][]int{}
	length := 1 << len(nums)
	for m := 0; m < length; m++ {
		tmp := []int{}
		for i := 0; i < len(nums); i++ {
			if (m & (1 << i)) != 0 {
				tmp = append(tmp, nums[i])
			}
		}
		result = append(result, tmp)
	}
	return result
}
