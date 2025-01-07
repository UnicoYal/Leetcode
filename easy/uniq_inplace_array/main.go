package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1,1,2}))
}

func removeDuplicates(nums []int) int {
	var i int

	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
