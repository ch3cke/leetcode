package main

func compareVersion(version1 string, version2 string) int {
	v1, v2 := version1, version2
	s1, s2 := -1, -1
	for {
		v1, s1 = splitInt(v1)
		v2, s2 = splitInt(v2)
		if s1 == s2 {
			if len(v1) == 0 && len(v2) != 0 {
				for len(v2) > 0 {
					v2, s2 = splitInt(v2)
					if s2 != 0 {
						return -1
					}
				}
				return 0
			} else if len(v2) == 0 && len(v1) != 0 {
				for len(v1) > 0 {
					v1, s1 = splitInt(v1)
					if s1 != 0 {
						return 1
					}
				}
				return 0
			} else if len(v2) == 0 && len(v1) == 0 {
				return 0
			}
		} else if s1 > s2 {
			return 1
		} else if s1 < s2 {
			return -1
		}
	}
}

func splitInt(version string) (string, int) {
	if len(version) == 0 {
		return version, 0
	} else {
		result := 0
		result_str := ""
		for j := 0; j < len(version); j++ {
			if version[j] == '.' {
				return version[j+1:], result
				break
			} else {
				result = result*10 + int(version[j]) - 48
			}
		}
		return result_str, result
	}
}
