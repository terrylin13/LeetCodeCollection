package _144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	return traverse(root)
}

func traverse(node *TreeNode) []int {
	if node == nil {
		return nil
	}

	arr := []int{node.Val}

	arr = append(arr, traverse(node.Left)...)

	arr = append(arr, traverse(node.Right)...)

	return arr
}
