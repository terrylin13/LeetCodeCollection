package _19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	node := findNthFromEnd(dummy, n+1)
	node.Next = node.Next.Next
	return dummy.Next
}

func findNthFromEnd(head *ListNode, k int) *ListNode {
	p1 := head
	for i := 0; i < k; i++ {
		if p1 == nil {
			return nil
		}
		p1 = p1.Next
	}

	p2 := head

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}
