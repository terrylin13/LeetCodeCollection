package sort

import "testing"

func TestQuickSort(t *testing.T) {
	nums := []int{4, 6, 6, 7, 9, 0, 1}
	quickSort(nums)
	t.Log(nums)
}
