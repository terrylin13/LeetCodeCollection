package _63

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	lengthX, lengthY := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, lengthY)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	for x := 0; x < lengthX; x++ {
		for y := 0; y < lengthY; y++ {
			if obstacleGrid[x][y] == 1 {
				dp[y] = 0
				continue
			}
			if y-1 >= 0 && obstacleGrid[x][y-1] == 0 {
				dp[y] += dp[y-1]
			}
		}
	}
	return dp[len(dp)-1]
}
