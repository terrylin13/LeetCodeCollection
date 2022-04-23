package _515

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelMax := math.MinInt
		qLen := len(queue)

		for i := 0; i < qLen; i++ {
			if levelMax < queue[i].Val {
				levelMax = queue[i].Val
			}

			// 在遍历本层的基础上，同时把下一层的节点记录到queue中
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		// 把每一层的最大值，添加到result
		res = append(res, levelMax)
		// 更新queue，保证数据为下一层所需的节点数据
		queue = queue[qLen:]
	}

	return res
}
