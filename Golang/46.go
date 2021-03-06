package main

import (
	"fmt"
	"reflect"
	"sort"
)

func permute(nums []int) [][]int {
	var results [][]int
	tmp := append([]int{}, nums...)
	results = append(results, nums)
	tmpresult := getNext(tmp)
	for {
		fmt.Println(tmpresult)
		if reflect.DeepEqual(tmpresult, results[0]) || len(tmpresult) == 0 {
			break
		} else {
			results = append(results, append([]int{}, tmpresult...))
			tmpresult = getNext(tmpresult)
		}
	}
	return results
}

func getNext(nums []int) []int {
	var result []int
	if len(nums) == 1 {
		return result
	}
	index := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			index = i - 1
			break
		}
	}
	if index == -1 {
		sort.Ints(nums)
		result = nums
		return result
	}
	for j := len(nums) - 1; j > index; j-- {
		if nums[j] > nums[index] {
			tmp := nums[j]
			nums[j] = nums[index]
			nums[index] = tmp
			sort.Ints(nums[index+1:])
			result = nums
			break
		}
	}
	return result
}
