package _46

func permute(nums []int) [][]int {
	res := [][]int{}

	var backtrack func(nums []int, res *[][]int, first int)
	backtrack = func(nums []int, res *[][]int, first int) {
		n := len(nums)
		if first == n {
			cp := make([]int, n)
			copy(cp, nums)
			*res = append(*res, cp)
		}
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(nums, res, first+1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrack(nums, &res, 0)

	return res
}

func permute2(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, n := range nums {
			if visited[n] {
				continue
			}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}
	dfs([]int{})
	return res
}
