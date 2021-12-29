package _23

import "testing"

func TestMergeKLists(t *testing.T) {
	lists := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{Val: 2, Next: &ListNode{Val: 6}},
	}
	l := mergeKLists(lists)
	for l != nil {
		t.Log(l.Val)
		l = l.Next
	}

}
