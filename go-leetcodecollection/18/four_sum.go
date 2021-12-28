package _18

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	length := len(nums)
	if length < 4 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 3Sum
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			lo := j + 1
			hi := length - 1
			// 2Sum
			for lo < hi {
				sum := nums[i] + nums[j] + nums[lo] + nums[hi]
				left := nums[lo]
				right := nums[hi]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[lo], nums[hi]})
					for lo < hi && nums[lo] == left {
						lo++
					}
					for lo < hi && nums[hi] == right {
						hi--
					}
				} else if sum < target {
					for lo < hi && nums[lo] == left {
						lo++
					}
				} else {
					for lo < hi && nums[hi] == right {
						hi--
					}
				}
			}
		}
	}
	return res
}
