package _1254

func closedIsland(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		dfs(grid, i, 0)   // 左边
		dfs(grid, i, n-1) // 右边
	}

	for j := 0; j < n; j++ {
		dfs(grid, 0, j)   // 上边
		dfs(grid, m-1, j) // 下边
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}

	if grid[i][j] == 1 {
		return
	}

	grid[i][j] = 1
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
