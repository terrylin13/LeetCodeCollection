package _111

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minDepth := math.MaxInt
	var dfs func(root *TreeNode, depth int)

	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		pDepth := depth + 1
		dfs(root.Left, pDepth)
		dfs(root.Right, pDepth)
		if root.Left == nil && root.Right == nil {
			minDepth = min(minDepth, pDepth)
		}
	}
	dfs(root, 0)
	return minDepth
}

// bfs
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1

	q := []*TreeNode{root}

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]

			if cur.Left == nil && cur.Right == nil {
				return depth
			}

			if cur.Left != nil {
				q = append(q, cur.Left)
			}

			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}

	return depth
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
