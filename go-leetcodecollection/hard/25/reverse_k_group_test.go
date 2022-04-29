package _25

import "testing"

func TestReverse(t *testing.T) {
	tree := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}

	res := reverseKGroup(tree, 2)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}

}
