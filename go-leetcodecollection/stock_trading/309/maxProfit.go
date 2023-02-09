package _309

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		if i-2 == -1 {
			// base case 2
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			// i - 2 小于 0 时根据状态转移方程推出对应 base case
			dp[i][1] = max(dp[i-1][1], -prices[i])
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
