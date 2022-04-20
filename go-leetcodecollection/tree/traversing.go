package tree

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   string
}

// 前序遍历是指，对于树中的任意节点来说，先打印这个节点，然后再打印它的左子树，最后打印它的右子树。
func preTraversing(t *TreeNode) {
	if t == nil {
		return
	}
	fmt.Println(t.Val)
	preTraversing(t.Left)
	preTraversing(t.Right)
}

// 中序遍历是指，对于树中的任意节点来说，先打印它的左子树，然后再打印它本身，最后打印它的右子树。
func inTraversing(t *TreeNode) {
	if t == nil {
		return
	}
	inTraversing(t.Left)
	fmt.Println(t.Val)
	inTraversing(t.Right)
}

// 后序遍历是指，对于树中的任意节点来说，先打印它的左子树，然后再打印它的右子树，最后打印这个节点本身。
func postTraversing(t *TreeNode) {
	if t == nil {
		return
	}
	postTraversing(t.Left)
	postTraversing(t.Right)
	fmt.Println(t.Val)
}
