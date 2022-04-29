package _20

// byte 表示一个字节，rune 表示四个字节
func isValid(s string) bool {
	mapping := map[rune]rune{']': '[', '}': '{', ')': '('}
	stack := make([]rune, 0)
	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack = append(stack, c)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != mapping[c] {
				return false
			}
			// Go的字符串截取和切片是一样的 s [n:m] 左闭右开的原则
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
