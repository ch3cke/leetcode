package main

func findMin(nums []int) int {
	m, n := 0, len(nums)-1
	s := (m + n) / 2
	if nums[m] < nums[n] {
		return nums[m]
	}
	for {
		if n-m == 1 || n == m {
			break
		}
		if nums[m] > nums[s] {
			n = s
		} else {
			m = s
		}
		s = (m + n) / 2
	}
	return nums[n]
}
