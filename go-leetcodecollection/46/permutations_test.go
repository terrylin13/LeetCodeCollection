package _46

import "testing"

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	// res := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	t.Log(permute(nums))
	t.Log(permute2(nums))
}
