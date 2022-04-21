package _144

import "testing"

func TestPreorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	t.Log(preorderTraversal(root))
}
