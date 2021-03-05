package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	len1, len2, len3 := len(s1), len(s2), len(s3)
	if len1+len2 != len3 {
		return false
	}
	result := make([][]bool, len1+1)
	for i := 0; i < len1+1; i++ {
		result[i] = make([]bool, len2+1)
	}
	result[0][0] = true
	for i := 0; i < len1+1; i++ {
		for j := 0; j < len2+1; j++ {
			if i > 0 && s3[i+j-1] == s1[i-1] {
				result[i][j] = result[i][j] || result[i-1][j]
			}
			if j > 0 && s3[i+j-1] == s2[j-1] {
				result[i][j] = result[i][j] || result[i][j-1]
			}
		}
	}
	return result[len1][len2]
}
