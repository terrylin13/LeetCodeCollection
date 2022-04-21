package _543

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	res := diameterOfBinaryTree(root)
	if res != 3 {
		t.Errorf("diameterOfBinaryTree 结果不为3! 结果为:%d", res)
	} else {
		t.Log("diameterOfBinaryTree 测试成功!")
	}
}
