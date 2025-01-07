package main

import "fmt"

func main() {
	fmt.Println(strStr("lililoc lililac", "lililas"))
}

func strStr(haystack string, needle string) int {
	p := generateP(needle)
	var i, j int

	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				return i - j
			}
		} else {
			if j == 0 {
				i++
			} else {
				j = p[j-1]
			}
		}
	}

	return -1
}

func generateP(needle string) []int {
	var j int
	i := 1
	slice := make([]int, 1, 10)

	for i < len(needle) {
		if needle[i] == needle[j] {
			slice = append(slice, slice[i-1]+1)
			j++
			i++
		} else {
			if j == 0 {
				slice = append(slice, 0)
				i++
			} else {
				j--
			}
		}
	}

	return slice
}
