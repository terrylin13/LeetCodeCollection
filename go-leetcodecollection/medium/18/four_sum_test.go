package _18

import "testing"

func TestThreeSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	t.Log(fourSum(nums, target))
}
