package _120

import "math"

func minimumTotal(triangle [][]int) int {
	h := len(triangle)
	dp := make([][]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}

	for i := h - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == h-1 {
				dp[i][j] = triangle[i][j]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}
	return dp[0][0]
}

func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	f := make([][]int, n)

	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}

	f[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		f[i][0] = f[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			// core
			f[i][j] = min(f[i-1][j-1], f[i-1][j]) + triangle[i][j]
		}
		f[i][i] = f[i-1][i-1] + triangle[i][i]
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, f[n-1][i])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
