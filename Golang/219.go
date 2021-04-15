package main

func containsNearbyDuplicate(nums []int, k int) bool {
	tmp := map[int]int{}
	for i, i2 := range nums {
		s1, ok := tmp[i2]
		if ok {
			if i-s1 <= k {
				return true
			} else {
				tmp[i2] = i
			}
		} else {
			tmp[i2] = i
		}
	}
	return false
}
