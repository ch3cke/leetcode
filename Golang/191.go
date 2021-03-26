package main

func hammingWeight(num uint32) int {
	limit := uint32(2147483648)
	result := 0
	if num == 0 {
		return result
	}
	for tmp := uint32(1); ; tmp = tmp * 2 {
		if tmp&num != 0 {
			result += 1
		}
		if num < tmp {
			break
		}
		if tmp == limit {
			break
		}
	}
	return result
}
