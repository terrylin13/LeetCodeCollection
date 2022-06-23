package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	head := dummy
	preAdd := 0

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + preAdd
		sum, preAdd = sum%10, sum/10

		head.Next = &ListNode{Val: sum}
		head = head.Next
	}

	if preAdd > 0 {
		head.Next = &ListNode{Val: preAdd}
	}

	return dummy.Next
}
