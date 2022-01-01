package _1905

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				dfs(grid2, i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res++
				dfs(grid2, i, j)
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

	if grid[i][j] == 0 {
		return
	}

	grid[i][j] = 0

	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
