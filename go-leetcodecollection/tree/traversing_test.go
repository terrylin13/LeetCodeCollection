package tree

import "testing"

var root = &TreeNode{
	Val: "A",
	Left: &TreeNode{
		Val: "B",
		Left: &TreeNode{
			Val: "D",
		},
		Right: &TreeNode{
			Val: "E",
		},
	},
	Right: &TreeNode{
		Val: "C",
		Left: &TreeNode{
			Val: "F",
		},
		Right: &TreeNode{
			Val: "G",
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
