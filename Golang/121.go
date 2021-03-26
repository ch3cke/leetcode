package main

func maxProfit(prices []int) int {
	Min := 100000
	result := 0
	for i := 0; i < len(prices); i++ {
		if Min > prices[i] {
			Min = prices[i]
		} else if result < (prices[i] - Min) {
			result = (prices[i] - Min)
		}
	}
	return result
}
