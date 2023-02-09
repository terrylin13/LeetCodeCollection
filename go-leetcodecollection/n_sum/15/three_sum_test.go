package _15

import (
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	t.Log(threeSum(nums))
	sort.Ints(nums)
	t.Log(nSum(nums, 3, 0, 0))
}
