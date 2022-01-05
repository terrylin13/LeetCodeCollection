package _560

func subarraySum(nums []int, k int) int {
	res := 0
	length := len(nums)
	preNums := make([]int, length+1)

	for i := 0; i < length; i++ {
		preNums[i+1] = preNums[i] + nums[i]
	}

	for i := 1; i <= length; i++ {
		for j := 0; j < i; j++ {
			if preNums[i]-preNums[j] == k {
				res++
			}
		}
	}
	return res
}

func subarraySum2(nums []int, k int) int {
	res := 0
	mapping := make(map[int]int)
	mapping[0] = 1
	sum := 0
	for _, v := range nums {
		sum += v
		if mapping[sum-k] > 0 {
			res += mapping[sum-k]
		}
		mapping[sum]++
	}
	return res
}
