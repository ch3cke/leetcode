package main

import "strings"

func lengthOfLastWord(s string) int {
	result := 0
	s = strings.TrimSpace(s)
	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			result = 0
		} else {
			result++
		}
	}
	return result
}
