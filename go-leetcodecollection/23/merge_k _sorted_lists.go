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

	var head *ListNode
	var tail *ListNode
	var next *ListNode

	for mergeHeap.Len() > 0 {
		if tail != nil {
			next = heap.Pop(mergeHeap).(*ListNode)
			if next.Next != nil {
				heap.Push(mergeHeap, next.Next)
			}
			tail.Next = next
			tail = next
		} else {
			head = heap.Pop(mergeHeap).(*ListNode)
			tail = head
			if head.Next != nil {
				heap.Push(mergeHeap, head.Next)
			}
		}
	}

	return head
}
