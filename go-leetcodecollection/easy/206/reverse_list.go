package _206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var rhead *ListNode
	for head != nil {
		temp := head
		head = head.Next
		temp.Next = rhead
		rhead = temp
	}
	return rhead
}
