package main

import "fmt"

func main() {
	fmt.Println(longestPalindrom("aabb"))
	fmt.Println(longestPalindrom("aabbb"))
	fmt.Println(longestPalindrom("abcd"))
	fmt.Println(longestPalindrom("aaabbbccccdd"))
}

func longestPalindrom(s string) int {
	m := make(map[byte]int, 10)
	for _, ch := range s {
		m[byte(ch)]++
	}

	var length int
	var isOdd bool

	for _, counter := range m {
		if counter%2 == 1 {
			isOdd = true
		}

		length += counter / 2 * 2
	}

	if isOdd {
		length++
	}

	return length
}
