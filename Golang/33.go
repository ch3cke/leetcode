package main

func search(nums []int, target int) int {
	length := len(nums)
	if length == 1 {
		if nums[0] != target {
			return -1
		} else {
			return 0
		}
	}
	k := length
	for {
		j := (k + 0) % length
		if nums[j] > target {
			if nums[(k-1)%length] < target {
				return -1
			} else {
				k = k - 1
			}
		}
		if nums[j] < target {
			if nums[(k+1)%length] > target {
				return -1
			} else {
				k = k + 1
			}
		}
		if nums[j] == target {
			return j
		}
		if k <= 0 || k > length*2 {
			return -1
		}
	}

}
