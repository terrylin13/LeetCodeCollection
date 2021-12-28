package _46

func permute(nums []int) [][]int {
	res := [][]int{}
	backtrack(nums, &res, 0)
	return res
}

func backtrack(nums []int, res *[][]int, first int) {
	// All the integers are used up.
	n := len(nums)
	if first == n {
		cp := make([]int, n)
		copy(cp, nums)
		*res = append(*res, cp)
	}
	for i := first; i < n; i++ {
		// Place i-th integer first in the current permutation.
		nums[first], nums[i] = nums[i], nums[first]
		// Use next integers to complete the permutations.
		backtrack(nums, res, first+1)
		// Backtrack.
		nums[first], nums[i] = nums[i], nums[first]
	}
}
