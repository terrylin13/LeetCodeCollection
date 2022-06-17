package _70

// 暴力递归
func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 带备忘录的递归解法
func climbStairs2(n int) int {
	mem := make([]int, n+1)
	var recursion func(n int) int
	recursion = func(n int) int {
		if n == 0 || n == 1 || n == 2 {
			return n
		}
		if mem[n] != 0 {
			return mem[n]
		}
		mem[n] = recursion(n-1) + recursion(n-2)
		return mem[n]
	}

	return recursion(n)
}

func climbStairs3(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	mem := make([]int, n+1)
	mem[0] = 1
	mem[1] = 2
	for i := 2; i < n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}
	return mem[n-1]
}
