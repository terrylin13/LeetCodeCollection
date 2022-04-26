package _3

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	str := "abcabcbb"
	longest := lengthOfLongestSubstring(str)
	t.Logf("longest:%d", longest)
}
