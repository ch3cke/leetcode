package main

type NumArray struct {
	nums []int
}

//func Constructor(nums []int) NumArray {
//	tmp := NumArray{
//		nums: nums,
//	}
//	return tmp
//}

func (this *NumArray) SumRange(i int, j int) int {
	tmp := 0
	for ; i <= j; i++ {
		tmp += this.nums[i]
	}
	return tmp
}
