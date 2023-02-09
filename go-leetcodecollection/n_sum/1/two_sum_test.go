package _1

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)
	t.Log(res)
}
