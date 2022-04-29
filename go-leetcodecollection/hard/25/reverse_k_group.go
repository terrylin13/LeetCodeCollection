package _25

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	// 区间 [start, end) 包含 k 个待反转元素
	start := head
	end := head
	for i := 0; i < k; i++ {

		if end == nil {
			return head
		}
		end = end.Next
	}

	newHead := reverse(start, end)
	start.Next = reverseKGroup(end, k)
	return newHead
}

/* 反转区间 [start, end) 的元素，注意是左闭右开 */
func reverse(startHead, end *ListNode) *ListNode {
	var cur, pre *ListNode
	cur = startHead
	for cur != end {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
