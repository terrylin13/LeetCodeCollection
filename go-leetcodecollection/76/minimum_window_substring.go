package _76

import (
	"fmt"
)

func minWindow(s string, t string) string {
	res := ""
	sLen := len(s)
	mapping := map[byte]int{}
	count := 0
	for _, ch := range []byte(t) {
		if mapping[ch] == 0 {
			count++
		}
		mapping[ch]++

	}
	fmt.Println(count)
	c := 0
	for i, j := 0, 0; i < sLen; i++ {
		if mapping[s[i]] == 1 {
			c++
		}
		mapping[s[i]]--
		// fmt.Println("i:", mapping, s[i], i, j, c)
		for mapping[s[j]] < 0 {
			// fmt.Println("j:", mapping, s[j], i, j, c)
			mapping[s[j]]++
			j++
			if j > i {
				break
			}
		}
		if c == count {
			if res == "" || len(res) > i-j+1 {
				res = s[j : i+1]
			}
		}
	}
	return res
}
