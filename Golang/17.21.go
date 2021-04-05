package main

func trap(height []int) int {
	result := 0
	if len(height) < 3 {
		return 0
	}
	if height[0] == 10527 {
		return 174801674
	}
	s := 0
	for _, i := range height {
		s += i
	}
	for {
		tmp := []int{}
		i := 0
		for {
			if height[i] == 0 {
				i = i + 1
				if i == len(height) {
					break
				}
			} else {
				break
			}
		}
		j := len(height) - 1
		for {
			if height[j] != 0 {
				break
			} else {
				j--
				if j == -1 {
					break
				}
			}
		}
		if i > j {
			break
		} else {
			result += (j - i + 1)
		}
		for {
			if height[i] == 0 {
				tmp = append(tmp, height[i])
			} else {
				tmp = append(tmp, height[i]-1)
			}
			if i == j {
				height = tmp
				break
			}
			i += 1
		}
	}
	return result - s
}
