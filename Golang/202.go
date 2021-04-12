package main

func isHappy(n int) bool {
	cord := map[int]int{}
	cord[n] = 0
	s := to_list(n)
	for {
		if s == 1 {
			return true
		}
		_, ok := cord[s]
		if ok {
			return false
		} else {
			cord[s] = 0
		}
		s = to_list(s)
	}
}

func to_list(n int) int {
	result := 0
	for n > 0 {
		s := n % 10
		result += (s * s)
		n = n / 10
	}
	return result
}
