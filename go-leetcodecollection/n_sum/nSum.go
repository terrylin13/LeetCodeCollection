package n_sum

// before call this func, u need to motherfucker sort the nums!!!
func nSum(nums []int, n, start, target int) [][]int {
	res := [][]int{}
	length := len(nums)
	if length < 2 {
		return res
	}
	//base 2Sum
	if n == 2 {
		lo := start
		hi := length - 1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			left := nums[lo]
			right := nums[hi]
			if sum == target {
				res = append(res, []int{left, right})
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
	} else {
		// when n > 2, recursion (n-1)Sum
		for i := start; i < length; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			// var sub [][]int
			sub := nSum(nums, n-1, i+1, target-nums[i])
			for _, item := range sub {
				// (n-1)Sum + nums[i] = nSum
				item = append([]int{nums[i]}, item...)
				res = append(res, item)
			}
		}
	}
	return res
}
