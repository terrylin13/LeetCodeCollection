package sort

func quickSort(nums []int) {
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, left int, right int) {
	if right <= left {
		return
	}

	p := partition(nums, left, right) // 获取分区点
	sort(nums, left, p-1)
	sort(nums, p+1, right)

}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	for j := i; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	// 最后将 pivot 与 i 下标对应数据值互换
	// 这样一来，pivot 就位于当前数据序列中间，i 也就是 pivot 值对应的下标
	nums[i], nums[right] = pivot, nums[i]

	// 返回 i 作为 pivot 分区位置
	return i
}
