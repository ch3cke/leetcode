package main

import "fmt"

//func rotate_4(nums []int, k int)  {
//	k  = k % len(nums)
//	length := len(nums)-1
//	if k==0{
//		return
//	}else {
//		tmp := []int{}
//		for i := 0; i < k; i++ {
//			s := []int{nums[length-i]}
//			tmp = append(s, tmp...)
//		}
//		tmp = append(tmp, nums[:length-k+1]...)
//		for i := 0; i <= length; i++ {
//			nums[i]=tmp[i]
//		}
//	}
//}

func rotate_4(nums []int, k int) {
	k = k % len(nums)
	length := len(nums) - 1
	sd := make([]int, len(nums)-k)
	if k == 0 {
		return
	} else {
		for i := 0; i < k; i++ {
			sd[i] = nums[i]
			nums[i] = nums[i+(length-k)+1]
		}
		for i := k; i <= length-k; i++ {
			sd[i] = nums[i]
		}
		fmt.Println(nums)
		for i := k; i <= length; i++ {
			nums[i] = sd[i-k]
		}
	}
}
