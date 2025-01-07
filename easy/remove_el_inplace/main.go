package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	var write int

	for read := 0; read < len(nums); read++ {
		if nums[read] != val {
			nums[write] = nums[read]
			write++
		}
	}

	return write
}
