package _188

import "math"

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	k = min(k, n/2)
	buy := make([][]int, n)
	sell := make([][]int, n)
	for i := range buy {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	buy[0][0] = -prices[0]

	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64 / 2
		sell[0][i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}
	return max(sell[n-1]...)
}

func maxProfit2(k int, prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}

	k = min(k, n/2)

	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0][1] = math.MinInt
		dp[i][0][0] = math.MinInt
	}
	for i := 0; i < n; i++ {
		for j := k; j >= 1; j-- {
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

func maxProfit3(k int, prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	k = min(k, n/2)
	//go不能定义变长数组，所以这里直接定义k的最大值100+1空间大小的数组，后面可以再优化
	var dp = make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
	}
	//初始化dp数组，第0天的k次卖出都为0，第0天的k次买入均为-prices[0]
	for i := 0; i <= k; i++ {
		dp[0][i][0] = 0
		dp[0][i][1] = -prices[0]
	}
	//还是动态规划思想，核心有两点，1是每个步骤的状态定义，2是状态间的转换方程，首先要理清这两点
	//复杂的是有了k次交易的限制，所以之前二维的dp不能再满足所有信息的记录，但再加一维即可记录所有信息。
	//每天的状态有两个，持有股票或不持有股票，这里，0：代表不持有股票，可以理解为卖出，1：持有股票，可以理解为买入
	//原先用dp[i][0]代表第i天不持有股票，dp[i][1]代表第i天持有股票，但这里需要考虑交易次数，所以，dp[i][ki][0]表示第i天第ki次交易不持有股票，dp[i][ki][1]表示第i天第ki次交易持有股票
	//dp[i][ki][0]的也是两种情况，1：第i-1天卖出，第i天不操作,这里需要考虑卖出次数ki，记上一次卖出时的ki为ki-1。2：第i-1天买入，然后第i天卖出，这里也要考虑卖出次数ki,但需要注意的是，买入时不用考虑ki，只有在卖出时才算一次完整的交易。
	//所以，dp[i][ki][0]=max(dp[i-1][ki-1][0],dp[i-1][ki][1]+prices[i]),但这里还有一种情况，这里计算出的dp[i][ki][0]并不总是历史最大值，所以需要与dp[i-1][ki][0]比较，将两者较大值重新赋值给dp[i-1][ki][0]即可
	//dp[i][ki][1]类似，也是两种情况，1:第i-1天买入，第i天不操作,买入不需要考虑交易此时ki，2：第i-1天卖出，然后第i天买入,卖出需要考虑交易次数ki，第i-1天卖出，交易次数应为ki-1。所以，dp[i][ki][1] = max(dp[i-1][ki][1],dp[i-1][ki-1][0]-prices[i])
	for i := 1; i < len(prices); i++ {
		for ki := 1; ki <= k; ki++ {
			dp[i][ki][0] = max(dp[i-1][ki][0], max(dp[i-1][ki-1][0], dp[i-1][ki][1]+prices[i]))
			dp[i][ki][1] = max(dp[i-1][ki][1], dp[i-1][ki-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][k][0]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
