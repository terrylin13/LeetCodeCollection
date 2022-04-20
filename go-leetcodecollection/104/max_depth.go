package _104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return traverse(root, 0)
}

func traverse(root *TreeNode, deep int) int {
	if root == nil {
		return deep
	}

	lDeep := traverse(root.Left, deep+1)
	rDeep := traverse(root.Right, deep+1)

	if lDeep > rDeep {
		return lDeep
	}
	return rDeep
}
