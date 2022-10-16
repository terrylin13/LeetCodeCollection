package _23

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val <= h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	mergeHeap := &NodeHeap{}
	heap.Init(mergeHeap)

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(mergeHeap, lists[i])
		}
	}

	// var head *ListNode
	// var tail *ListNode
	// for mergeHeap.Len() > 0 {
	// 	if tail != nil {
	// 		next := heap.Pop(mergeHeap).(*ListNode)
	// 		if next.Next != nil {
	// 			heap.Push(mergeHeap, next.Next)
	// 		}
	// 		tail.Next = next
	// 		tail = next
	// 	} else {
	// 		head = heap.Pop(mergeHeap).(*ListNode)
	// 		tail = head
	// 		if head.Next != nil {
	// 			heap.Push(mergeHeap, head.Next)
	// 		}
	// 	}
	// }

	dummy := new(ListNode)
	cur := dummy
	for mergeHeap.Len() > 0 {
		next := heap.Pop(mergeHeap).(*ListNode)
		if next.Next != nil {
			heap.Push(mergeHeap, next.Next)
		}
		cur.Next = next
		cur = next
	}
	return dummy.Next
}

// dp depend on merge2List func
func mergeKLists2(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	merge2Lists := func(a, b *ListNode) *ListNode {
		if a == nil && b == nil {
			return nil
		} else if a == nil && b != nil {
			return b
		} else if b == nil && a != nil {
			return a
		}

		dummy := &ListNode{}
		cur := dummy
		for a != nil && b != nil {
			if a.Val > b.Val {
				cur.Next = b
				b = b.Next
			} else {
				cur.Next = a
				a = a.Next
			}
			cur = cur.Next
		}

		if a != nil {
			cur.Next = a
		} else if b != nil {
			cur.Next = b
		}
		return dummy.Next
	}

	for i := 0; i < len(lists); i++ {
		cur.Next = merge2Lists(cur.Next, lists[i])
	}

	return dummy.Next
}
