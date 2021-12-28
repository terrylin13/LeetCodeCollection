package _1

import "sort"

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for idx, num := range nums {
		if val, ok := mapping[num]; ok {
			return []int{val, idx}
		}
		mapping[target-num] = idx
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	sort.Ints(nums)
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum == target {
			return []int{lo, hi}
		} else if sum < target {
			lo++
		} else if sum > target {
			hi--
		}
	}
	return []int{}
}
