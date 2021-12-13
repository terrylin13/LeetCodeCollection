package _141

import "testing"

func TestHasCycle(t *testing.T) {
	tail := &ListNode{Val: -1}

	head := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  2,
			Next: tail,
		},
	}

	loopNode := &ListNode{
		Val:  0,
		Next: head,
	}

	tail.Next = loopNode

	t.Log(hasCycle(head))
	t.Log(hasCycle2(head))
}
