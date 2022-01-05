package _303

type NumArray2 struct {
	preSum []int
}

func Constructor2(nums []int) NumArray2 {
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArray2{
		preSum: preSum,
	}
}

func (this *NumArray2) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}
