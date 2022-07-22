package _5

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var max string

	var maxPalindrome func(s string, left, right int, max string) string

	maxPalindrome = func(s string, left, right int, max string) string {
		var sub string
		for left >= 0 && len(s) > right && s[left] == s[right] {
			sub = s[left : right+1]
			left--
			right++
		}

		if len(max) < len(sub) {
			return sub
		}
		return max
	}

	for i := 0; i < len(s); i++ {
		// 找到以 s[i] 为中心的回文串
		max = maxPalindrome(s, i, i, max)
		// 找到以 s[i] 和 s[i+1] 为中心的回文串
		max = maxPalindrome(s, i, i+1, max)
	}
	return max
}

//dp
func LongestPalindrome2(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	begin := 0
	maxNum := 1
	table := make([][]bool, length)
	for i := 0; i < length; i++ {
		table[i] = make([]bool, length)
	}
	for i := 0; i < length; i++ {
		table[i][i] = true
	}

	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] != s[j] {
				table[i][j] = false
			} else {
				if i-j >= 2 {
					table[j][i] = table[j+1][i-1]
				} else {
					table[j][i] = true
				}
			}
			if table[j][i] && i-j+1 > maxNum {
				maxNum = i - j + 1
				begin = j
			}
		}
	}
	return s[begin : begin+maxNum]
}
