package _703

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	sort.IntSlice
	K int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{K: k}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (this *KthLargest) Push(v interface{}) {
	this.IntSlice = append(this.IntSlice, v.(int))
}

func (this *KthLargest) Pop() interface{} {
	num := this.IntSlice[len(this.IntSlice)-1]
	this.IntSlice = this.IntSlice[:len(this.IntSlice)-1]
	return num
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.K {
		heap.Pop(this)
	}
	return this.IntSlice[0]

}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
