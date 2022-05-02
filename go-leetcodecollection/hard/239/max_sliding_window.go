package _239

import (
	"container/heap"
	"sort"
)

// Monotonic Queue
func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

//Priority Queue
type hp struct {
	sort.IntSlice
	a []int
}

func (this *hp) Less(i, j int) bool {
	return this.a[this.IntSlice[i]] > this.a[this.IntSlice[j]]
}

func (this *hp) Pop() interface{} {
	num := this.IntSlice[len(this.IntSlice)-1]
	this.IntSlice = this.IntSlice[:len(this.IntSlice)-1]
	return num
}

func (this *hp) Push(v interface{}) {
	this.IntSlice = append(this.IntSlice, v.(int))
}

func maxSlidingWindow2(nums []int, k int) []int {

	q := &hp{make([]int, k), nums}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}
