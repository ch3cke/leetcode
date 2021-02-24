package main

func numDecodings(s string) int {
	if s[0] == 48 {
		return 0
	}
	if len(s) < 3 {
		if len(s) == 1 {
			if s[0] == 48 {
				return 0
			}
		}
		if len(s) == 2 {
			if s[1] == 48 {
				if s[0] <= 50 {
					return 1
				} else {
					return 0
				}
			} else {
				if s[0] == 48 {
					return 0
				} else if s[0] == 49 {
					return 2
				} else if s[0] == 50 {
					if s[1] >= 49 && s[1] <= 54 {
						return 2
					} else {
						return 1
					}
				}
			}
		}
	}
	tmp := make([]int, len(s)+2)
	tmp[0] = 1
	tmp[1] = 1
	for i := 1; i < len(s); i++ {
		if s[i] == 48 {
			if s[i-1] <= 50 && s[i-1] > 48 {
				tmp[i] = tmp[i-1]
				tmp[i+1] = tmp[i-1]
			} else {
				return 0
			}
		} else {
			if s[i-1] == 49 {
				tmp[i+1] = tmp[i-1] + tmp[i]
			} else if s[i-1] == 50 {
				if s[i] <= '6' && s[i] >= '1' {
					tmp[i+1] = tmp[i-1] + tmp[i]
				} else {
					tmp[i+1] = tmp[i]
				}
			} else {
				tmp[i+1] = tmp[i]
			}
		}
	}
	return tmp[len(s)]
}
