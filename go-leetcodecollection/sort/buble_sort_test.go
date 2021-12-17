package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	nums := []int{4, 6, 6, 7, 9, 0, 1}
	bubbleSort(nums)
	t.Log(nums)
}
