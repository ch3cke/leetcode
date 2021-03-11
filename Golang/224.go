package main

func calculate(s string) int {
	result := 0
	ops := []int{1}
	sign := 1
	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i += 1
			continue
		} else if s[i] == '+' {
			sign = ops[len(ops)-1]
			i += 1
		} else if s[i] == '-' {
			sign = -ops[len(ops)-1]
			i += 1
		} else if s[i] == '(' {
			ops = append(ops, sign)
			i += 1
		} else if s[i] == ')' {
			ops = ops[:len(ops)-1]
			i += 1
		} else {
			tmp := 0
			for {
				if i >= len(s) {
					break
				}
				if s[i] > '9' {
					break
				}
				if s[i] < '0' {
					break
				}
				tmp = tmp*10 + int(s[i]-48)
				i += 1
			}
			result += sign * tmp
		}
	}
	return result
}
