package main

func removeDuplicates_1(S string) string {
	result := ""
	i := 0
	for {
		if i == len(S) {
			break
		}
		if len(result) == 0 {
			result += string(S[i])
		} else {
			if S[i] == result[len(result)-1] {
				result = result[:len(result)-1]
			} else {
				result += string(S[i])
			}
		}
		i = i + 1
	}
	return result
}
