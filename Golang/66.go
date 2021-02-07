package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1
		if digits[i] == 10 {
			if i == 0 {
				digits[i] = 0
				s := []int{1}
				s = append(s, digits...)
				digits = s
			} else {
				digits[i] = 0
			}
		} else {
			break
		}
	}
	return digits
}
