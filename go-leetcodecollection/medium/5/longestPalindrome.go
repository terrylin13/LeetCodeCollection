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
