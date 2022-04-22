package _144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func preorderTraversal(root *TreeNode) []int {
// 	arr := []int{}

// 	var traverse func(node *TreeNode) []int
// 	traverse = func(node *TreeNode) []int {
// 		if node == nil {
// 			return nil
// 		}

// 		arr = append(arr, node.Val)

// 		traverse(node.Left)
// 		traverse(node.Right)

// 		return arr
// 	}
// 	traverse(root)
// 	return arr
// }

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	arr := []int{}
	arr = append(arr, root.Val)
	arr = append(arr, preorderTraversal(root.Left)...)
	arr = append(arr, preorderTraversal(root.Right)...)

	return arr
}
