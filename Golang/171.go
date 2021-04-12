package main

func titleToNumber(columnTitle string) int {
	result := 0
	s := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		result += ((int(columnTitle[i]) - 64) * s)
		s = s * 26
	}
	return result
}
