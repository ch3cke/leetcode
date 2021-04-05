package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	result := ""
	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
		} else {
			j := i
			for {
				if j >= len(s) || s[j] == ' ' {
					break
				} else {
					j++
				}
			}
			fmt.Println(s[i:])
			if j == len(s) {
				result = s[i:] + result
			} else {
				result = s[i:j] + result
			}
			result = " " + result
			i = j
		}
	}
	return result[1:]
}
