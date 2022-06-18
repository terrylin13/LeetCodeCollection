package _152

import "math"

func maxProduct(nums []int) int {
	ans, preMax, preMin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		a := preMax * nums[i]
		b := preMin * nums[i]
		preMax = max(nums[i], max(a, b))
		preMin = min(nums[i], min(a, b))
		ans = max(ans, preMax)
	}
	return ans
}

func maxProduct2(nums []int) int {
	preMax, preMin, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		preMax, preMin = max(preMax*nums[i], preMin*nums[i], nums[i]), min(preMax*nums[i], preMin*nums[i], nums[i])
		ans = max(ans, preMax)
	}
	return ans
}

func max(vals ...int) int {
	ans := math.MinInt
	for _, v := range vals {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func min(vals ...int) int {
	ans := math.MaxInt
	for _, v := range vals {
		if v < ans {
			ans = v
		}
	}
	return ans
}
