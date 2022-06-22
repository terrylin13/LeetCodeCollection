package _121

import "math"

func maxProfit(prices []int) int {
	min, ans := math.MaxInt, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > ans {
			ans = prices[i] - min
		}
	}
	return ans
}
