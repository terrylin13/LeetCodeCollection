package _15

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	length := len(nums)

	if length < 3 {
		return res
	}

	target := 0
	sort.Ints(nums)

	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		want := target - nums[i]
		lo := i + 1
		hi := length - 1

		//2sum
		for lo < hi {
			sum := nums[lo] + nums[hi]
			left := nums[lo]
			right := nums[hi]
			if sum == want {
				res = append(res, []int{nums[i], nums[lo], nums[hi]})
				for lo < hi && left == nums[lo] {
					lo++
				}
				for lo < hi && right == nums[hi] {
					hi--
				}
			} else if sum < want {
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
	return res
}
