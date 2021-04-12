package main

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp2, dp3, dp5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[dp2]*2, dp[dp3]*3, dp[dp5]*5
		dp[i] = min_num(min_num(x2, x3), x5)
		if dp[i] == x2 {
			dp2++
		}
		if dp[i] == x3 {
			dp3++
		}
		if dp[i] == x5 {
			dp5++
		}
	}
	return dp[n]
}

func min_num(a, b int) int {
	if a > b {
		return b
	}
	return a
}
