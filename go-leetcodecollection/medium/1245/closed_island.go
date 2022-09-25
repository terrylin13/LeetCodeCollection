package _1245

func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var dfs func(grid [][]int, i, j int)
	dfs = func(grid [][]int, i, j int) {
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

	for i := 0; i < m; i++ {
		dfs(grid, i, 0)
		dfs(grid, i, n-1)
	}

	for j := 0; j < n; j++ {
		dfs(grid, 0, j)
		dfs(grid, m-1, j)
	}

	res := 0
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
