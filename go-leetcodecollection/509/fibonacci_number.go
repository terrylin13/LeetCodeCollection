package _509

// 暴力递归
func fib1(n int) int {
	if n <= 1 {
		return n
	}
	return fib1(n-2) + fib1(n-1)
}

// 带备忘录的递归解法
// 时间复杂度是 O(n)
func fib2(n int) int {
	mem := make([]int, n+1)
	return helper(mem, n)
}

func helper(mem []int, n int) int {
	if n <= 1 {
		return n
	}

	if mem[n] != 0 {
		return mem[n]
	}

	mem[n] = helper(mem, n-2) + helper(mem, n-1)
	return mem[n]
}

// dp 数组的迭代解法
// 上一步「备忘录」的启发，我们可以把这个「备忘录」独立出来成为一张表，就叫做 DP table 吧，在这张表上完成「自底向上」的推算
// 「状态转移方程」？
// 其实就是为了听起来高端。把 f(n) 想做一个状态 n，这个状态 n 是由状态 n - 1 和状态 n - 2 相加转移而来，这就叫状态转移，仅此而已。
func fib3(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	// base case
	dp[0] = 0
	dp[1] = 1
	// 状态转移
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}

// 状态压缩
// 我们发现每次状态转移只需要 DP table 中的一部分，那么可以尝试用状态压缩来缩小 DP table 的大小，
// 只记录必要的数据，上述例子就相当于把DP table 的大小从 n 缩小到 2
func fib(n int) int {
	if n < 1 {
		return 0
	}

	if n == 2 || n == 1 {
		return 1
	}

	prev := 1
	curr := 1
	
	for i := 3; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}
