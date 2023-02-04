package _268

func missingNumber(nums []int) int {
	ans := 0
	n := len(nums)
	ans ^= n
	for i, num := range nums {
		ans ^= i ^ num
	}
	return ans
}
