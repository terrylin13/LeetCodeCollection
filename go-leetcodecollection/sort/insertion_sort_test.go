package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	nums := []int{4, 6, 6, 7, 9, 0, 1}
	insertionSort(nums, len(nums))
	t.Log(nums)
}
