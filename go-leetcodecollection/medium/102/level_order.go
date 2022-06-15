package _102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	q := []*TreeNode{root}

	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return res
}

//dfs
func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) < depth+1 {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], root.Val)
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return res
}
