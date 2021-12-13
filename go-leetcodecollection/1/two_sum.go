package _1

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for idx, num := range nums {
		if val, ok := mapping[num]; ok {
			return []int{val, idx}
		}
		mapping[target-num] = idx
	}
	return []int{}
}
