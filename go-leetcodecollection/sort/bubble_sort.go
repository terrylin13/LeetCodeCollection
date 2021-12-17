package sort

func bubbleSort(nums []int) {
	length := len(nums)

	if length < 1 {
		return
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
