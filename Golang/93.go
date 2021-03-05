package main

import (
	"strconv"
	"strings"
)

var result []string

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	result = []string{}
	track := make([]string, 0)
	backtrack(s, track, 1)
	return result
}

func backtrack(s string, track []string, key int) {
	if key == 5 {
		if len(s) == 0 {
			str := strings.Join(track, ".")
			result = append(result, str)
		}
		return
	}

	for j := 1; j <= 3; j++ {
		if j <= len(s) {
			v, err := strconv.Atoi(s[:j])
			if err == nil && v <= 255 {
				track = append(track, s[:j])
				str := s[j:]

				backtrack(str, track, key+1)

				track = track[:len(track)-1]
			}
			if v == 0 {
				break
			}
		}
	}
}
