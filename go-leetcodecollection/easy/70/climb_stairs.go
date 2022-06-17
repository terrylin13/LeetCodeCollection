package _70

// 暴力递归
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 带备忘录的递归解法
func climbStairs2(n int) int {
	mem := make([]int, n+1)
	var recursion func(n int) int
	recursion = func(n int) int {
		if n == 1 {
			return 1
		}
		if n == 2 {
			return 2
		}
		if mem[n] != 0 {
			return mem[n]
		}
		mem[n] = recursion(n-1) + recursion(n-2)
		return mem[n]
	}

	return recursion(n)
}
