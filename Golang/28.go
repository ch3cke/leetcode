package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == len(needle) {
		if haystack == needle {
			return 0
		} else {
			return -1
		}
	}
	for i := 0; i < (len(haystack)-len(needle))+1; i++ {
		if haystack[i] == needle[0] {
			flag := i
			j := 0
			for {
				if haystack[i] == needle[j] {
					if j == len(needle)-1 {
						return flag
					}
					i++
					j++
				} else {
					i = flag
					break
				}
			}
		}
	}
	return -1
}
