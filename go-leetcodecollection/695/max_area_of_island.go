package _695

import (
	"math"
)

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			temp := int(math.Max(float64(res), float64(dfs(grid, i, j))))
			if temp > res {
				res = temp
			}

		}
	}
	return res
}

func dfs(grid [][]int, i, j int) int {
	m, n := len(grid), len(grid[0])

	if i < 0 || j < 0 || i >= m || j >= n {
		// 超出索引边界
		return 0
	}

	if grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0

	return dfs(grid, i-1, j) + dfs(grid, i+1, j) +
		dfs(grid, i, j-1) + dfs(grid, i, j+1) + 1
}
