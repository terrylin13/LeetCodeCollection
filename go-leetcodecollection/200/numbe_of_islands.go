package _200

func numIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		// 超出索引边界
		return
	}

	if grid[i][j] == '0' {
		return
	}

	// 将 (i, j) 变成海水
	grid[i][j] = '0'
	// 淹没上下左右的陆地
	dfs(grid, i-1, j) // 上
	dfs(grid, i+1, j) // 下
	dfs(grid, i, j-1) // 左
	dfs(grid, i, j+1) // 右

}
