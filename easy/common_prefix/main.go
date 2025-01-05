package main

import "fmt"

func main() {
	strs1 := []string{"flower", "flow", "flight"}
	strs2 := []string{"dog", "racecar", "car"}
	
	fmt.Println(longestCommonPrefix(strs1))
	fmt.Println(longestCommonPrefix(strs2))
}

func longestCommonPrefix(strs []string) string {
	firstStr := strs[0]
	for i, ch := range firstStr {
		for _, str := range strs[1:] {
			if i == len(str) || str[i] != byte(ch) {
				return firstStr[:i]
			}
		}
	}

	return firstStr
}
