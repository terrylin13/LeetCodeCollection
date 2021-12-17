package _32

func longestValidParentheses(s string) int {
	stack := []int{-1}
	max := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' { // string start with (
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 { // string start with )
				stack = append(stack, i)
			} else {
				if i-stack[len(stack)-1] > max {
					max = i - stack[len(stack)-1]
				}
			}
		}
	}
	return max
}

func longestValidParentheses2(s string) int {
	max := 0
	l := 0
	r := 0
	for i := 0; i < len(s); i++ { // traversing the string from left towards right
		if s[i] == '(' {
			l++
		}
		if s[i] == ')' {
			r++
		}
		if r > l {
			r = 0
			l = 0
		} else if r == l {
			if r+l > max {
				max = r + l
			}
		}
	}
	l = 0
	r = 0
	for i := len(s) - 1; i >= 0; i-- { // traversing the string from right towards left
		if s[i] == ')' {
			l++
		}
		if s[i] == '(' {
			r++
		}
		if r > l {
			r = 0
			l = 0
		} else if r == l {
			if r+l > max {
				max = r + l
			}
		}
	}
	return max
}
