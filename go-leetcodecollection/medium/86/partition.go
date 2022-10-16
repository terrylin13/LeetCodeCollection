package _86

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	cur1, cur2 := dummy1, dummy2
	for head != nil {
		if head.Val >= x {
			cur2.Next = head
			cur2 = cur2.Next
		} else {
			cur1.Next = head
			cur1 = cur1.Next
		}
		temp := head.Next
		head.Next = nil
		head = temp
	}
	cur1.Next = dummy2.Next
	return dummy1.Next
}
