package _322

import (
	"math"
)

func coinChange1(coins []int, amount int) int {
	return dp(coins, amount)
}

func dp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	res := math.MaxInt32

	for _, coin := range coins {
		sub := dp(coins, amount-coin)
		if sub == -1 {
			continue
		}

		if res > sub+1 {
			res = sub + 1
		}

	}

	if res == math.MaxInt32 {
		return -1
	} else {
		return res
	}
}

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coins[j]]+1)))
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
