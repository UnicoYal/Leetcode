package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	var mx, left, right int
	m := make(map[byte]int)

	for right < len(s) {
		m[s[right]]++
		for m[s[right]] > 1 {
			m[s[left]]--
			left++
		}

		mx = max(mx, right-left+1)
		right++
	}

	return mx
}
