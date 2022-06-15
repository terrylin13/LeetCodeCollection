package _104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func maxDepth(root *TreeNode) int {
// 	return traverse(root, 0)
// }

// func traverse(root *TreeNode, deep int) int {
// 	if root == nil {
// 		return deep
// 	}

// 	lDeep := traverse(root.Left, deep+1)
// 	rDeep := traverse(root.Right, deep+1)

// 	if lDeep > rDeep {
// 		return lDeep
// 	}
// 	return rDeep
// }

func maxDepth(root *TreeNode) int {
	maxDepth := 0

	var traverse func(root *TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lDepth := traverse(root.Left)
		rDepth := traverse(root.Right)
		pDepth := max(lDepth, rDepth) + 1
		maxDepth = max(maxDepth, pDepth)
		return pDepth
	}

	traverse(root)
	return maxDepth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
