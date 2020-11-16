package main

func searchInsert(nums []int, target int) int {
	head := 0
	end := len(nums) - 1
	if nums[head] > target {
		tmp := append([]int{}, target)
		nums = append(nums, tmp...)
		return 0
	}
	if nums[end] < target {
		nums = append(nums, target)
		return end + 1
	}
	for {
		mid := (end + head) / 2
		if (end - head) == 1 {
			if nums[head] == target {
				return head
			} else if nums[end] == target {
				return end
			} else if nums[end] < target {
				tmp := append([]int{}, nums[:end]...)
				tmp = append(tmp, target)
				nums = append(tmp, nums[end:]...)
				return end + 1
			} else if nums[head] > target {
				tmp := append([]int{}, nums[:head-1]...)
				tmp = append(tmp, target)
				nums = append(tmp, nums[head-1:]...)
				return head
			} else if nums[head] < target && nums[end] > end {
				tmp := append([]int{}, nums[:head]...)
				tmp = append(tmp, target)
				nums = append(tmp, nums[head:]...)
				return head + 1
			}
		}
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			head = mid
		} else if nums[mid] == target {
			return mid
		}
	}
}
