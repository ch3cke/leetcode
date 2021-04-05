package main

func numRabbits(answers []int) int {
	result := len(answers)
	if result == 0 {
		return result
	}
	tmp_map := map[int]int{}
	for _, i2 := range answers {
		if i2 == 0 {
			continue
		} else {
			_, ok := tmp_map[i2]
			if ok {
				if tmp_map[i2] > 0 {
					tmp_map[i2] = tmp_map[i2] - 1
				} else {
					tmp_map[i2] = i2
				}
			} else {
				tmp_map[i2] = i2
			}
		}
	}
	for _, value := range tmp_map {
		result += value
	}
	return result
}
