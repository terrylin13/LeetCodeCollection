package _237

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	value := node.Next.Val
	node.Next = node.Next.Next
	node.Val = value
}
