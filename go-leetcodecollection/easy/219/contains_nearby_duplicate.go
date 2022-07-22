package _219

func containsNearbyDuplicate(nums []int, k int) bool {
	pos := map[int]int{}
	for i, num := range nums {
		if p, ok := pos[num]; ok && i-p <= k {
			return true
		}
		pos[num] = i
	}
	return false
}
