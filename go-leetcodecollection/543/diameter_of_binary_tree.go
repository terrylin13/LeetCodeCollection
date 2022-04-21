package _543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		lHeight := dfs(root.Left)
		rHeight := dfs(root.Right)
		// maxDiameter即为该二叉树的最大直径
		maxDiameter = max(maxDiameter, lHeight+rHeight)
		// 该节点为根的二叉树的最大高度，也就是1+max(lh+rh)
		return 1 + max(lHeight, rHeight)
	}

	dfs(root)
	return maxDiameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
