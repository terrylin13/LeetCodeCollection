package _49

import "sort"

// sorted
func groupAnagrams(strs []string) [][]string {
	mapping := map[string][]string{}

	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mapping[sortedStr] = append(mapping[sortedStr], str)
	}

	ans := make([][]string, 0, len(mapping))

	for _, v := range mapping {
		ans = append(ans, v)
	}
	return ans
}

//counting
func groupAnagrams2(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
