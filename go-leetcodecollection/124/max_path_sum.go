package _124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		lMax := max(0, dfs(root.Left))
		rMax := max(0, dfs(root.Right))
		maxSum = max(lMax+rMax+root.Val, maxSum)
		return max(lMax, rMax) + root.Val
	}
	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
