package _122

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return ans
}

/**
DP:
考虑到「不能同时参与多笔交易」，因此每天交易结束后只可能存在手里有一支股票或者没有股票的状态。
定义状态 dp[i][0] 表示第 i 天交易完后手里没有股票的最大利润
dp[i][1] 表示第 i 天交易完后手里持有一支
*/
func maxProfit2(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
