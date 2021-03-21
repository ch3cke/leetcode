package main

import "strconv"

func evalRPN(tokens []string) int {
	result := 0
	if len(tokens) == 0 {
		return result
	}
	nums := []int{}
	for _, i := range tokens {
		if i != "+" && i != "-" && i != "/" && i != "*" {
			tmp, _ := strconv.Atoi(i)
			nums = append(nums, tmp)
		} else {
			a := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			b := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			switch i {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, b-a)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, b/a)
			}
		}
	}
	return nums[0]
}
