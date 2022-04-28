package _141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// with hash
func hasCycle2(head *ListNode) bool {
	dict := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := dict[head]; ok {
			return true
		}
		dict[head] = struct{}{}
		head = head.Next
	}
	return false
}
