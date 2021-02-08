package main

//func sortColors(nums []int)  {
//	sort.Ints(nums)
//}

func sortColors(nums []int) {
	num_0 := make([]int, 0)
	num_1 := make([]int, 0)
	num_2 := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			num_0 = append(num_0, nums[i])
		}
		if nums[i] == 1 {
			num_1 = append(num_1, nums[i])
		}
		if nums[i] == 2 {
			num_2 = append(num_2, nums[i])
		}
	}
	tmp := []int{}
	tmp = append(tmp, num_0...)
	tmp = append(tmp, num_1...)
	tmp = append(tmp, num_2...)
	for i := 0; i < len(nums); i++ {
		nums[i] = tmp[i]
	}
}
