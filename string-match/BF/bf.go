package bf

import "fmt"

func BF(s, target string) bool {
	for i := range s {
		for j := range target {
			if i+j == len(s) {
				return false
			}
			if s[i+j] != target[j] {
				break
			}
			if j == len(target)-1 {
				fmt.Println(i, s[i:i+j+1])
				return true
			}
		}
	}
	return false
}

func BFAll(s, target string) []int {
	res := make([]int, 0, len(s)-len(target)+1)
	for i := range s {
		for j := range target {
			if i+j == len(s) {
				break
			}
			if s[i+j] != target[j] {
				break
			}
			if j == len(target)-1 {
				fmt.Println(i, s[i:i+j+1])
				res = append(res, i)
				break
			}
		}
	}
	return res
}

func Replace(s, target, want string) string {
	for i := range s {
		for j := range target {
			if i+j == len(s) {
				break
			}
			if s[i+j] != target[j] {
				break
			}
			if j == len(target)-1 {
				return s[:i] + want + s[j+1:]
			}
		}
	}
	return ""
}
