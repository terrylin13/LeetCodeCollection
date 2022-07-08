package binary_search

func BinarySearch(nums []int, k int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if k < nums[mid] {
			right = mid - 1
		} else if k > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
