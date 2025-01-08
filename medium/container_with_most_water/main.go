package main

import "fmt"

func main() {
	sl1 := []int{1,8,6,2,5,4,8,3,7}
	sl2 := []int{1,1}
	
	fmt.Println(maxArea(sl1))
	fmt.Println(maxArea(sl2))
}

func maxArea(height []int) int {
	var left, mx int
	right := len(height) - 1

	for left != right {
		curr := (right - left) * min(height[left], height[right])
		mx = max(mx, curr)
		if left < right {
			left++
		} else {
			right--
		}
	}

	return mx
}
