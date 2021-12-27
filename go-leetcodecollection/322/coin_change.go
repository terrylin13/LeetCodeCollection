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

// 剪枝
func coinChange3(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for idx := range memo {
		memo[idx] = math.MaxInt32
	}
	return dp2(memo, coins, amount)
}

func dp2(memo, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	if memo[amount] != math.MaxInt32 {
		return memo[amount]
	}

	res := math.MaxInt32

	for _, coin := range coins {
		sub := dp2(memo, coins, amount-coin)
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
