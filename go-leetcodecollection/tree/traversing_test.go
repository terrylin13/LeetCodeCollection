package tree

import "testing"

var root = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val: 9,
	},
	Right: &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val: 15,
		},
		Right: &TreeNode{
			Val: 7,
		},
	},
}

func TestPreTraversing(t *testing.T) {
	preTraversing(root)
	t.Log("done\n")
	inTraversing(root)
	t.Log("done\n")
	postTraversing(root)

}
